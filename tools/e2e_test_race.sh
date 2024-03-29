#!/bin/bash

# Copyright 2019 The Vitess Authors.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

source build.env

temp_log_file="$(mktemp --suffix .unit_test_race.log)"
trap '[ -f "$temp_log_file" ] && rm $temp_log_file' EXIT

# Wrapper around go test -race.

# This script exists because the -race test doesn't allow to distinguish
# between a failed (e.g. flaky) unit test and a found data race.
# Although Go 1.5 says 'exit status 66' in case of a race, it exits with 1.
# Therefore, we manually check the output of 'go test' for data races and
# exit with an error if one was found.
# TODO(mberlin): Test all packages (go/... instead of go/vt/...) once
#                go/cgzip is moved into a separate repository. We currently
#                skip the cgzip package because -race takes >30 sec for it.

# All endtoend Go packages with test files.
# Output per line: <full Go package name> <all _test.go files in the package>*
packages_with_tests=$(go list -f '{{if len .TestGoFiles}}{{.ImportPath}} {{join .TestGoFiles " "}}{{end}}{{if len .XTestGoFiles}}{{.ImportPath}} {{join .XTestGoFiles " "}}{{end}}' ./go/.../endtoend/... | sort)
packages_with_tests=$(echo "$packages_with_tests" |  grep -vE "go/test/endtoend" | cut -d" " -f1)

# endtoend tests should be in a directory called endtoend
all_e2e_tests=$(echo "$packages_with_tests" | cut -d" " -f1)

set -exo pipefail

# Run all endtoend tests.
echo "$all_e2e_tests" | xargs go test $VT_GO_PARALLEL -race 2>&1 | tee $temp_log_file
if [ ${PIPESTATUS[1]} -ne 0 ]; then
  if grep "WARNING: DATA RACE" -q $temp_log_file; then
    echo
    echo "ERROR: go test -race found a data race. See log above."
    exit 2
  fi

  echo "ERROR: go test -race found NO data race, but failed. See log above."
  exit 1
fi

echo
echo "SUCCESS: No data race was found."
