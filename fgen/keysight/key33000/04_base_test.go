// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33000

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

func TestDriver_OutputCount(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if got := d.OutputCount(); got != 2 {
		t.Errorf("OutputCount() = %d, want 2", got)
	}
}

func TestDriver_Channel(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	ch, err := d.Channel(0)
	if err != nil {
		t.Fatalf("Channel(0) error: %v", err)
	}
	if got := ch.Name(); got != "Output1" {
		t.Errorf("Channel(0).Name() = %q, want %q", got, "Output1")
	}

	ch, err = d.Channel(1)
	if err != nil {
		t.Fatalf("Channel(1) error: %v", err)
	}
	if got := ch.Name(); got != "Output2" {
		t.Errorf("Channel(1).Name() = %q, want %q", got, "Output2")
	}

	_, err = d.Channel(2)
	if err == nil {
		t.Error("Channel(2) expected error, got nil")
	}
}

func TestChannel_srcPrefix(t *testing.T) {
	mock := &mockInst{}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	ch0, _ := d.Channel(0)
	if got := ch0.srcPrefix(); got != "SOUR1:" {
		t.Errorf("Channel(0).srcPrefix() = %q, want %q", got, "SOUR1:")
	}

	ch1, _ := d.Channel(1)
	if got := ch1.srcPrefix(); got != "SOUR2:" {
		t.Errorf("Channel(1).srcPrefix() = %q, want %q", got, "SOUR2:")
	}
}

func TestChannel_SetOutputEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		want    string
	}{
		{"enable ch1", true, "OUTP1 ON"},
		{"disable ch1", false, "OUTP1 OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, _ := New(mock)
			ch, _ := d.Channel(0)
			err := ch.SetOutputEnabled(tt.enabled)
			if err != nil {
				t.Fatalf("SetOutputEnabled() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.want)
			}
		})
	}
}

