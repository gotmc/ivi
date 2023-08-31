// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

// Base provides the interface required for the IviDMMBase capability group.
type Base interface {
	MeaurementFunction() (MeasurementFunction, error)
	SetMeasurementFunction(MeasurementFunction) error
	Range() (float64, error)
	SetRange(float64) error
	AutoRange() (AutoRange, error)
	SetAutoRange(AutoRange) error
	ResolutionAbsolute() (float64, error)
	SetResolutionAbsoluate(float64) error
	TriggerDelay() (float64, error)
	SetTriggerDelay(float64) error
}

// ACMeasurement provides the interface required for the IviDMMACMeasurement
// capability group.
type ACMeasurement interface {
}

// MeasurementFunction provides the defined values for the Measurement Function defined in
// Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
type MeasurementFunction int

// The MeasurementFunction defined values are the available measurement functions.
const (
	DCVolts MeasurementFunction = iota
	ACVolts
	DCCurrent
	ACCurrent
	TwoWireResistance
	FourWireResistance
	ACPlusDCVolts
	ACPlusDCCurrent
	Frequency
	Period
	Temperature
)

var measurementFunctions = map[MeasurementFunction]string{
	DCVolts:            "DC Volts",
	ACVolts:            "AC Volts",
	DCCurrent:          "DC Current",
	ACCurrent:          "AC Current",
	TwoWireResistance:  "2-wire Resistance",
	FourWireResistance: "4-wire Resistance",
	ACPlusDCVolts:      "AC Plus DC Volts",
	ACPlusDCCurrent:    "AC Plus DC Current",
	Frequency:          "Frequency",
	Period:             "Period",
	Temperature:        "Temperature",
}

func (f MeasurementFunction) String() string {
	return measurementFunctions[f]
}

// AutoRange provides the defined values for the AutoRange function defined in
// Section 4.2.3 of IVI-4.2: IviDmm Class Specification.
type AutoRange int

// The AutoRange defined values are the available auto range settings.
const (
	AutoOn AutoRange = iota
	AutoOff
	AutoOnce
)
