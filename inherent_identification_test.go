// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"errors"
	"testing"
	"time"
)

// mockIDNInstrument returns a configurable *IDN? response.
type mockIDNInstrument struct {
	idnResponse string
	shouldError bool
}

func (m *mockIDNInstrument) ReadBinary(_ context.Context, p []byte) (int, error) {
	if m.shouldError {
		return 0, errors.New("mock read error")
	}
	n := copy(p, []byte(m.idnResponse+"\n"))
	return n, nil
}

func (m *mockIDNInstrument) WriteBinary(_ context.Context, p []byte) (int, error) {
	return len(p), nil
}

func (m *mockIDNInstrument) Close() error { return nil }

func (m *mockIDNInstrument) Command(_ context.Context, _ string, _ ...any) error {
	return nil
}

func (m *mockIDNInstrument) Query(_ context.Context, _ string) (string, error) {
	if m.shouldError {
		return "", errors.New("mock query error")
	}
	return m.idnResponse, nil
}

func TestInherent_InstrumentManufacturer(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01-03.01-03.01-00.52-04-02",
	}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)
	mfr, err := inherent.InstrumentManufacturer()
	if err != nil {
		t.Fatalf("InstrumentManufacturer() error: %v", err)
	}
	if mfr != "KEYSIGHT TECHNOLOGIES" {
		t.Errorf("InstrumentManufacturer() = %q, want %q", mfr, "KEYSIGHT TECHNOLOGIES")
	}
}

func TestInherent_InstrumentModel(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01-03.01-03.01-00.52-04-02",
	}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)
	model, err := inherent.InstrumentModel()
	if err != nil {
		t.Fatalf("InstrumentModel() error: %v", err)
	}
	if model != "34465A" {
		t.Errorf("InstrumentModel() = %q, want %q", model, "34465A")
	}
}

func TestInherent_InstrumentSerialNumber(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01",
	}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)
	sn, err := inherent.InstrumentSerialNumber()
	if err != nil {
		t.Fatalf("InstrumentSerialNumber() error: %v", err)
	}
	if sn != "MY54505281" {
		t.Errorf("InstrumentSerialNumber() = %q, want %q", sn, "MY54505281")
	}
}

func TestInherent_FirmwareRevision(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01",
	}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)
	fw, err := inherent.FirmwareRevision()
	if err != nil {
		t.Fatalf("FirmwareRevision() error: %v", err)
	}
	if fw != "A.03.01" {
		t.Errorf("FirmwareRevision() = %q, want %q", fw, "A.03.01")
	}
}

func TestInherent_Identification_QueryError(t *testing.T) {
	mock := &mockIDNInstrument{shouldError: true}
	inherent := NewInherent(mock, InherentBase{ReturnToLocal: true}, 0)
	_, err := inherent.InstrumentManufacturer()
	if err == nil {
		t.Error("expected error from query failure, got nil")
	}
}

func TestInherent_CheckID(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01",
	}
	inherent := NewInherent(mock, InherentBase{
		SupportedInstrumentModels: []string{"34460A", "34461A", "34465A", "34470A"},
	}, 0)
	model, err := inherent.CheckID()
	if err != nil {
		t.Fatalf("CheckID() error: %v", err)
	}
	if model != "34465A" {
		t.Errorf("CheckID() model = %q, want %q", model, "34465A")
	}
	if inherent.IDNString == "" {
		t.Error("CheckID() did not populate IDNString")
	}
}

func TestInherent_CheckID_UnsupportedModel(t *testing.T) {
	mock := &mockIDNInstrument{
		idnResponse: "KEYSIGHT TECHNOLOGIES,34465A,MY54505281,A.03.01",
	}
	inherent := NewInherent(mock, InherentBase{
		SupportedInstrumentModels: []string{"33220A"},
	}, 0)
	_, err := inherent.CheckID()
	if err == nil {
		t.Fatal("CheckID() expected error for unsupported model, got nil")
	}
	if !errors.Is(err, ErrUnsupportedModel) {
		t.Errorf("CheckID() error = %v, want ErrUnsupportedModel", err)
	}
}

func TestInherent_CheckID_QueryError(t *testing.T) {
	mock := &mockIDNInstrument{shouldError: true}
	inherent := NewInherent(mock, InherentBase{
		SupportedInstrumentModels: []string{"34465A"},
	}, 0)
	_, err := inherent.CheckID()
	if err == nil {
		t.Error("CheckID() expected error from query failure, got nil")
	}
}

func TestInherent_Reset(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{
		ReturnToLocal: true,
		ResetDelay:    1 * time.Millisecond,
	}, 0)

	err := inherent.Reset()
	if err != nil {
		t.Errorf("Reset() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "*rst" {
		t.Errorf("Reset() sent %v, want [\"*rst\"]", mock.commandsSent)
	}
}

func TestInherent_Clear(t *testing.T) {
	mock := &mockInstrumentWithClose{}
	inherent := NewInherent(mock, InherentBase{
		ReturnToLocal: true,
		ClearDelay:    1 * time.Millisecond,
	}, 0)

	err := inherent.Clear()
	if err != nil {
		t.Errorf("Clear() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "*cls" {
		t.Errorf("Clear() sent %v, want [\"*cls\"]", mock.commandsSent)
	}
}
