// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

import "time"

/*

The DSA (Dynamic Signal Analyzer) class is not part of the IVI Foundation
specifications. This interface is modeled on IVI-4.8: IviSpecAn, adapted for
the characteristics of dynamic signal analyzers which perform FFT-based
frequency analysis on time-domain input signals.

Key differences from spectrum analyzers:
  - DSAs use FFT analysis rather than swept-tuned receivers
  - Resolution is specified in spectral lines, not bandwidth
  - DSAs have configurable window functions for FFT processing
  - DSAs support multiple averaging types (RMS, time, peak hold)
  - DSAs often include a built-in source output

*/

// Base provides the interface for the core DSA capability group, covering
// frequency configuration, averaging, input range, measurement control,
// and trace data retrieval.
type Base interface {
	// Frequency
	FrequencyStart() (float64, error)
	SetFrequencyStart(freq float64) error
	FrequencyStop() (float64, error)
	SetFrequencyStop(freq float64) error
	FrequencySpan() (float64, error)
	SetFrequencySpan(span float64) error
	FrequencyCenter() (float64, error)
	SetFrequencyCenter(freq float64) error
	ConfigureFrequencyStartStop(startFreq, stopFreq float64) error
	ConfigureFrequencyCenterSpan(centerFreq, span float64) error

	// Resolution (spectral lines)
	SpectralLines() (int, error)
	SetSpectralLines(lines int) error

	// Window function
	WindowType() (WindowType, error)
	SetWindowType(window WindowType) error

	// Averaging
	AveragingEnabled() (bool, error)
	SetAveragingEnabled(enabled bool) error
	AveragingCount() (int, error)
	SetAveragingCount(count int) error
	AveragingType() (AveragingType, error)
	SetAveragingType(avgType AveragingType) error

	// Input range
	InputRange(channel int) (float64, error)
	SetInputRange(channel int, rangeDBVrms float64) error
	InputAutoRange(channel int) (bool, error)
	SetInputAutoRange(channel int, auto bool) error
	InputCoupling(channel int) (InputCoupling, error)
	SetInputCoupling(channel int, coupling InputCoupling) error

	// Measurement control
	MeasurementMode() (MeasurementMode, error)
	SetMeasurementMode(mode MeasurementMode) error
	ChannelCount() (int, error)
	SetChannelCount(count int) error

	// Acquisition control
	Abort() error
	Initiate() error
	SweepModeContinuous() (bool, error)
	SetSweepModeContinuous(continuous bool) error

	// Trace data
	FetchYTrace(traceName string) ([]float64, error)
	ReadYTrace(traceName string, maxTime time.Duration) ([]float64, error)
}

// Source provides the interface for DSA source output control. Not all DSAs
// have a built-in source; drivers for instruments without a source should
// return [ivi.ErrFunctionNotSupported].
type Source interface {
	SourceEnabled() (bool, error)
	SetSourceEnabled(enabled bool) error
	SourceShape() (SourceShape, error)
	SetSourceShape(shape SourceShape) error
	SourceFrequency() (float64, error)
	SetSourceFrequency(freq float64) error
	SourceOutputLevel() (float64, error)
	SetSourceOutputLevel(level float64) error
}
