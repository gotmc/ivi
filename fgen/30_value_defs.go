// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// OutputMode determines how the function generator produces waveforms. This
// attribute determines which extension group's functions and attributes are
// used to configure the waveform the function generator produces. OutputMode
// implements the defined values for the Output Mode read-write attribute
// defined in Section 4.2.5 and Section 30 of IVI-4.3: IviFgen Class
// Specification.
type OutputMode int

// Function indicates the IVI driver uses the attributes and functions of the
// IviFgenStdFunc extension group. Arbitrary indicates the IVI driver uses the
// attributes and functions of the IviFgenArbWfm extension group.
const (
	OutputModeFunction OutputMode = iota
	OutputModeArbitrary
	OutputModeSequence
	OutputModeNoise
)

// String implements the Stringer interface for OutputMode.
func (om OutputMode) String() string {
	outputModes := map[OutputMode]string{
		OutputModeFunction:  "function",
		OutputModeArbitrary: "arbitrary",
		OutputModeSequence:  "sequence",
		OutputModeNoise:     "noise",
	}

	return outputModes[om]
}

// ClockSource models the defined values for the Reference Clock Source defined
// in Section 4.2.7 and Section 30 of IVI-4.3 IviFgenClass Specification.
type ClockSource int

// Available reference clock source.
const (
	RefClockInternal ClockSource = iota
	RefClockExternal
	RefClockRTSIClock
)

// OperationMode provides the defined values for the Operation Mode defined in
// Section 4.2.2 and Section 30 of IVI-4.3: IviFgen Class Specification.
type OperationMode int

// ContinuousMode and BurstMode are the available Operation Modes. "A burst
// consists of a discrete number of waveform cycles. the user uses the
// attribute of the IviFgenTrigger Extension Group to configure the trigger,
// and the attributes of the IviFgenBurst extension group to configure how the
// function generator produces bursts.
const (
	ContinuousMode OperationMode = iota
	BurstMode
)

var operationModes = map[OperationMode]string{
	ContinuousMode: "continuous mode",
	BurstMode:      "burst mode",
}

// String implements the Stringer interface for OperationMode.
func (om OperationMode) String() string {
	return operationModes[om]
}

// StandardWaveform models the defined values for the Standard Waveform defined
// in Section 5.2.6 and Section 30 of IVI-4.3: IviFgen Class Specification.
type StandardWaveform int

// These are the available standard waveforms.
const (
	Sine StandardWaveform = iota
	Square
	Triangle
	RampUp
	RampDown
	DC
)

var standardWaveforms = map[StandardWaveform]string{
	Sine:     "Sine",
	Square:   "Square",
	Triangle: "Triangle",
	RampUp:   "Ramp Up",
	RampDown: "Ramp Down",
	DC:       "DC",
}

func (wave StandardWaveform) String() string {
	return standardWaveforms[wave]
}

// OldTriggerSource models the defined values for the Trigger Source defined in
// Section 9.2.1 and Section 30 of IVI-4.3: IviFgenClass Specification.
type OldTriggerSource int

// These are the available trigger sources. These have been deprecated, which
// is why they are referred to as OldTriggerSource.
const (
	OldTriggerSourceInternal OldTriggerSource = iota
	OldTriggerSourceExternal
	OldTriggerSourceSoftware
	OldTriggerSourceTTL0
	OldTriggerSourceTTL1
	OldTriggerSourceTTL2
	OldTriggerSourceTTL3
	OldTriggerSourceTTL4
	OldTriggerSourceTTL5
	OldTriggerSourceTTL6
	OldTriggerSourceTTL7
	OldTriggerSourceECL0
	OldTriggerSourceECL1
	OldTriggerSourcePXIStar
	OldTriggerSourceRTSI0
	OldTriggerSourceRTSI1
	OldTriggerSourceRTSI2
	OldTriggerSourceRTSI3
	OldTriggerSourceRTSI4
	OldTriggerSourceRTSI5
	OldTriggerSourceRTSI6
)

