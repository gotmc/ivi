// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
)

// mockInst captures commands and returns configurable query responses.
type mockInst struct {
	commandsSent []string
	queryResp    string
	shouldError  bool
}

func (m *mockInst) Read(p []byte) (int, error) {
	return 0, nil
}

func (m *mockInst) Write(p []byte) (int, error) {
	return len(p), nil
}

func (m *mockInst) WriteString(s string) (int, error) {
	return len(s), nil
}

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

func (m *mockInst) Query(_ context.Context, s string) (string, error) {
	if m.shouldError {
		return "", errors.New("mock query error")
	}
	return m.queryResp, nil
}

func newTestDriver(mock *mockInst) *Driver {
	channels := []Channel{
		{name: "CHAN1", num: 1, inst: mock},
		{name: "CHAN2", num: 2, inst: mock},
		{name: "CHAN3", num: 3, inst: mock},
		{name: "CHAN4", num: 4, inst: mock},
	}
	inherent := ivi.NewInherent(mock, ivi.InherentBase{ReturnToLocal: true})
	return &Driver{
		inst:     mock,
		Channels: channels,
		Inherent: inherent,
	}
}

func TestDriver_ChannelCount(t *testing.T) {
	d := newTestDriver(&mockInst{})
	if got := d.ChannelCount(); got != 4 {
		t.Errorf("ChannelCount() = %d, want 4", got)
	}
}

func TestChannel_Name(t *testing.T) {
	d := newTestDriver(&mockInst{})
	want := []string{"CH1", "CH2", "CH3", "CH4"}
	for i, w := range want {
		if got := d.Channels[i].Name(); got != w {
			t.Errorf("Channel[%d].Name() = %q, want %q", i, got, w)
		}
	}
}

