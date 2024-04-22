// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

# Section 7 IviFgenArbFrequency Extension Group

## Section 7.1 IviFgenArbFrequency Overview

The IviFgenArbFrequency extension group supports function generators capable of
producing arbitrary waveforms that allow the user to set the rate at which an
entire waveform buffer is generated. In order to support this extension, a
driver must first support the IviFgenArbWfm extension group. This extension
uses the IviFgenArbWfm extension group’s attributes of Arbitrary Waveform
Handle, Arbitrary Gain, and Arbitrary Offset to configure an arbitrary
waveform.

This extension affects instrument behavior when the Output Mode attribute is
set to Output Arbitrary.


## Section 7.2 IviFgenArbFrequency Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|   7.2.1 | Arbitrary Freq   | Real64 | R/W    | Channel   |


## Section 7.3 IviFgenArbFrequency Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// ArbFrequencyChannel provides the interface required for the channel repeated
// capability group for the IviFgenArbFrequency extension group.
type ArbFrequencyChannel interface {
	ArbitraryFrequency() (float64, error)
	SetArbitraryFrequency(freq float64) error
}
