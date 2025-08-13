// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"testing"
)

// mockInstrumentWithClose simulates an instrument that implements Close()
type mockInstrumentWithClose struct {
	commandsSent     []string
	closeCalled      bool
	shouldError      bool
	errorOnFirstOnly bool // Only error on first command, succeed on second
}

func (m *mockInstrumentWithClose) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (m *mockInstrumentWithClose) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockInstrumentWithClose) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockInstrumentWithClose) Command(format string, a ...any) error {
	m.commandsSent = append(m.commandsSent, format)

	if m.shouldError && !m.errorOnFirstOnly {
		return ErrUnexpectedResponse
	}

	if m.errorOnFirstOnly && len(m.commandsSent) == 1 {
		return ErrUnexpectedResponse
	}

	return nil
}

func (m *mockInstrumentWithClose) Query(s string) (value string, err error) {
	return "", nil
}

func (m *mockInstrumentWithClose) Close() error {
	m.closeCalled = true
	return nil
}

func TestInherent_Disable(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{})

	err := inherent.Disable()
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

func TestInherent_Disable_Fallback(t *testing.T) {
	mock := &mockInstrumentWithClose{errorOnFirstOnly: true}
	inherent := NewInherent(mock, InherentBase{})

	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error after fallback, got %v", err)
	}

	// Should have tried both SYST:LOC and SYSTem:LOCal
	if len(mock.commandsSent) != 2 {
		t.Errorf("Expected 2 commands (with fallback), got %d", len(mock.commandsSent))
		return
	}

	if mock.commandsSent[0] != "SYST:LOC" {
		t.Errorf("Expected SYST:LOC as first command, got %s", mock.commandsSent[0])
	}

	if mock.commandsSent[1] != "SYSTem:LOCal" {
		t.Errorf("Expected SYSTem:LOCal as fallback command, got %s", mock.commandsSent[1])
	}
}

func TestInherent_Close(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{})

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

func TestInherent_Close_WithoutCloser(t *testing.T) {
	// Use the regular mock that doesn't implement Close()
	mock := &mockInstrument{}
	inherent := NewInherent(mock, InherentBase{})

	err := inherent.Close()
	if err != nil {
		t.Errorf("Expected no error when underlying instrument doesn't implement Close(), got %v", err)
	}
}

func TestInherent_ReturnToLocal_Control(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{})

	// Should default to true
	if !inherent.GetReturnToLocal() {
		t.Error("Expected ReturnToLocal to default to true")
	}

	// Test disabling return to local
	inherent.SetReturnToLocal(false)
	if inherent.GetReturnToLocal() {
		t.Error("Expected ReturnToLocal to be false after setting")
	}

	// Disable should not send commands when ReturnToLocal is false
	err := inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 0 {
		t.Errorf("Expected no commands when ReturnToLocal is false, got %d", len(mock.commandsSent))
	}

	// Re-enable and verify it works
	inherent.SetReturnToLocal(true)
	err = inherent.Disable()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(mock.commandsSent) != 1 {
		t.Errorf("Expected 1 command when ReturnToLocal is true, got %d", len(mock.commandsSent))
	}
}

