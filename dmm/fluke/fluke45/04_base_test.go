// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fluke45

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
)

// mockInst captures commands and returns configurable query responses.
type mockInst struct {
	commandsSent []string
	queryResp    string
	shouldError  bool
}

func (m *mockInst) ReadBinary(_ context.Context, p []byte) (int, error) {
	return 0, nil
}

func (m *mockInst) WriteBinary(_ context.Context, p []byte) (int, error) {
	return len(p), nil
}

func (m *mockInst) Close() error { return nil }

func (m *mockInst) Command(_ context.Context, format string, a ...any) error {
	if m.shouldError {
		return errors.New("mock command error")
	}
	cmd := fmt.Sprintf(format, a...)
	m.commandsSent = append(m.commandsSent, cmd)
	return nil
}

func (m *mockInst) Query(_ context.Context, s string) (string, error) {
	if m.shouldError {
		return "", errors.New("mock query error")
	}
	return m.queryResp, nil
}

func TestDetermineRangeCommand(t *testing.T) {
	tests := []struct {
		name       string
		rate       string
		fcn        dmm.MeasurementFunction
		rangeValue float64
		want       string
		wantErr    bool
	}{
		// DC Volts, Fast/Medium rate
		{"dcv fast 0.3V", "F", dmm.DCVolts, 0.3, "1", false},
		{"dcv medium 3V", "M", dmm.DCVolts, 3.0, "2", false},
		{"dcv fast 30V", "F", dmm.DCVolts, 30.0, "3", false},
		{"dcv fast 300V", "F", dmm.DCVolts, 300.0, "4", false},
		{"dcv fast 1000V", "F", dmm.DCVolts, 1000.0, "5", false},

		// DC Volts, Slow rate
		{"dcv slow 0.1V", "S", dmm.DCVolts, 0.1, "1", false},
		{"dcv slow 1V", "S", dmm.DCVolts, 1.0, "2", false},
		{"dcv slow 10V", "S", dmm.DCVolts, 10.0, "3", false},
		{"dcv slow 100V", "S", dmm.DCVolts, 100.0, "4", false},
		{"dcv slow 1000V", "S", dmm.DCVolts, 1000.0, "5", false},

		// AC Volts, Fast rate
		{"acv fast 0.3V", "F", dmm.ACVolts, 0.3, "1", false},
		{"acv fast 3V", "F", dmm.ACVolts, 3.0, "2", false},
		{"acv fast 750V", "F", dmm.ACVolts, 750.0, "5", false},

		// AC Volts, Slow rate
		{"acv slow 0.1V", "S", dmm.ACVolts, 0.1, "1", false},
		{"acv slow 750V", "S", dmm.ACVolts, 750.0, "5", false},

		// 2-wire Resistance, Fast rate
		{"ohms fast 300", "F", dmm.TwoWireResistance, 300, "1", false},
		{"ohms fast 3k", "F", dmm.TwoWireResistance, 3e3, "2", false},
		{"ohms fast 300M", "F", dmm.TwoWireResistance, 300e6, "7", false},

		// 2-wire Resistance, Slow rate
		{"ohms slow 100", "S", dmm.TwoWireResistance, 100, "1", false},
		{"ohms slow 100M", "S", dmm.TwoWireResistance, 100e6, "7", false},

		// Error: out of range
		{"dcv fast out of range", "F", dmm.DCVolts, 1001.0, "", true},
		// Error: unsupported function
		{"frequency unsupported", "F", dmm.Frequency, 100.0, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := determineRangeCommand(tt.rate, tt.fcn, tt.rangeValue)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDriver_MeasurementFunction(t *testing.T) {
	tests := []struct {
		name    string
		resp    string
		want    dmm.MeasurementFunction
		wantErr bool
	}{
		{"dc volts", "VDC", dmm.DCVolts, false},
		{"ac volts", "VAC", dmm.ACVolts, false},
		{"dc current", "ADC", dmm.DCCurrent, false},
		{"resistance", "OHMS", dmm.TwoWireResistance, false},
		{"frequency", "FREQ", dmm.Frequency, false},
		{"unknown", "UNKNOWN", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, err := New(mock, ivi.WithoutIDQuery())
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			got, err := d.MeasurementFunction()
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetMeasurementFunction(t *testing.T) {
	tests := []struct {
		name    string
		fcn     dmm.MeasurementFunction
		wantCmd string
		wantErr bool
	}{
		{"dc volts", dmm.DCVolts, `"VDC"`, false},
		{"ac volts", dmm.ACVolts, `"VAC"`, false},
		{"dc current", dmm.DCCurrent, `"ADC"`, false},
		{"unsupported", dmm.Temperature, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, err := New(mock, ivi.WithoutIDQuery())
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			err = d.SetMeasurementFunction(tt.fcn)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestDriver_TriggerSource(t *testing.T) {
	tests := []struct {
		name    string
		resp    string
		want    dmm.TriggerSource
		wantErr bool
	}{
		{"internal", "1", dmm.TriggerSourceImmediate, false},
		{"external", "2", dmm.TriggerSourceExternal, false},
		{"external with delay", "3", dmm.TriggerSourceExternal, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, err := New(mock, ivi.WithoutIDQuery())
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			got, err := d.TriggerSource()
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetTriggerSource(t *testing.T) {
	tests := []struct {
		name    string
		src     dmm.TriggerSource
		wantCmd string
		wantErr bool
	}{
		{"immediate", dmm.TriggerSourceImmediate, "TRIGGER 1", false},
		{"external", dmm.TriggerSourceExternal, "TRIGGER 2", false},
		{"software", dmm.TriggerSourceSoftware, "TRIGGER 2", false},
		{"unsupported", dmm.TriggerSource(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, err := New(mock, ivi.WithoutIDQuery())
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			err = d.SetTriggerSource(tt.src)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestDriver_ResolutionAbsolute_NotSupported(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	_, err = d.ResolutionAbsolute()
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("ResolutionAbsolute() = %v, want ErrFunctionNotSupported", err)
	}

	err = d.SetResolutionAbsolute(1.0)
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("SetResolutionAbsolute() = %v, want ErrFunctionNotSupported", err)
	}
}

func TestDriver_Abort_NotSupported(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	err = d.Abort()
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("Abort() = %v, want ErrFunctionNotSupported", err)
	}
}

func TestDriver_InitiateMeasurement(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	err = d.InitiateMeasurement()
	if err != nil {
		t.Errorf("InitiateMeasurement() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "*TRG" {
		t.Errorf("sent %v, want [\"*TRG\"]", mock.commandsSent)
	}
}

func TestDriver_IsOutOfRange_NotSupported(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	_, err = d.IsOutOfRange(1.0)
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("IsOutOfRange() = %v, want ErrFunctionNotSupported", err)
	}

	_, err = d.IsOverRange(1.0)
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("IsOverRange() = %v, want ErrFunctionNotSupported", err)
	}

	_, err = d.IsUnderRange(1.0)
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("IsUnderRange() = %v, want ErrFunctionNotSupported", err)
	}
}

func TestDriver_TriggerDelay(t *testing.T) {
	tests := []struct {
		name         string
		resp         string
		wantHasDelay bool
		wantDuration time.Duration
		wantErr      bool
	}{
		{"internal no delay", "1", false, 0, false},
		{"external no delay", "2", false, 0, false},
		{"external with delay", "3", true, 0, false},
		{"external with delay type 5", "5", true, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, err := New(mock, ivi.WithoutIDQuery())
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			hasDelay, dur, err := d.TriggerDelay()
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if hasDelay != tt.wantHasDelay {
				t.Errorf("hasDelay = %v, want %v", hasDelay, tt.wantHasDelay)
			}
			if dur != tt.wantDuration {
				t.Errorf("duration = %v, want %v", dur, tt.wantDuration)
			}
		})
	}
}

func TestDriver_FetchMeasurement(t *testing.T) {
	mock := &mockInst{queryResp: "1.234"}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	got, err := d.FetchMeasurement(0)
	if err != nil {
		t.Errorf("FetchMeasurement() error: %v", err)
	}
	if got != 1.234 {
		t.Errorf("FetchMeasurement() = %f, want 1.234", got)
	}
}

func TestDriver_ReadMeasurement(t *testing.T) {
	mock := &mockInst{queryResp: "5.678"}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	got, err := d.ReadMeasurement(0)
	if err != nil {
		t.Errorf("ReadMeasurement() error: %v", err)
	}
	if got != 5.678 {
		t.Errorf("ReadMeasurement() = %f, want 5.678", got)
	}
}
