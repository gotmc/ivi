// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package specan

import (
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
	AmplitudeUnits() (AmplitudeUnits, error)
	SetAmplitudeUnits(units AmplitudeUnits) error
	ReferenceLevel() (float64, error)
	SetReferenceLevel(level float64) error
	ReferenceLevelOffset() (float64, error)
	SetReferenceLevelOffset(offset float64) error
	InputImpedance() (float64, error)
	SetInputImpedance(impedance float64) error
	VerticalScale() (VerticalScale, error)
	SetVerticalScale(scale VerticalScale) error

	// Attenuation
	Attenuation() (float64, error)
	SetAttenuation(attenuation float64) error
	AttenuationAuto() (bool, error)
	SetAttenuationAuto(auto bool) error

	// Frequency
	FrequencyStart() (float64, error)
	SetFrequencyStart(freq float64) error
	FrequencyStop() (float64, error)
	SetFrequencyStop(freq float64) error
	FrequencyOffset() (float64, error)
	SetFrequencyOffset(offset float64) error

	// Bandwidth
	ResolutionBandwidth() (float64, error)
	SetResolutionBandwidth(bw float64) error
	ResolutionBandwidthAuto() (bool, error)
	SetResolutionBandwidthAuto(auto bool) error
	VideoBandwidth() (float64, error)
	SetVideoBandwidth(bw float64) error
	VideoBandwidthAuto() (bool, error)
	SetVideoBandwidthAuto(auto bool) error

	// Sweep
	SweepModeContinuous() (bool, error)
	SetSweepModeContinuous(continuous bool) error
	SweepTime() (float64, error)
	SetSweepTime(sweepTime float64) error
	SweepTimeAuto() (bool, error)
	SetSweepTimeAuto(auto bool) error
	NumberOfSweeps() (int, error)
	SetNumberOfSweeps(num int) error

	// Trace
	TraceCount() int
	TraceType(traceName string) (TraceType, error)
	SetTraceType(traceName string, traceType TraceType) error
	DetectorType(traceName string) (DetectorType, error)
	SetDetectorType(traceName string, detector DetectorType) error
	DetectorTypeAuto(traceName string) (bool, error)
	SetDetectorTypeAuto(traceName string, auto bool) error

	// Acquisition control
	Abort() error
	AcquisitionStatus() (AcquisitionStatus, error)
	Initiate() error

	// Configuration helpers
	ConfigureFrequencyCenterSpan(centerFreq, span float64) error
	ConfigureFrequencyStartStop(startFreq, stopFreq float64) error
	ConfigureLevel(units AmplitudeUnits, refLevel float64) error
	ConfigureSweepCoupling(resBW, videoBW, sweepTime float64) error

	// Trace data
	FetchYTrace(traceName string) ([]float64, error)
	ReadYTrace(traceName string, maxTime time.Duration) ([]float64, error)
}
