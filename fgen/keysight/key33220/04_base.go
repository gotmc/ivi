// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviFgenBase Attribute Output
// Count described in Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputCount() int {
	return len(d.Channels)
}

// OutputMode returns the determines how the function generator produces
// waveforms. This attribute determines which extension group’s functions and
// attributes are used to configure the waveform the function generator
// produces.
//
// OutputMode is the getter for the read-only IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputMode() (fgen.OutputMode, error) {
	var outputMode fgen.OutputMode

	funcType, err := query.String(d.inst, "FUNC?")
	if err != nil {
		return outputMode, fmt.Errorf("error determining the output function type: %w", err)
	}

	switch funcType {
	case "SIN", "SQU", "RAMP":
		return fgen.OutputModeFunction, nil
	case "NOIS":
		return fgen.OutputModeNoise, nil
	case "USER":
		return fgen.OutputModeArbitrary, nil
	}

	return 0, fmt.Errorf("unknown output mode type")
}

// SetOutputMode sets how the function generator produces waveforms. This
// attribute determines which extension group’s functions and attributes are
// used to configure the waveform the function generator produces.
//
// OutputMode is the setter for the read-only IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) SetOutputMode(outputMode fgen.OutputMode) error {
	switch outputMode {
	case fgen.OutputModeFunction:
		return d.inst.Command("FUNC SIN")
	case fgen.OutputModeArbitrary:
		return d.inst.Command("FUNC USER")
	case fgen.OutputModeSequence:
		return fmt.Errorf("function generator does not support output mode sequency")
	case fgen.OutputModeNoise:
		return d.inst.Command("FUNC NOIS")
	}

	return fmt.Errorf("error setting output mode")
}

// InitiateGeneration initiates signal generation by enabling all outputs.
// Instead of calling this function, the user can simply enable outputs.
//
// InitiateGeneration implements the IviFgenBase function described in Section
// 4.3.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) InitiateGeneration() error {
	for _, channel := range d.Channels {
		if err := channel.EnableOutput(); err != nil {
			return err
		}
	}

	return nil
}

// AbortGeneration aborts a previously initiated signal generation by disabling
// all outputs.
//
// AbortGeneration implements the IviFgenBase function described in Section 4.3.1
// of IVI-4.3: IviFgen Class Specification.
func (d *Driver) AbortGeneration() error {
	for _, channel := range d.Channels {
		if err := channel.DisableOutput(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Driver) ReferenceClockSource() (fgen.ClockSource, error) {
	return fgen.RefClockInternal, nil
}

func (d *Driver) SetReferenceClockSource(_ fgen.ClockSource) error {
	return nil
}

func (ch *Channel) Name() string {
	return "output"
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel.
//
// OperationMode implements the getter for the read-write IviFgenBase Attribute
// Operation Mode described in Section 4.2.2 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OperationMode() (fgen.OperationMode, error) {
	var mode fgen.OperationMode

	s, err := query.String(ch.inst, "BURS:STAT?")
	if err != nil {
		return mode, fmt.Errorf("error getting operation mode: %s", err)
	}

	switch strings.TrimSpace(s) {
	case "0":
		return fgen.ContinuousMode, nil
	case "1":
		return fgen.BurstMode, nil
	default:
		return mode, fmt.Errorf("error determining operation mode; received: %s", s)
	}
}

// SetOperationMode specifies whether the function generator should produce a
// continuous or burst output on the channel. SetOperationMode implements the
// setter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetOperationMode(mode fgen.OperationMode) error {
	switch mode {
	case fgen.BurstMode:
		return ch.inst.Command("BURS:MODE TRIG;STAT ON")
	case fgen.ContinuousMode:
		return ch.inst.Command("BURS:STAT OFF")
	}

	return errors.New("bad fgen operation mode")
}

// OutputEnabled determines if the output channel is enabled or disabled.
// OutputEnabled is the getter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return query.Bool(ch.inst, "OUTP?")
}

// SetOutputEnabled sets the output channel to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(b bool) error {
	if b {
		return ch.inst.Command("OUTP ON")
	}

	return ch.inst.Command("OUTP OFF")
}

// DisableOutput is a convenience function for setting the Output Enabled
// attribute to false.
func (ch *Channel) DisableOutput() error {
	return ch.SetOutputEnabled(false)
}

// EnableOutput is a convenience function for setting the Output Enabled
// attribute to true.
func (ch *Channel) EnableOutput() error {
	return ch.SetOutputEnabled(true)
}

// OutputImpedance return the output channel's impedance in ohms.
// OutputImpedance is the getter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputImpedance() (float64, error) {
	return query.Float64(ch.inst, "OUTP:LOAD?")
}

// SetOutputImpedance sets the output channel's impedance in ohms.
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(impedance float64) error {
	return ch.inst.Command("OUTP:LOAD %f", impedance)
}

// AbortGeneration Aborts a previously initiated signal generation. If the
// function generator is in the Output Generation State, this function moves
// the function generator to the Configuration State. If the function generator
// is already in the Configuration State, the function does nothing and returns
// Success. AbortGeneration implements the IviFgenBase function described in
// Section 4.3 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) AbortGeneration() error {
	return ch.DisableOutput()
}
