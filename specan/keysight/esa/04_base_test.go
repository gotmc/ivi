// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package esa

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gotmc/ivi/specan"
)

type mockInst struct {
	commandsSent []string
	queryResp    string
	shouldError  bool
}

func (m *mockInst) ReadBinary(_ context.Context, _ []byte) (int, error) {
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

func TestNew(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if d == nil {
		t.Fatal("New() returned nil")
	}
}

func TestDriver_TraceCount(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if got := d.TraceCount(); got != 3 {
		t.Errorf("TraceCount() = %d, want 3", got)
	}
}

func TestDriver_SetFrequencyStart(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.SetFrequencyStart(1e6); err != nil {
		t.Fatalf("SetFrequencyStart() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "FREQ:STAR") {
		t.Errorf("sent %v, want FREQ:STAR command", mock.commandsSent)
	}
}

func TestDriver_SetFrequencyStop(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.SetFrequencyStop(1.5e9); err != nil {
		t.Fatalf("SetFrequencyStop() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "FREQ:STOP") {
		t.Errorf("sent %v, want FREQ:STOP command", mock.commandsSent)
	}
}

func TestDriver_ConfigureFrequencyCenterSpan(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.ConfigureFrequencyCenterSpan(750e6, 1.5e9); err != nil {
		t.Fatalf("ConfigureFrequencyCenterSpan() error: %v", err)
	}
	if len(mock.commandsSent) != 2 {
		t.Fatalf("sent %d commands, want 2", len(mock.commandsSent))
	}
	if !strings.HasPrefix(mock.commandsSent[0], "FREQ:CENT") {
		t.Errorf("cmd[0] = %q, want FREQ:CENT", mock.commandsSent[0])
	}
	if !strings.HasPrefix(mock.commandsSent[1], "FREQ:SPAN") {
		t.Errorf("cmd[1] = %q, want FREQ:SPAN", mock.commandsSent[1])
	}
}

func TestDriver_SetReferenceLevel(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.SetReferenceLevel(-20.0); err != nil {
		t.Fatalf("SetReferenceLevel() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "DISP:WIND:TRAC:Y:RLEV") {
		t.Errorf("sent %v, want DISP:WIND:TRAC:Y:RLEV command", mock.commandsSent)
	}
}

func TestDriver_SetResolutionBandwidth(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.SetResolutionBandwidth(10e3); err != nil {
		t.Fatalf("SetResolutionBandwidth() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "BAND ") {
		t.Errorf("sent %v, want BAND command", mock.commandsSent)
	}
}

func TestDriver_SetSweepModeContinuous(t *testing.T) {
	tests := []struct {
		name       string
		continuous bool
		want       string
	}{
		{"continuous", true, "INIT:CONT ON"},
		{"single", false, "INIT:CONT OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, _ := New(mock)
			if err := d.SetSweepModeContinuous(tt.continuous); err != nil {
				t.Fatalf("SetSweepModeContinuous() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.want)
			}
		})
	}
}

func TestDriver_SetAmplitudeUnits(t *testing.T) {
	tests := []struct {
		name  string
		units specan.AmplitudeUnits
		want  string
	}{
		{"dBm", specan.AmplitudeUnitsDBM, "UNIT:POW DBM"},
		{"dBuV", specan.AmplitudeUnitsDBUV, "UNIT:POW DBUV"},
		{"Volt", specan.AmplitudeUnitsVolt, "UNIT:POW V"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, _ := New(mock)
			if err := d.SetAmplitudeUnits(tt.units); err != nil {
				t.Fatalf("SetAmplitudeUnits() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.want)
			}
		})
	}
}

func TestDriver_AmplitudeUnits(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want specan.AmplitudeUnits
	}{
		{"dBm", "DBM", specan.AmplitudeUnitsDBM},
		{"dBmV", "DBMV", specan.AmplitudeUnitsDBMV},
		{"Volt", "V", specan.AmplitudeUnitsVolt},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, _ := New(mock)
			got, err := d.AmplitudeUnits()
			if err != nil {
				t.Fatalf("AmplitudeUnits() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("AmplitudeUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_VerticalScale(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want specan.VerticalScale
	}{
		{"log", "LOG", specan.VerticalScaleLogarithmic},
		{"lin", "LIN", specan.VerticalScaleLinear},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, _ := New(mock)
			got, err := d.VerticalScale()
			if err != nil {
				t.Fatalf("VerticalScale() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("VerticalScale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetAttenuationAuto(t *testing.T) {
	tests := []struct {
		name string
		auto bool
		want string
	}{
		{"auto on", true, "POW:ATT:AUTO ON"},
		{"auto off", false, "POW:ATT:AUTO OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, _ := New(mock)
			if err := d.SetAttenuationAuto(tt.auto); err != nil {
				t.Fatalf("SetAttenuationAuto() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.want)
			}
		})
	}
}

func TestDriver_Abort(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.Abort(); err != nil {
		t.Fatalf("Abort() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "ABOR" {
		t.Errorf("sent %v, want [\"ABOR\"]", mock.commandsSent)
	}
}

func TestDriver_Initiate(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	if err := d.Initiate(); err != nil {
		t.Fatalf("Initiate() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "INIT" {
		t.Errorf("sent %v, want [\"INIT\"]", mock.commandsSent)
	}
}

func TestDriver_CommandError(t *testing.T) {
	mock := &mockInst{shouldError: true}
	d, _ := New(mock)
	if err := d.SetFrequencyStart(1e6); err == nil {
		t.Error("expected error, got nil")
	}
}