func (ts OldTriggerSource) String() string {
	oldTriggerSources := map[OldTriggerSource]string{
		OldTriggerSourceInternal: "internal trigger",
		OldTriggerSourceExternal: "external trigger",
		OldTriggerSourceSoftware: "software trigger",
		OldTriggerSourceTTL0:     "TTL0 trigger",
		OldTriggerSourceTTL1:     "TTL1 trigger",
		OldTriggerSourceTTL2:     "TTL2 trigger",
		OldTriggerSourceTTL3:     "TTL3 trigger",
		OldTriggerSourceTTL4:     "TTL4 trigger",
		OldTriggerSourceTTL5:     "TTL5 trigger",
		OldTriggerSourceTTL6:     "TTL6 trigger",
		OldTriggerSourceTTL7:     "TTL7 trigger",
		OldTriggerSourceECL0:     "ECL0 trigger",
		OldTriggerSourceECL1:     "ECL1 trigger",
		OldTriggerSourcePXIStar:  "PXI star trigger",
		OldTriggerSourceRTSI0:    "RTSI0 trigger",
		OldTriggerSourceRTSI1:    "RTSI1 trigger",
		OldTriggerSourceRTSI2:    "RTSI2 trigger",
		OldTriggerSourceRTSI3:    "RTSI3 trigger",
		OldTriggerSourceRTSI4:    "RTSI4 trigger",
		OldTriggerSourceRTSI5:    "RTSI5 trigger",
		OldTriggerSourceRTSI6:    "RTSI6 trigger",
	}

	return oldTriggerSources[ts]
}

// TriggerSource models the defined values for the Start Trigger Source, Stop
// Trigger Source, Hold Trigger Source, Resume Trigger Source, Advanced Trigger
// Source, Data Marker Destination, and Sparse Marker Destination defined in
// Section 30 IviFgen Attribute Value Definitions of IVI-4.3: IviFgenClass
// Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	TriggerSourceNone TriggerSource = iota
	TriggerSourceImmediate
	TriggerSourceExternal
	TriggerSourceInternal
	TriggerSourceSoftware
	TriggerSourceLAN0
	TriggerSourceLAN1
	TriggerSourceLAN2
	TriggerSourceLAN3
	TriggerSourceLAN4
	TriggerSourceLAN5
	TriggerSourceLAN6
	TriggerSourceLAN7
	TriggerSourceLXI0
	TriggerSourceLXI1
	TriggerSourceLXI2
	TriggerSourceLXI3
	TriggerSourceLXI4
	TriggerSourceLXI5
	TriggerSourceLXI6
	TriggerSourceLXI7
	TriggerSourceTTL0
	TriggerSourceTTL1
	TriggerSourceTTL2
	TriggerSourceTTL3
	TriggerSourceTTL4
	TriggerSourceTTL5
	TriggerSourceTTL6
	TriggerSourceTTL7
	TriggerSourcePXIStar
	TriggerSourcePXITrig0
	TriggerSourcePXITrig1
	TriggerSourcePXITrig2
	TriggerSourcePXITrig3
	TriggerSourcePXITrig4
	TriggerSourcePXITrig5
	TriggerSourcePXITrig6
	TriggerSourcePXITrig7
	TriggerSourcePXIeDStarA
	TriggerSourcePXIeDStarB
	TriggerSourcePXIeDStarC
	TriggerSourceRTSI0
	TriggerSourceRTSI1
	TriggerSourceRTSI2
	TriggerSourceRTSI3
	TriggerSourceRTSI4
	TriggerSourceRTSI5
	TriggerSourceRTSI6
)

