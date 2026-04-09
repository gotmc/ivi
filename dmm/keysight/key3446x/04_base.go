// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/convert"
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

// MeasurementFunction returns the currently specified measurement function.
//
// MeasurementFunction is the getter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) MeasurementFunction(ctx context.Context) (dmm.MeasurementFunction, error) {
	response, err := query.String(ctx, d.inst, "FUNC?")
	if err != nil {
		return 0, err
	}

	response = convert.StripDoubleQuotes(response)

	fcn, err := ivi.ReverseLookup(cmdToMsrFunc, response)
	if err != nil {
		return 0, fmt.Errorf("invalid measurement function %q: %w", response, err)
	}

	return fcn, nil
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
	cmd := fmt.Sprintf("FUNC %q", scpiCmd)

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
	fcn, err := d.MeasurementFunction(ctx)
	if err != nil {
		return 0, 0.0, err
	}

	scpiFunc, err := ivi.LookupSCPI(msrFuncToCmd, fcn)
	if err != nil {
		return 0, 0.0, err
	}

	isAutoRange, err := query.Boolf(
		ctx, d.inst, "%s:rang:auto?", scpiFunc,
	)
	if err != nil {
		return 0, 0.0, err
	}

	autoRange := dmm.AutoOn
	if !isAutoRange {
		autoRange = dmm.AutoOff
	}

	rng, err := query.Float64f(ctx, d.inst, "%s:rang?", scpiFunc)
	if err != nil {
		return 0, 0.0, err
	}

	return autoRange, rng, nil
}

// SetRange sets the range corresponding to the maximum input value based on
// the rest of the instrument configuration (the same as the IVI.NET behavior).
// Setting this property sets AutoRange to Auto.Off If the property is set to a
// negative value and that negative value is valid for the current function
// (for instance DC Volts) the instrument will configure to measure that value.
//
// SetRange is the setter for the read-write IviDmmBase Attribute
// Range described in Section 4.2.2 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetRange(ctx context.Context, autoRange dmm.AutoRange, rangeValue float64) error {
	fcn, err := d.MeasurementFunction(ctx)
	if err != nil {
		return err
	}

	scpiFunc, err := ivi.LookupSCPI(msrFuncToCmd, fcn)
	if err != nil {
		return err
	}

	// Set the range to auto if appropriate.
	if autoRange == dmm.AutoOn {
		return d.inst.Command(ctx, "%s:rang:auto on", scpiFunc)
	}

	// Not auto ranging, so  we need to determine the appropriate SCPI range
	// string for the given range value and measurement function.
	var rng string

	switch fcn {
	case dmm.DCVolts, dmm.ACVolts, dmm.ACPlusDCVolts:
		// 100 mV|1 V|10 V|100 V|1000 V
		rng, err = determineManualVoltageRange(rangeValue)
		if err != nil {
			return err
		}
	case dmm.DCCurrent, dmm.ACPlusDCCurrent:
		// 1 µA|10 µA|100 µA|1 mA|10 mA|100 mA|1 A|3 A
		rng, err = determineManualDCCurrentRange(rangeValue)
		if err != nil {
			return err
		}
	case dmm.ACCurrent:
		// 100 µA|1 mA|10 mA|100 mA|1 A|3 A
		rng, err = determineManualACCurrentRange(rangeValue)
		if err != nil {
			return err
		}
	case dmm.TwoWireResistance, dmm.FourWireResistance:
		// 100 Ω|1 kΩ|10 kΩ|100 kΩ|1 MΩ|10 MΩ|100 MΩ|1 GΩ
		rng, err = determineManualResistanceRange(rangeValue)
		if err != nil {
			return err
		}
	case dmm.Frequency, dmm.Period:
		// Frequency/period range is a voltage range: 100 mV|1 V|10 V|100 V|750 V
		rng, err = determineManualFrequencyVoltageRange(rangeValue)
		if err != nil {
			return err
		}
	case dmm.Temperature:
		// Temperature has no user-selectable range; it is determined by probe type.
		return fmt.Errorf("SetRange: %w: temperature range is probe-determined", ivi.ErrFunctionNotSupported)
	}

	return d.inst.Command(ctx, "%s:rang %s", scpiFunc, rng)
}

// ResolutionAbsolute returns the measurement resolution in absolute units (e.g.,
// Volts for voltage measurements).
//
// ResolutionAbsolute is the getter for the read-write IviDmmBase Attribute
// Resolution Absolute described in Section 4.2.3 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) ResolutionAbsolute(ctx context.Context) (float64, error) {
	fcn, err := d.MeasurementFunction(ctx)
	if err != nil {
		return 0.0, err
	}

	scpiFunc, err := ivi.LookupSCPI(msrFuncToCmd, fcn)
	if err != nil {
		return 0.0, err
	}

	return query.Float64f(ctx, d.inst, "%s:RES?", scpiFunc)
}

