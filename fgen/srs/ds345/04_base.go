// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"context"
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

var outputModeToSCPI = map[fgen.OutputMode]string{
	fgen.OutputModeFunction:  "FUNC0",
	fgen.OutputModeArbitrary: "FUNC5",
	fgen.OutputModeNoise:     "FUNC4",
}

var operationModeToSCPI = map[fgen.OperationMode]string{
	fgen.BurstMode:      "MTYP5;MENA1",
	fgen.ContinuousMode: "MENA0",
}

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
func (d *Driver) OutputMode(ctx context.Context) (fgen.OutputMode, error) {
	var outputMode fgen.OutputMode

	funcType, err := query.Int(ctx, d.inst, "FUNC?")
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
func (d *Driver) SetOutputMode(ctx context.Context, outputMode fgen.OutputMode) error {
	cmd, err := ivi.LookupSCPI(outputModeToSCPI, outputMode)
	if err != nil {
		return fmt.Errorf("SetOutputMode: %w", err)
	}

	return d.inst.Command(ctx, cmd)
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel. OperationMode implements the
// getter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OperationMode(ctx context.Context) (fgen.OperationMode, error) {
	var mode fgen.OperationMode

	isModulationEnabled, err := query.Bool(ctx, ch.inst, "MENA?")
	if err != nil {
		return mode, fmt.Errorf("error determining if modulation is enabled: %w", err)
	}

	if !isModulationEnabled {
		return fgen.ContinuousMode, nil
	}

	modType, err := query.Int(ctx, ch.inst, "MTYP?")
	if err != nil {
		return mode, fmt.Errorf("error determining modulation type: %w", err)
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
func (ch *Channel) SetOperationMode(ctx context.Context, mode fgen.OperationMode) error {
	cmd, err := ivi.LookupSCPI(operationModeToSCPI, mode)
	if err != nil {
		return fmt.Errorf("SetOperationMode: %w", err)
	}

	return ch.inst.Command(ctx, cmd)
}

// OutputEnabled determines if the output channel is enabled or disabled.
//
// OutputEnabled is the getter for the read-write IviFgenBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OutputEnabled(_ context.Context) (bool, error) {
	return false, ivi.ErrFunctionNotSupported
}

// SetOutputEnabled sets the output channel to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(_ context.Context, _ bool) error {
	return ivi.ErrFunctionNotSupported
}

// DisableOutput is a convenience function for setting the Output Enabled
// attribute to false.
func (ch *Channel) DisableOutput(ctx context.Context) error {
	return ch.SetOutputEnabled(ctx, false)
}

// EnableOutput is a convenience function for setting the Output Enabled
// attribute to true.
func (ch *Channel) EnableOutput(ctx context.Context) error {
	return ch.SetOutputEnabled(ctx, true)
}

// OutputImpedance return the output channel's impedance in ohms.
//
// OutputImpedance is the getter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputImpedance(_ context.Context) (float64, error) {
	return outputImpedance, nil
}

// SetOutputImpedance sets the output channel's impedance in ohms.
//
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(_ context.Context, impedance float64) error {
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
func (d *Driver) AbortGeneration(_ context.Context) error {
	return nil
}

// InitiateGeneration initiates signal generation. Since the DS345 constantly
// generates a signal and the output cannot be disabled, this function does
// nothing and returns no error.
//
// InitiateGeneration implements the IviFgenBase function described in Section
// 4.3.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) InitiateGeneration(_ context.Context) error {
	return nil
}

func (d *Driver) ReferenceClockSource(_ context.Context) (fgen.ClockSource, error) {
	return fgen.RefClockInternal, nil
}

func (d *Driver) SetReferenceClockSource(_ context.Context, src fgen.ClockSource) error {
	return nil
}

func (ch *Channel) Name() string {
	return ch.name
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
