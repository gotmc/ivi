// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36000

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/ivi/internal/ivitest"
)

func newTestDriver(mock *ivitest.Mock) *Driver {
	supply, err := supplyForModel("E3631A")
	if err != nil {
		panic(err)
	}

	channels := make([]Channel, len(supply.channels))
	for i, name := range supply.channels {
		channels[i] = Channel{
			inst:       mock,
			name:       name,
			family:     supply.family,
			protection: supply.protection,
		}
	}

	inherent := ivi.NewInherent(mock, ivi.InherentBase{ReturnToLocal: true}, 0)

	return &Driver{
		inst:     mock,
		supply:   supply,
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
			ch := channelForModel(t, mock, "E3631A", 0)
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

// TestChannel_SetCommandsPerFamily checks that the setters are spelled for
// the connected model's command set. The E36100B series has no INSTrument
// subsystem, so an "INST Output; " prefix would be rejected by the
// instrument.
func TestChannel_SetCommandsPerFamily(t *testing.T) {
	tests := []struct {
		name  string
		model string
		send  func(ch *Channel) error
		want  string
	}{
		{
			"E3631A voltage", "E3631A",
			func(ch *Channel) error { return ch.SetVoltageLevel(5.0) },
			"INST P6V; VOLT 5.0000",
		},
		{
			"E36102B voltage", "E36102B",
			func(ch *Channel) error { return ch.SetVoltageLevel(3.3) },
			"VOLT 3.3000",
		},
		{
			"E3631A current", "E3631A",
			func(ch *Channel) error { return ch.SetCurrentLimit(0.5) },
			"INST P6V; CURR 0.5000",
		},
		{
			"E36102B current", "E36102B",
			func(ch *Channel) error { return ch.SetCurrentLimit(0.5) },
			"CURR 0.5000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			ch := channelForModel(t, mock, tt.model, 0)
			if err := tt.send(ch); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != tt.want {
				t.Errorf("sent %v, want [%q]", mock.CommandsSent, tt.want)
			}
		})
	}
}

// TestChannel_QueriesPerFamily is the query-side counterpart to
// TestChannel_SetCommandsPerFamily.
func TestChannel_QueriesPerFamily(t *testing.T) {
	tests := []struct {
		name  string
		model string
		read  func(ch *Channel) error
		want  string
	}{
		{
			"E3631A voltage", "E3631A",
			func(ch *Channel) error { _, err := ch.VoltageLevel(); return err },
			"INST P6V; VOLT?",
		},
		{
			"E36102B voltage", "E36102B",
			func(ch *Channel) error { _, err := ch.VoltageLevel(); return err },
			"VOLT?",
		},
		{
			"E3631A measure voltage", "E3631A",
			func(ch *Channel) error { _, err := ch.MeasureVoltage(); return err },
			"MEAS:VOLT? P6V",
		},
		{
			"E36102B measure voltage", "E36102B",
			func(ch *Channel) error { _, err := ch.MeasureVoltage(); return err },
			"MEAS:VOLT?",
		},
		{
			"E36102B measure current", "E36102B",
			func(ch *Channel) error { _, err := ch.MeasureCurrent(); return err },
			"MEAS:CURR?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &queryRecorder{Mock: ivitest.Mock{QueryResp: "1.0"}}
			ch := channelForModel(t, rec, tt.model, 0)
			if err := tt.read(ch); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(rec.QueriesSent) != 1 || rec.QueriesSent[0] != tt.want {
				t.Errorf("queried %v, want [%q]", rec.QueriesSent, tt.want)
			}
		})
	}
}

