// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 20 IviFgenSampleClock Extension Group

## Section 20.1 IviFgenSampleClock Overview

The IviFgenSampleClock extension group supports arbitrary waveform generators
with the ability to use (or provide) an external sample clock. Note that when
using an external sample clock, the Arbitrary Sample Rate attribute must be set
to the corresponding frequency of the external sample clock.


## Section 20.2 IviFgenSampleClock Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                 | Type   | Access | AppliesTo |
| ------- | ------------------------- | ------ | ------ | --------- |
|  20.2.1 | Sample Clk Source         | Int32  | R/W    | N/A       |
|  20.2.2 | Sample Clk Output Enabled | Bool   | R/W    | N/A       |


## Section 20.3 IviFgenSampleClock Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// SampleClock provides the interface required for the IviFgenSampleClock
// extension group.
type SampleClock interface {
	SampleClockSource() (SampleClockSource, error)
	SetSampleClockSource(SampleClockSource) error
	SampleClockOutputEnabled() (bool, error)
	SetSampleClockOutputEnabled(b bool) error
}