func TestDriver_AcquisitionType(t *testing.T) {
	tests := []struct {
		name    string
		resp    string
		want    scope.AcquisitionType
		wantErr bool
	}{
		{"normal", "NORM", scope.NormalAcquisition, false},
		{"average", "AVER", scope.AverageAcquisition, false},
		{"high res", "HRES", scope.HighResolutionAcquisition, false},
		{"peak detect", "PEAK", scope.PeakDetectAcquisition, false},
		{"unknown", "UNKNOWN", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d := newTestDriver(mock)

			got, err := d.AcquisitionType(context.Background())
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

func TestDriver_SetAcquisitionType(t *testing.T) {
	tests := []struct {
		name    string
		acType  scope.AcquisitionType
		wantCmd string
		wantErr bool
	}{
		{"normal", scope.NormalAcquisition, ":ACQ:TYPE NORM", false},
		{"average", scope.AverageAcquisition, ":ACQ:TYPE AVER", false},
		{"high res", scope.HighResolutionAcquisition, ":ACQ:TYPE HRES", false},
		{"peak detect", scope.PeakDetectAcquisition, ":ACQ:TYPE PEAK", false},
		{"unsupported", scope.AcquisitionType(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d := newTestDriver(mock)

			err := d.SetAcquisitionType(context.Background(), tt.acType)
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

func TestDriver_TriggerType(t *testing.T) {
	tests := []struct {
		name    string
		resp    string
		want    scope.TriggerType
		wantErr bool
	}{
		{"edge", "EDGE", scope.EdgeTrigger, false},
		{"glitch", "GLIT", scope.GlitchTrigger, false},
		{"width", "PATT", scope.WidthTrigger, false},
		{"tv", "TV", scope.TVTrigger, false},
		{"runt", "RUNT", scope.RuntTrigger, false},
		{"unknown", "UNKNOWN", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{queryResp: tt.resp}
			d := newTestDriver(mock)

			got, err := d.TriggerType(context.Background())
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

func TestDriver_SetTriggerType(t *testing.T) {
	tests := []struct {
		name    string
		trig    scope.TriggerType
		wantCmd string
		wantErr bool
	}{
		{"edge", scope.EdgeTrigger, ":TRIG:MODE EDGE", false},
		{"glitch", scope.GlitchTrigger, ":TRIG:MODE GLIT", false},
		{"immediate", scope.ImmediateTrigger, ":TRIG:FORC", false},
		{"unsupported", scope.TriggerType(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d := newTestDriver(mock)

			err := d.SetTriggerType(context.Background(), tt.trig)
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

func TestDriver_SetTriggerLevel(t *testing.T) {
	mock := &mockInst{}
	d := newTestDriver(mock)
	err := d.SetTriggerLevel(context.Background(), 1.5)
	if err != nil {
		t.Errorf("SetTriggerLevel() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], ":TRIG:EDGE:LEV") {
		t.Errorf("sent %v, want :TRIG:EDGE:LEV command", mock.commandsSent)
	}
}

func TestDriver_SetTriggerHoldoff(t *testing.T) {
	tests := []struct {
		name    string
		holdoff time.Duration
		wantErr bool
	}{
		{"valid", 100 * time.Microsecond, false},
		{"minimum", 40 * time.Nanosecond, false},
		{"maximum", 10 * time.Second, false},
		{"too small", 1 * time.Nanosecond, true},
		{"too large", 11 * time.Second, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			d := newTestDriver(mock)
			err := d.SetTriggerHoldoff(context.Background(), tt.holdoff)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestChannel_SetChannelEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		wantCmd string
	}{
		{"enable", true, ":CHAN1:DISP 1"},
		{"disable", false, ":CHAN1:DISP 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "CHAN1", num: 1}
			err := ch.SetChannelEnabled(context.Background(), tt.enabled)
			if err != nil {
				t.Errorf("SetChannelEnabled() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetInputImpedance(t *testing.T) {
	tests := []struct {
		name      string
		impedance float64
		wantCmd   string
		wantErr   bool
	}{
		{"50 ohm", 50.0, ":CHAN1:IMP FIFT", false},
		{"1M ohm", 1e6, ":CHAN1:IMP ONEM", false},
		{"unsupported", 75.0, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "CHAN1", num: 1}
			err := ch.SetInputImpedance(context.Background(), tt.impedance)
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

func TestChannel_SetVerticalRange(t *testing.T) {
	tests := []struct {
		name    string
		rng     float64
		wantErr bool
	}{
		{"valid", 10.0, false},
		{"min", 0.008, false},
		{"max", 40.0, false},
		{"too small", 0.001, true},
		{"too large", 50.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "CHAN1", num: 1}
			err := ch.SetVerticalRange(context.Background(), tt.rng)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if !errors.Is(err, ivi.ErrValueNotSupported) {
					t.Errorf("got %v, want ErrValueNotSupported", err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestChannel_SetProbeAttenuation(t *testing.T) {
	tests := []struct {
		name    string
		atten   float64
		wantErr bool
	}{
		{"1x", 1.0, false},
		{"10x", 10.0, false},
		{"too small", 0.0001, true},
		{"too large", 20000.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "CHAN1", num: 1}
			err := ch.SetProbeAttenuation(context.Background(), tt.atten)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestChannel_SetProbeAttenuationAuto(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "CHAN1", num: 1}

	// false should succeed
	err := ch.SetProbeAttenuationAuto(context.Background(), false)
	if err != nil {
		t.Errorf("SetProbeAttenuationAuto(false) error: %v", err)
	}

	// true should fail (not supported)
	err = ch.SetProbeAttenuationAuto(context.Background(), true)
	if !errors.Is(err, ivi.ErrValueNotSupported) {
		t.Errorf("SetProbeAttenuationAuto(true) = %v, want ErrValueNotSupported", err)
	}
}

func TestChannel_SetVerticalCoupling(t *testing.T) {
	tests := []struct {
		name     string
		coupling scope.VerticalCoupling
		wantCmd  string
		wantErr  bool
	}{
		{"AC", scope.ACVerticalCoupling, ":CHAN1:COUP AC", false},
		{"DC", scope.DCVerticalCoupling, ":CHAN1:COUP DC", false},
		{"unsupported", scope.VerticalCoupling(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{inst: mock, name: "CHAN1", num: 1}
			err := ch.SetVerticalCoupling(context.Background(), tt.coupling)
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

func TestDecodeTimebase(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    timebase
		wantErr bool
	}{
		{
			"valid",
			":TIM:MODE MAIN;REF CENT;MAIN:RANG 1.000000E-03;POS 0.000000E+00",
			timebase{mode: "MAIN", reference: "CENT", rng: 1e-3, position: 0.0},
			false,
		},
		{
			"invalid parts count",
			":TIM:MODE MAIN;REF CENT",
			timebase{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeTimebase(tt.input)
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
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestDurationFromSeconds(t *testing.T) {
	tests := []struct {
		name    string
		seconds float64
		want    time.Duration
	}{
		{"one second", 1.0, time.Second},
		{"one millisecond", 0.001, time.Millisecond},
		{"negative", -0.5, -500 * time.Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := durationFromSeconds(tt.seconds)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_NotImplemented(t *testing.T) {
	mock := &mockInst{}
	d := newTestDriver(mock)
	ctx := context.Background()

	_, err := d.AcquisitionStatus(ctx)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("AcquisitionStatus() = %v, want ErrNotImplemented", err)
	}

	_, err = d.TriggerSlope(ctx)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("TriggerSlope() = %v, want ErrNotImplemented", err)
	}

	_, err = d.TriggerSource(ctx)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("TriggerSource() = %v, want ErrNotImplemented", err)
	}

	err = d.AbortMeasurement(ctx)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("AbortMeasurement() = %v, want ErrNotImplemented", err)
	}

	err = d.InitiateMeasurement(ctx)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("InitiateMeasurement() = %v, want ErrNotImplemented", err)
	}
}

func TestDriver_SetAcquisitionStartTime_NotSupported(t *testing.T) {
	mock := &mockInst{}
	d := newTestDriver(mock)
	err := d.SetAcquisitionStartTime(context.Background(), 100*time.Microsecond)
	if !errors.Is(err, ivi.ErrFunctionNotSupported) {
		t.Errorf("SetAcquisitionStartTime() = %v, want ErrFunctionNotSupported", err)
	}
}

func TestChannel_SetVerticalOffset(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{inst: mock, name: "CHAN1", num: 1}
	err := ch.SetVerticalOffset(context.Background(), 2.5)
	if err != nil {
		t.Errorf("SetVerticalOffset() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], ":CHAN1:OFFS") {
		t.Errorf("sent %v, want :CHAN1:OFFS command", mock.commandsSent)
	}
}

func TestDriver_ConfigureTrigger(t *testing.T) {
	mock := &mockInst{}
	d := newTestDriver(mock)
	err := d.ConfigureTrigger(context.Background(), scope.EdgeTrigger, 100*time.Microsecond)
	if err != nil {
		t.Errorf("ConfigureTrigger() error: %v", err)
	}
	if len(mock.commandsSent) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(mock.commandsSent))
	}
	if mock.commandsSent[0] != ":TRIG:MODE EDGE" {
		t.Errorf("first command = %q, want \":TRIG:MODE EDGE\"", mock.commandsSent[0])
	}
	if !strings.HasPrefix(mock.commandsSent[1], ":TRIG:HOLD") {
		t.Errorf("second command = %q, want :TRIG:HOLD command", mock.commandsSent[1])
	}
}

func TestDriver_SetAcquisitionTimePerRecord(t *testing.T) {
	mock := &mockInst{}
	d := newTestDriver(mock)
	err := d.SetAcquisitionTimePerRecord(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Errorf("SetAcquisitionTimePerRecord() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.HasPrefix(mock.commandsSent[0], ":TIM:RANG") {
		t.Errorf("sent %v, want :TIM:RANG command", mock.commandsSent)
	}
}