// SetResolutionAbsolute sets the measurement resolution in absolute units.
//
// SetResolutionAbsolute is the setter for the read-write IviDmmBase Attribute
// Resolution Absolute described in Section 4.2.3 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetResolutionAbsolute(
	ctx context.Context,
	resolution float64,
) error {
	fcn, err := d.MeasurementFunction(ctx)
	if err != nil {
		return err
	}

	scpiFunc, err := ivi.LookupSCPI(msrFuncToCmd, fcn)
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, "%s:RES %g", scpiFunc, resolution)
}

// TriggerDelay returns whether auto delay is enabled and the trigger delay
// duration.
//
// TriggerDelay is the getter for the read-write IviDmmBase Attribute Trigger
// Delay described in Section 4.2.5 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) TriggerDelay(ctx context.Context) (bool, time.Duration, error) {
	autoDelay, err := query.Bool(ctx, d.inst, "TRIG:DEL:AUTO?")
	if err != nil {
		return false, 0, fmt.Errorf("TriggerDelay: %w", err)
	}

	seconds, err := query.Float64(ctx, d.inst, "TRIG:DEL?")
	if err != nil {
		return false, 0, fmt.Errorf("TriggerDelay: %w", err)
	}

	delay := time.Duration(seconds * float64(time.Second))

	return autoDelay, delay, nil
}

// SetTriggerDelay sets the trigger delay. If autoDelay is true, the instrument
// determines the delay automatically; otherwise, the specified delay is used.
//
// SetTriggerDelay is the setter for the read-write IviDmmBase Attribute
// Trigger Delay described in Section 4.2.5 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetTriggerDelay(
	ctx context.Context,
	autoDelay bool,
	delay time.Duration,
) error {
	if autoDelay {
		return d.inst.Command(ctx, "TRIG:DEL:AUTO ON")
	}

	seconds := delay.Seconds()

	return d.inst.Command(ctx, "TRIG:DEL %f", seconds)
}

// TriggerSource returns the current trigger source.
//
// TriggerSource is the getter for the read-write IviDmmBase Attribute Trigger
// Source described in Section 4.2.6 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) TriggerSource(ctx context.Context) (dmm.TriggerSource, error) {
	s, err := query.String(ctx, d.inst, "TRIG:SOUR?")
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	src, err := ivi.ReverseLookup(scpiToTriggerSource, convert.StripDoubleQuotes(s))
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	return src, nil
}

// SetTriggerSource sets the trigger source.
//
// SetTriggerSource is the setter for the read-write IviDmmBase Attribute
// Trigger Source described in Section 4.2.6 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetTriggerSource(ctx context.Context, src dmm.TriggerSource) error {
	cmd, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("SetTriggerSource: %w", err)
	}

	return d.inst.Command(ctx, "TRIG:SOUR %s", cmd)
}

// Abort aborts a measurement in progress, returning the instrument to the
// trigger idle state.
//
// Abort implements the IviDmmBase function described in Section 4.3.1 of
// IVI-4.2: IviDmm Class Specification.
func (d *Driver) Abort(ctx context.Context) error {
	return d.inst.Command(ctx, "ABOR")
}

func (d *Driver) ConfigureMeasurement(
	ctx context.Context,
	msrFunc dmm.MeasurementFunction,
	autoRange dmm.AutoRange,
	rangeValue float64,
	resolution float64,
) error {
	cmd, err := createConfigureMeasurementCommand(msrFunc, autoRange, rangeValue, resolution)
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, cmd)
}

// ConfigureTrigger configures the trigger source and trigger delay for the
// DMM.
//
// ConfigureTrigger implements the IviDmmBase function described in Section
// 4.3.3 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) ConfigureTrigger(
	ctx context.Context,
	src dmm.TriggerSource,
	delay time.Duration,
) error {
	if err := d.SetTriggerSource(ctx, src); err != nil {
		return err
	}

	return d.SetTriggerDelay(ctx, false, delay)
}

// FetchMeasurement returns the measured value from a measurement that the
// Initiate function initiates. After this function executes, the Reading
// parameter contains an actual reading or a value indicating that an overrange
// condition occurred.
//
// Currently, the maxTime is ignored.
//
// FetchMeasurement implements the IviDmmBase function described in Section
// 4.3.4 of the IVI-4.2 IviDmm Class Specification.
func (d *Driver) FetchMeasurement(ctx context.Context, _ time.Duration) (float64, error) {
	return query.Float64(ctx, d.inst, "FETC?")
}

