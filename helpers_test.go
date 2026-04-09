// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type mockCommander struct {
	commandsSent []string
	err          error
}

func (m *mockCommander) Command(_ context.Context, format string, a ...any) error {
	if m.err != nil {
		return m.err
	}
	cmd := fmt.Sprintf(format, a...)
	m.commandsSent = append(m.commandsSent, cmd)
	return nil
}

func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		args     []any
		expected string
	}{
		{"no args", "OUTP ON", nil, "OUTP ON"},
		{"with float", "VOLT %f", []any{1.5}, "VOLT 1.500000"},
		{"with string", "INST %s", []any{"P6V"}, "INST P6V"},
		{
			"with multiple args",
			"CURR %f;:CURR:PROT %f",
			[]any{0.5, 0.5},
			"CURR 0.500000;:CURR:PROT 0.500000",
		},
	}

	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockCommander{}
			err := Set(ctx, mock, tt.format, tt.args...)
			if err != nil {
				t.Errorf("Set() unexpected error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.expected {
				t.Errorf("Set() sent %v, want [%q]", mock.commandsSent, tt.expected)
			}
		})
	}
}

func TestSet_Error(t *testing.T) {
	mock := &mockCommander{err: errors.New("mock error")}
	err := Set(context.Background(), mock, "VOLT %f", 1.0)
	if err == nil {
		t.Error("Set() expected error, got nil")
	}
}

func TestQueryID(t *testing.T) {
	mock := &mockInstrument{}
	ctx := context.Background()
	result, err := QueryID(ctx, mock)
	if err != nil {
		t.Errorf("QueryID() unexpected error: %v", err)
	}
	if result != "query response" {
		t.Errorf("QueryID() = %q, want %q", result, "query response")
	}
}

func TestQueryID_Error(t *testing.T) {
	mock := &mockInstrument{shouldError: true}
	ctx := context.Background()
	_, err := QueryID(ctx, mock)
	if err == nil {
		t.Error("QueryID() expected error, got nil")
	}
}
