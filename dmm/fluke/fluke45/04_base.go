// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fluke45

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

// MeasurementFunction returns the currently specified measurement function.
//
// MeasurementFunction is the getter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) MeasurementFunction(ctx context.Context) (dmm.MeasurementFunction, error) {
	fcn, err := query.String(ctx, d.inst, "FUNC1?")
	if err != nil {
		return 0, err
	}

	msrFunc, err := ivi.ReverseLookup(cmdToMsrFunc, strings.TrimSpace(fcn))
	if err != nil {
		return 0, fmt.Errorf("invalid measurement function %q: %w", fcn, err)
	}

	return msrFunc, nil
}

// SetMeasurementFunction specifies the measurement function.
//
// SetMeasurementFunction is the setter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetMeasurementFunction(
	ctx context.Context,
	msrFunc dmm.MeasurementFunction,
) error {
	scpiCmd, err := ivi.LookupSCPI(msrFuncToCmd, msrFunc)
	if err != nil {
		return fmt.Errorf("measurement function %v not supported: %w", msrFunc, err)
	}

	// Need to return a quoted string, so use %q in the fmt.Sprintf
	cmd := fmt.Sprintf("%q", scpiCmd)

	return d.inst.Command(ctx, cmd)
}

// Range returns the measurement range and whether auto range is enabled,
// disabled, or enabled for one measurement.
//
// There is a dependency between the Range attribute and the Resolution
// Absolute attribute. The allowed values of Resolution Absolute attribute
// depend on the Range attribute. Typically, when the value of the Range
// attribute changes, the instrument settings that correspond to the Resolution
// Absolute attribute change as well. This is true regardless of how the change
// of measurement range occurs.
//
// There are two possible ways that the measurement range can change. The
// application program can set the value of the Range attribute. Or, the
// instrument changes the measurement range because Range attribute is set to
// Auto Range On and the input signal changes. In both cases, the instrument
// resolution is likely to change.
//
// The value of the MeasurementFunction attribute determines the units for this
// attribute as follows:
//
// DC Volts = Volts
// AC Volts = Volts RMS
// DC Current = Amps
// AC Current = Amps
// 2-Wire Resistance = Ohms
// 4-Wire Resistance = Ohms
// AC Plus DC Volts = Volts
// AC Plus DC Current = Amps
// Frequency = Hertz
// Period = Seconds
// Temperature = Degrees Celsius
//
// Range is the getter for the read-write IviDmmBase Attribute Range described
// in Section 4.2.2 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) Range(ctx context.Context) (dmm.AutoRange, float64, error) {
	isAutoRange, err := query.Bool(ctx, d.inst, "AUTO?")
	if err != nil {
		return 0, 0.0, err
	}

	autoRange := dmm.AutoOn
	if !isAutoRange {
		autoRange = dmm.AutoOff
	}

	rng, err := query.Float64(ctx, d.inst, "RANG1?")
	if err != nil {
		return 0, 0.0, err
	}

	return autoRange, rng, nil
}

// SetRange sets the range corresponding to the maximum input value based on
// the rest of hte instrument configuration (the same as the IVI.NET behavior).
// Setting this property sets AutoRange to Auto.Off If the property is set to a
// negative value and that negative value is valid for the current function
// (for instance DC Volts) the instrument will configure to measure that value.
//
// SetRange is the setter for the read-write IviDmmBase Attribute
// Range described in Section 4.2.2 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetRange(ctx context.Context, autoRange dmm.AutoRange, rangeValue float64) error {
	// Set the range to auto if appropriate.
	if autoRange == dmm.AutoOn {
		return d.inst.Command(ctx, "auto")
	}

	// Not auto ranging, so  we need to determine the appropriate SCPI range
	// string for the given range value and measurement function.

	// The Fluke 45 has three sampling rates—slow (2.5 readings/sec), medium (5.0
	// readings/sec), and fast (20 readings/sec). The ranges differ based on the
	// selected sampling rate, so we need to query the rate first.
	rate, err := query.String(ctx, d.inst, "RATE?")
	if err != nil {
		return err
	}

	fcn, err := d.MeasurementFunction(ctx)
	if err != nil {
		return err
	}

	rangeCmd, err := determineRangeCommand(rate, fcn, rangeValue)
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, "RANG %s", rangeCmd)
}

