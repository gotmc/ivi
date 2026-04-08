// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dp800

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
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
	m.commandsSent = append(m.commandsSent, s)
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
		{name: "CH1", idx: 1, inst: mock, maxVoltage: 32.0, maxCurrent: 3.2},
		{name: "CH2", idx: 2, inst: mock, maxVoltage: 32.0, maxCurrent: 3.2},
		{name: "CH3", idx: 3, inst: mock, maxVoltage: 5.3, maxCurrent: 3.2},
	}
	inherent := ivi.NewInherent(mock, ivi.InherentBase{ReturnToLocal: true})
	return &Driver{
		inst:     mock,
		Channels: channels,
		Inherent: inherent,
	}
}

func TestDriver_OutputChannelCount(t *testing.T) {
	d := newTestDriver(&mockInst{})
	if got := d.OutputChannelCount(); got != 3 {
		t.Errorf("OutputChannelCount() = %d, want 3", got)
	}
}

func TestChannel_Name(t *testing.T) {
	d := newTestDriver(&mockInst{})
	names := []string{"CH1", "CH2", "CH3"}
	for i, want := range names {
		if got := d.Channels[i].Name(); got != want {
			t.Errorf("Channel[%d].Name() = %q, want %q", i, got, want)
		}
	}
}

func TestChannel_SetCurrentLimit(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{name: "CH1", idx: 1, inst: mock}
	err := ch.SetCurrentLimit(context.Background(), 0.5)
	if err != nil {
		t.Errorf("SetCurrentLimit() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.Contains(mock.commandsSent[0], ":SOUR1:CURR") {
		t.Errorf("sent %v, want :SOUR1:CURR command", mock.commandsSent)
	}
}

func TestChannel_SetVoltageLevel(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{name: "CH1", idx: 1, inst: mock}
	err := ch.SetVoltageLevel(context.Background(), 5.0)
	if err != nil {
		t.Errorf("SetVoltageLevel() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.Contains(mock.commandsSent[0], ":SOUR1:VOLT") {
		t.Errorf("sent %v, want :SOUR1:VOLT command", mock.commandsSent)
	}
}

func TestChannel_SetOutputEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		wantCmd string
	}{
		{"enable", true, ":OUTP CH1,ON"},
		{"disable", false, ":OUTP CH1,OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{name: "CH1", idx: 1, inst: mock}
			err := ch.SetOutputEnabled(context.Background(), tt.enabled)
			if err != nil {
				t.Errorf("SetOutputEnabled() error: %v", err)
			}
			// SetOutputEnabled uses ivi.Set which calls WriteString
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetCurrentLimitBehavior(t *testing.T) {
	tests := []struct {
		name     string
		behavior dcpwr.CurrentLimitBehavior
		wantCmd  string
		wantErr  bool
	}{
		{"regulate", dcpwr.CurrentRegulate, ":OUTP:OCP CH1,OFF", false},
		{"trip", dcpwr.CurrentTrip, ":OUTP:OCP CH1,ON", false},
		{"unsupported", dcpwr.CurrentLimitBehavior(99), "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{name: "CH1", idx: 1, inst: mock}
			err := ch.SetCurrentLimitBehavior(context.Background(), tt.behavior)
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

func TestChannel_SetOVPEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		wantCmd string
	}{
		{"enable", true, ":OUTP:OVP CH1,ON"},
		{"disable", false, ":OUTP:OVP CH1,OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{name: "CH1", idx: 1, inst: mock}
			err := ch.SetOVPEnabled(context.Background(), tt.enabled)
			if err != nil {
				t.Errorf("SetOVPEnabled() error: %v", err)
			}
			if len(mock.commandsSent) != 1 || mock.commandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.commandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_SetOVPLimit(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{name: "CH1", idx: 1, inst: mock}
	err := ch.SetOVPLimit(context.Background(), 10.0)
	if err != nil {
		t.Errorf("SetOVPLimit() error: %v", err)
	}
	if len(mock.commandsSent) != 1 || !strings.Contains(mock.commandsSent[0], ":OUTP:OVP:VAL CH1") {
		t.Errorf("sent %v, want :OUTP:OVP:VAL command", mock.commandsSent)
	}
}

func TestChannel_QueryCurrentLimitMax(t *testing.T) {
	d := newTestDriver(&mockInst{})
	got, err := d.Channels[0].QueryCurrentLimitMax(context.Background(), 5.0)
	if err != nil {
		t.Errorf("QueryCurrentLimitMax() error: %v", err)
	}
	if got != 3.2 {
		t.Errorf("QueryCurrentLimitMax() = %f, want 3.2", got)
	}
}

func TestChannel_QueryVoltageLevelMax(t *testing.T) {
	d := newTestDriver(&mockInst{})
	got, err := d.Channels[0].QueryVoltageLevelMax(context.Background(), 1.0)
	if err != nil {
		t.Errorf("QueryVoltageLevelMax() error: %v", err)
	}
	if got != 32.0 {
		t.Errorf("QueryVoltageLevelMax() = %f, want 32.0", got)
	}
}

func TestChannel_ConfigureOutputRange_NoOp(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{name: "CH1", idx: 1, inst: mock}
	err := ch.ConfigureOutputRange(context.Background(), dcpwr.VoltageRange, 10.0)
	if err != nil {
		t.Errorf("ConfigureOutputRange() error: %v", err)
	}
	if len(mock.commandsSent) != 0 {
		t.Errorf("expected no commands, got %v", mock.commandsSent)
	}
}

func TestChannel_ResetOutputProtection(t *testing.T) {
	mock := &mockInst{}
	ch := Channel{name: "CH1", idx: 1, inst: mock}
	err := ch.ResetOutputProtection(context.Background())
	if err != nil {
		t.Errorf("ResetOutputProtection() error: %v", err)
	}
	if len(mock.commandsSent) != 2 {
		t.Fatalf("expected 2 commands, got %d", len(mock.commandsSent))
	}
	if !strings.Contains(mock.commandsSent[0], ":OUTP:OVP:CLEAR CH1") {
		t.Errorf("first command = %q, want OVP clear", mock.commandsSent[0])
	}
	if !strings.Contains(mock.commandsSent[1], ":OUTP:OCP:CLEAR CH1") {
		t.Errorf("second command = %q, want OCP clear", mock.commandsSent[1])
	}
}

func TestChannel_ConfigureOVP(t *testing.T) {
	tests := []struct {
		name         string
		enabled      bool
		limit        float64
		wantCmdCount int
	}{
		{"enabled with limit", true, 10.0, 2},
		{"disabled", false, 10.0, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockInst{}
			ch := Channel{name: "CH1", idx: 1, inst: mock}
			err := ch.ConfigureOVP(context.Background(), tt.enabled, tt.limit)
			if err != nil {
				t.Errorf("ConfigureOVP() error: %v", err)
			}
			if len(mock.commandsSent) != tt.wantCmdCount {
				t.Errorf("sent %d commands, want %d", len(mock.commandsSent), tt.wantCmdCount)
			}
		})
	}
}
