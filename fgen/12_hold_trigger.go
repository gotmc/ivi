// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "time"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 10 IviFgenHoldTrigger Extension Group

## Section 12.1 IviFgenHoldTrigger Overview

The IviFgenHoldTrigger Extension Group supports function generators capable of
configuring a hold trigger. A hold trigger pauses generation. From the paused
state, a resume trigger resumes generation; a stop trigger terminates
generation; start trigger behavior is vendor defined.

Setting the Hold Trigger Source attribute to a value other than None enables
the hold trigger. To disable the hold trigger, set the Hold Trigger Source to
None.

## Section 11.2 IviFgenHoldTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | ------   | ------ | --------- |
|  12.2.1 | Hold Trigger Delay      | TimeSpan | R/W    | Channel   |
|  12.2.2 | Hold Trigger Slope      | Int32    | R/W    | Channel   |
|  12.2.3 | Hold Trigger Source     | String   | R/W    | Channel   |
|  12.2.4 | Hold Trigger Threshold  | Real64   | R/W    | Channel   |


## Section 11.3 IviFgenHoldTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

12.3.1 void Trigger.Hold.Configure(String channelName,
                                   String source,
                                   TriggerSlope slope);
12.3.2 void Trigger.Hold.SendSoftwareTrigger();

*/

// HoldTrigger provides the interface required for the IviFgenHoldTrigger
// extension group.
type HoldTrigger interface {
	SendSoftwareHoldTrigger() error
}

// HoldTriggerChannel provides the interface for the channel repeated
// capability for the IviFgenHoldTrigger extension group.
type HoldTriggerChannel interface {
	HoldTriggerDelay() (time.Duration, error)
	SetHoldTriggerDelay(delay time.Duration) error
	HoldTriggerSlope() (TriggerSlope, error)
	SetHoldTriggerSlope(slope TriggerSlope) error
	HoldTriggerSource() (TriggerSource, error)
	SetHoldTriggerSource(source TriggerSource) error
	HoldTriggerThreshold() (float64, error)
	SetHoldTriggerThreshold(threshold float64) error
	HoldTriggerConfigure(source TriggerSource, slope TriggerSlope) error
}