func determineRangeCommand(
	rate string,
	fcn dmm.MeasurementFunction,
	rangeValue float64,
) (string, error) {
	switch {
	case fcn == dmm.DCVolts && (rate == "F" || rate == "M"):
		switch {
		case rangeValue <= 0.3:
			return "1", nil
		case rangeValue <= 3.0:
			return "2", nil
		case rangeValue <= 30.0:
			return "3", nil
		case rangeValue <= 300.0:
			return "4", nil
		case rangeValue <= 1000.0:
			return "5", nil
		}
	case fcn == dmm.DCVolts && rate == "S":
		switch {
		case rangeValue <= 0.1:
			return "1", nil
		case rangeValue <= 1.0:
			return "2", nil
		case rangeValue <= 10.0:
			return "3", nil
		case rangeValue <= 100.0:
			return "4", nil
		case rangeValue <= 1000.0:
			return "5", nil
		}
	case fcn == dmm.ACVolts && (rate == "F" || rate == "M"):
		switch {
		case rangeValue <= 0.3:
			return "1", nil
		case rangeValue <= 3.0:
			return "2", nil
		case rangeValue <= 30.0:
			return "3", nil
		case rangeValue <= 300.0:
			return "4", nil
		case rangeValue <= 750.0:
			return "5", nil
		}
	case fcn == dmm.ACVolts && rate == "S":
		switch {
		case rangeValue <= 0.1:
			return "1", nil
		case rangeValue <= 1.0:
			return "2", nil
		case rangeValue <= 10.0:
			return "3", nil
		case rangeValue <= 100.0:
			return "4", nil
		case rangeValue <= 750.0:
			return "5", nil
		}
	case fcn == dmm.TwoWireResistance && (rate == "F" || rate == "M"):
		switch {
		case rangeValue <= 300:
			return "1", nil
		case rangeValue <= 3e3:
			return "2", nil
		case rangeValue <= 30e3:
			return "3", nil
		case rangeValue <= 300e3:
			return "4", nil
		case rangeValue <= 3e6:
			return "5", nil
		case rangeValue <= 30e6:
			return "6", nil
		case rangeValue <= 300e6:
			return "7", nil
		}
	case fcn == dmm.TwoWireResistance && rate == "S":
		switch {
		case rangeValue <= 100:
			return "1", nil
		case rangeValue <= 1e3:
			return "2", nil
		case rangeValue <= 10e3:
			return "3", nil
		case rangeValue <= 100e3:
			return "4", nil
		case rangeValue <= 1e6:
			return "5", nil
		case rangeValue <= 10e6:
			return "6", nil
		case rangeValue <= 100e6:
			return "7", nil
		}
	}
	return "", fmt.Errorf(
		"error determining range for rate = %s / fcn = %s / rangeValue = %g",
		rate,
		fcn,
		rangeValue,
	)
}

// ResolutionAbsolute is not directly queryable on the Fluke 45. Resolution is
// determined by the measurement rate (S/M/F).
//
// ResolutionAbsolute is the getter for the read-write IviDmmBase Attribute
// Resolution Absolute described in Section 4.2.3 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) ResolutionAbsolute(_ context.Context) (float64, error) {
	return 0.0, fmt.Errorf(
		"ResolutionAbsolute: %w", ivi.ErrFunctionNotSupported,
	)
}

// SetResolutionAbsolute is not directly supported on the Fluke 45. Resolution
// is controlled by the measurement rate (RATE S/M/F).
func (d *Driver) SetResolutionAbsolute(
	_ context.Context,
	_ float64,
) error {
	return fmt.Errorf(
		"SetResolutionAbsolute: %w", ivi.ErrFunctionNotSupported,
	)
}

// TriggerDelay returns the trigger configuration. The Fluke 45 does not have a
// separate delay setting; settling delay is part of the trigger type. autoDelay
// is true when the trigger type includes a settling delay (types 3 or 5).
//
// TriggerDelay is the getter for the read-write IviDmmBase Attribute Trigger
// Delay described in Section 4.2.5 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) TriggerDelay(
	ctx context.Context,
) (bool, time.Duration, error) {
	trigType, err := query.Int(ctx, d.inst, "TRIGGER?")
	if err != nil {
		return false, 0, fmt.Errorf("TriggerDelay: %w", err)
	}

	// Trigger types 3 and 5 include settling delay.
	hasDelay := trigType == 3 || trigType == 5

	return hasDelay, 0, nil
}

// SetTriggerDelay sets the trigger type on the Fluke 45. The Fluke 45 does not
// support a configurable delay duration; the settling delay is either enabled
// or disabled as part of the trigger type. When autoDelay is true, a trigger
// type with settling delay is selected.
//
// SetTriggerDelay is the setter for the read-write IviDmmBase Attribute
// Trigger Delay described in Section 4.2.5 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetTriggerDelay(
	ctx context.Context,
	autoDelay bool,
	_ time.Duration,
) error {
	// Get current trigger type to determine if internal or external.
	trigType, err := query.Int(ctx, d.inst, "TRIGGER?")
	if err != nil {
		return fmt.Errorf("SetTriggerDelay: %w", err)
	}

	isInternal := trigType == 1

	if isInternal {
		// Internal trigger (type 1) does not support settling delay.
		if autoDelay {
			return fmt.Errorf(
				"SetTriggerDelay: settling delay not available with internal trigger: %w",
				ivi.ErrValueNotSupported,
			)
		}

		return nil
	}

	// External trigger: select type with or without settling delay.
	if autoDelay {
		return d.inst.Command(ctx, "TRIGGER 3")
	}

	return d.inst.Command(ctx, "TRIGGER 2")
}

