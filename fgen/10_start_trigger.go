// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "time"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 10 IviFgenStartTrigger Extension Group

## Section 10.1 IviFgenStartTrigger Overview

The IviFgenStartTrigger Extension Group supports function generators capable of
configuring a start trigger. A start trigger initiates generation of a waveform
or sequence. This Extension group deprecates the IviFgenTrigger extension
group. Drivers that implement this extension group shall implement the
IviFgenTrigger extension group as well, to ensure that applications based on
previous versions of the IviFgen class specification continue to work with
version 5.0.

Setting the Start Trigger Source attribute to a value other than None enables
the start trigger. To disable the start trigger, set the Start Trigger Source
to None.


## Section 10.2 IviFgenStartTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | ------   | ------ | --------- |
|  10.2.1 | Start Trigger Delay     | TimeSpan | R/W    | Channel   |
|  10.2.2 | Start Trigger Slope     | Int32    | R/W    | Channel   |
|  10.2.3 | Start Trigger Source    | String   | R/W    | Channel   |
|  10.2.4 | Start Trigger Threshold | Real64   | R/W    | Channel   |


## Section 10.3 IviFgenStartTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

10.3.1 void Trigger.Start.Configure(String channelName,
                                    String source,
                                    TriggerSlope slope);

*/

// StartTriggerChannel provides the interface for the channel repeated
// capability for the IviFgenStartTrigger extension group.
type StartTriggerChannel interface {
	StartTriggerDelay() (time.Duration, error)
	SetStartTriggerDelay(delay time.Duration) error
	StartTriggerSlope() (TriggerSlope, error)
	SetStartTriggerSlope(slope TriggerSlope) error
	StartTriggerSource() (TriggerSource, error)
	SetStartTriggerSource(source TriggerSource) error
	StartTriggerThreshold() (float64, error)
	SetStartTriggerThreshold(threshold float64) error
	StartTriggerConfigure(source TriggerSource, slope TriggerSlope) error
}