// InitiateMeasurement initiates a measurement. When this function executes,
// the DMM leaves the idle state and waits for a trigger.
//
// This function does not check the instrument status. Typically, the end-user
// calls this function only in a sequence of calls to other low-level driver
// functions. The sequence performs one operation. The end-user uses the
// low-level functions to optimize one or more aspects of interaction with the
// instrument. To check the instrument status, call the Error Query function at
// the conclusion of the sequence.
//
// InitiateMeasurement implements the IviDmmBase function described in Section
// 4.3.5 of the IVI-4.2 IviDmm Class Specification.
func (d *Driver) InitiateMeasurement(ctx context.Context) error {
	return d.inst.Command(ctx, "init")
}

// overRangeValue is the value returned by Keysight 3446x DMMs when a
// measurement exceeds the selected range. Positive for over-range, negative
// for under-range.
const overRangeValue = 9.9e37

// IsOutOfRange returns true if the given value indicates an over-range or
// under-range condition.
//
// IsOutOfRange implements the IviDmmBase function described in Section 4.3.6
// of IVI-4.2: IviDmm Class Specification.
func (d *Driver) IsOutOfRange(_ context.Context, value float64) (bool, error) {
	return value >= overRangeValue || value <= -overRangeValue, nil
}

// IsOverRange returns true if the given value indicates a positive over-range
// condition.
//
// IsOverRange implements the IviDmmBase function described in Section 4.3.7
// of IVI-4.2: IviDmm Class Specification.
func (d *Driver) IsOverRange(_ context.Context, value float64) (bool, error) {
	return value >= overRangeValue, nil
}

// IsUnderRange returns true if the given value indicates a negative
// under-range condition.
//
// IsUnderRange implements the IviDmmBase function described in Section 4.3.8
// of IVI-4.2: IviDmm Class Specification.
func (d *Driver) IsUnderRange(_ context.Context, value float64) (bool, error) {
	return value <= -overRangeValue, nil
}

func (d *Driver) ReadMeasurement(ctx context.Context, _ time.Duration) (float64, error) {
	return query.Float64(ctx, d.inst, "read?")
}

// cmdToMsrFunc maps the SCPI command string name of a measurement function to
// the MeasurementFunction.
var cmdToMsrFunc = map[string]dmm.MeasurementFunction{
	"VOLT":    dmm.DCVolts,
	"VOLT:DC": dmm.DCVolts,
	"VOLT:AC": dmm.ACVolts,
	"CURR":    dmm.DCCurrent,
	"CURR:DC": dmm.DCCurrent,
	"CURR:AC": dmm.ACCurrent,
	"RES":     dmm.TwoWireResistance,
	"FRES":    dmm.FourWireResistance,
	"FREQ":    dmm.Frequency,
	"PER":     dmm.Period,
	"TEMP":    dmm.Temperature,
}

// msrFuncToCmd maps the MeasurementFunction to the SCPI command string
var msrFuncToCmd = map[dmm.MeasurementFunction]string{
	dmm.DCVolts:            "VOLT",
	dmm.ACVolts:            "VOLT:AC",
	dmm.DCCurrent:          "CURR",
	dmm.ACCurrent:          "CURR:AC",
	dmm.TwoWireResistance:  "RES",
	dmm.FourWireResistance: "FRES",
	dmm.ACPlusDCVolts:      "VOLT:AC",
	dmm.ACPlusDCCurrent:    "CURR:AC",
	dmm.Frequency:          "FREQ",
	dmm.Period:             "PER",
	dmm.Temperature:        "TEMP",
}

var scpiToTriggerSource = map[string]dmm.TriggerSource{
	"IMM": dmm.Immediate,
	"EXT": dmm.External,
	"BUS": dmm.SoftwareTrigger,
}

var triggerSourceToSCPI = map[dmm.TriggerSource]string{
	dmm.Immediate:       "IMM",
	dmm.External:        "EXT",
	dmm.SoftwareTrigger: "BUS",
}

func createConfigureVoltageDCCommand(
	autoRange dmm.AutoRange,
	rangeValue float64,
	resolution float64,
) (string, error) {
	rng, err := determineVoltageRange(autoRange, rangeValue)
	if err != nil {
		return "", ivi.ErrNotImplemented
	}

	if autoRange == dmm.AutoOff {
		return fmt.Sprintf("CONF:VOLT:DC %s,%f", rng, resolution), nil
	}

	return fmt.Sprintf("CONF:VOLT:DC %s", rng), nil
}

