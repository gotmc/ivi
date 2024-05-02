// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

import "time"

/*

# Section 11 IviScopeWaveformMeasurement Extension Group

## Section 11.1 IviScopeWaveformMeasurement Overview

The IviScopeWaveformMeasurement extension group defines extensions for
oscilloscopes capable of calculating various measurements such as rise-time,
fall-time, period, and frequency from an acquired waveform. In .NET, see the
Fetch Waveform Measurement and Read Waveform Measurement methods for an
explanation of the differences between measurement functions and time
measurement enumerations.

## Section 11.2 IviScopeWaveformMeasurement Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|  11.2.1 | Measurement High Ref   | Real64   | R/W    | N/A         |
|  11.2.2 | Measurement Low Ref    | Real64   | R/W    | N/A         |
|  11.2.3 | Measurement Middle Ref | Real64   | R/W    | N/A         |

## Section 11.3 IviScopeWaveformMeasurement Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

11.3.1 void ReferenceLevel.Configure (Double low,
                                      Double middle,
                                      Double high);

11.3.2 Double Channels[].Measurement.FetchWaveformMeasurement (
                                  MeasurementFunction measurementFunction);

       PrecisionTimeSpan Channels[].Measurement.FetchWaveformMeasurement (
                                  TimeMeasurementFunction measurementFunction);

11.3.3 Double Channels[].Measurement.ReadWaveformMeasurement (
                                  MeasurementFunction measurementFunction,
                                  PrecisionTimeSpan maximumTime);

       PrecisionTimeSpan Channels[].Measurement.ReadWaveformMeasurement (
                                  TimeMeasurementFunction measurementFunction,
                                  PrecisionTimeSpan maximumTime);

*/

// WaveformMeasurement provides the interface required for the
// IviScopeWaveformMeasurement extension group.
type WaveformMeasurer interface {
	HighReferenceLevel() (float64, error)
	SetHighReferenceLevel(high float64) error
	LowReferenceLevel() (float64, error)
	SetLowReferenceLevel(low float64) error
	MiddleReferenceLevel() (float64, error)
	SetMiddleReferenceLevel(mid float64) error
	ConfigureReferenceLevels(low, mid, high float64) error
}

type WaveformMeasurerChannel interface {
	FetchWaveformMeasurement(msrmnt WaveformMeasurement) (float64, error)
	ReadWaveformMeasurement(msrmnt WaveformMeasurement, maxTime time.Duration) (float64, error)
}