// TriggerSource returns the current trigger source. The Fluke 45 uses trigger
// types 1-5: type 1 is internal, types 2-5 are external.
//
// TriggerSource is the getter for the read-write IviDmmBase Attribute Trigger
// Source described in Section 4.2.6 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) TriggerSource(
	ctx context.Context,
) (dmm.TriggerSource, error) {
	trigType, err := query.Int(ctx, d.inst, "TRIGGER?")
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	if trigType == 1 {
		return dmm.TriggerSourceImmediate, nil
	}

	return dmm.TriggerSourceExternal, nil
}

// SetTriggerSource sets the trigger source. The Fluke 45 supports internal
// (Immediate) and external trigger sources.
//
// SetTriggerSource is the setter for the read-write IviDmmBase Attribute
// Trigger Source described in Section 4.2.6 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetTriggerSource(
	ctx context.Context,
	src dmm.TriggerSource,
) error {
	switch src {
	case dmm.TriggerSourceImmediate:
		return d.inst.Command(ctx, "TRIGGER 1")
	case dmm.TriggerSourceExternal:
		return d.inst.Command(ctx, "TRIGGER 2")
	case dmm.TriggerSourceSoftware:
		// External trigger with *TRG for software-initiated triggers.
		return d.inst.Command(ctx, "TRIGGER 2")
	default:
		return fmt.Errorf(
			"SetTriggerSource %v: %w", src, ivi.ErrValueNotSupported,
		)
	}
}

// Abort is not supported on the Fluke 45.
func (d *Driver) Abort(_ context.Context) error {
	return fmt.Errorf("Abort: %w", ivi.ErrFunctionNotSupported)
}

// ConfigureMeasurement configures the measurement function and range.
//
// ConfigureMeasurement implements the IviDmmBase function described in Section
// 4.3.2 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) ConfigureMeasurement(
	ctx context.Context,
	msrFunc dmm.MeasurementFunction,
	autoRange dmm.AutoRange,
	rangeValue float64,
	_ float64,
) error {
	if err := d.SetMeasurementFunction(ctx, msrFunc); err != nil {
		return err
	}

	return d.SetRange(ctx, autoRange, rangeValue)
}

// ConfigureTrigger configures the trigger source. The Fluke 45 does not
// support a separate trigger delay duration; the delay parameter is ignored.
//
// ConfigureTrigger implements the IviDmmBase function described in Section
// 4.3.3 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) ConfigureTrigger(
	ctx context.Context,
	src dmm.TriggerSource,
	_ time.Duration,
) error {
	return d.SetTriggerSource(ctx, src)
}

// FetchMeasurement returns the value shown on the primary display without
// triggering a new measurement. Uses the VAL1? query.
//
// FetchMeasurement implements the IviDmmBase function described in Section
// 4.3.4 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) FetchMeasurement(
	ctx context.Context,
	_ time.Duration,
) (float64, error) {
	return query.Float64(ctx, d.inst, "VAL1?")
}

// InitiateMeasurement triggers a measurement using the *TRG command. The
// instrument must be in an external trigger mode (TRIGGER 2-5) for *TRG to
// take effect.
//
// InitiateMeasurement implements the IviDmmBase function described in Section
// 4.3.5 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) InitiateMeasurement(ctx context.Context) error {
	return d.inst.Command(ctx, "*TRG")
}

func (d *Driver) IsOutOfRange(_ context.Context, _ float64) (bool, error) {
	return false, fmt.Errorf(
		"IsOutOfRange: %w", ivi.ErrFunctionNotSupported,
	)
}

func (d *Driver) IsOverRange(_ context.Context, _ float64) (bool, error) {
	return false, fmt.Errorf(
		"IsOverRange: %w", ivi.ErrFunctionNotSupported,
	)
}

func (d *Driver) IsUnderRange(_ context.Context, _ float64) (bool, error) {
	return false, fmt.Errorf(
		"IsUnderRange: %w", ivi.ErrFunctionNotSupported,
	)
}

func (d *Driver) ReadMeasurement(ctx context.Context, maxTime time.Duration) (float64, error) {
	return query.Float64(ctx, d.inst, "meas1?")
}

// cmdToMsrFunc maps the SCPI command string name of a measurement function to
// the MeasurementFunction.
var cmdToMsrFunc = map[string]dmm.MeasurementFunction{
	"VDC":   dmm.DCVolts,
	"VAC":   dmm.ACVolts,
	"ADC":   dmm.DCCurrent,
	"AAC":   dmm.ACCurrent,
	"OHMS":  dmm.TwoWireResistance,
	"FREQ":  dmm.Frequency,
	"VACDC": dmm.ACPlusDCVolts,
	"AACDC": dmm.ACPlusDCCurrent,
}

// msrFuncToCmd maps the MeasurementFunction to the SCPI command string
var msrFuncToCmd = map[dmm.MeasurementFunction]string{
	dmm.DCVolts:           "VDC",
	dmm.ACVolts:           "VAC",
	dmm.DCCurrent:         "ADC",
	dmm.ACCurrent:         "AAC",
	dmm.TwoWireResistance: "OHMS",
	dmm.ACPlusDCVolts:     "VACDC",
	dmm.ACPlusDCCurrent:   "AACDC",
	dmm.Frequency:         "FREQ",
}
