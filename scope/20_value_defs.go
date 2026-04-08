// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
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

var acquisitionTypes = map[AcquisitionType]string{
	NormalAcquisition:         "normal",
	PeakDetectAcquisition:     "peak detect",
	HighResolutionAcquisition: "high resolution",
	EnvelopeAcquisition:       "envelope",
	AverageAcquisition:        "average",
}

func (at AcquisitionType) String() string {
	return acquisitionTypes[at]
}

type VerticalCoupling int

const (
	ACVerticalCoupling VerticalCoupling = iota
	DCVerticalCoupling
	GndVerticalCoupling
)

var verticalCouplings = map[VerticalCoupling]string{
	ACVerticalCoupling:  "AC",
	DCVerticalCoupling:  "DC",
	GndVerticalCoupling: "GND",
}

// String implements the Stringer interface for VerticalCoupling.
func (vc VerticalCoupling) String() string {
	return verticalCouplings[vc]
}

type TriggerCoupling int

const (
	ACTriggerCoupling TriggerCoupling = iota
	DCTriggerCoupling
	HFRejectTriggerCoupling
	LFRejectTriggerCoupling
	NoiseRejectTriggerCoupling
)

var triggerCouplings = map[TriggerCoupling]string{
	ACTriggerCoupling:          "AC",
	DCTriggerCoupling:          "DC",
	HFRejectTriggerCoupling:    "HF reject",
	LFRejectTriggerCoupling:    "LF reject",
	NoiseRejectTriggerCoupling: "noise reject",
}

