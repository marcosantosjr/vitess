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

package helpers

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/vt/sqlparser"

	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
)

func TestTee(t *testing.T) {
	ctx := context.Background()

	// create the setup, copy the data
	fromTS, toTS := createSetup(ctx, t)
	CopyKeyspaces(ctx, fromTS, toTS, sqlparser.NewTestParser())
	CopyShards(ctx, fromTS, toTS)
	CopyTablets(ctx, fromTS, toTS)

	// create a tee and check it implements the interface.
	teeTS, err := NewTee(fromTS, toTS, true)
	require.NoError(t, err)

	// create a keyspace, make sure it is on both sides
	if err := teeTS.CreateKeyspace(ctx, "keyspace2", &topodatapb.Keyspace{}); err != nil {
		t.Fatalf("tee.CreateKeyspace(keyspace2) failed: %v", err)
	}
	teeKeyspaces, err := teeTS.GetKeyspaces(ctx)
	if err != nil {
		t.Fatalf("tee.GetKeyspaces() failed: %v", err)
	}
	expected := []string{"keyspace2", "test_keyspace"}
	if !reflect.DeepEqual(expected, teeKeyspaces) {
		t.Errorf("teeKeyspaces mismatch, got %+v, want %+v", teeKeyspaces, expected)
	}
	fromKeyspaces, err := fromTS.GetKeyspaces(ctx)
	if err != nil {
		t.Fatalf("fromTS.GetKeyspaces() failed: %v", err)
	}
	expected = []string{"keyspace2", "test_keyspace"}
	if !reflect.DeepEqual(expected, fromKeyspaces) {
		t.Errorf("fromKeyspaces mismatch, got %+v, want %+v", fromKeyspaces, expected)
	}
	toKeyspaces, err := toTS.GetKeyspaces(ctx)
	if err != nil {
		t.Fatalf("toTS.GetKeyspaces() failed: %v", err)
	}
	expected = []string{"keyspace2", "test_keyspace"}
	if !reflect.DeepEqual(expected, toKeyspaces) {
		t.Errorf("toKeyspaces mismatch, got %+v, want %+v", toKeyspaces, expected)
	}
}
