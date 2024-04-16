// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key3446x implements the IVI Instrument driver for the Keysight 3446x
family of DMM.

The Keysight 3446x family of DMMs use LAN port 5025 for SCPI Telnet sessions
and port 5025 for SCPI Socket sessions (confirmed for the Keysight 34461A and
assumed for the others). The default GPIB address for the 34461A is 22 (per p.
475 of the manual).

State Caching: Not implemented
*/
package key3446x

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/convert"
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

// Driver provides the IVI driver for the Keysight 3446x family of DMMs.
type Driver struct {
	inst        ivi.Instrument
	outputCount int
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 3446x series of DMMs.
func New(inst ivi.Instrument, reset bool) (*Driver, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		GroupCapabilities: []string{
			"IviDmmBase",
			"IviDmmACMeasurement",
			"IviDmmFrequencyMeasurement",
			// "IviDmmTemperatureMeasurement",
			// "IviDmmResistanceTemperatureDevice",
			// "IviDmmThermistor",
			// "IviDmmMultiPoint",
			// "IviDmmTriggerSlope",
			// "IviDmmSoftwareTrigger",
			// "IviDmmDeviceInfo",
			// "IviDmmAutoRangeValue",
			// "IviDmmAutoZero",
			// "IviDmmPowerLineFrequency",
		},
		SupportedInstrumentModels: []string{
			"34460A",
			"34461A",
			"34465A",
			"34470A",
		},
		SupportedBusInterfaces: []string{
			"USB",
			"GPIB",
			"LAN",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
		inst:        inst,
		outputCount: 1,
		Inherent:    inherent,
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
	conf, err := d.QueryString("CONF?")
	if err != nil {
		return 0, err
	}

	conf = convert.StripDoubleQuotes(conf)
	fcnString := strings.Split(conf, " ")[0]
	fcn, ok := stringToMeasurementFunction[fcnString]

	if !ok {
		return 0, fmt.Errorf("%s is not a valid measurement function", fcnString)
	}

	return fcn, nil
}

// SetMeasurementFunction specifies the measurement function.
//
// SetMeasurementFunction is the setter for the read-write IviDmmBase Attribute
// Function described in Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetMeasurementFunction(msrFunc dmm.MeasurementFunction) error {
	cmd := msrFuncToConfigCommand[msrFunc]
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
	return 0, 0.0, dmm.ErrNotImplemented
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
	return dmm.ErrNotImplemented
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
	cmd, err := createConfigureMeasurementCommand(msrFunc, autoRange, rangeValue, resolution)
	if err != nil {
		return err
	}

	return d.inst.Command(cmd)
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
		return "", dmm.ErrNotImplemented
	case dmm.ACCurrent:
		// return createConfigureCurrentACCommand(autoRange, rangeValue)
		return "", dmm.ErrNotImplemented
	case dmm.TwoWireResistance:
		return "", dmm.ErrNotImplemented
	case dmm.FourWireResistance:
		return "", dmm.ErrNotImplemented
	case dmm.ACPlusDCVolts:
		return "", dmm.ErrNotImplemented
	case dmm.ACPlusDCCurrent:
		return "", dmm.ErrNotImplemented
	case dmm.Frequency:
		return "", dmm.ErrNotImplemented
	case dmm.Period:
		return "", dmm.ErrNotImplemented
	case dmm.Temperature:
		return "", dmm.ErrNotImplemented
	}

	return "", dmm.ErrNotImplemented
}

func createConfigureVoltageDCCommand(
	autoRange dmm.AutoRange,
	rangeValue float64,
	resolution float64,
) (string, error) {
	rng, err := determineVoltageRange(autoRange, rangeValue)
	if err != nil {
		return "", dmm.ErrNotImplemented
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
		return "", dmm.ErrNotImplemented
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
		return "", dmm.ErrNotImplemented
	}
	return "", dmm.ErrNotImplemented
}

func determineManualVoltageRange(rangeValue float64) (string, error) {
	switch {
	case rangeValue <= 0.1:
		return "100 mV", nil
	case rangeValue <= 1.0:
		return "1 V", nil
	case rangeValue <= 10.0:
		return "10 V", nil
	case rangeValue <= 100.0:
		return "100 V", nil
	case rangeValue <= 1000.0:
		return "1000 V", nil
	}
	return "", dmm.ErrNotImplemented
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
	return 0.0, dmm.ErrNotImplemented
}

// stringToMeasurementFunction maps the string name of a measurement function to the
// MeasurementFunction.
var stringToMeasurementFunction = map[string]dmm.MeasurementFunction{
	"VOLT":    dmm.DCVolts,
	"VOLT:DC": dmm.DCVolts,
	"VOLT:AC": dmm.ACVolts,
	"CURR":    dmm.DCCurrent,
	"CURR:DC": dmm.DCCurrent,
	"CURR:AC": dmm.ACCurrent,
	"RES":     dmm.TwoWireResistance,
	"FRES":    dmm.FourWireResistance,
	"FREQ":    dmm.Frequency,
	"TEMP":    dmm.Temperature,
}

// msrFuncToConfigCommand maps the MeasurementFunction to the SCPI
// CONFigure command.
var msrFuncToConfigCommand = map[dmm.MeasurementFunction]string{
	dmm.DCVolts:            "CONF:VOLT:DC",
	dmm.ACVolts:            "CONF:VOLT:AC",
	dmm.DCCurrent:          "CONF:CURR:DC",
	dmm.ACCurrent:          "CURR:AC",
	dmm.TwoWireResistance:  "CONF:RES",
	dmm.FourWireResistance: "CONF:FRES",
	dmm.Frequency:          "CONF:FREQ",
	dmm.Temperature:        "CONF:TEMP",
}
