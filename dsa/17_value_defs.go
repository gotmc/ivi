// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

// WindowType specifies the FFT window function applied to time-domain data
// before frequency analysis.
type WindowType int

const (
	WindowHanning WindowType = iota
	WindowUniform
	WindowFlatTop
	WindowForce
	WindowExponential
)

var windowTypeDesc = map[WindowType]string{
	WindowHanning:     "Hanning",
	WindowUniform:     "Uniform",
	WindowFlatTop:     "Flat Top",
	WindowForce:       "Force",
	WindowExponential: "Exponential",
}

func (wt WindowType) String() string { return windowTypeDesc[wt] }

// AveragingType specifies the averaging method used to combine multiple
// measurements.
type AveragingType int

const (
	AveragingRMS AveragingType = iota
	AveragingRMSExponential
	AveragingTime
	AveragingTimeExponential
	AveragingPeakHold
)

var averagingTypeDesc = map[AveragingType]string{
	AveragingRMS:             "RMS",
	AveragingRMSExponential:  "RMS Exponential",
	AveragingTime:            "Time",
	AveragingTimeExponential: "Time Exponential",
	AveragingPeakHold:        "Peak Hold",
}

func (at AveragingType) String() string { return averagingTypeDesc[at] }

// InputCoupling specifies the input coupling mode.
type InputCoupling int

const (
	InputCouplingAC InputCoupling = iota
	InputCouplingDC
)

var inputCouplingDesc = map[InputCoupling]string{
	InputCouplingAC: "AC",
	InputCouplingDC: "DC",
}

func (ic InputCoupling) String() string { return inputCouplingDesc[ic] }

// MeasurementMode specifies the analysis mode of the DSA.
type MeasurementMode int

const (
	MeasurementModeFFT MeasurementMode = iota
	MeasurementModeCorrelation
	MeasurementModeHistogram
	MeasurementModeOctave
	MeasurementModeOrder
	MeasurementModeSweptSine
)

var measurementModeDesc = map[MeasurementMode]string{
	MeasurementModeFFT:         "FFT",
	MeasurementModeCorrelation: "Correlation",
	MeasurementModeHistogram:   "Histogram",
	MeasurementModeOctave:      "Octave",
	MeasurementModeOrder:       "Order",
	MeasurementModeSweptSine:   "Swept Sine",
}

func (mm MeasurementMode) String() string { return measurementModeDesc[mm] }
