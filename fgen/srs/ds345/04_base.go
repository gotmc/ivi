// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"errors"
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

const (
	outputImpedance = 50.0
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

	funcType, err := query.Int(d.inst, "FUNC?")
	if err != nil {
		return outputMode, fmt.Errorf("error determining the output function type: %w", err)
	}

	switch funcType {
	case sine, square, triangle, ramp:
		return fgen.OutputModeFunction, nil
	case noise:
		return fgen.OutputModeNoise, nil
	case arbitrary:
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
		return d.inst.Command("FUNC0")
	case fgen.OutputModeArbitrary:
		return d.inst.Command("FUNC5")
	case fgen.OutputModeSequence:
		return fmt.Errorf("function generator does not support output mode sequency")
	case fgen.OutputModeNoise:
		return d.inst.Command("FUNC4")
	}

	return fmt.Errorf("error setting output mode")
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel. OperationMode implements the
// getter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OperationMode() (fgen.OperationMode, error) {
	var mode fgen.OperationMode

	isModulationEnabled, err := query.Bool(ch.inst, "MENA?")
	if err != nil {
		return mode, fmt.Errorf("error determining if modulation is enabled: %s", err)
	}

	if !isModulationEnabled {
		return fgen.ContinuousMode, nil
	}

	modType, err := query.Int(ch.inst, "MTYP?")
	if err != nil {
		return mode, fmt.Errorf("error determining modulation type: %s", err)
	}

	switch modType {
	case arbitrary:
		return fgen.BurstMode, nil
	default:
		return mode, fmt.Errorf("error determining operation mode, mtyp = %v", modType)
	}
}

// SetOperationMode specifies whether the function generator should produce a
// continuous or burst output on the channel.
//
// SetOperationMode implements the setter for the read-write IviFgenBase
// Attribute Operation Mode described in Section 4.2.2 of IVI-4.3: IviFgen
// Class Specification.
func (ch *Channel) SetOperationMode(mode fgen.OperationMode) error {
	switch mode {
	case fgen.BurstMode:
		// Set the modulation type to burst (MTYP5) and enable modulation (MENA1).
		return ch.inst.Command("MTYP5;MENA1")
	case fgen.ContinuousMode:
		// Disable modulation (MENA0).
		return ch.inst.Command("MENA0")
	}

	return errors.New("bad fgen operation mode")
}

// OutputEnabled determines if the output channel is enabled or disabled.
//
// OutputEnabled is the getter for the read-write IviFgenBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return false, ivi.ErrFunctionNotSupported
}

// SetOutputEnabled sets the output channel to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(_ bool) error {
	return ivi.ErrFunctionNotSupported
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
//
// OutputImpedance is the getter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputImpedance() (float64, error) {
	return outputImpedance, nil
}

// SetOutputImpedance sets the output channel's impedance in ohms.
//
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(impedance float64) error {
	if impedance != outputImpedance {
		return ivi.ErrFunctionNotSupported
	}

	return nil
}

// AbortGeneration aborts a previously initiated signal generation. Since the
// DS345 constantly generates a signal and the output cannot be aborted, this
// function does nothing and returns no error.
//
// AbortGeneration implements the IviFgenBase function described in Section 4.3.1
// of IVI-4.3: IviFgen Class Specification.
func (d *Driver) AbortGeneration() error {
	return nil
}

// InitiateGeneration initiates signal generation. Since the DS345 constantly
// generates a signal and the output cannot be disabled, this function does
// nothing and returns no error.
//
// InitiateGeneration implements the IviFgenBase function described in Section
// 4.3.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) InitiateGeneration() error {
	return nil
}

func (d *Driver) ReferenceClockSource() (fgen.ClockSource, error) {
	return fgen.RefClockInternal, nil
}

func (d *Driver) SetReferenceClockSource(src fgen.ClockSource) error {
	return nil
}

func (ch *Channel) Name() string {
	return "output"
}

const (
	sine      int = 0
	square    int = 1
	triangle  int = 2
	ramp      int = 3
	noise     int = 4
	arbitrary int = 5
)

const (
	linearSweep int = 0
	logSweep    int = 1
	internalAM  int = 2
	fm          int = 3
	phasem      int = 4
	burst       int = 5
)