func TestChannel_SetOutputEnabled_Ch2(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(1)
	if err := ch.SetOutputEnabled(true); err != nil {
		t.Fatalf("SetOutputEnabled() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || mock.commandsSent[0] != "OUTP2 ON" {
		t.Errorf("sent %v, want [\"OUTP2 ON\"]", mock.commandsSent)
	}
}

func TestChannel_SetFrequency(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetFrequency(1000.0); err != nil {
		t.Fatalf("SetFrequency() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "SOUR1:FREQ") {
		t.Errorf("sent %v, want SOUR1:FREQ command", mock.commandsSent)
	}
}

func TestChannel_SetAmplitude(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(1)
	if err := ch.SetAmplitude(2.5); err != nil {
		t.Fatalf("SetAmplitude() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "SOUR2:VOLT") {
		t.Errorf("sent %v, want SOUR2:VOLT command", mock.commandsSent)
	}
}

func TestChannel_SetStandardWaveform(t *testing.T) {
	tests := []struct {
		name string
		wave fgen.StandardWaveform
		want string
	}{
		{"sine", fgen.Sine, "SOUR1:FUNC SIN"},
		{"square", fgen.Square, "SOUR1:FUNC SQU"},
		{"dc", fgen.DC, "SOUR1:FUNC DC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d, _ := New(mock)
			ch, _ := d.Channel(0)
			if err := ch.SetStandardWaveform(tt.wave); err != nil {
				t.Fatalf("SetStandardWaveform() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.want)
			}
		})
	}
}

func TestChannel_SetStandardWaveform_Triangle(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetStandardWaveform(fgen.Triangle); err != nil {
		t.Fatalf("SetStandardWaveform(Triangle) error: %v", err)
	}
	// Triangle sends two separate commands (no semicolons).
	if len(mock.commandsSent) != 2 {
		t.Fatalf("sent %d commands, want 2", len(mock.commandsSent))
	}
	if mock.commandsSent[0] != "SOUR1:FUNC RAMP" {
		t.Errorf("cmd[0] = %q, want %q", mock.commandsSent[0], "SOUR1:FUNC RAMP")
	}
	if mock.commandsSent[1] != "SOUR1:FUNC:RAMP:SYMM 50" {
		t.Errorf(
			"cmd[1] = %q, want %q",
			mock.commandsSent[1], "SOUR1:FUNC:RAMP:SYMM 50",
		)
	}
}

func TestChannel_SetOperationMode_Burst(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetOperationMode(fgen.BurstMode); err != nil {
		t.Fatalf("SetOperationMode(BurstMode) error: %v", err)
	}
	// Burst mode sends two separate commands (no semicolons).
	if len(mock.commandsSent) != 2 {
		t.Fatalf("sent %d commands, want 2", len(mock.commandsSent))
	}
	if mock.commandsSent[0] != "SOUR1:BURS:MODE TRIG" {
		t.Errorf(
			"cmd[0] = %q, want %q",
			mock.commandsSent[0], "SOUR1:BURS:MODE TRIG",
		)
	}
	if mock.commandsSent[1] != "SOUR1:BURS:STAT ON" {
		t.Errorf(
			"cmd[1] = %q, want %q",
			mock.commandsSent[1], "SOUR1:BURS:STAT ON",
		)
	}
}

func TestChannel_SetOperationMode_Continuous(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetOperationMode(fgen.ContinuousMode); err != nil {
		t.Fatalf("SetOperationMode(ContinuousMode) error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		mock.commandsSent[0] != "SOUR1:BURS:STAT OFF" {
		t.Errorf(
			"sent %v, want [\"SOUR1:BURS:STAT OFF\"]", mock.commandsSent,
		)
	}
}

func TestChannel_SetBurstCount(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetBurstCount(10); err != nil {
		t.Fatalf("SetBurstCount() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		mock.commandsSent[0] != "SOUR1:BURS:NCYC 10" {
		t.Errorf(
			"sent %v, want [\"SOUR1:BURS:NCYC 10\"]", mock.commandsSent,
		)
	}
}

func TestChannel_SetOutputImpedance(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(1)
	if err := ch.SetOutputImpedance(50.0); err != nil {
		t.Fatalf("SetOutputImpedance() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "OUTP2:LOAD") {
		t.Errorf("sent %v, want OUTP2:LOAD command", mock.commandsSent)
	}
}

func TestChannel_InternalTriggerRate(t *testing.T) {
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
			d, _ := New(mock)
			ch, _ := d.Channel(0)
			got, err := ch.InternalTriggerRate()
			if err != nil {
				t.Fatalf("InternalTriggerRate() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("InternalTriggerRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_SetInternalTriggerRate(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetInternalTriggerRate(1000.0); err != nil {
		t.Fatalf("SetInternalTriggerRate() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		mock.commandsSent[0] != "SOUR1:BURS:INT:PER 0.001" {
		t.Errorf(
			"sent %v, want [\"SOUR1:BURS:INT:PER 0.001\"]",
			mock.commandsSent,
		)
	}
}

func TestChannel_SetDCOffset(t *testing.T) {
	mock := &mockInst{}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetDCOffset(0.5); err != nil {
		t.Fatalf("SetDCOffset() error: %v", err)
	}
	if len(mock.commandsSent) != 1 ||
		!strings.HasPrefix(mock.commandsSent[0], "SOUR1:VOLT:OFFS") {
		t.Errorf("sent %v, want SOUR1:VOLT:OFFS command", mock.commandsSent)
	}
}

func TestChannel_StandardWaveform(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want fgen.StandardWaveform
	}{
		{"sine", "SIN", fgen.Sine},
		{"square", "SQU", fgen.Square},
		{"dc", "DC", fgen.DC},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, _ := New(mock)
			ch, _ := d.Channel(0)
			got, err := ch.StandardWaveform()
			if err != nil {
				t.Fatalf("StandardWaveform() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("StandardWaveform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_OutputMode(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want fgen.OutputMode
	}{
		{"function sine", "SIN", fgen.OutputModeFunction},
		{"function square", "SQU", fgen.OutputModeFunction},
		{"noise", "NOIS", fgen.OutputModeNoise},
		{"arbitrary", "ARB", fgen.OutputModeArbitrary},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d, _ := New(mock)
			got, err := d.OutputMode()
			if err != nil {
				t.Fatalf("OutputMode() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("OutputMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_CommandError(t *testing.T) {
	mock := &mockInst{shouldError: true}
	d, _ := New(mock)
	ch, _ := d.Channel(0)
	if err := ch.SetFrequency(1000); err == nil {
		t.Error("expected error, got nil")
	}
}
