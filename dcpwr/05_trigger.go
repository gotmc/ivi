// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

/*

# Section 5 IviDCPwrTrigger Extension Group

## Section 5.1 IviDCPwrTrigger Overview

The IviDCPwrTrigger extension group defines extensions for DC power supplies
capable of changing the output signal based on a trigger event.

## Section 5.2 IviDCPwrTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type        | Access | AppliesTo   |
| ------- | ----------------------- | ----------- | ------ | ----------- |
|   5.2.1 | Trigger Source          | String      | R/W    | Channel     |
|   5.2.2 | Triggered Current Limit | Real64      | R/W    | Channel     |
|   5.2.3 | Triggered Voltage Level | Real64      | R/W    | Channel     |

## Section 5.3 IviDCPwrTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

5.3.1 void Trigger.Abort();
5.3.5 void Trigger.Initiate();

*/

// Trigger provides the interface for the IviDCPwrTrigger extension group.
type Trigger interface {
	AbortTrigger() error
	InitiateTrigger() error
}

// TriggerChannel provides the interface for the channel repeated capability
// for the IviDCPwrTrigger extension group.
type TriggerChannel interface {
	TriggerSource() (TriggerSource, error)
	SetTriggerSource(source TriggerSource) error
	TriggeredCurrentLimit(float64, error)
	SetTriggeredCurrentLimit(limit float64) error
	TriggeredVoltageLevel(float64, error)
	SetTriggeredVoltageLevel(level float64) error
}
