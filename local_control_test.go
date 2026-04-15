// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"testing"
)

// mockLocalControlInst simulates an instrument for local control testing.
type mockLocalControlInst struct {
	commandsSent []string
	shouldError  bool
}

func (m *mockLocalControlInst) ReadBinary(_ context.Context, p []byte) (int, error) {
	return 0, nil
}

func (m *mockLocalControlInst) WriteBinary(_ context.Context, p []byte) (int, error) {
	return len(p), nil
}

func (m *mockLocalControlInst) Command(_ context.Context, format string, a ...any) error {
	m.commandsSent = append(m.commandsSent, format)

	if m.shouldError {
		return ErrUnexpectedResponse
	}

	return nil
}

func (m *mockLocalControlInst) Query(_ context.Context, s string) (value string, err error) {
	return "", nil
}

func TestInherent_Disable(t *testing.T) {
	mock := &mockLocalControlInst{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)

	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command, got %d", len(mock.commandsSent))
	}

	if mock.commandsSent[0] != "SYST:LOC" {
		t.Errorf("Expected SYST:LOC command, got %s", mock.commandsSent[0])
	}
}

func TestInherent_Disable_CustomCommand(t *testing.T) {
	mock := &mockLocalControlInst{}
	inherent := NewInherent(mock, InherentBase{
		ReturnToLocal:       true,
		LocalControlCommand: "SYST:COMM:RLST LOC",
	}, 0)

	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Fatalf("Expected 1 command, got %d", len(mock.commandsSent))
	}

	if mock.commandsSent[0] != "SYST:COMM:RLST LOC" {
		t.Errorf(
			"Expected SYST:COMM:RLST LOC command, got %s",
			mock.commandsSent[0],
		)
	}
}

func TestInherent_Disable_ErrorIgnored(t *testing.T) {
	mock := &mockLocalControlInst{shouldError: true}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)

	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error (best-effort), got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command attempt, got %d", len(mock.commandsSent))
	}
}

func TestInherent_Close(t *testing.T) {
	mock := &mockLocalControlInst{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)

	err := inherent.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Close should send the local control command but not close the transport.
	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command, got %d", len(mock.commandsSent))
	}

	if mock.commandsSent[0] != "SYST:LOC" {
		t.Errorf("Expected SYST:LOC command, got %s", mock.commandsSent[0])
	}
}

func TestInherent_ReturnToLocal_Control(t *testing.T) {
	mock := &mockLocalControlInst{}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)

	if !inherent.IsReturnToLocal() {
		t.Error("Expected ReturnToLocal to be true")
	}

	inherent.SetReturnToLocal(false)
	if inherent.IsReturnToLocal() {
		t.Error("Expected ReturnToLocal to be false after setting")
	}

	// Disable should not send commands when ReturnToLocal is false.
	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 0 {
		t.Errorf(
			"Expected no commands when ReturnToLocal is false, got %d",
			len(mock.commandsSent),
		)
	}

	// Re-enable and verify it works.
	inherent.SetReturnToLocal(true)
	err = inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Errorf(
			"Expected 1 command when ReturnToLocal is true, got %d",
			len(mock.commandsSent),
		)
	}
}