func createConfigureVoltageACCommand(
	autoRange dmm.AutoRange,
	rangeValue float64,
) (string, error) {
	rng, err := determineVoltageRange(autoRange, rangeValue)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CONF:VOLT:AC %s", rng), nil
}

func determineVoltageRange(autoRange dmm.AutoRange, rangeValue float64) (string, error) {
	switch autoRange {
	case dmm.AutoOn:
		return "AUTO", nil
	case dmm.AutoOff:
		return determineManualVoltageRange(rangeValue)
	case dmm.AutoOnce:
		return "", ivi.ErrNotImplemented
	}

	return "", ivi.ErrNotImplemented
}

func determineManualVoltageRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 0.1:
		return "0.1", nil
	case rangeValue <= 1.0:
		return "1", nil
	case rangeValue <= 10.0:
		return "10", nil
	case rangeValue <= 100.0:
		return "100", nil
	case rangeValue <= 1000.0:
		return "1000", nil
	}

	return "", ivi.ErrValueNotSupported
}

func determineManualResistanceRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 100:
		return "100", nil
	case rangeValue <= 1.0e3:
		return "1e3", nil
	case rangeValue <= 10e3:
		return "10e3", nil
	case rangeValue <= 100e3:
		return "100e3", nil
	case rangeValue <= 1e6:
		return "1e6", nil
	case rangeValue <= 10e6:
		return "10e6", nil
	case rangeValue <= 100e6:
		return "100e6", nil
	case rangeValue <= 1e9:
		return "1e9", nil
	}

	return "", ivi.ErrValueNotSupported
}

func determineManualDCCurrentRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 1e-6:
		return "1e-6", nil
	case rangeValue <= 10e-6:
		return "10e-6", nil
	case rangeValue <= 100e-6:
		return "100e-6", nil
	case rangeValue <= 1e-3:
		return "1e-3", nil
	case rangeValue <= 10e-3:
		return "10e-3", nil
	case rangeValue <= 100e-3:
		return "100e-3", nil
	case rangeValue <= 1.0:
		return "1", nil
	case rangeValue <= 3.0:
		return "3", nil
	}

	return "", ivi.ErrValueNotSupported
}

func determineManualACCurrentRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 100e-6:
		return "100e-6", nil
	case rangeValue <= 1e-3:
		return "1e-3", nil
	case rangeValue <= 10e-3:
		return "10e-3", nil
	case rangeValue <= 100e-3:
		return "100e-3", nil
	case rangeValue <= 1.0:
		return "1", nil
	case rangeValue <= 3.0:
		return "3", nil
	}

	return "", ivi.ErrValueNotSupported
}

func determineManualFrequencyVoltageRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 0.1:
		return "0.1", nil
	case rangeValue <= 1.0:
		return "1", nil
	case rangeValue <= 10.0:
		return "10", nil
	case rangeValue <= 100.0:
		return "100", nil
	case rangeValue <= 750.0:
		return "750", nil
	}

	return "", ivi.ErrValueNotSupported
}

func createConfigureMeasurementCommand(
	msrFunc dmm.MeasurementFunction,
	autoRange dmm.AutoRange,
	rangeValue float64,
	resolution float64,
) (string, error) {
	switch msrFunc {
	case dmm.DCVolts:
		return createConfigureVoltageDCCommand(autoRange, rangeValue, resolution)
	case dmm.ACVolts:
		return createConfigureVoltageACCommand(autoRange, rangeValue)
	case dmm.DCCurrent:
		// return createConfigureCurrentDCCommand(autoRange, rangeValue, resolution)
		return "", ivi.ErrNotImplemented
	case dmm.ACCurrent:
		// return createConfigureCurrentACCommand(autoRange, rangeValue)
		return "", ivi.ErrNotImplemented
	case dmm.TwoWireResistance:
		return "", ivi.ErrNotImplemented
	case dmm.FourWireResistance:
		return "", ivi.ErrNotImplemented
	case dmm.ACPlusDCVolts:
		return "", ivi.ErrNotImplemented
	case dmm.ACPlusDCCurrent:
		return "", ivi.ErrNotImplemented
	case dmm.Frequency:
		return "", ivi.ErrNotImplemented
	case dmm.Period:
		return "", ivi.ErrNotImplemented
	case dmm.Temperature:
		return "", ivi.ErrNotImplemented
	}

	return "", ivi.ErrNotImplemented
}
