// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

type AcquisitionType int

const (
	NormalAcquisition AcquisitionType = iota
	PeakDetectAcquisition
	HighResolutionAcquisition
	EnvelopeAcquisition
	AverageAcquisition
)

func (at AcquisitionType) String() string {
	return map[AcquisitionType]string{
		NormalAcquisition:         "normal",
		PeakDetectAcquisition:     "peak detect",
		HighResolutionAcquisition: "high resolution",
		EnvelopeAcquisition:       "evnvelope",
		AverageAcquisition:        "average",
	}[at]
}

type VerticalCoupling int

const (
	ACVerticalCoupling VerticalCoupling = iota
	DCVerticalCoupling
	GndVerticalCoupling
)

type TriggerCoupling int

const (
	ACTriggerCoupling TriggerCoupling = iota
	DCTriggerCoupling
	HFRejectTriggerCoupling
	LFRejectTriggerCoupling
	NoiseRejectTriggerCoupling
)

type TriggerSlope int

const (
	PositiveTriggerSlope TriggerSlope = iota
	NegativeTriggerSlope
)

// TriggerSource models the defined values for the Start Trigger Source, Stop
// Trigger Source, Hold Trigger Source, Resume Trigger Source, Advanced Trigger
// Source, Data Marker Destination, and Sparse Marker Destination defined in
// Section 30 IviFgen Attribute Value Definitions of IVI-4.3: IviFgenClass
// Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	TriggerSourceExternal TriggerSource = iota
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
		TriggerSourceExternal: "external",
		TriggerSourceTTL0:     "ttl0",
		TriggerSourceTTL1:     "ttl1",
		TriggerSourceTTL2:     "ttl2",
		TriggerSourceTTL3:     "ttl3",
		TriggerSourceTTL4:     "ttl4",
		TriggerSourceTTL5:     "ttl5",
		TriggerSourceTTL6:     "ttl6",
		TriggerSourceTTL7:     "ttl7",
		TriggerSourceECL0:     "ecl0",
		TriggerSourceECL1:     "ecl1",
		TriggerSourceRTSI0:    "rtsi0",
		TriggerSourceRTSI1:    "rtsi1",
		TriggerSourceRTSI2:    "rtsi2",
		TriggerSourceRTSI3:    "rtsi3",
		TriggerSourceRTSI4:    "rtsi4",
		TriggerSourceRTSI5:    "rtsi5",
		TriggerSourceRTSI6:    "rtsi6",
	}

	return triggerSources[ts]
}

type TriggerType int

const (
	EdgeTrigger TriggerType = iota
	WidthTrigger
	RuntTrigger
	GlitchTrigger
	TVTrigger
	ImmediateTrigger
	ACLineTrigger
)

// String implements the Stringer interface for TriggerType.
func (tt TriggerType) String() string {
	return map[TriggerType]string{
		EdgeTrigger:      "Edge trigger",
		WidthTrigger:     "Width trigger",
		RuntTrigger:      "Runt trigger",
		GlitchTrigger:    "Glitch trigger",
		TVTrigger:        "TV trigger",
		ImmediateTrigger: "Immediate trigger",
		ACLineTrigger:    "A/C line trigger",
	}[tt]
}

// InterpolationType models the defined values for the available interpolation
// methods defined in Section 20 IviScope Attribute Value Definitions and used
// in the Interpolation function in Section 5.2.1 of IVI-4.1: IviScopeClass
// Specification.
type InterpolationMethod int

const (
	NoInterpolation InterpolationMethod = iota
	SineXOverXInterpolation
	LinearInterpolation
)

// TVTriggerEvent models the defined values for the available events on which
// the oscilloscope triggers defined in Section 20 IviScope Attribute Value
// Definitions and used in the TV Trigger Event attribute in Section 6.2.1 of
// IVI-4.1: IviScopeClass Specification.
type TVTriggerEvent int

const (
	TVTriggerEventField1 TVTriggerEvent = iota
	TVTriggerEventField2
	TVTriggerEventAnyField
	TVTriggerEventAnyLine
	TVTriggerEventLineNumber
)

type TVTriggerSignalFormat int

const (
	TVSignalFormatNTSC TVTriggerSignalFormat = iota
	TVSignalFormatPAL
	TVSignalFormatSECAM
)

type TVTriggerPolarity int

const (
	TVTriggerPositive TVTriggerPolarity = iota
	TVTriggerNegative
)

type Polarity int

const (
	PositivePolarity Polarity = iota
	NegativePolarity
	EitherPolarity
)

type GlitchCondition int

const (
	GlitchLessThan GlitchCondition = iota
	GlitchGreaterThan
)

type WidthCondition int

const (
	WidthWithin WidthCondition = iota
	WidthOutside
)

type ACLineTriggerSlope int

const (
	ACLinePositive ACLineTriggerSlope = iota
	ACLineNegative
	ACLineEither
)

type SampleMode int

const (
	RealTimeSampleMode SampleMode = iota
	EquivalentTimeSampleMode
)

type TriggerModifier int

const (
	TriggerModifierNone TriggerModifier = iota
	TriggerModifierAuto
	TriggerModifierAutoLevel
)

type MaximumTime int

const (
	Zero MaximumTime = iota
	MaxTimeValue
)

type WaveformMeasurement int

const (
	RiseTime WaveformMeasurement = iota
	FallTime
	Frequency
	Period
	VoltageRMS
	VoltageCycleRMS
	VoltageMax
	VoltageMin
	VoltagePeakToPeak
	VoltageHigh
	VoltageLow
	VoltageAverage
	VoltageCycleAverage
	WidthNegative
	WidthPositive
	DutyCycleNegative
	DutyCyclePositive
	Amplitude
	Overshoot
	Preshoot
)
