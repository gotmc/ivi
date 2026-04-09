// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
)

func TestCreateConfigureVoltageACCommand(t *testing.T) {
	testCases := []struct {
		autoRange   dmm.AutoRange
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{dmm.AutoOn, 0.0, "CONF:VOLT:AC AUTO", nil},
		{dmm.AutoOff, 0.09, "CONF:VOLT:AC 0.1", nil},
		{dmm.AutoOn, 10.0, "CONF:VOLT:AC AUTO", nil},
		{dmm.AutoOff, 9.0, "CONF:VOLT:AC 10", nil},
		{dmm.AutoOff, 10.0, "CONF:VOLT:AC 10", nil},
		{dmm.AutoOff, 10.001, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 99.99, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 100.00, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 100.001, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 999.999, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 1000.0, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 1000.01, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := createConfigureVoltageACCommand(tc.autoRange, tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf("wanted err %s / got err %s", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Errorf("wanted %v / got %v", tc.expected, got)
		}
	}
}

func TestDetermineManualDCCurrentRange(t *testing.T) {
	testCases := []struct {
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{0.5e-6, "1e-6", nil},
		{1e-6, "1e-6", nil},
		{5e-6, "10e-6", nil},
		{50e-6, "100e-6", nil},
		{500e-6, "1e-3", nil},
		{5e-3, "10e-3", nil},
		{50e-3, "100e-3", nil},
		{0.5, "1", nil},
		{2.0, "3", nil},
		{3.0, "3", nil},
		{3.01, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := determineManualDCCurrentRange(tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf(
				"rangeValue=%g: wanted err %v / got err %v",
				tc.rangeValue, tc.expectedErr, err,
			)
		}

		if got != tc.expected {
			t.Errorf("rangeValue=%g: wanted %v / got %v", tc.rangeValue, tc.expected, got)
		}
	}
}

func TestDetermineManualACCurrentRange(t *testing.T) {
	testCases := []struct {
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{50e-6, "100e-6", nil},
		{100e-6, "100e-6", nil},
		{500e-6, "1e-3", nil},
		{5e-3, "10e-3", nil},
		{50e-3, "100e-3", nil},
		{0.5, "1", nil},
		{2.0, "3", nil},
		{3.0, "3", nil},
		{3.01, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := determineManualACCurrentRange(tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf(
				"rangeValue=%g: wanted err %v / got err %v",
				tc.rangeValue, tc.expectedErr, err,
			)
		}

		if got != tc.expected {
			t.Errorf("rangeValue=%g: wanted %v / got %v", tc.rangeValue, tc.expected, got)
		}
	}
}

func TestDetermineManualFrequencyVoltageRange(t *testing.T) {
	testCases := []struct {
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{0.05, "0.1", nil},
		{0.1, "0.1", nil},
		{0.5, "1", nil},
		{5.0, "10", nil},
		{50.0, "100", nil},
		{500.0, "750", nil},
		{750.0, "750", nil},
		{751.0, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := determineManualFrequencyVoltageRange(tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf(
				"rangeValue=%g: wanted err %v / got err %v",
				tc.rangeValue, tc.expectedErr, err,
			)
		}

		if got != tc.expected {
			t.Errorf("rangeValue=%g: wanted %v / got %v", tc.rangeValue, tc.expected, got)
		}
	}
}

func TestDetermineVoltageRange(t *testing.T) {
	testCases := []struct {
		autoRange   dmm.AutoRange
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{dmm.AutoOn, 0.0, "AUTO", nil},
		{dmm.AutoOff, 0.09, "0.1", nil},
		{dmm.AutoOn, 10.0, "AUTO", nil},
		{dmm.AutoOff, 9.0, "10", nil},
		{dmm.AutoOff, 10.0, "10", nil},
		{dmm.AutoOff, 10.001, "100", nil},
		{dmm.AutoOff, 99.99, "100", nil},
		{dmm.AutoOff, 100.00, "100", nil},
		{dmm.AutoOff, 100.001, "1000", nil},
		{dmm.AutoOff, 999.999, "1000", nil},
		{dmm.AutoOff, 1000.0, "1000", nil},
		{dmm.AutoOff, 1000.01, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := determineVoltageRange(tc.autoRange, tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf("wanted err %s / got err %s", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Errorf("wanted %v / got %v", tc.expected, got)
		}
	}
}

// mockInst implements the ivi.Instrument interface for unit testing.
type mockInst struct {
	commandsSent []string
	queryResp    string
	shouldError  bool
}

func (m *mockInst) Read(p []byte) (int, error)        { return 0, nil }
func (m *mockInst) Write(p []byte) (int, error)       { return len(p), nil }
func (m *mockInst) WriteString(s string) (int, error) { return len(s), nil }

func (m *mockInst) ReadContext(_ context.Context, p []byte) (int, error) {
	return m.Read(p)
}

func (m *mockInst) WriteContext(_ context.Context, p []byte) (int, error) {
	return m.Write(p)
}

func (m *mockInst) Command(_ context.Context, format string, a ...any) error {
	if m.shouldError {
		return errors.New("mock command error")
	}

	cmd := fmt.Sprintf(format, a...)
	m.commandsSent = append(m.commandsSent, cmd)

	return nil
}

func (m *mockInst) Query(_ context.Context, _ string) (string, error) {
	if m.shouldError {
		return "", errors.New("mock query error")
	}

	return m.queryResp, nil
}

// newTestDriver creates a Driver with the given mock instrument for testing.
func newTestDriver(t *testing.T, mock *mockInst) *Driver {
	t.Helper()

	d, err := New(mock, false)
	if err != nil {
		t.Fatalf("New() returned unexpected error: %v", err)
	}

	return d
}

func TestMeasurementFunction(t *testing.T) {
	testCases := []struct {
		name        string
		queryResp   string
		shouldError bool
		expected    dmm.MeasurementFunction
		expectErr   bool
	}{
		{"dc volts", "VOLT", false, dmm.DCVolts, false},
		{"ac volts", "VOLT:AC", false, dmm.ACVolts, false},
		{"dc current", "CURR", false, dmm.DCCurrent, false},
		{"ac current", "CURR:AC", false, dmm.ACCurrent, false},
		{"two wire resistance", "RES", false, dmm.TwoWireResistance, false},
		{"four wire resistance", "FRES", false, dmm.FourWireResistance, false},
		{"frequency", "FREQ", false, dmm.Frequency, false},
		{"period", "PER", false, dmm.Period, false},
		{"temperature", "TEMP", false, dmm.Temperature, false},
		{"dc volts alternate", "VOLT:DC", false, dmm.DCVolts, false},
		{"dc current alternate", "CURR:DC", false, dmm.DCCurrent, false},
		{"unknown function", "INVALID", false, 0, true},
		{"query error", "", true, 0, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockInst{
				queryResp:   tc.queryResp,
				shouldError: tc.shouldError,
			}
			d := newTestDriver(t, mock)
			ctx := context.Background()

			got, err := d.MeasurementFunction(ctx)

			if tc.expectErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectErr && got != tc.expected {
				t.Errorf("wanted %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSetMeasurementFunction(t *testing.T) {
	testCases := []struct {
		name        string
		msrFunc     dmm.MeasurementFunction
		shouldError bool
		expectedCmd string
		expectErr   bool
	}{
		{"dc volts", dmm.DCVolts, false, `FUNC "VOLT"`, false},
		{"ac volts", dmm.ACVolts, false, `FUNC "VOLT:AC"`, false},
		{"dc current", dmm.DCCurrent, false, `FUNC "CURR"`, false},
		{"ac current", dmm.ACCurrent, false, `FUNC "CURR:AC"`, false},
		{"two wire resistance", dmm.TwoWireResistance, false, `FUNC "RES"`, false},
		{"four wire resistance", dmm.FourWireResistance, false, `FUNC "FRES"`, false},
		{"frequency", dmm.Frequency, false, `FUNC "FREQ"`, false},
		{"period", dmm.Period, false, `FUNC "PER"`, false},
		{"temperature", dmm.Temperature, false, `FUNC "TEMP"`, false},
		{"command error", dmm.DCVolts, true, "", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockInst{shouldError: tc.shouldError}
			d := newTestDriver(t, mock)
			ctx := context.Background()

			err := d.SetMeasurementFunction(ctx, tc.msrFunc)

			if tc.expectErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectErr {
				if len(mock.commandsSent) != 1 {
					t.Fatalf("expected 1 command sent, got %d", len(mock.commandsSent))
				}

				if mock.commandsSent[0] != tc.expectedCmd {
					t.Errorf("wanted command %q, got %q", tc.expectedCmd, mock.commandsSent[0])
				}
			}
		})
	}
}

func TestTriggerSource(t *testing.T) {
	testCases := []struct {
		name        string
		queryResp   string
		shouldError bool
		expected    dmm.TriggerSource
		expectErr   bool
	}{
		{"immediate", "IMM", false, dmm.Immediate, false},
		{"external", "EXT", false, dmm.External, false},
		{"software trigger", "BUS", false, dmm.SoftwareTrigger, false},
		{"unknown source", "UNKNOWN", false, 0, true},
		{"query error", "", true, 0, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockInst{
				queryResp:   tc.queryResp,
				shouldError: tc.shouldError,
			}
			d := newTestDriver(t, mock)
			ctx := context.Background()

			got, err := d.TriggerSource(ctx)

			if tc.expectErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectErr && got != tc.expected {
				t.Errorf("wanted %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestSetTriggerSource(t *testing.T) {
	testCases := []struct {
		name        string
		src         dmm.TriggerSource
		shouldError bool
		expectedCmd string
		expectErr   bool
	}{
		{"immediate", dmm.Immediate, false, "TRIG:SOUR IMM", false},
		{"external", dmm.External, false, "TRIG:SOUR EXT", false},
		{"software trigger", dmm.SoftwareTrigger, false, "TRIG:SOUR BUS", false},
		{"command error", dmm.Immediate, true, "", true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockInst{shouldError: tc.shouldError}
			d := newTestDriver(t, mock)
			ctx := context.Background()

			err := d.SetTriggerSource(ctx, tc.src)

			if tc.expectErr && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tc.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !tc.expectErr {
				if len(mock.commandsSent) != 1 {
					t.Fatalf("expected 1 command sent, got %d", len(mock.commandsSent))
				}

				if mock.commandsSent[0] != tc.expectedCmd {
					t.Errorf("wanted command %q, got %q", tc.expectedCmd, mock.commandsSent[0])
				}
			}
		})
	}
}

func TestAbort(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := &mockInst{}
		d := newTestDriver(t, mock)
		ctx := context.Background()

		err := d.Abort(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(mock.commandsSent) != 1 {
			t.Fatalf("expected 1 command sent, got %d", len(mock.commandsSent))
		}

		if mock.commandsSent[0] != "ABOR" {
			t.Errorf("wanted command %q, got %q", "ABOR", mock.commandsSent[0])
		}
	})

	t.Run("error", func(t *testing.T) {
		mock := &mockInst{shouldError: true}
		d := newTestDriver(t, mock)
		ctx := context.Background()

		err := d.Abort(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestInitiateMeasurement(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mock := &mockInst{}
		d := newTestDriver(t, mock)
		ctx := context.Background()

		err := d.InitiateMeasurement(ctx)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(mock.commandsSent) != 1 {
			t.Fatalf("expected 1 command sent, got %d", len(mock.commandsSent))
		}

		if mock.commandsSent[0] != "init" {
			t.Errorf("wanted command %q, got %q", "init", mock.commandsSent[0])
		}
	})

	t.Run("error", func(t *testing.T) {
		mock := &mockInst{shouldError: true}
		d := newTestDriver(t, mock)
		ctx := context.Background()

		err := d.InitiateMeasurement(ctx)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
