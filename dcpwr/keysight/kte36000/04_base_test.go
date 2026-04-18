// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kte36000

import (
	"errors"
	"strings"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/ivi/internal/ivitest"
)

func newTestDriver(mock *ivitest.Mock) *Driver {
	channels := []Channel{
		{inst: mock, name: "P6V"},
		{inst: mock, name: "P25V"},
		{inst: mock, name: "N25V"},
	}
	inherent := ivi.NewInherent(mock, ivi.InherentBase{ReturnToLocal: true}, 0)
	return &Driver{
		inst:     mock,
		channels: channels,
		Inherent: inherent,
	}
}

func TestDriver_OutputChannelCount(t *testing.T) {
	d := newTestDriver(&ivitest.Mock{})
	if got := d.OutputChannelCount(); got != 3 {
		t.Errorf("OutputChannelCount() = %d, want 3", got)
	}
}

func TestChannel_Name(t *testing.T) {
	d := newTestDriver(&ivitest.Mock{})
	names := []string{"P6V", "P25V", "N25V"}
	for i, want := range names {
		ch, err := d.Channel(i)
		if err != nil {
			t.Fatalf("Channel(%d) error: %v", i, err)
		}
		if got := ch.Name(); got != want {
			t.Errorf("Channel(%d).Name() = %q, want %q", i, got, want)
		}
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
			mock := &ivitest.Mock{}
			ch := Channel{inst: mock, name: "P6V"}
			err := ch.SetOutputEnabled(tt.enabled)
			if err != nil {
				t.Errorf("SetOutputEnabled() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.wantCmd {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.wantCmd)
			}
		})
	}
}

func TestChannel_CurrentLimitBehavior(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}
	got, err := ch.CurrentLimitBehavior()
	if err != nil {
		t.Errorf("CurrentLimitBehavior() error: %v", err)
	}
	if got != dcpwr.CurrentRegulate {
		t.Errorf("CurrentLimitBehavior() = %v, want CurrentRegulate", got)
	}
}

func TestChannel_SetCurrentLimitBehavior(t *testing.T) {
	tests := []struct {
		name     string
		behavior dcpwr.CurrentLimitBehavior
		wantErr  bool
	}{
		{"regulate", dcpwr.CurrentRegulate, false},
		{"trip", dcpwr.CurrentTrip, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			ch := Channel{inst: mock, name: "P6V"}
			err := ch.SetCurrentLimitBehavior(tt.behavior)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				if !errors.Is(err, ivi.ErrValueNotSupported) {
					t.Errorf("expected ErrValueNotSupported, got %v", err)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestChannel_SetVoltageLevel(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}
	err := ch.SetVoltageLevel(5.0)
	if err != nil {
		t.Errorf("SetVoltageLevel() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || !strings.Contains(mock.CommandsSent[0], "volt") {
		t.Errorf("sent %v, want volt command", mock.CommandsSent)
	}
}

func TestChannel_SetCurrentLimit(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}
	err := ch.SetCurrentLimit(0.5)
	if err != nil {
		t.Errorf("SetCurrentLimit() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || !strings.Contains(mock.CommandsSent[0], "CURR") {
		t.Errorf("sent %v, want CURR command", mock.CommandsSent)
	}
}

func TestChannel_OVP_NotSupported(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}

	// OVPEnabled should return false with no error
	enabled, err := ch.OVPEnabled()
	if err != nil {
		t.Errorf("OVPEnabled() error: %v", err)
	}
	if enabled {
		t.Error("OVPEnabled() = true, want false")
	}

	// SetOVPEnabled should return ErrOVPUnsupported
	err = ch.SetOVPEnabled(true)
	if !errors.Is(err, dcpwr.ErrOVPUnsupported) {
		t.Errorf("SetOVPEnabled() = %v, want ErrOVPUnsupported", err)
	}

	// EnableOVP should return ErrOVPUnsupported
	err = ch.EnableOVP()
	if !errors.Is(err, dcpwr.ErrOVPUnsupported) {
		t.Errorf("EnableOVP() = %v, want ErrOVPUnsupported", err)
	}

	// OVPLimit should return ErrOVPUnsupported
	_, err = ch.OVPLimit()
	if !errors.Is(err, dcpwr.ErrOVPUnsupported) {
		t.Errorf("OVPLimit() = %v, want ErrOVPUnsupported", err)
	}

	// SetOVPLimit should return ErrOVPUnsupported
	err = ch.SetOVPLimit(10.0)
	if !errors.Is(err, dcpwr.ErrOVPUnsupported) {
		t.Errorf("SetOVPLimit() = %v, want ErrOVPUnsupported", err)
	}
}

func TestChannel_NotImplemented_WrapsCorrectError(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}

	// Verify unimplemented methods return errors.Is-compatible ivi.ErrNotImplemented
	err := ch.ConfigureCurrentLimit(dcpwr.CurrentRegulate, 1.0)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("ConfigureCurrentLimit() = %v, want ErrNotImplemented", err)
	}

	err = ch.ConfigureOutputRange(dcpwr.CurrentRange, 1.0)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("ConfigureOutputRange() = %v, want ErrNotImplemented", err)
	}

	err = ch.ResetOutputProtection()
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("ResetOutputProtection() = %v, want ErrNotImplemented", err)
	}

	_, err = ch.QueryCurrentLimitMax(5.0)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("QueryCurrentLimitMax() = %v, want ErrNotImplemented", err)
	}

	_, err = ch.QueryOutputState(dcpwr.OverCurrent)
	if !errors.Is(err, ivi.ErrNotImplemented) {
		t.Errorf("QueryOutputState() = %v, want ErrNotImplemented", err)
	}
}

func TestChannel_DisableOutput(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}
	err := ch.DisableOutput()
	if err != nil {
		t.Errorf("DisableOutput() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "OUTP OFF" {
		t.Errorf("sent %v, want [\"OUTP OFF\"]", mock.CommandsSent)
	}
}

func TestChannel_EnableOutput(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := Channel{inst: mock, name: "P6V"}
	err := ch.EnableOutput()
	if err != nil {
		t.Errorf("EnableOutput() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "OUTP ON" {
		t.Errorf("sent %v, want [\"OUTP ON\"]", mock.CommandsSent)
	}
}
