// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "context"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 5 IviFgenStdFunc Extension Group

## Section 5.1 IviFgenStdFunc Overview

The IviFgenStdFunc Extension Group supports function generators that can
produce manufacturer-supplied periodic waveforms. The user can modify
properties of the waveform such as frequency, amplitude, DC offset, and phase
offset.

This extension affects instrument behavior when the Output Mode attribute is
set to Output Function.

## Section 5.2 IviFgenStdFunc Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|   5.2.1 | Amplitude        | Real64 | R/W    | Channel   |
|   5.2.2 | DC Offset        | Real64 | R/W    | Channel   |
|   5.2.3 | Duty Cycle High  | Real64 | R/W    | Channel   |
|   5.2.4 | Frequency        | Real64 | R/W    | Channel   |
|   5.2.5 | Start Phase      | Real64 | R/W    | Channel   |
|   5.2.6 | Waveform         | Int32  | R/W    | Channel   |


## Section 5.3 IviFgenStdFunc Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

5.3.1 void StandardWaveform.Configure (String channelName,
                                       StandardWaveformFunction waveformFunction,
                                       Double amplitude,
                                       Double dcOffset,
                                       Double frequency,
                                       Double startPhase);

*/

// StdFuncChannel provides the interface for the channel repeated capability
// for the IviFgenStdFunc extension group.
type StdFuncChannel interface {
	Amplitude(ctx context.Context) (float64, error)
	SetAmplitude(ctx context.Context, amp float64) error
	DCOffset(ctx context.Context) (float64, error)
	SetDCOffset(ctx context.Context, offset float64) error
	DutyCycleHigh(ctx context.Context) (float64, error)
	SetDutyCycleHigh(ctx context.Context, duty float64) error
	Frequency(ctx context.Context) (float64, error)
	SetFrequency(ctx context.Context, freq float64) error
	StartPhase(ctx context.Context) (float64, error)
	SetStartPhase(ctx context.Context, start float64) error
	StandardWaveform(ctx context.Context) (StandardWaveform, error)
	SetStandardWaveform(ctx context.Context, wave StandardWaveform) error
	ConfigureStandardWaveform(
		ctx context.Context,
		wave StandardWaveform,
		amp, offset, freq, phase float64,
	) error
}
