// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gotmc/ivi/fgen"
)

// mockInst captures commands and returns configurable query responses.
type mockInst struct {
	commandsSent []string
	queryResp    string
	shouldError  bool
}

func (m *mockInst) ReadContext(_ context.Context, p []byte) (int, error) {
	return 0, nil
}

func (m *mockInst) WriteContext(_ context.Context, p []byte) (int, error) {
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

func TestDriver_OutputCount(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if got := d.OutputCount(); got != 1 {
		t.Errorf("OutputCount() = %d, want 1", got)
	}
}

func TestChannel_Name(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	ch, err := d.Channel(0)
	if err != nil {
		t.Fatalf("Channel(0) error: %v", err)
	}
	if got := ch.Name(); got != "Output" {
		t.Errorf("Name() = %q, want %q", got, "Output")
	}
}

func TestDriver_OutputMode(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     fgen.OutputMode
		wantErr  bool
	}{
		{"sine", "SIN", fgen.OutputModeFunction, false},
		{"square", "SQU", fgen.OutputModeFunction, false},
		{"noise", "NOIS", fgen.OutputModeNoise, false},
		{"user", "USER", fgen.OutputModeArbitrary, false},
		{"unknown", "PULSE", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			d, err := New(mock)
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			got, err := d.OutputMode(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Error("OutputMode() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("OutputMode() unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("OutputMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetOutputMode(t *testing.T) {
	tests := []struct {
		name    string
		mode    fgen.OutputMode
		wantCmd string
		wantErr bool
	}{
		{"function", fgen.OutputModeFunction, "FUNC SIN", false},
		{"arbitrary", fgen.OutputModeArbitrary, "FUNC USER", false},
		{"noise", fgen.OutputModeNoise, "FUNC NOIS", false},
		{"unsupported", fgen.OutputMode(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
			d, err := New(mock)
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			err = d.SetOutputMode(context.Background(), tt.mode)
			if tt.wantErr {
				if err == nil {
					t.Error("SetOutputMode() expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("SetOutputMode() unexpected error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("SetOutputMode() sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetOutputEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		wantCmd string
	}{
		{"enable", true, "OUTP ON"},
		{"disable", false, "OUTP OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "Output"}
			err := ch.SetOutputEnabled(context.Background(), tt.enabled)
			if err != nil {
				t.Errorf("SetOutputEnabled() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("SetOutputEnabled() sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetStandardWaveform(t *testing.T) {
	tests := []struct {
		name    string
		wave    fgen.StandardWaveform
		wantCmd string
		wantErr bool
	}{
		{"sine", fgen.Sine, "FUNC SIN", false},
		{"square", fgen.Square, "FUNC SQU", false},
		{"dc", fgen.DC, "FUNC DC", false},
		{"triangle", fgen.Triangle, "FUNC RAMP; RAMP:SYMM 50", false},
		{"ramp_up", fgen.RampUp, "FUNC RAMP; RAMP:SYMM 100", false},
		{"ramp_down", fgen.RampDown, "FUNC RAMP; RAMP:SYMM 0", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "Output"}
			err := ch.SetStandardWaveform(context.Background(), tt.wave)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("SetStandardWaveform() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetAmplitude(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.SetAmplitude(context.Background(), 2.5)
	if err != nil {
		t.Errorf("SetAmplitude() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], "VOLT ") {
		t.Errorf("SetAmplitude() sent %v, want VOLT command", mock.commandsSent)
	}
}

func TestChannel_SetFrequency(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.SetFrequency(context.Background(), 1000.0)
	if err != nil {
		t.Errorf("SetFrequency() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], "FREQ ") {
		t.Errorf("SetFrequency() sent %v, want FREQ command", mock.commandsSent)
	}
}

func TestChannel_SetDCOffset(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.SetDCOffset(context.Background(), 0.5)
	if err != nil {
		t.Errorf("SetDCOffset() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], "VOLT:OFFS ") {
		t.Errorf("SetDCOffset() sent %v, want VOLT:OFFS command", mock.commandsSent)
	}
}

func TestChannel_OperationMode(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     fgen.OperationMode
		wantErr  bool
	}{
		{"continuous", "0", fgen.ContinuousMode, false},
		{"burst", "1", fgen.BurstMode, false},
		{"unknown", "9", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.OperationMode(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("OperationMode() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("OperationMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_SetOperationMode(t *testing.T) {
	tests := []struct {
		name    string
		mode    fgen.OperationMode
		wantCmd string
		wantErr bool
	}{
		{"burst", fgen.BurstMode, "BURS:MODE TRIG;STAT ON", false},
		{"continuous", fgen.ContinuousMode, "BURS:STAT OFF", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "Output"}
			err := ch.SetOperationMode(context.Background(), tt.mode)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("SetOperationMode() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetOutputImpedance(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.SetOutputImpedance(context.Background(), 50.0)
	if err != nil {
		t.Errorf("SetOutputImpedance() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], "OUTP:LOAD ") {
		t.Errorf("sent %v, want OUTP:LOAD command", mock.commandsSent)
	}
}

func TestChannel_OutputEnabled(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     bool
		wantErr  bool
	}{
		{"enabled", "1", true, false},
		{"disabled", "0", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.OutputEnabled(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("OutputEnabled() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("OutputEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_OutputImpedance(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
	}{
		{"50 ohm", "50.0", 50.0},
		{"high-z", "9.9E+37", 9.9e+37},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.OutputImpedance(context.Background())
			if err != nil {
				t.Errorf("OutputImpedance() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("OutputImpedance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_Amplitude(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
	}{
		{"2.5 vpp", "2.500000", 2.5},
		{"0.1 vpp", "0.100000", 0.1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.Amplitude(context.Background())
			if err != nil {
				t.Errorf("Amplitude() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Amplitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_DCOffset(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
	}{
		{"positive", "0.500000", 0.5},
		{"zero", "0.000000", 0.0},
		{"negative", "-1.250000", -1.25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.DCOffset(context.Background())
			if err != nil {
				t.Errorf("DCOffset() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("DCOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_Frequency(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
	}{
		{"1 kHz", "1000.0", 1000.0},
		{"1 MHz", "1000000.0", 1000000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.Frequency(context.Background())
			if err != nil {
				t.Errorf("Frequency() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_StandardWaveform(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     fgen.StandardWaveform
		wantErr  bool
	}{
		{"sine", "SIN", fgen.Sine, false},
		{"square", "SQU", fgen.Square, false},
		{"dc", "DC", fgen.DC, false},
		{"unknown", "PULS", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.StandardWaveform(context.Background())
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("StandardWaveform() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("StandardWaveform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_DisableOutput(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.DisableOutput(context.Background())
	if err != nil {
		t.Errorf("DisableOutput() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "OUTP OFF" {
		t.Errorf("DisableOutput() sent %v, want [%q]", mock.commandsSent, "OUTP OFF")
	}
}

func TestChannel_EnableOutput(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.EnableOutput(context.Background())
	if err != nil {
		t.Errorf("EnableOutput() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "OUTP ON" {
		t.Errorf("EnableOutput() sent %v, want [%q]", mock.commandsSent, "OUTP ON")
	}
}

func TestChannel_BurstCount(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     int
	}{
		{"5 cycles", "5", 5},
		{"1 cycle", "1", 1},
		{"100 cycles", "100", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			ch := Channel{inst: mock, name: "Output"}
			got, err := ch.BurstCount(context.Background())
			if err != nil {
				t.Errorf("BurstCount() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("BurstCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestChannel_SetBurstCount(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "Output"}
	err := ch.SetBurstCount(context.Background(), 5)
	if err != nil {
		t.Errorf("SetBurstCount() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "BURS:NCYC 5" {
		t.Errorf("SetBurstCount() sent %v, want [%q]", mock.commandsSent, "BURS:NCYC 5")
	}
}

func TestDriver_InternalTriggerRate(t *testing.T) {
	tests := []struct {
		name     string
		response string
		want     float64
	}{
		{"1000 Hz", "0.001", 1000.0},
		{"100 Hz", "0.01", 100.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.response}
			d, err := New(mock)
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}
			got, err := d.InternalTriggerRate(context.Background())
			if err != nil {
				t.Errorf("InternalTriggerRate() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("InternalTriggerRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetInternalTriggerRate(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	err = d.SetInternalTriggerRate(context.Background(), 1000.0)
	if err != nil {
		t.Errorf("SetInternalTriggerRate() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "BURS:INT:PER 0.001" {
		t.Errorf("SetInternalTriggerRate() sent %v, want [%q]", mock.commandsSent, "BURS:INT:PER 0.001")
	}
}

func TestDriver_ReferenceClockSource(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	got, err := d.ReferenceClockSource(context.Background())
	if err != nil {
		t.Errorf("ReferenceClockSource() error: %v", err)
	}
	if got != fgen.RefClockInternal {
		t.Errorf("ReferenceClockSource() = %v, want %v", got, fgen.RefClockInternal)
	}
}

func TestDriver_InitiateGeneration(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	err = d.InitiateGeneration(context.Background())
	if err != nil {
		t.Errorf("InitiateGeneration() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "OUTP ON" {
		t.Errorf("InitiateGeneration() sent %v, want [%q]", mock.commandsSent, "OUTP ON")
	}
}

func TestDriver_AbortGeneration(t *testing.T) {
	mock := &mockInst{queryResp: "KEYSIGHT,33220A,0,1.0"}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	err = d.AbortGeneration(context.Background())
	if err != nil {
		t.Errorf("AbortGeneration() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "OUTP OFF" {
		t.Errorf("AbortGeneration() sent %v, want [%q]", mock.commandsSent, "OUTP OFF")
	}
}
