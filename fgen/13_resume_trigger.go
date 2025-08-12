// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "time"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 13 IviFgenResumeTrigger Extension Group

## Section 13.1 IviFgenResumeTrigger Overview

The IviFgenResumeTrigger Extension Group supports function generators capable
of configuring a resume trigger. A resume trigger resumes generation after it
has been paused by a hold trigger, starting with the next point.

Setting the Resume Trigger Source attribute to a value other than None enables
the resume trigger. To disable the resume trigger, set the Resume Trigger
Source to None.


## Section 13.2 IviFgenResumeTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                | Type     | Access | AppliesTo |
| ------- | ------------------------ | ------   | ------ | --------- |
|  13.2.1 | Resume Trigger Delay     | TimeSpan | R/W    | Channel   |
|  13.2.2 | Resume Trigger Slope     | Int32    | R/W    | Channel   |
|  13.2.3 | Resume Trigger Source    | String   | R/W    | Channel   |
|  13.2.4 | Resume Trigger Threshold | Real64   | R/W    | Channel   |


## Section 13.3 IviFgenResumeTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

13.3.1 void Trigger.Resume.Configure(String channelName,
                                     String source,
                                     TriggerSlope slope);
13.3.2 void Trigger.Resume.SendSoftwareTrigger();

*/

// ResumeTrigger provides the interface required for the IviFgenHoldTrigger
// extension group.
type ResumeTrigger interface {
	SendSoftwareResumeTrigger() error
}

// ResumeTriggerChannel provides the interface for the channel repeated
// capability for the IviFgenResumeTrigger extension group.
type ResumeTriggerChannel interface {
	ResumeTriggerDelay() (time.Duration, error)
	SetResumeTriggerDelay(delay time.Duration) error
	ResumeTriggerSlope() (TriggerSlope, error)
	SetResumeTriggerSlope(slope TriggerSlope) error
	ResumeTriggerSource() (TriggerSource, error)
	SetResumeTriggerSource(source TriggerSource) error
	ResumeTriggerThreshold() (float64, error)
	SetResumeTriggerThreshold(threshold float64) error
	ResumeTriggerConfigure(source TriggerSource, slope TriggerSlope) error
}
