// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package specan

// AmplitudeUnits specifies the amplitude units for input, output, and display.
type AmplitudeUnits int

// Available AmplitudeUnits values.
const (
	AmplitudeUnitsDBM AmplitudeUnits = iota
	AmplitudeUnitsDBMV
	AmplitudeUnitsDBUV
	AmplitudeUnitsVolt
	AmplitudeUnitsWatt
)

var amplitudeUnitsDesc = map[AmplitudeUnits]string{
	AmplitudeUnitsDBM:  "dBm",
	AmplitudeUnitsDBMV: "dBmV",
	AmplitudeUnitsDBUV: "dBuV",
	AmplitudeUnitsVolt: "Volt",
	AmplitudeUnitsWatt: "Watt",
}

func (au AmplitudeUnits) String() string { return amplitudeUnitsDesc[au] }

// DetectorType specifies the detection method used to capture and process
// signal data during each trace point.
type DetectorType int

// Available DetectorType values.
const (
	DetectorTypeAutoPeak DetectorType = iota
	DetectorTypeAverage
	DetectorTypeMaxPeak
	DetectorTypeMinPeak
	DetectorTypeSample
	DetectorTypeRMS
)

var detectorTypeDesc = map[DetectorType]string{
	DetectorTypeAutoPeak: "Auto Peak",
	DetectorTypeAverage:  "Average",
	DetectorTypeMaxPeak:  "Maximum Peak",
	DetectorTypeMinPeak:  "Minimum Peak",
	DetectorTypeSample:   "Sample",
	DetectorTypeRMS:      "RMS",
}

func (dt DetectorType) String() string { return detectorTypeDesc[dt] }

// TraceType specifies how the spectrum analyzer combines the current sweep
// data with the existing trace data.
type TraceType int

// Available TraceType values.
const (
	TraceTypeClearWrite TraceType = iota
	TraceTypeMaxHold
	TraceTypeMinHold
	TraceTypeVideoAverage
	TraceTypeView
	TraceTypeStore
)

var traceTypeDesc = map[TraceType]string{
	TraceTypeClearWrite:   "Clear Write",
	TraceTypeMaxHold:      "Maximum Hold",
	TraceTypeMinHold:      "Minimum Hold",
	TraceTypeVideoAverage: "Video Average",
	TraceTypeView:         "View",
	TraceTypeStore:        "Store",
}

func (tt TraceType) String() string { return traceTypeDesc[tt] }

// VerticalScale specifies the vertical scale of the measurement display.
type VerticalScale int

// Available VerticalScale values.
const (
	VerticalScaleLinear VerticalScale = iota
	VerticalScaleLogarithmic
)

var verticalScaleDesc = map[VerticalScale]string{
	VerticalScaleLinear:      "Linear",
	VerticalScaleLogarithmic: "Logarithmic",
}

func (vs VerticalScale) String() string { return verticalScaleDesc[vs] }

// TriggerSource specifies the source of the trigger.
type TriggerSource int

// Available TriggerSource values.
const (
	TriggerSourceExternal TriggerSource = iota
	TriggerSourceImmediate
	TriggerSourceSoftware
	TriggerSourceACLine
	TriggerSourceVideo
)

var triggerSourceDesc = map[TriggerSource]string{
	TriggerSourceExternal:  "External",
	TriggerSourceImmediate: "Immediate",
	TriggerSourceSoftware:  "Software",
	TriggerSourceACLine:    "AC Line",
	TriggerSourceVideo:     "Video",
}

func (ts TriggerSource) String() string { return triggerSourceDesc[ts] }

// AcquisitionStatus indicates the status of the acquisition.
type AcquisitionStatus int

// Available AcquisitionStatus values; integer values match the status codes
// returned by the instrument.
const (
	AcquisitionStatusComplete   AcquisitionStatus = 1
	AcquisitionStatusInProgress AcquisitionStatus = 0
	AcquisitionStatusUnknown    AcquisitionStatus = -1
)

var acquisitionStatusDesc = map[AcquisitionStatus]string{
	AcquisitionStatusComplete:   "Complete",
	AcquisitionStatusInProgress: "In Progress",
	AcquisitionStatusUnknown:    "Unknown",
}

func (as AcquisitionStatus) String() string { return acquisitionStatusDesc[as] }

// MarkerSearch specifies the type of marker search.
type MarkerSearch int

// Available MarkerSearch values.
const (
	MarkerSearchHighest MarkerSearch = iota
	MarkerSearchMinimum
	MarkerSearchNextPeak
	MarkerSearchNextPeakLeft
	MarkerSearchNextPeakRight
)

var markerSearchDesc = map[MarkerSearch]string{
	MarkerSearchHighest:       "Highest",
	MarkerSearchMinimum:       "Minimum",
	MarkerSearchNextPeak:      "Next Peak",
	MarkerSearchNextPeakLeft:  "Next Peak Left",
	MarkerSearchNextPeakRight: "Next Peak Right",
}

func (ms MarkerSearch) String() string { return markerSearchDesc[ms] }
