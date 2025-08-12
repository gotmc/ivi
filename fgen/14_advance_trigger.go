// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "time"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 14 IviFgenAdvanceTrigger Extension Group

## Section 14.1 IviFgenAdvanceTrigger Overview

The IviFgenAdvanceTrigger Extension Group supports function generators capable
of configuring an advance trigger. An advance trigger advances generation to
the end of the current waveform, where generation proceeds according to the
current configuration.

Setting the Advance Trigger Source attribute to a value other than None enables
the advance trigger. To disable the advance trigger, set the Advance Trigger
Source to None.


## Section 14.2 IviFgenAdvanceTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                 | Type     | Access | AppliesTo |
| ------- | ------------------------- | ------   | ------ | --------- |
|  14.2.1 | Advance Trigger Delay     | TimeSpan | R/W    | Channel   |
|  14.2.2 | Advance Trigger Slope     | Int32    | R/W    | Channel   |
|  14.2.3 | Advance Trigger Source    | String   | R/W    | Channel   |
|  14.2.4 | Advance Trigger Threshold | Real64   | R/W    | Channel   |


## Section 14.3 IviFgenAdvanceTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

14.3.1 void Trigger.Advance.Configure(String channelName,
                                      String source,
                                      TriggerSlope slope);
14.3.2 void Trigger.Advance.SendSoftwareTrigger();

*/

// AdvanceTrigger provides the interface required for the IviFgenAdvanceTrigger
// extension group.
type AdvanceTrigger interface {
	SendSoftwareAdvanceTrigger() error
}

// AdvanceTriggerChannel provides the interface for the channel repeated
// capability for the IviFgenAdvanceTrigger extension group.
type AdvanceTriggerChannel interface {
	AdvanceTriggerDelay() (time.Duration, error)
	SetAdvanceTriggerDelay(delay time.Duration) error
	AdvanceTriggerSlope() (TriggerSlope, error)
	SetAdvanceTriggerSlope(slope TriggerSlope) error
	AdvanceTriggerSource() (TriggerSource, error)
	SetAdvanceTriggerSource(source TriggerSource) error
	AdvanceTriggerThreshold() (float64, error)
	SetAdvanceTriggerThreshold(threshold float64) error
	AdvanceTriggerConfigure(source TriggerSource, slope TriggerSlope) error
}