func TestChannel_CurrentLimitBehavior(t *testing.T) {
	tests := []struct {
		name  string
		model string
		resp  string
		want  dcpwr.CurrentLimitBehavior
	}{
		{"E3631A regulates only", "E3631A", "", dcpwr.CurrentRegulate},
		{"E36102B ocp off", "E36102B", "0", dcpwr.CurrentRegulate},
		{"E36102B ocp on", "E36102B", "1", dcpwr.CurrentTrip},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: tt.resp}
			ch := channelForModel(t, mock, tt.model, 0)
			got, err := ch.CurrentLimitBehavior()
			if err != nil {
				t.Fatalf("CurrentLimitBehavior() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("CurrentLimitBehavior() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_SetCurrentLimitBehavior(t *testing.T) {
	tests := []struct {
		name     string
		model    string
		behavior dcpwr.CurrentLimitBehavior
		wantErr  bool
		wantCmds []string
	}{
		{"E3631A regulate", "E3631A", dcpwr.CurrentRegulate, false, nil},
		{"E3631A trip", "E3631A", dcpwr.CurrentTrip, true, nil},
		{
			"E36102B regulate", "E36102B", dcpwr.CurrentRegulate, false,
			[]string{"CURR:PROT:STAT OFF"},
		},
		{
			"E36102B trip", "E36102B", dcpwr.CurrentTrip, false,
			[]string{"CURR:PROT:STAT ON"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			ch := channelForModel(t, mock, tt.model, 0)
			err := ch.SetCurrentLimitBehavior(tt.behavior)
			if tt.wantErr {
				if !errors.Is(err, ivi.ErrValueNotSupported) {
					t.Fatalf("expected ErrValueNotSupported, got %v", err)
				}

				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !equalCommands(mock.CommandsSent, tt.wantCmds) {
				t.Errorf("sent %v, want %v", mock.CommandsSent, tt.wantCmds)
			}
		})
	}
}

func TestChannel_ConfigureCurrentLimit(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := channelForModel(t, mock, "E36102B", 0)
	if err := ch.ConfigureCurrentLimit(dcpwr.CurrentTrip, 0.25); err != nil {
		t.Fatalf("ConfigureCurrentLimit() error: %v", err)
	}
	want := []string{"CURR:PROT:STAT ON", "CURR 0.2500"}
	if !equalCommands(mock.CommandsSent, want) {
		t.Errorf("sent %v, want %v", mock.CommandsSent, want)
	}
}

// TestChannel_ConfigureCurrentLimit_Unsupported checks that a model without
// over-current protection still reports CurrentTrip as unavailable rather
// than silently programming only the limit.
func TestChannel_ConfigureCurrentLimit_Unsupported(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := channelForModel(t, mock, "E3631A", 0)
	err := ch.ConfigureCurrentLimit(dcpwr.CurrentTrip, 0.25)
	if !errors.Is(err, ivi.ErrValueNotSupported) {
		t.Errorf("ConfigureCurrentLimit() = %v, want ErrValueNotSupported", err)
	}
	if len(mock.CommandsSent) != 0 {
		t.Errorf("sent %v, want no commands", mock.CommandsSent)
	}
}

func TestChannel_OVP_Unsupported(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := channelForModel(t, mock, "E3631A", 0)

	enabled, err := ch.OVPEnabled()
	if err != nil {
		t.Errorf("OVPEnabled() error: %v", err)
	}
	if enabled {
		t.Error("OVPEnabled() = true, want false")
	}

	if err := ch.DisableOVP(); err != nil {
		t.Errorf("DisableOVP() = %v, want nil", err)
	}

	unsupported := map[string]error{
		"SetOVPEnabled": ch.SetOVPEnabled(true),
		"EnableOVP":     ch.EnableOVP(),
		"SetOVPLimit":   ch.SetOVPLimit(10.0),
		"ConfigureOVP":  ch.ConfigureOVP(true, 10.0),
	}
	for name, err := range unsupported {
		if !errors.Is(err, dcpwr.ErrOVPUnsupported) {
			t.Errorf("%s() = %v, want ErrOVPUnsupported", name, err)
		}
	}

	if _, err := ch.OVPLimit(); !errors.Is(err, dcpwr.ErrOVPUnsupported) {
		t.Errorf("OVPLimit() = %v, want ErrOVPUnsupported", err)
	}

	if len(mock.CommandsSent) != 0 {
		t.Errorf("sent %v, want no commands", mock.CommandsSent)
	}
}

func TestChannel_OVP_Supported(t *testing.T) {
	tests := []struct {
		name string
		call func(ch *Channel) error
		want []string
	}{
		{
			"enable",
			func(ch *Channel) error { return ch.EnableOVP() },
			[]string{"VOLT:PROT:STAT ON"},
		},
		{
			"disable",
			func(ch *Channel) error { return ch.DisableOVP() },
			[]string{"VOLT:PROT:STAT OFF"},
		},
		{
			"set limit",
			func(ch *Channel) error { return ch.SetOVPLimit(6.5) },
			[]string{"VOLT:PROT 6.5000"},
		},
		{
			"configure enabled",
			func(ch *Channel) error { return ch.ConfigureOVP(true, 6.5) },
			[]string{"VOLT:PROT 6.5000", "VOLT:PROT:STAT ON"},
		},
		{
			// With OVP being turned off the limit is not applied, per
			// Section 4.3.4 of the IviDCPwr class specification.
			"configure disabled",
			func(ch *Channel) error { return ch.ConfigureOVP(false, 6.5) },
			[]string{"VOLT:PROT:STAT OFF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			ch := channelForModel(t, mock, "E36102B", 0)
			if err := tt.call(ch); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !equalCommands(mock.CommandsSent, tt.want) {
				t.Errorf("sent %v, want %v", mock.CommandsSent, tt.want)
			}
		})
	}
}

func TestChannel_OVPEnabled_Supported(t *testing.T) {
	rec := &queryRecorder{Mock: ivitest.Mock{QueryResp: "1"}}
	ch := channelForModel(t, rec, "E36102B", 0)
	enabled, err := ch.OVPEnabled()
	if err != nil {
		t.Fatalf("OVPEnabled() error: %v", err)
	}
	if !enabled {
		t.Error("OVPEnabled() = false, want true")
	}
	want := "VOLT:PROT:STAT?"
	if len(rec.QueriesSent) != 1 || rec.QueriesSent[0] != want {
		t.Errorf("queried %v, want [%q]", rec.QueriesSent, want)
	}
}

func TestChannel_QueryOutputState(t *testing.T) {
	tests := []struct {
		name      string
		state     dcpwr.OutputState
		resp      string
		wantQuery string
		want      bool
	}{
		{"cv set", dcpwr.ConstantVoltage, "256", "STAT:OPER:COND?", true},
		{"cv clear", dcpwr.ConstantVoltage, "1024", "STAT:OPER:COND?", false},
		{"cc set", dcpwr.ConstantCurrent, "1024", "STAT:OPER:COND?", true},
		{"ov set", dcpwr.OverVoltage, "1", "STAT:QUES:COND?", true},
		{"oc set", dcpwr.OverCurrent, "2", "STAT:QUES:COND?", true},
		{"oc clear", dcpwr.OverCurrent, "1", "STAT:QUES:COND?", false},
		{"unregulated", dcpwr.Unregulated, "1024", "STAT:QUES:COND?", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &queryRecorder{Mock: ivitest.Mock{QueryResp: tt.resp}}
			ch := channelForModel(t, rec, "E36102B", 0)
			got, err := ch.QueryOutputState(tt.state)
			if err != nil {
				t.Fatalf("QueryOutputState() error: %v", err)
			}
			if got != tt.want {
				t.Errorf("QueryOutputState(%v) = %t, want %t", tt.state, got, tt.want)
			}
			if len(rec.QueriesSent) != 1 || rec.QueriesSent[0] != tt.wantQuery {
				t.Errorf("queried %v, want [%q]", rec.QueriesSent, tt.wantQuery)
			}
		})
	}
}

func TestChannel_ResetOutputProtection(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := channelForModel(t, mock, "E36102B", 0)
	if err := ch.ResetOutputProtection(); err != nil {
		t.Fatalf("ResetOutputProtection() error: %v", err)
	}
	want := []string{"OUTP:PROT:CLE"}
	if !equalCommands(mock.CommandsSent, want) {
		t.Errorf("sent %v, want %v", mock.CommandsSent, want)
	}
}

func TestChannel_NotImplemented_WrapsCorrectError(t *testing.T) {
	mock := &ivitest.Mock{}
	ch := channelForModel(t, mock, "E3631A", 0)

	err := ch.ConfigureOutputRange(dcpwr.CurrentRange, 1.0)
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
	ch := channelForModel(t, mock, "E3631A", 0)
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
	ch := channelForModel(t, mock, "E3631A", 0)
	err := ch.EnableOutput()
	if err != nil {
		t.Errorf("EnableOutput() error: %v", err)
	}
	if len(mock.CommandsSent) != 1 || mock.CommandsSent[0] != "OUTP ON" {
		t.Errorf("sent %v, want [\"OUTP ON\"]", mock.CommandsSent)
	}
}

// equalCommands reports whether the commands the mock captured match want,
// treating a nil want as "no commands sent".
func equalCommands(got, want []string) bool {
	return slices.Equal(got, want)
}

// scriptedMock answers each query from a table keyed by the exact query
// string, so a single test can serve a measurement and an output state query
// with different values. Unlisted queries fail the call rather than returning
// a stale value from a previous step.
type scriptedMock struct {
	ivitest.Mock
	responses map[string]string
	Queries   []string
}

func (m *scriptedMock) Query(_ context.Context, cmd string) (string, error) {
	m.Queries = append(m.Queries, cmd)

	resp, ok := m.responses[cmd]
	if !ok {
		return "", fmt.Errorf("scriptedMock: unexpected query %q", cmd)
	}

	return resp, nil
}

// TestE3631A_ConfigureEnableAndReadBack walks the P6V output through the
// sequence a bench user would run by hand: set the output to 4.1 V with a
// 1.2 A limit, turn the output on, read the voltage back, and confirm the
// output is enabled.
//
// The driver has no APPLy method, so the first step is the IviDCPwr pair
// SetVoltageLevel and SetCurrentLimit, which the E3631A accepts as the
// equivalent of "APPL P6V,4.1,1.2". The read back uses the explicit
// MEAS:VOLT? spelling rather than the "MEAS?" short form; VOLTage is the
// default MEASure function on this supply, so the two are equivalent.
func TestE3631A_ConfigureEnableAndReadBack(t *testing.T) {
	mock := &scriptedMock{
		responses: map[string]string{
			"MEAS:VOLT? P6V": "+4.10000000E+00",
			"OUTP?":          "1",
		},
	}

	ch := channelForModel(t, mock, "E3631A", 0)
	if got := ch.Name(); got != "P6V" {
		t.Fatalf("Channel(0).Name() = %q, want %q", got, "P6V")
	}

	// Configure 4.1 V at a 1.2 A limit on the P6V output.
	if err := ch.SetVoltageLevel(4.1); err != nil {
		t.Fatalf("SetVoltageLevel() error: %v", err)
	}

	if err := ch.SetCurrentLimit(1.2); err != nil {
		t.Fatalf("SetCurrentLimit() error: %v", err)
	}

	// Turn the output on. OUTPut:STATe is global on the E3631A, so this
	// command carries no INSTrument prefix.
	if err := ch.EnableOutput(); err != nil {
		t.Fatalf("EnableOutput() error: %v", err)
	}

	wantCommands := []string{
		"INST P6V; VOLT 4.1000",
		"INST P6V; CURR 1.2000",
		"OUTP ON",
	}
	if !slices.Equal(mock.CommandsSent, wantCommands) {
		t.Errorf("sent %q, want %q", mock.CommandsSent, wantCommands)
	}

	// Read the voltage back from the P6V output.
	volts, err := ch.MeasureVoltage()
	if err != nil {
		t.Fatalf("MeasureVoltage() error: %v", err)
	}

	if volts != 4.1 {
		t.Errorf("MeasureVoltage() = %v, want %v", volts, 4.1)
	}

	// Confirm the output is enabled.
	enabled, err := ch.OutputEnabled()
	if err != nil {
		t.Fatalf("OutputEnabled() error: %v", err)
	}

	if !enabled {
		t.Error("OutputEnabled() = false, want true")
	}

	wantQueries := []string{"MEAS:VOLT? P6V", "OUTP?"}
	if !slices.Equal(mock.Queries, wantQueries) {
		t.Errorf("queried %q, want %q", mock.Queries, wantQueries)
	}
}
