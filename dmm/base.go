// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

import (
	"errors"
	"time"
)

// Base provides the interface required for the IviDMMBase capability group
// described in Sectino 4 of IVI-4.2 IviDmm Class Specification.
type Base interface {
	MeasurementFunction() (MeasurementFunction, error)
	SetMeasurementFunction(msrFunc MeasurementFunction) error
	Range() (AutoRange, float64, error)
	SetRange(autoRange AutoRange, rangeValue float64) error
	ResolutionAbsolute() (float64, error)
	SetResolutionAbsolute(resolution float64) error
	TriggerDelay() (bool, float64, error)
	SetTriggerDelay(autoDelay bool, delay float64) error
	TriggerSource() (TriggerSource, error)
	SetTriggerSource(src TriggerSource) error
	Abort() error
	ConfigureMeasurement(
		msrFunc MeasurementFunction,
		autoRange AutoRange,
		rangeValue float64,
		resolution float64,
	) error
	ConfigureTrigger(src TriggerSource, delay time.Duration) error
	FetchMeasurement(maxTime time.Duration) (float64, error)
	InitiateMeasurement() error
	IsOverRange(value float64) bool
	ReadMeasurement(maxTime time.Duration) (float64, error)
}

// Error codes related to the IviDCPwr Class Specification.
var (
	ErrNotImplemented     = errors.New("not implemented in ivi driver")
	ErrOVPUnsupported     = errors.New("ovp not supported")
	ErrTriggerNotSoftware = errors.New("trigger source is not set to software trigger")
)

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
	AutoOn   AutoRange = iota // IVIDMM_VAL_AUTO_RANGE_ON / IviDmmAutoRangeOn
	AutoOff                   // IVIDMM_VAL_AUTO_RANGE_OFF / IviDmmAutoRangeOff
	AutoOnce                  // IVIDMM_VAL_AUTO_RANGE_ONCE / IviDammAutoRangeOnce
)

// TriggerSource provides the defined values for the Trigger Source Attribute
// defined in Sectino 4.2.7 of IVI-4.2 IviDmm Class Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	Immediate TriggerSource = iota
	External
	SoftwareTrigger
	TTL0
	TTL1
	TTL2
	TTL3
	TTL4
	TTL5
	TTL6
	TTL7
	ECL0
	ECL1
	PXIStar
	RTSI0
	RTSI1
	RTSI2
	RTSI3
	RTSI4
	RTSI5
	RTSI6
)
