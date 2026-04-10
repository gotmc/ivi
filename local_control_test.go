// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"testing"
)

// mockInstrumentWithClose simulates an instrument that implements Close()
type mockInstrumentWithClose struct {
	commandsSent []string
	closeCalled  bool
	shouldError  bool
}

func (m *mockInstrumentWithClose) ReadContext(_ context.Context, p []byte) (int, error) {
	return 0, nil
}

func (m *mockInstrumentWithClose) WriteContext(_ context.Context, p []byte) (int, error) {
	return len(p), nil
}

func (m *mockInstrumentWithClose) Command(_ context.Context, format string, a ...any) error {
	m.commandsSent = append(m.commandsSent, format)

	if m.shouldError {
		return ErrUnexpectedResponse
	}

	return nil
}

func (m *mockInstrumentWithClose) Query(_ context.Context, s string) (value string, err error) {
	return "", nil
}

func (m *mockInstrumentWithClose) Close() error {
	m.closeCalled = true
	return nil
}

func TestInherent_Disable(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true})

	err := inherent.Disable(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Should have sent SYST:LOC command
	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command, got %d", len(mock.commandsSent))
	}

	if mock.commandsSent[0] != "SYST:LOC" {
		t.Errorf("Expected SYST:LOC command, got %s", mock.commandsSent[0])
	}
}

func TestInherent_Disable_ErrorIgnored(t *testing.T) {
	mock := &mockInstrumentWithClose{shouldError: true}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true})

	err := inherent.Disable(context.Background())
	if err != nil {
		t.Errorf("Expected no error (best-effort), got %v", err)
	}

	// Should have attempted SYST:LOC even though it failed
	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command attempt, got %d", len(mock.commandsSent))
	}
}

func TestInherent_Close(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true})

	err := inherent.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Should have sent SYST:LOC command (from Disable)
	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command, got %d", len(mock.commandsSent))
	}

	if mock.commandsSent[0] != "SYST:LOC" {
		t.Errorf("Expected SYST:LOC command, got %s", mock.commandsSent[0])
	}

	// Should have called Close on the underlying instrument
	if !mock.closeCalled {
		t.Error("Expected Close() to be called on underlying instrument")
	}
}

func TestInherent_ReturnToLocal_Control(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true})

	// Should be true when explicitly set
	if !inherent.IsReturnToLocal() {
		t.Error("Expected ReturnToLocal to be true")
	}

	// Test disabling return to local
	inherent.SetReturnToLocal(false)
	if inherent.IsReturnToLocal() {
		t.Error("Expected ReturnToLocal to be false after setting")
	}

	// Disable should not send commands when ReturnToLocal is false
	err := inherent.Disable(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 0 {
		t.Errorf("Expected no commands when ReturnToLocal is false, got %d", len(mock.commandsSent))
	}

	// Re-enable and verify it works
	inherent.SetReturnToLocal(true)
	err = inherent.Disable(context.Background())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command when ReturnToLocal is true, got %d", len(mock.commandsSent))
	}
}
