// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 17 IviFgenBurst Extension Group

## Section 17.1 IviFgenBurst Overview

The IviFgenBurst Extension Group supports function generators capable of
generating a discrete number of waveform cycles based on a trigger. The trigger
is configured with the IviFgenTrigger or IviFgenStartTrigger extension group.
The user can specify the number of waveform cycles to generate when a trigger
event occurs.

For standard and arbitrary waveforms, a cycle is one period of the waveform.
For arbitrary sequences, a cycle is one complete progression through the
generation of all iterations of all waveforms in the sequence.

This extension affects instrument behavior when the Operation Mode attribute is
set to Operate Burst.


## Section 17.2 IviFgenBurst Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|  17.2.1 | Burst Count      | Int32  | R/W    | Channel   |


## Section 17.3 IviFgenBurst Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// BurstChannel provides the interface for the channel repeated capability for
// the IviFgenBurst capability group.
type BurstChannel interface {
	BurstCount() (int, error)
	SetBurstCount(count int) error
}
