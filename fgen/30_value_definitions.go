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
	OutputFunction OutputMode = iota
	OutputArbitrary
	OutputSequence
)

func (om OutputMode) String() string {
	switch om {
	case OutputFunction:
		return "function"
	case OutputArbitrary:
		return "arbitrary"
	case OutputSequence:
		return "sequence"
	default:
		return ""
	}
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

func (om OperationMode) String() string {
	switch om {
	case ContinuousMode:
		return "continuous mode"
	case BurstMode:
		return "burst mode"
	default:
		return ""
	}
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
)

func (ts OldTriggerSource) String() string {
	switch ts {
	case OldTriggerSourceInternal:
		return "internal trigger"
	case OldTriggerSourceExternal:
		return "external trigger"
	case OldTriggerSourceSoftware:
		return "software trigger"
	}

	return ""
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

type SampleClock int

const (
	SampleClockInternal SampleClock = iota
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

type TerminalConfiguration int

const (
	TerminalConfigurationSingleEnded TerminalConfiguration = iota
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
