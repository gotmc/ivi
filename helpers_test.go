// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"testing"
)

type mockStringWriter struct {
	written string
	err     error
}

func (m *mockStringWriter) WriteString(s string) (int, error) {
	if m.err != nil {
		return 0, m.err
	}
	m.written = s
	return len(s), nil
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockStringWriter{}
			err := Set(mock, tt.format, tt.args...)
			if err != nil {
				t.Errorf("Set() unexpected error: %v", err)
			}
			if mock.written != tt.expected {
				t.Errorf("Set() wrote %q, want %q", mock.written, tt.expected)
			}
		})
	}
}

func TestSet_Error(t *testing.T) {
	mock := &mockStringWriter{err: ErrUnexpectedResponse}
	err := Set(mock, "VOLT %f", 1.0)
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
