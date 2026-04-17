// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package esa

import (
	"strings"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/internal/ivitest"
	"github.com/gotmc/ivi/specan"
)

func TestNew(t *testing.T) {
	mock := &ivitest.Mock{}
	d, err := New(mock, ivi.WithoutIDQuery())
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if d == nil {
		t.Fatal("New() returned nil")
	}
}

func TestDriver_TraceCount(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if got := d.TraceCount(); got != 3 {
		t.Errorf("TraceCount() = %d, want 3", got)
	}
}

func TestDriver_SetFrequencyStart(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetFrequencyStart(1e6); err != nil {
		t.Fatalf("SetFrequencyStart() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "FREQ:STAR") {
		t.Errorf("sent %v, want FREQ:STAR command", mock.CommandsSent)
	}
}

func TestDriver_SetFrequencyStop(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetFrequencyStop(1.5e9); err != nil {
		t.Fatalf("SetFrequencyStop() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "FREQ:STOP") {
		t.Errorf("sent %v, want FREQ:STOP command", mock.CommandsSent)
	}
}

func TestDriver_ConfigureFrequencyCenterSpan(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.ConfigureFrequencyCenterSpan(750e6, 1.5e9); err != nil {
		t.Fatalf("ConfigureFrequencyCenterSpan() error: %v", err)
	}
	if len(mock.CommandsSent) != 2 {
		t.Fatalf("sent %d commands, want 2", len(mock.CommandsSent))
	}
	if !strings.HasPrefix(mock.CommandsSent[0], "FREQ:CENT") {
		t.Errorf("cmd[0] = %q, want FREQ:CENT", mock.CommandsSent[0])
	}
	if !strings.HasPrefix(mock.CommandsSent[1], "FREQ:SPAN") {
		t.Errorf("cmd[1] = %q, want FREQ:SPAN", mock.CommandsSent[1])
	}
}

func TestDriver_SetReferenceLevel(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetReferenceLevel(-20.0); err != nil {
		t.Fatalf("SetReferenceLevel() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "DISP:WIND:TRAC:Y:RLEV") {
		t.Errorf("sent %v, want DISP:WIND:TRAC:Y:RLEV command", mock.CommandsSent)
	}
}

func TestDriver_SetResolutionBandwidth(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetResolutionBandwidth(10e3); err != nil {
		t.Fatalf("SetResolutionBandwidth() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "BAND ") {
		t.Errorf("sent %v, want BAND command", mock.CommandsSent)
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
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetSweepModeContinuous(tt.continuous); err != nil {
				t.Fatalf("SetSweepModeContinuous() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
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
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetAmplitudeUnits(tt.units); err != nil {
				t.Fatalf("SetAmplitudeUnits() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
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
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
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
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
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
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetAttenuationAuto(tt.auto); err != nil {
				t.Fatalf("SetAttenuationAuto() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestDriver_Abort(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.Abort(); err != nil {
		t.Fatalf("Abort() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "ABOR" {
		t.Errorf("sent %v, want [\"ABOR\"]", mock.CommandsSent)
	}
}

func TestDriver_Initiate(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.Initiate(); err != nil {
		t.Fatalf("Initiate() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "INIT" {
		t.Errorf("sent %v, want [\"INIT\"]", mock.CommandsSent)
	}
}

func TestDriver_CommandError(t *testing.T) {
	mock := &ivitest.Mock{ShouldError: true}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetFrequencyStart(1e6); err == nil {
		t.Error("expected error, got nil")
	}
}
