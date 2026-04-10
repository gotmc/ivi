// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

type ApertureTimeUnits int

const (
	Seconds ApertureTimeUnits = iota
	PowerLineCycles
)

var apertureTimeUnits = map[ApertureTimeUnits]string{
	Seconds:         "seconds",
	PowerLineCycles: "power line cycles",
}

// String implements the Stringer interface for ApertureTimeUnits.
func (atu ApertureTimeUnits) String() string {
	return apertureTimeUnits[atu]
}

// AutoRange provides the defined values for the AutoRange function defined in
// Section 4.2.3 and Section 19 of IVI-4.2: IviDmm Class Specification.
type AutoRange int

// The AutoRange defined values are the available auto range settings.
const (
	AutoOff AutoRange = iota
	AutoOn
	AutoOnce
)

var autoRanges = map[AutoRange]string{
	AutoOff:  "Auto Range Off",
	AutoOn:   "Auto Range On",
	AutoOnce: "Auto Range Once",
}

// String implements the Stringer interface for AutoRange.
func (ar AutoRange) String() string {
	return autoRanges[ar]
}

type AutoZero int

const (
	AutoZeroOff AutoZero = iota
	AutoZeroOn
	AutoZeroOnce
)

var autoZeros = map[AutoZero]string{
	AutoZeroOff:  "auto zero off",
	AutoZeroOn:   "auto zero on",
	AutoZeroOnce: "auto zero once",
}

// String implements the Stringer interface for AutoZero.
func (az AutoZero) String() string {
	return autoZeros[az]
}

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

// String implements the Stringer interface for MeasurementFunction.
func (fcn MeasurementFunction) String() string {
	return measurementFunctions[fcn]
}

// TriggerSource provides the defined values for the Trigger Source Attribute
// defined in Section 4.2.7 of IVI-4.2 IviDmm Class Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	TriggerSourceImmediate TriggerSource = iota
	TriggerSourceExternal
	TriggerSourceSoftware
	TriggerSourceInterval
	TriggerSourceTTL0
	TriggerSourceTTL1
	TriggerSourceTTL2
	TriggerSourceTTL3
	TriggerSourceTTL4
	TriggerSourceTTL5
	TriggerSourceTTL6
	TriggerSourceTTL7
	TriggerSourceECL0
	TriggerSourceECL1
	TriggerSourcePXIStar
	TriggerSourceRTSI0
	TriggerSourceRTSI1
	TriggerSourceRTSI2
	TriggerSourceRTSI3
	TriggerSourceRTSI4
	TriggerSourceRTSI5
	TriggerSourceRTSI6
)

var triggerSources = map[TriggerSource]string{
	TriggerSourceImmediate: "immediate",
	TriggerSourceExternal:  "external",
	TriggerSourceSoftware:  "software trigger",
	TriggerSourceInterval:  "interval",
	TriggerSourceTTL0:      "ttl0",
	TriggerSourceTTL1:      "ttl1",
	TriggerSourceTTL2:      "ttl2",
	TriggerSourceTTL3:      "ttl3",
	TriggerSourceTTL4:      "ttl4",
	TriggerSourceTTL5:      "ttl5",
	TriggerSourceTTL6:      "ttl6",
	TriggerSourceTTL7:      "ttl7",
	TriggerSourceECL0:      "ecl0",
	TriggerSourceECL1:      "ecl1",
	TriggerSourcePXIStar:   "pxi star",
	TriggerSourceRTSI0:     "rtsi0",
	TriggerSourceRTSI1:     "rtsi1",
	TriggerSourceRTSI2:     "rtsi2",
	TriggerSourceRTSI3:     "rtsi3",
	TriggerSourceRTSI4:     "rtsi4",
	TriggerSourceRTSI5:     "rtsi5",
	TriggerSourceRTSI6:     "rtsi6",
}

// String implements the Stringer interface for TriggerSource.
func (ts TriggerSource) String() string {
	return triggerSources[ts]
}

type TempTransducerType int

const (
	Thermocouple TempTransducerType = iota
	Thermistor
	TwoWireRTD
	FourWireRTD
)

var tempTransducerTypes = map[TempTransducerType]string{
	Thermocouple: "thermocouple",
	Thermistor:   "thermistor",
	TwoWireRTD:   "2-wire RTD",
	FourWireRTD:  "4-wire RTD",
}

// String implements the Stringer interface for TempTransducerType.
func (ttt TempTransducerType) String() string {
	return tempTransducerTypes[ttt]
}

type ReferenceJunctionType int

const (
	InternalReferenceJunction ReferenceJunctionType = iota
	FixedReferenceJunction
)

var referenceJunctionTypes = map[ReferenceJunctionType]string{
	InternalReferenceJunction: "internal",
	FixedReferenceJunction:    "fixed",
}

// String implements the Stringer interface for ReferenceJunctionType.
func (rjt ReferenceJunctionType) String() string {
	return referenceJunctionTypes[rjt]
}

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

var thermocoupleTypes = map[ThermocoupleType]string{
	ThermocoupleB: "B",
	ThermocoupleC: "C",
	ThermocoupleD: "D",
	ThermocoupleE: "E",
	ThermocoupleG: "G",
	ThermocoupleJ: "J",
	ThermocoupleK: "K",
	ThermocoupleN: "N",
	ThermocoupleR: "R",
	ThermocoupleS: "S",
	ThermocoupleT: "T",
	ThermocoupleU: "U",
	ThermocoupleV: "V",
}

// String implements the Stringer interface for ThermocoupleType.
func (tt ThermocoupleType) String() string {
	return thermocoupleTypes[tt]
}

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

var measurementDestinations = map[MeasurementDestination]string{
	MsrDestinationNone:     "none",
	MsrDestinationExternal: "external",
	MsrDestinationTTL0:     "ttl0",
	MsrDestinationTTL1:     "ttl1",
	MsrDestinationTTL2:     "ttl2",
	MsrDestinationTTL3:     "ttl3",
	MsrDestinationTTL4:     "ttl4",
	MsrDestinationTTL5:     "ttl5",
	MsrDestinationTTL6:     "ttl6",
	MsrDestinationTTL7:     "ttl7",
	MsrDestinationECL0:     "ecl0",
	MsrDestinationECL1:     "ecl1",
	MsrDestinationPXIStar:  "pxi star",
	MsrDestinationRTSI0:    "rtsi0",
	MsrDestinationRTSI1:    "rtsi1",
	MsrDestinationRTSI2:    "rtsi2",
	MsrDestinationRTSI3:    "rtsi3",
	MsrDestinationRTSI4:    "rtsi4",
	MsrDestinationRTSI5:    "rtsi5",
	MsrDestinationRTSI6:    "rtsi6",
}

// String implements the Stringer interface for MeasurementDestination.
func (md MeasurementDestination) String() string {
	return measurementDestinations[md]
}

type TriggerSlope int

const (
	PositiveTriggerSlope TriggerSlope = iota
	NegativeTriggerSlope
)

var triggerSlopes = map[TriggerSlope]string{
	PositiveTriggerSlope: "positive",
	NegativeTriggerSlope: "negative",
}

// String implements the Stringer interface for TriggerSlope.
func (ts TriggerSlope) String() string {
	return triggerSlopes[ts]
}
