//go:build !windows

/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package events

import (
	"fmt"
	"log/syslog"

	"vitess.io/vitess/go/event/syslogger"
)

// Syslog writes the event to syslog.
func (sc *ShardChange) Syslog() (syslog.Priority, string) {
	return syslog.LOG_INFO, fmt.Sprintf("%s/%s [shard] %s value: %s",
		sc.KeyspaceName, sc.ShardName, sc.Status, sc.Shard.String())
}

var _ syslogger.Syslogger = (*ShardChange)(nil) // compile-time interface check
