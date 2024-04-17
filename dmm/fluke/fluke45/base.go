// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package fluke45 implements the IVI driver for the Fluke 45 DMM.

State Caching: Not implemented
*/
package fluke45

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

const (
	specMajorVersion = 4
	specMinorVersion = 2
	specRevision     = "4.1"
)

// Confirm the driver implements the interface for the IviDMMBase capability
// group.
var _ dmm.Base = (*Driver)(nil)

// Driver provides the IVI driver for the Fluke 45 DMM.
type Driver struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Driver, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		GroupCapabilities: []string{
			"IviDmmBase",
			"IviDmmACMeasurement",
			"IviDmmFrequencyMeasurement",
			"IviDmmDeviceInfo",
		},
		SupportedInstrumentModels: []string{
			"45",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"Serial",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
		inst:     inst,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// QueryString queries the DMM and returns a string.
func (d *Driver) QueryString(cmd string) (string, error) {
	return query.String(d.inst, cmd)
}

// MeasurementFunction returns the currently specified measurement function.
//
// MeasurementFunction is the getter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) MeasurementFunction() (dmm.MeasurementFunction, error) {
	fcn, err := query.String(d.inst, "FUNC1?")
	if err != nil {
		return 0, err
	}

	switch strings.TrimSpace(fcn) {
	case "VDC":
		return dmm.DCVolts, nil
	case "VAC":
		return dmm.ACVolts, nil
	case "ADC":
		return dmm.DCCurrent, nil
	case "AAC":
		return dmm.ACCurrent, nil
	case "OHMS":
		return dmm.TwoWireResistance, nil
	case "VACDC":
		return dmm.ACPlusDCVolts, nil
	case "AACDC":
		return dmm.ACPlusDCCurrent, nil
	default:
		return 0, fmt.Errorf("%s is not a valid measurement function", fcn)
	}
}

// SetMeasurementFunction specifies the measurement function.
//
// SetMeasurementFunction is the setter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetMeasurementFunction(msrFunc dmm.MeasurementFunction) error {
	// Need to return a quoted string, so use %q in the fmt.Sprintf
	cmd := fmt.Sprintf("%q", msrFuncToCmd[msrFunc])
	return d.inst.Command(cmd)
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
func (d *Driver) Range() (dmm.AutoRange, float64, error) {
	isAutoRange, err := query.Bool(d.inst, "AUTO?")
	if err != nil {
		return 0, 0.0, err
	}

	autoRange := dmm.AutoOn
	if !isAutoRange {
		autoRange = dmm.AutoOff
	}

	rng, err := query.Float64(d.inst, "RANG1?")
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
func (d *Driver) SetRange(autoRange dmm.AutoRange, rangeValue float64) error {
	// Set the range to auto if appropriate.
	if autoRange == dmm.AutoOn {
		return d.inst.Command("auto")
	}

	// Not auto ranging, so  we need to determine the appropriate SCPI range
	// string for the given range value and measurement function.

	// The Fluke 45 has three sampling rates—slow (2.5 readings/sec), medium (5.0
	// readings/sec), and fast (20 readings/sec). The ranges differ based on the
	// selected sampling rate, so we needf to query the rate first.
	rate, err := query.String(d.inst, "RATE?")

	fcn, err := d.MeasurementFunction()
	if err != nil {
		return err
	}

	rangeCmd, err := determineRangeCommand(rate, fcn, rangeValue)

	return d.inst.Command("RANG %d", rangeCmd)
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

func (d *Driver) ResolutionAbsolute() (float64, error) {
	return 0.0, dmm.ErrNotImplemented
}
func (d *Driver) SetResolutionAbsolute(resolution float64) error {
	return dmm.ErrNotImplemented
}
func (d *Driver) TriggerDelay() (bool, float64, error) {
	return false, 0.0, dmm.ErrNotImplemented
}
func (d *Driver) SetTriggerDelay(autoDelay bool, delay float64) error {
	return dmm.ErrNotImplemented
}

func (d *Driver) TriggerSource() (dmm.TriggerSource, error) {
	return 0, dmm.ErrNotImplemented
}

func (d *Driver) SetTriggerSource(src dmm.TriggerSource) error {
	return dmm.ErrNotImplemented
}

func (d *Driver) Abort() error {
	return dmm.ErrNotImplemented
}

func (d *Driver) ConfigureMeasurement(
	msrFunc dmm.MeasurementFunction,
	autoRange dmm.AutoRange,
	rangeValue float64,
	resolution float64,
) error {
	return dmm.ErrNotImplemented
}

func (d *Driver) ConfigureTrigger(src dmm.TriggerSource, delay time.Duration) error {
	return dmm.ErrNotImplemented
}

func (d *Driver) FetchMeasurement(maxTime time.Duration) (float64, error) {
	return 0.0, dmm.ErrNotImplemented
}

func (d *Driver) InitiateMeasurement() error {
	return dmm.ErrNotImplemented
}

func (d *Driver) IsOverRange(value float64) bool {
	return true
}

func (d *Driver) ReadMeasurement(maxTime time.Duration) (float64, error) {
	return query.Float64(d.inst, "meas1?")
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
