// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "time"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 11 IviFgenStopTrigger Extension Group

## Section 11.1 IviFgenStopTrigger Overview

The IviFgenStopTrigger Extension Group supports function generators capable of
configuring a stop trigger. A stop trigger terminates any generation and has
the same effect as calling the AbortGeneration function.

Setting the Stop Trigger Source attribute to a value other than None enables
the stop trigger. To disable the stop trigger, set the Stop Trigger Source to
None.


## Section 11.2 IviFgenStopTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | ------   | ------ | --------- |
|  11.2.1 | Stop Trigger Delay      | TimeSpan | R/W    | Channel   |
|  11.2.2 | Stop Trigger Slope      | Int32    | R/W    | Channel   |
|  11.2.3 | Stop Trigger Source     | String   | R/W    | Channel   |
|  11.2.4 | Stop Trigger Threshold  | Real64   | R/W    | Channel   |


## Section 11.3 IviFgenStopTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

11.3.1 void Trigger.Stop.Configure(String channelName,
                                   String source,
                                   TriggerSlope slope);
11.3.2 void Trigger.Stop.SendSoftwareTrigger();

*/

// StopTrigger provides the interface required for the IviFgenStopTrigger
// extension group.
type StopTrigger interface {
	SendSoftwareStopTrigger() error
}

// StopTriggerChannel provides the interface for the channel repeated
// capability for the IviFgenStopTrigger extension group.
type StopTriggerChannel interface {
	StopTriggerDelay() (time.Duration, error)
	SetStopTriggerDelay(delay time.Duration) error
	StopTriggerSlope() (TriggerSlope, error)
	SetStopTriggerSlope(slope TriggerSlope) error
	StopTriggerSource() (TriggerSource, error)
	SetStopTriggerSource(source TriggerSource) error
	StopTriggerThreshold() (float64, error)
	SetStopTriggerThreshold(threshold float64) error
	StopTriggerConfigure(source TriggerSource, slope TriggerSlope) error
}
