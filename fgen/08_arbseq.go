// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 8 IviFgenArbSeq Extension Group

## Section 8.1 IviFgenArbSeq Overview

The IviFgenArbSeq extension group supports function generators capable of
producing sequences of arbitrary waveforms. In order to support this extension,
a driver must first support the IviFgenArbWfm extension group. This extension
uses the IviFgenArbWfm extension group’s attributes of sample rate, gain, and
offset to configure a sequence. The IviFgenArbSeq extension group includes
functions for creating, configuring, and generating sequences, and for
returning information about arbitrary sequence creation.

This extension affects instrument behavior when the Output Mode attribute is
set to Output Sequence.


## Section 8.2 IviFgenArbSeq Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|   8.2.1 | Arb Seq Handle   | Int32  | R/W    | Channel   |
|   8.2.2 | Num Seq Max      | Int32  | RO     | N/A       |
|   8.2.3 | Loop Count Max   | Int32  | RO     | N/A       |
|   8.2.4 | Seq Length Max   | Int32  | RO     | N/A       |
|   8.2.5 | Seq Length Min   | Int32  | RO     | N/A       |


## Section 8.3 IviFgenArbSeq Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

8.3.1 void Arbitrary.ClearMemory ();
8.3.2 void Arbitrary.Sequence.Clear (Int32 handle);
8.3.3 void Arbitrary.Sequence.Configure (String channelName,
                                         Int32 handle,
                                         Double gain,
                                         Double offset);
8.3.4 Int32 Arbitrary.Sequence.Create (Int32[] waveformHandle,
                                       Int32[] loopCount);

*/

// ArbSeq provides the interface required for the IviFgenArbSeq extension
// group.
type ArbSeq interface {
	NumSeqMax() (int, error)
	LoopCountMax() (int, error)
	SeqLengthMax() (int, error)
	SeqLengthMin() (int, error)
	ArbSeqClearMemory() error
	ArbSeqClearSequence(handle int) error
}

// ArbSeqChannel provides the interface required for the channel repeated
// capability for the IviFgenArbSeq extension group.
type ArbSeqChannel interface {
	ArbSeqHandle(int, error)
	SetArbSeqHandle(handle int) error
	ArbSeqConfigure(handle int, gain, offset float64) error
}
