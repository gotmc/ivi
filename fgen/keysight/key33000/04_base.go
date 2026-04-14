// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33000

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

var scpiToOutputMode = map[string]fgen.OutputMode{
	"SIN":  fgen.OutputModeFunction,
	"SQU":  fgen.OutputModeFunction,
	"RAMP": fgen.OutputModeFunction,
	"TRI":  fgen.OutputModeFunction,
	"PULS": fgen.OutputModeFunction,
	"DC":   fgen.OutputModeFunction,
	"PRBS": fgen.OutputModeFunction,
	"NOIS": fgen.OutputModeNoise,
	"ARB":  fgen.OutputModeArbitrary,
}

var outputModeToSCPI = map[fgen.OutputMode]string{
	fgen.OutputModeFunction:  "FUNC SIN",
	fgen.OutputModeArbitrary: "FUNC ARB",
	fgen.OutputModeNoise:     "FUNC NOIS",
}

var scpiToOperationMode = map[string]fgen.OperationMode{
	"0": fgen.ContinuousMode,
	"1": fgen.BurstMode,
}

var operationModeToSCPI = map[fgen.OperationMode]string{
	fgen.BurstMode:      "BURS:MODE TRIG;STAT ON",
	fgen.ContinuousMode: "BURS:STAT OFF",
}

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviFgenBase Attribute Output
// Count described in Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputCount() int {
	return len(d.channels)
}

// OutputMode returns how the function generator produces waveforms. This
// attribute determines which extension group's functions and attributes are
// used to configure the waveform the function generator produces.
//
// OutputMode is the getter for the read-write IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputMode(ctx context.Context) (fgen.OutputMode, error) {
	s, err := query.String(ctx, d.inst, d.channels[0].srcPrefix()+"FUNC?")
	if err != nil {
		return 0, fmt.Errorf("OutputMode: %w", err)
	}

	mode, err := ivi.ReverseLookup(scpiToOutputMode, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("OutputMode: %w", err)
	}

	return mode, nil
}

// SetOutputMode sets how the function generator produces waveforms.
//
// SetOutputMode is the setter for the read-write IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) SetOutputMode(ctx context.Context, outputMode fgen.OutputMode) error {
	cmd, err := ivi.LookupSCPI(outputModeToSCPI, outputMode)
	if err != nil {
		return fmt.Errorf("SetOutputMode: %w", err)
	}

	return d.inst.Command(ctx, d.channels[0].srcPrefix()+cmd)
}

// InitiateGeneration initiates signal generation by enabling all outputs.
//
// InitiateGeneration implements the IviFgenBase function described in Section
// 4.3.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) InitiateGeneration(ctx context.Context) error {
	for _, channel := range d.channels {
		if err := channel.EnableOutput(ctx); err != nil {
			return err
		}
	}

	return nil
}

// AbortGeneration aborts a previously initiated signal generation by disabling
// all outputs.
//
// AbortGeneration implements the IviFgenBase function described in Section
// 4.3.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) AbortGeneration(ctx context.Context) error {
	for _, channel := range d.channels {
		if err := channel.DisableOutput(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (d *Driver) ReferenceClockSource(_ context.Context) (fgen.ClockSource, error) {
	return fgen.RefClockInternal, nil
}

func (d *Driver) SetReferenceClockSource(_ context.Context, _ fgen.ClockSource) error {
	return nil
}

func (ch *Channel) Name() string {
	return ch.name
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel.
//
// OperationMode implements the getter for the read-write IviFgenBase Attribute
// Operation Mode described in Section 4.2.2 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OperationMode(ctx context.Context) (fgen.OperationMode, error) {
	s, err := query.String(ctx, ch.inst, ch.srcPrefix()+"BURS:STAT?")
	if err != nil {
		return 0, fmt.Errorf("OperationMode: %w", err)
	}

	mode, err := ivi.ReverseLookup(scpiToOperationMode, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("OperationMode: %w", err)
	}

	return mode, nil
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

	return ch.inst.Command(ctx, ch.srcPrefix()+cmd)
}

// OutputEnabled determines if the output channel is enabled or disabled.
//
// OutputEnabled is the getter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputEnabled(ctx context.Context) (bool, error) {
	return query.Boolf(ctx, ch.inst, "OUTP%d?", ch.num)
}

// SetOutputEnabled sets the output channel to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(ctx context.Context, b bool) error {
	if b {
		return ch.inst.Command(ctx, "OUTP%d ON", ch.num)
	}

	return ch.inst.Command(ctx, "OUTP%d OFF", ch.num)
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
func (ch *Channel) OutputImpedance(ctx context.Context) (float64, error) {
	return query.Float64f(ctx, ch.inst, "OUTP%d:LOAD?", ch.num)
}

// SetOutputImpedance sets the output channel's impedance in ohms.
//
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(ctx context.Context, impedance float64) error {
	return ch.inst.Command(ctx, "OUTP%d:LOAD %f", ch.num, impedance)
}

// AbortGeneration aborts a previously initiated signal generation.
//
// AbortGeneration implements the IviFgenBase function described in Section
// 4.3.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) AbortGeneration(ctx context.Context) error {
	return ch.DisableOutput(ctx)
}
