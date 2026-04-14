// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 15 IviFgenInternalTrigger Extension Group

## Section 15.1 IviFgenInternalTrigger Overview

The IviFgenInternalTrigger Extension Group supports function generators that
can generate output based on an internally generated trigger signal. The user
can configure the rate at which internal triggers are generated.

This extension affects instrument behavior when the Trigger Source attribute is
set to Internal Trigger.


## Section 15.2 IviFgenInternalTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|  15.2.1 | Int Trigger Rate | Real64 | R/W    | N/A       |


## Section 15.3 IviFgenInternalTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// IntTriggerChannel provides the interface for the channel repeated capability
// to support the IviFgenInternalTrigger extension group.
//
// Deviation from IVI-4.3: The IVI specification defines InternalTriggerRate as
// a driver-level (non-channel) attribute. We place it on the channel because
// multi-channel function generators (e.g., Keysight 33500B) have independent
// burst periods per channel. Single-channel instruments are unaffected since
// they have only one channel.
type IntTriggerChannel interface {
	InternalTriggerRate() (float64, error)
	SetInternalTriggerRate(rate float64) error
}
