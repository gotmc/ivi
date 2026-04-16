// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package sr630

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/tempmon"
	"github.com/gotmc/query"
)

// --- Channel Count ---

func (d *Driver) ChannelCount() int {
	return numChannels
}

// --- Thermocouple Type ---

// The SR630 uses single-letter mnemonics for thermocouple types.
var tcTypeToSCPI = map[tempmon.ThermocoupleType]string{
	tempmon.TypeB: "B",
	tempmon.TypeE: "E",
	tempmon.TypeJ: "J",
	tempmon.TypeK: "K",
	tempmon.TypeR: "R",
	tempmon.TypeS: "S",
	tempmon.TypeT: "T",
}

var scpiToTCType = map[string]tempmon.ThermocoupleType{
	"B": tempmon.TypeB,
	"E": tempmon.TypeE,
	"J": tempmon.TypeJ,
	"K": tempmon.TypeK,
	"R": tempmon.TypeR,
	"S": tempmon.TypeS,
	"T": tempmon.TypeT,
}

func (d *Driver) ThermocoupleType(channel int) (tempmon.ThermocoupleType, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "TTYP? %d", channel)
	if err != nil {
		return 0, fmt.Errorf("ThermocoupleType: %w", err)
	}

	tt, err := ivi.ReverseLookup(scpiToTCType, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("ThermocoupleType: %w", err)
	}

	return tt, nil
}

func (d *Driver) SetThermocoupleType(
	channel int, tcType tempmon.ThermocoupleType,
) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(tcTypeToSCPI, tcType)
	if err != nil {
		return fmt.Errorf("SetThermocoupleType: %w", err)
	}

	return d.inst.Command(ctx, "TTYP %d, %s", channel, cmd)
}

// --- Temperature Units ---

// The SR630 uses multi-letter mnemonics for units.
var unitsToSCPI = map[tempmon.TemperatureUnits]string{
	tempmon.Celsius:    "CENT",
	tempmon.Fahrenheit: "FHRN",
	tempmon.Kelvin:     "ABS",
}

var scpiToUnits = map[string]tempmon.TemperatureUnits{
	"CENT": tempmon.Celsius,
	"FHRN": tempmon.Fahrenheit,
	"ABS":  tempmon.Kelvin,
}

func (d *Driver) TemperatureUnits(channel int) (tempmon.TemperatureUnits, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "UNIT? %d", channel)
	if err != nil {
		return 0, fmt.Errorf("TemperatureUnits: %w", err)
	}

	units, err := ivi.ReverseLookup(scpiToUnits, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("TemperatureUnits: %w", err)
	}

	return units, nil
}

func (d *Driver) SetTemperatureUnits(
	channel int, units tempmon.TemperatureUnits,
) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(unitsToSCPI, units)
	if err != nil {
		return fmt.Errorf("SetTemperatureUnits: %w", err)
	}

	return d.inst.Command(ctx, "UNIT %d, %s", channel, cmd)
}

// --- Measurement ---

func (d *Driver) MeasureTemperature(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "MEAS? %d", channel)
}

// --- Alarm Limits ---

func (d *Driver) AlarmEnabled(channel int) (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "ALRM? %d", channel)
	if err != nil {
		return false, fmt.Errorf("AlarmEnabled: %w", err)
	}

	return strings.TrimSpace(s) == "YES", nil
}

func (d *Driver) SetAlarmEnabled(channel int, enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "ALRM %d, YES", channel)
	}

	return d.inst.Command(ctx, "ALRM %d, NO", channel)
}

func (d *Driver) AlarmHighLimit(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "TMAX? %d", channel)
}

func (d *Driver) SetAlarmHighLimit(channel int, limit float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TMAX %d, %f", channel, limit)
}

func (d *Driver) AlarmLowLimit(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "TMIN? %d", channel)
}

func (d *Driver) SetAlarmLowLimit(channel int, limit float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TMIN %d, %f", channel, limit)
}

// --- Scanner ---

func (d *Driver) ScanEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "SCAN?")
}

func (d *Driver) SetScanEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "SCAN 1")
	}

	return d.inst.Command(ctx, "SCAN 0")
}

func (d *Driver) ChannelScanEnabled(channel int) (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "SCNE? %d", channel)
	if err != nil {
		return false, fmt.Errorf("ChannelScanEnabled: %w", err)
	}

	return strings.TrimSpace(s) == "YES", nil
}

func (d *Driver) SetChannelScanEnabled(channel int, enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "SCNE %d, YES", channel)
	}

	return d.inst.Command(ctx, "SCNE %d, NO", channel)
}

func (d *Driver) DwellTime() (time.Duration, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	sec, err := query.Float64(ctx, d.inst, "DWEL?")
	if err != nil {
		return 0, fmt.Errorf("DwellTime: %w", err)
	}

	return time.Duration(sec * float64(time.Second)), nil
}

func (d *Driver) SetDwellTime(dwell time.Duration) error {
	ctx, cancel := d.newContext()
	defer cancel()

	sec := int(dwell.Seconds())
	if sec < 10 || sec > 9999 {
		return fmt.Errorf(
			"SetDwellTime %v: %w (must be 10-9999 seconds)",
			dwell, ivi.ErrValueNotSupported,
		)
	}

	return d.inst.Command(ctx, "DWEL %d", sec)
}

// --- Relative Temperature ---

func (d *Driver) NominalTemperature(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "TNOM? %d", channel)
}

func (d *Driver) SetNominalTemperature(channel int, nominal float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TNOM %d, %f", channel, nominal)
}

func (d *Driver) DeltaTemperature(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "TDLT? %d", channel)
}
