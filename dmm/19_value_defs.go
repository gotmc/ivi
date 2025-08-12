// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

type ApertureTimeUnits int

const (
	Seconds ApertureTimeUnits = iota
	PowerLineCycles
)

// AutoRange provides the defined values for the AutoRange function defined in
// Section 4.2.3 and Section 19 of IVI-4.2: IviDmm Class Specification.
type AutoRange int

// The AutoRange defined values are the available auto range settings.
const (
	AutoOff AutoRange = iota
	AutoOn
	AutoOnce
)

// String implements the Stringer interface for AutoRange.
func (ar AutoRange) String() string {
	return map[AutoRange]string{
		AutoOff:  "Auto Range Off",
		AutoOn:   "Auto Range On",
		AutoOnce: "Auto Range Once",
	}[ar]
}

type AutoZero int

const (
	AutoZeroOff AutoZero = iota
	AutoZeroOn
	AutoZeroOnce
)

// MeasurementFunction provides the defined values for the Measurement Function
// defined in Section 4.2.1 and Section 19 of IVI-4.2: IviDmm Class
// Specification.
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

// String implements the Stringer interface for MeasurementFunction.
func (fcn MeasurementFunction) String() string {
	return map[MeasurementFunction]string{
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
	}[fcn]
}

// TriggerSource provides the defined values for the Trigger Source Attribute
// defined in Section 4.2.7 of IVI-4.2 IviDmm Class Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	Immediate TriggerSource = iota
	External
	SoftwareTrigger
	Interval
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

type TempTransducerType int

const (
	Thermocouple TempTransducerType = iota
	Thermistor
	TwoWireRTD
	FourWireRTD
)

type ReferenceJunctionType int

const (
	InternalReferenceJunction ReferenceJunctionType = iota
	FixedReferenceJunction
)

type ThermocoupleType int

const (
	ThermocoupleB ThermocoupleType = iota
	ThermocoupleC
	ThermocoupleD
	ThermocoupleE
	ThermocoupleG
	ThermocoupleJ
	ThermocoupleK
	ThermocoupleN
	ThermocoupleR
	ThermocoupleS
	ThermocoupleT
	ThermocoupleU
	ThermocoupleV
)

type MeasurementDestination int

const (
	MsrDestinationNone MeasurementDestination = iota
	MsrDestinationExternal
	MsrDestinationTTL0
	MsrDestinationTTL1
	MsrDestinationTTL2
	MsrDestinationTTL3
	MsrDestinationTTL4
	MsrDestinationTTL5
	MsrDestinationTTL6
	MsrDestinationTTL7
	MsrDestinationECL0
	MsrDestinationECL1
	MsrDestinationPXIStar
	MsrDestinationRTSI0
	MsrDestinationRTSI1
	MsrDestinationRTSI2
	MsrDestinationRTSI3
	MsrDestinationRTSI4
	MsrDestinationRTSI5
	MsrDestinationRTSI6
)

type TriggerSlope int

const (
	PositiveTriggerSlope TriggerSlope = iota
	NegativeTriggerSlope
)
