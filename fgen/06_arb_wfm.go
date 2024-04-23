// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

# Section 6 IviFgenArbWfm Extension Group

## Section 6.1 IviFgenArbWfm Overview

The IviFgenArbWfm Extension Group supports function generators capable of
producing user- defined arbitrary waveforms. The user can modify parameters of
the arbitrary waveform such as sample rate, waveform gain, and waveform offset.
The IviFgenArbWfm extension group includes functions for creating, configuring,
and generating arbitrary waveforms, and for returning information about
arbitrary waveform creation.This extension affects instrument behavior when the
Output Mode attribute is set to Output Arbitrary or Output Sequence.

## Section 6.2 IviFgenArbWfm Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                 | Type   | Access | AppliesTo |
| ------- | ------------------------- | ------ | ------ | --------- |
|   6.2.1 | Arbitrary Gain            | Real64 | R/W    | Channel   |
|   6.2.2 | Arbitrary Offset          | Real64 | R/W    | Channel   |
|   6.2.3 | Arbitrary Sample Rate     | Real64 | R/W    | N/A       |
|   6.2.4 | Arbitrary Waveform Handle | Int32  | R/W    | Channel   |
|   6.2.5 | Number Waveforms Max      | Int32  | RO     | N/A       |
|   6.2.6 | Waveform Size Max         | Int64  | RO     | N/A       |
|   6.2.7 | Waveform Size Min         | Int64  | RO     \ N/A       |
|   6.2.8 | Waveform Quantum          | Int32  | RO     | N/A       |

### Section 6.3 IviFgenArbWfm Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

6.3.1 void Arbitrary.Waveform.Clear (Int32 handle);
6.3.2 void Arbitrary.Waveform.Condigure (String channelName,
                                         Int32 handle,
                                         Double gain,
                                         Double offset);

*/

// ArbWfm provides the interface required to support the IviFgenArbWfm
// extension group as described in Section 6 of the IVI-4.3: IviFgen Class
// Specification.
type ArbWfm interface {
	ArbitrarySampleRate() (float64, error)
	SetArbitrarySampleRate(rate float64) error
	ArbWfmNumberWaveformsMax() int
	ArbWfmMaxSize() int
	ArbWfmMinSize() int
	ArbWfmQuantum() int
}

// ArbWfmChannel provides the interface required for the channel repeated
// capability to support the IviFgenArbWfm extension group as
// described in Section 6 of the IVI-4.3: IviFgen Class Specification.
type ArbWfmChannel interface {
	ArbitraryGain() (float64, error)
	SetArbitraryGain(gain float64) error
	ArbitraryOffset() (float64, error)
	SetArbitraryOffset(offset float64) error
	ArbitraryWaveformHandle() (int, error)
	SetArbitraryWaveformHandle(handle int) error
}