// String implements the Stringer interface for TriggerSource.
func (ts TriggerSource) String() string {
	triggerSources := map[TriggerSource]string{
		TriggerSourceNone:       "none",
		TriggerSourceImmediate:  "immediate",
		TriggerSourceExternal:   "external",
		TriggerSourceInternal:   "internal",
		TriggerSourceSoftware:   "software",
		TriggerSourceLAN0:       "lan0",
		TriggerSourceLAN1:       "lan1",
		TriggerSourceLAN2:       "lan2",
		TriggerSourceLAN3:       "lan3",
		TriggerSourceLAN4:       "lan4",
		TriggerSourceLAN5:       "lan5",
		TriggerSourceLAN6:       "lan6",
		TriggerSourceLAN7:       "lan7",
		TriggerSourceLXI0:       "lxi0",
		TriggerSourceLXI1:       "lxi1",
		TriggerSourceLXI2:       "lxi2",
		TriggerSourceLXI3:       "lxi3",
		TriggerSourceLXI4:       "lxi4",
		TriggerSourceLXI5:       "lxi5",
		TriggerSourceLXI6:       "lxi6",
		TriggerSourceLXI7:       "lxi7",
		TriggerSourceTTL0:       "ttl0",
		TriggerSourceTTL1:       "ttl1",
		TriggerSourceTTL2:       "ttl2",
		TriggerSourceTTL3:       "ttl3",
		TriggerSourceTTL4:       "ttl4",
		TriggerSourceTTL5:       "ttl5",
		TriggerSourceTTL6:       "ttl6",
		TriggerSourceTTL7:       "ttl7",
		TriggerSourcePXIStar:    "pxi star",
		TriggerSourcePXITrig0:   "pxi trigger 0",
		TriggerSourcePXITrig1:   "pxi trigger 1",
		TriggerSourcePXITrig2:   "pxi trigger 2",
		TriggerSourcePXITrig3:   "pxi trigger 3",
		TriggerSourcePXITrig4:   "pxi trigger 4",
		TriggerSourcePXITrig5:   "pxi trigger 5",
		TriggerSourcePXITrig6:   "pxi trigger 6",
		TriggerSourcePXITrig7:   "pxi trigger 7",
		TriggerSourcePXIeDStarA: "pxied star a",
		TriggerSourcePXIeDStarB: "pxied star b",
		TriggerSourcePXIeDStarC: "pxied stard c",
		TriggerSourceRTSI0:      "rtsi0",
		TriggerSourceRTSI1:      "rtsi1",
		TriggerSourceRTSI2:      "rtsi2",
		TriggerSourceRTSI3:      "rtsi3",
		TriggerSourceRTSI4:      "rtsi4",
		TriggerSourceRTSI5:      "rtsi5",
		TriggerSourceRTSI6:      "rtsi6",
	}

	return triggerSources[ts]
}

type SampleClockSource int

const (
	SampleClockInternal SampleClockSource = iota
	SampleClockExternal
)

type MarkerPolarity int

const (
	MarkerActiveHigh MarkerPolarity = iota
	MarkerActiveLow
)

type AMSource int

const (
	AMSourceInternal AMSource = iota
	AMSourceExternal
)

// FIXME: I'm going to try to use the StandardWaveform instead, since that's
// what the standard calls for. However, the AM Modulation doesn't allow a DC
// standard waveform, whereas the StdFunc does.
type AMWaveform int

const (
	AMInternalSine AMWaveform = iota
	AMInternalSquare
	AMInternalTriangle
	AMInternalRampUp
	AMInternalRampDown
)

type FMSource int

const (
	FMSourceInternal FMSource = iota
	FMSourceExternal
)

type BinaryAlignment int

const (
	BinaryAlignmentLeft BinaryAlignment = iota
	BinaryAlignmentRight
)

type TerminalConfigurationType int

const (
	TerminalConfigurationSingleEnded TerminalConfigurationType = iota
	TerminalConfigurationDifferential
)

// TriggerSlope models the defined values for the Trigger Slope defined in
// Section 10.2.2 and Section 30 of IVI-4.3: IviFgenClass Specification.
type TriggerSlope int

const (
	TriggerSlopePositive TriggerSlope = iota
	TriggerSlopeNegative
	TriggerSlopeEither
)

// String implements the Stringer interface for TriggerSlope.
func (ts TriggerSlope) String() string {
	triggerSlopes := map[TriggerSlope]string{
		TriggerSlopePositive: "positive",
		TriggerSlopeNegative: "negative",
		TriggerSlopeEither:   "either",
	}

	return triggerSlopes[ts]
}