// String implements the Stringer interface for TriggerCoupling.
func (tc TriggerCoupling) String() string {
	return triggerCouplings[tc]
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

// TriggerSource models the defined values for the Trigger Source defined in
// Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope Class
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

var triggerSources = map[TriggerSource]string{
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

// String implements the Stringer interface for TriggerSource.
func (ts TriggerSource) String() string {
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

var triggerTypes = map[TriggerType]string{
	EdgeTrigger:      "Edge trigger",
	WidthTrigger:     "Width trigger",
	RuntTrigger:      "Runt trigger",
	GlitchTrigger:    "Glitch trigger",
	TVTrigger:        "TV trigger",
	ImmediateTrigger: "Immediate trigger",
	ACLineTrigger:    "A/C line trigger",
}

// String implements the Stringer interface for TriggerType.
func (tt TriggerType) String() string {
	return triggerTypes[tt]
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

var interpolationMethods = map[InterpolationMethod]string{
	NoInterpolation:         "none",
	SineXOverXInterpolation: "sinc",
	LinearInterpolation:     "linear",
}

// String implements the Stringer interface for InterpolationMethod.
func (im InterpolationMethod) String() string {
	return interpolationMethods[im]
}

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

var tvTriggerEvents = map[TVTriggerEvent]string{
	TVTriggerEventField1:     "field 1",
	TVTriggerEventField2:     "field 2",
	TVTriggerEventAnyField:   "any field",
	TVTriggerEventAnyLine:    "any line",
	TVTriggerEventLineNumber: "line number",
}

// String implements the Stringer interface for TVTriggerEvent.
func (te TVTriggerEvent) String() string {
	return tvTriggerEvents[te]
}

type TVTriggerSignalFormat int

const (
	TVSignalFormatNTSC TVTriggerSignalFormat = iota
	TVSignalFormatPAL
	TVSignalFormatSECAM
)

var tvTriggerSignalFormats = map[TVTriggerSignalFormat]string{
	TVSignalFormatNTSC:  "NTSC",
	TVSignalFormatPAL:   "PAL",
	TVSignalFormatSECAM: "SECAM",
}

// String implements the Stringer interface for TVTriggerSignalFormat.
func (sf TVTriggerSignalFormat) String() string {
	return tvTriggerSignalFormats[sf]
}

type TVTriggerPolarity int

const (
	TVTriggerPositive TVTriggerPolarity = iota
	TVTriggerNegative
)

var tvTriggerPolarities = map[TVTriggerPolarity]string{
	TVTriggerPositive: "positive",
	TVTriggerNegative: "negative",
}

// String implements the Stringer interface for TVTriggerPolarity.
func (tp TVTriggerPolarity) String() string {
	return tvTriggerPolarities[tp]
}

type Polarity int

const (
	PositivePolarity Polarity = iota
	NegativePolarity
	EitherPolarity
)

var polarities = map[Polarity]string{
	PositivePolarity: "positive",
	NegativePolarity: "negative",
	EitherPolarity:   "either",
}

// String implements the Stringer interface for Polarity.
func (p Polarity) String() string {
	return polarities[p]
}

type GlitchCondition int

const (
	GlitchLessThan GlitchCondition = iota
	GlitchGreaterThan
)

var glitchConditions = map[GlitchCondition]string{
	GlitchLessThan:    "less than",
	GlitchGreaterThan: "greater than",
}

// String implements the Stringer interface for GlitchCondition.
func (gc GlitchCondition) String() string {
	return glitchConditions[gc]
}

type WidthCondition int

const (
	WidthWithin WidthCondition = iota
	WidthOutside
)

var widthConditions = map[WidthCondition]string{
	WidthWithin:  "within",
	WidthOutside: "outside",
}

// String implements the Stringer interface for WidthCondition.
func (wc WidthCondition) String() string {
	return widthConditions[wc]
}

type ACLineTriggerSlope int

const (
	ACLinePositive ACLineTriggerSlope = iota
	ACLineNegative
	ACLineEither
)

var acLineTriggerSlopes = map[ACLineTriggerSlope]string{
	ACLinePositive: "positive",
	ACLineNegative: "negative",
	ACLineEither:   "either",
}

// String implements the Stringer interface for ACLineTriggerSlope.
func (s ACLineTriggerSlope) String() string {
	return acLineTriggerSlopes[s]
}

type SampleMode int

const (
	RealTimeSampleMode SampleMode = iota
	EquivalentTimeSampleMode
)

var sampleModes = map[SampleMode]string{
	RealTimeSampleMode:       "real time",
	EquivalentTimeSampleMode: "equivalent time",
}

// String implements the Stringer interface for SampleMode.
func (sm SampleMode) String() string {
	return sampleModes[sm]
}

type TriggerModifier int

const (
	TriggerModifierNone TriggerModifier = iota
	TriggerModifierAuto
	TriggerModifierAutoLevel
)

var triggerModifiers = map[TriggerModifier]string{
	TriggerModifierNone:      "none",
	TriggerModifierAuto:      "auto",
	TriggerModifierAutoLevel: "auto level",
}

// String implements the Stringer interface for TriggerModifier.
func (tm TriggerModifier) String() string {
	return triggerModifiers[tm]
}

type MaximumTime int

const (
	Zero MaximumTime = iota
	MaxTimeValue
)

var maximumTimes = map[MaximumTime]string{
	Zero:         "zero",
	MaxTimeValue: "max time value",
}

// String implements the Stringer interface for MaximumTime.
func (mt MaximumTime) String() string {
	return maximumTimes[mt]
}

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

var waveformMeasurements = map[WaveformMeasurement]string{
	RiseTime:            "rise time",
	FallTime:            "fall time",
	Frequency:           "frequency",
	Period:              "period",
	VoltageRMS:          "voltage RMS",
	VoltageCycleRMS:     "voltage cycle RMS",
	VoltageMax:          "voltage max",
	VoltageMin:          "voltage min",
	VoltagePeakToPeak:   "voltage peak-to-peak",
	VoltageHigh:         "voltage high",
	VoltageLow:          "voltage low",
	VoltageAverage:      "voltage average",
	VoltageCycleAverage: "voltage cycle average",
	WidthNegative:       "width negative",
	WidthPositive:       "width positive",
	DutyCycleNegative:   "duty cycle negative",
	DutyCyclePositive:   "duty cycle positive",
	Amplitude:           "amplitude",
	Overshoot:           "overshoot",
	Preshoot:            "preshoot",
}

// String implements the Stringer interface for WaveformMeasurement.
func (wm WaveformMeasurement) String() string {
	return waveformMeasurements[wm]
}
