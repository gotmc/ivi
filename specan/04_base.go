// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package specan

import (
	"context"
	"time"
)

/*

# Section 4 IviSpecAnBase Capability Group

## Section 4.1 IviSpecAnBase Overview

The IviSpecAnBase capability group supports spectrum analyzers that configure
and take a frequency sweep. A frequency sweep adjusts the frequency of a tuner
from the start frequency to the stop frequency in a defined amount of time.
While the tuner is being adjusted, power levels are measured. The result is an
array of amplitude versus frequency data.

## Section 4.2 IviSpecAnBase Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute            | Type   | Access | AppliesTo |
| ------- | -------------------- | ------ | ------ | --------- |
|   4.2.1 | Amplitude Units      | Int32  | R/W    | N/A       |
|   4.2.2 | Attenuation          | Real64 | R/W    | N/A       |
|   4.2.3 | Attenuation Auto     | Bool   | R/W    | N/A       |
|   4.2.4 | Detector Type        | Int32  | R/W    | Trace     |
|   4.2.5 | Detector Type Auto   | Bool   | R/W    | Trace     |
|   4.2.6 | Frequency Start      | Real64 | R/W    | N/A       |
|   4.2.7 | Frequency Stop       | Real64 | R/W    | N/A       |
|   4.2.8 | Frequency Offset     | Real64 | R/W    | N/A       |
|   4.2.9 | Input Impedance      | Real64 | R/W    | N/A       |
|  4.2.10 | Number Of Sweeps     | Int32  | R/W    | N/A       |
|  4.2.11 | Reference Level      | Real64 | R/W    | N/A       |
|  4.2.12 | Reference Level Offset | Real64 | R/W  | N/A       |
|  4.2.13 | Resolution Bandwidth | Real64 | R/W    | N/A       |
|  4.2.14 | Resolution BW Auto   | Bool   | R/W    | N/A       |
|  4.2.15 | Sweep Mode Continuous| Bool   | R/W    | N/A       |
|  4.2.16 | Sweep Time           | Real64 | R/W    | N/A       |
|  4.2.17 | Sweep Time Auto      | Bool   | R/W    | N/A       |
|  4.2.18 | Trace Count          | Int32  | RO     | N/A       |
|  4.2.22 | Trace Type           | Int32  | R/W    | Trace     |
|  4.2.23 | Vertical Scale       | Int32  | R/W    | N/A       |
|  4.2.24 | Video Bandwidth      | Real64 | R/W    | N/A       |
|  4.2.25 | Video Bandwidth Auto | Bool   | R/W    | N/A       |

## Section 4.3 IviSpecAnBase Functions

4.3.1  void Abort();
4.3.2  AcquisitionStatus AcquisitionStatus();
4.3.3  void ConfigureAcquisition(Boolean sweepModeContinuous, Int32 numberOfSweeps)
4.3.4  void ConfigureFrequencyCenterSpan(Real64 centerFreq, Real64 span)
4.3.6  void ConfigureFrequencyStartStop(Real64 startFreq, Real64 stopFreq)
4.3.7  void ConfigureLevel(AmplitudeUnits units, Real64 refLevel)
4.3.8  void ConfigureSweepCoupling(Real64 resBW, Real64 videoBW, Real64 sweepTime)
4.3.10 Real64[] FetchYTrace(String traceName)
4.3.12 void Initiate()
4.3.14 Real64[] ReadYTrace(String traceName, TimeSpan maxTime)

*/

// Base provides the interface required for the IviSpecAnBase capability group.
type Base interface {
	// Amplitude and level
	AmplitudeUnits(ctx context.Context) (AmplitudeUnits, error)
	SetAmplitudeUnits(ctx context.Context, units AmplitudeUnits) error
	ReferenceLevel(ctx context.Context) (float64, error)
	SetReferenceLevel(ctx context.Context, level float64) error
	ReferenceLevelOffset(ctx context.Context) (float64, error)
	SetReferenceLevelOffset(ctx context.Context, offset float64) error
	InputImpedance(ctx context.Context) (float64, error)
	SetInputImpedance(ctx context.Context, impedance float64) error
	VerticalScale(ctx context.Context) (VerticalScale, error)
	SetVerticalScale(ctx context.Context, scale VerticalScale) error

	// Attenuation
	Attenuation(ctx context.Context) (float64, error)
	SetAttenuation(ctx context.Context, attenuation float64) error
	AttenuationAuto(ctx context.Context) (bool, error)
	SetAttenuationAuto(ctx context.Context, auto bool) error

	// Frequency
	FrequencyStart(ctx context.Context) (float64, error)
	SetFrequencyStart(ctx context.Context, freq float64) error
	FrequencyStop(ctx context.Context) (float64, error)
	SetFrequencyStop(ctx context.Context, freq float64) error
	FrequencyOffset(ctx context.Context) (float64, error)
	SetFrequencyOffset(ctx context.Context, offset float64) error

	// Bandwidth
	ResolutionBandwidth(ctx context.Context) (float64, error)
	SetResolutionBandwidth(ctx context.Context, bw float64) error
	ResolutionBandwidthAuto(ctx context.Context) (bool, error)
	SetResolutionBandwidthAuto(ctx context.Context, auto bool) error
	VideoBandwidth(ctx context.Context) (float64, error)
	SetVideoBandwidth(ctx context.Context, bw float64) error
	VideoBandwidthAuto(ctx context.Context) (bool, error)
	SetVideoBandwidthAuto(ctx context.Context, auto bool) error

	// Sweep
	SweepModeContinuous(ctx context.Context) (bool, error)
	SetSweepModeContinuous(ctx context.Context, continuous bool) error
	SweepTime(ctx context.Context) (float64, error)
	SetSweepTime(ctx context.Context, sweepTime float64) error
	SweepTimeAuto(ctx context.Context) (bool, error)
	SetSweepTimeAuto(ctx context.Context, auto bool) error
	NumberOfSweeps(ctx context.Context) (int, error)
	SetNumberOfSweeps(ctx context.Context, num int) error

	// Trace
	TraceCount() int
	TraceType(ctx context.Context, traceName string) (TraceType, error)
	SetTraceType(ctx context.Context, traceName string, traceType TraceType) error
	DetectorType(ctx context.Context, traceName string) (DetectorType, error)
	SetDetectorType(ctx context.Context, traceName string, detector DetectorType) error
	DetectorTypeAuto(ctx context.Context, traceName string) (bool, error)
	SetDetectorTypeAuto(ctx context.Context, traceName string, auto bool) error

	// Acquisition control
	Abort(ctx context.Context) error
	AcquisitionStatus(ctx context.Context) (AcquisitionStatus, error)
	Initiate(ctx context.Context) error

	// Configuration helpers
	ConfigureFrequencyCenterSpan(ctx context.Context, centerFreq, span float64) error
	ConfigureFrequencyStartStop(ctx context.Context, startFreq, stopFreq float64) error
	ConfigureLevel(ctx context.Context, units AmplitudeUnits, refLevel float64) error
	ConfigureSweepCoupling(ctx context.Context, resBW, videoBW, sweepTime float64) error

	// Trace data
	FetchYTrace(ctx context.Context, traceName string) ([]float64, error)
	ReadYTrace(ctx context.Context, traceName string, maxTime time.Duration) ([]float64, error)
}
