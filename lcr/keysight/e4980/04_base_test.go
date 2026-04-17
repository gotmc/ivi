// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e4980

import (
	"strings"
	"testing"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/internal/ivitest"
	"github.com/gotmc/ivi/lcr"
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

func TestDriver_SetMeasurementFunction(t *testing.T) {
	tests := []struct {
		name string
		fcn  lcr.MeasurementFunction
		want string
	}{
		{"CpD", lcr.CpD, "FUNC:IMP:TYPE CPD"},
		{"CsRs", lcr.CsRs, "FUNC:IMP:TYPE CSRS"},
		{"LsD", lcr.LsD, "FUNC:IMP:TYPE LSD"},
		{"RX", lcr.RX, "FUNC:IMP:TYPE RX"},
		{"ZThetaDeg", lcr.ZThetaDeg, "FUNC:IMP:TYPE ZTD"},
		{"GB", lcr.GB, "FUNC:IMP:TYPE GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetMeasurementFunction(tt.fcn); err != nil {
				t.Fatalf("SetMeasurementFunction() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestDriver_MeasurementFunction(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want lcr.MeasurementFunction
	}{
		{"CpD", "CPD", lcr.CpD},
		{"LsRs", "LSRS", lcr.LsRs},
		{"ZThetaRad", "ZTR", lcr.ZThetaRad},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
			got, err := d.MeasurementFunction()
			if err != nil {
				t.Fatalf("MeasurementFunction() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("MeasurementFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetFrequency(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetFrequency(1000.0); err != nil {
		t.Fatalf("SetFrequency() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "FREQ:CW") {
		t.Errorf("sent %v, want FREQ:CW command", mock.CommandsSent)
	}
}

func TestDriver_SetTestVoltageLevel(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetTestVoltageLevel(1.0); err != nil {
		t.Fatalf("SetTestVoltageLevel() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "VOLT:LEV") {
		t.Errorf("sent %v, want VOLT:LEV command", mock.CommandsSent)
	}
}

func TestDriver_SetImpedanceAutoRange(t *testing.T) {
	tests := []struct {
		name string
		auto bool
		want string
	}{
		{"auto on", true, "FUNC:IMP:RANG:AUTO ON"},
		{"auto off", false, "FUNC:IMP:RANG:AUTO OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetImpedanceAutoRange(tt.auto); err != nil {
				t.Fatalf("SetImpedanceAutoRange() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestDriver_SetTriggerSource(t *testing.T) {
	tests := []struct {
		name string
		src  lcr.TriggerSource
		want string
	}{
		{"internal", lcr.TriggerSourceInternal, "TRIG:SOUR INT"},
		{"external", lcr.TriggerSourceExternal, "TRIG:SOUR EXT"},
		{"bus", lcr.TriggerSourceBus, "TRIG:SOUR BUS"},
		{"hold", lcr.TriggerSourceHold, "TRIG:SOUR HOLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetTriggerSource(tt.src); err != nil {
				t.Fatalf("SetTriggerSource() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestDriver_TriggerSource(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want lcr.TriggerSource
	}{
		{"internal", "INT", lcr.TriggerSourceInternal},
		{"bus", "BUS", lcr.TriggerSourceBus},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
			got, err := d.TriggerSource()
			if err != nil {
				t.Fatalf("TriggerSource() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("TriggerSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_SetTriggerDelay(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetTriggerDelay(100 * time.Millisecond); err != nil {
		t.Fatalf("SetTriggerDelay() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		!strings.HasPrefix(mock.CommandsSent[0], "TRIG:DEL") {
		t.Errorf("sent %v, want TRIG:DEL command", mock.CommandsSent)
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
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "INIT:IMM" {
		t.Errorf("sent %v, want [\"INIT:IMM\"]", mock.CommandsSent)
	}
}

func TestDriver_Trigger(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.Trigger(); err != nil {
		t.Fatalf("Trigger() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "TRIG:IMM" {
		t.Errorf("sent %v, want [\"TRIG:IMM\"]", mock.CommandsSent)
	}
}

func TestDriver_FetchMeasurement(t *testing.T) {
	mock := &ivitest.Mock{QueryResp: "+1.23456E-09,+3.45678E-03,+0"}
	d, _ := New(mock, ivi.WithoutIDQuery())
	pri, sec, status, err := d.FetchMeasurement()
	if err != nil {
		t.Fatalf("FetchMeasurement() error: %v", err)
	}
	if pri != 1.23456e-09 {
		t.Errorf("primary = %g, want 1.23456e-09", pri)
	}
	if sec != 3.45678e-03 {
		t.Errorf("secondary = %g, want 3.45678e-03", sec)
	}
	if status != lcr.MeasurementStatusNormal {
		t.Errorf("status = %v, want Normal", status)
	}
}

func TestDriver_FetchMeasurement_Overload(t *testing.T) {
	mock := &ivitest.Mock{QueryResp: "+9.99999E+37,+9.99999E+37,+1"}
	d, _ := New(mock, ivi.WithoutIDQuery())
	_, _, status, err := d.FetchMeasurement()
	if err != nil {
		t.Fatalf("FetchMeasurement() error: %v", err)
	}
	if status != lcr.MeasurementStatusOverload {
		t.Errorf("status = %v, want Overload", status)
	}
}

func TestDriver_SetDCBiasEnabled(t *testing.T) {
	tests := []struct {
		name    string
		enabled bool
		want    string
	}{
		{"enable", true, "BIAS:STAT ON"},
		{"disable", false, "BIAS:STAT OFF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			d, _ := New(mock, ivi.WithoutIDQuery())
			if err := d.SetDCBiasEnabled(tt.enabled); err != nil {
				t.Fatalf("SetDCBiasEnabled() error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestDriver_SetOpenCorrectionEnabled(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetOpenCorrectionEnabled(true); err != nil {
		t.Fatalf("SetOpenCorrectionEnabled() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		mock.CommandsSent[0] != "CORR:OPEN:STAT ON" {
		t.Errorf(
			"sent %v, want [\"CORR:OPEN:STAT ON\"]", mock.CommandsSent,
		)
	}
}

func TestDriver_ExecuteOpenCorrection(t *testing.T) {
	mock := &ivitest.Mock{}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.ExecuteOpenCorrection(); err != nil {
		t.Fatalf("ExecuteOpenCorrection() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 ||
		mock.CommandsSent[0] != "CORR:OPEN" {
		t.Errorf("sent %v, want [\"CORR:OPEN\"]", mock.CommandsSent)
	}
}

func TestDriver_MeasurementSpeed(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want lcr.MeasurementSpeed
	}{
		{"short", "SHOR,1", lcr.MeasurementSpeedShort},
		{"medium", "MED,10", lcr.MeasurementSpeedMedium},
		{"long", "LONG,4", lcr.MeasurementSpeedLong},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
			got, err := d.MeasurementSpeed()
			if err != nil {
				t.Fatalf("MeasurementSpeed() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("MeasurementSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_AveragingCount(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want int
	}{
		{"1", "MED,1", 1},
		{"10", "LONG,10", 10},
		{"256", "SHOR,256", 256},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: tt.resp}
			d, _ := New(mock, ivi.WithoutIDQuery())
			got, err := d.AveragingCount()
			if err != nil {
				t.Fatalf("AveragingCount() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("AveragingCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestDriver_CommandError(t *testing.T) {
	mock := &ivitest.Mock{ShouldError: true}
	d, _ := New(mock, ivi.WithoutIDQuery())
	if err := d.SetFrequency(1000); err == nil {
		t.Error("expected error, got nil")
	}
}

func TestParseMeasurementResult(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantPri   float64
		wantSec   float64
		wantStat  lcr.MeasurementStatus
		wantError bool
	}{
		{
			"normal",
			"+1.00000E-09,+5.00000E-03,+0",
			1e-09, 5e-03, lcr.MeasurementStatusNormal, false,
		},
		{
			"overload",
			"+9.99999E+37,+9.99999E+37,+1",
			9.99999e+37, 9.99999e+37, lcr.MeasurementStatusOverload, false,
		},
		{
			"with bin",
			"+1.00000E-09,+5.00000E-03,+0,+3",
			1e-09, 5e-03, lcr.MeasurementStatusNormal, false,
		},
		{
			"too few fields",
			"+1.00000E-09,+5.00000E-03",
			0, 0, lcr.MeasurementStatusNoData, true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri, sec, status, err := parseMeasurementResult(tt.input)
			if tt.wantError {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if pri != tt.wantPri {
				t.Errorf("primary = %g, want %g", pri, tt.wantPri)
			}
			if sec != tt.wantSec {
				t.Errorf("secondary = %g, want %g", sec, tt.wantSec)
			}
			if status != tt.wantStat {
				t.Errorf("status = %v, want %v", status, tt.wantStat)
			}
		})
	}
}
