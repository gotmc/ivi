// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

// AcquisitionType models the defined values for the acquisition type defined
// in Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope
// Class Specification.
type AcquisitionType int

// Available AcquisitionType values.
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
	EnvelopeAcquisition:       "average",
	AverageAcquisition:        "average",
}

// String implements the Stringer interface for AcquisitionType.
func (at AcquisitionType) String() string {
	return acquisitionTypes[at]
}

// VerticalCoupling models the defined values for a channel's vertical input
// coupling defined in Section 20 IviScope Attribute Value Definitions of
// IVI-4.1: IviScope Class Specification.
type VerticalCoupling int

// Available VerticalCoupling values.
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

// TriggerCoupling models the defined values for the trigger coupling defined
// in Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope
// Class Specification.
type TriggerCoupling int

// Available TriggerCoupling values.
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

// TriggerSlope models the defined values for the trigger slope defined in
// Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope Class
// Specification.
type TriggerSlope int

// Available TriggerSlope values.
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

// TriggerType models the defined values for the kinds of triggers defined in
// Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope Class
// Specification.
type TriggerType int

// Available TriggerType values.
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

// InterpolationMethod models the defined values for the available
// interpolation methods defined in Section 20 IviScope Attribute Value
// Definitions and used in the Interpolation function in Section 5.2.1 of
// IVI-4.1: IviScopeClass Specification.
type InterpolationMethod int

// Available InterpolationMethod values.
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

// Available TVTriggerEvent values.
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

// TVTriggerSignalFormat models the defined values for the supported TV signal
// formats defined in Section 20 IviScope Attribute Value Definitions of
// IVI-4.1: IviScope Class Specification.
type TVTriggerSignalFormat int

// Available TVTriggerSignalFormat values.
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

// TVTriggerPolarity models the defined values for the TV trigger polarity
// defined in Section 20 IviScope Attribute Value Definitions of IVI-4.1:
// IviScope Class Specification.
type TVTriggerPolarity int

// Available TVTriggerPolarity values.
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

// Polarity models the defined values for the runt-trigger polarity defined
// in Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope
// Class Specification.
type Polarity int

// Available Polarity values.
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

// GlitchCondition models the defined values for the glitch-trigger condition
// defined in Section 20 IviScope Attribute Value Definitions of IVI-4.1:
// IviScope Class Specification.
type GlitchCondition int

// Available GlitchCondition values.
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

// WidthCondition models the defined values for the width-trigger condition
// defined in Section 20 IviScope Attribute Value Definitions of IVI-4.1:
// IviScope Class Specification.
type WidthCondition int

// Available WidthCondition values.
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

// ACLineTriggerSlope models the defined values for the AC-line trigger slope
// defined in Section 20 IviScope Attribute Value Definitions of IVI-4.1:
// IviScope Class Specification.
type ACLineTriggerSlope int

// Available ACLineTriggerSlope values.
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

// SampleMode models the defined values for the acquisition sample mode
// defined in Section 20 IviScope Attribute Value Definitions of IVI-4.1:
// IviScope Class Specification.
type SampleMode int

// Available SampleMode values.
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

// TriggerModifier models the defined values for the trigger modifier defined
// in Section 20 IviScope Attribute Value Definitions of IVI-4.1: IviScope
// Class Specification.
type TriggerModifier int

// Available TriggerModifier values.
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

// MaximumTime models the special maximum-time values accepted by fetch/read
// methods that take a timeout: Zero for "do not wait" and MaxTimeValue for
// "wait indefinitely."
type MaximumTime int

// Available MaximumTime values.
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

// WaveformMeasurement models the defined values for the supported waveform
// measurements defined in Section 20 IviScope Attribute Value Definitions of
// IVI-4.1: IviScope Class Specification and used by the
// IviScopeWaveformMeasurement extension group.
type WaveformMeasurement int

// Available WaveformMeasurement values.
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
