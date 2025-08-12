// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

/*

# Section 6 IviScopeTVTrigger Extension Group

## Section 6.1 IviScopeTVTrigger Overview

The IviScopeTVTrigger extension group defines extensions for oscilloscopes
capable of triggering on TV signals.


## Section 6.2 IviScopeTVTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   6.2.1 | TV Trigger Event       | Int32    | R/W    | N/A         |
|   6.2.2 | TV Trigger Line Number | Int32    | R/W    | N/A         |
|   6.2.3 | TV Trigger Polarity    | Int32    | R/W    | N/A         |
|   6.2.4 | TV Trigger Sig Format  | Int32    | R/W    | N/A         |

## Section 6.3 IviScopeTVTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

6.3.2 void Trigger.TV.Configure (String source,
                                 TVSignalFormat signalFormat,
                                 TVTriggerEvent event,
                                 TVTriggerPolarity polarity);

*/

// TVTriggerer provides the interface required for the IviScopeTVTrigger
// extension group.
type TVTriggerer interface {
	TVTriggerEvent() (TVTriggerEvent, error)
	SetTVTriggerEvent(event TVTriggerEvent) error
	TVTriggerLineNumber() (int, error)
	SetTVTriggerLineNumber(line int) error
	TVTriggerPolarity() (TVTriggerPolarity, error)
	SetTVTriggerPolarity(polarity TVTriggerPolarity) error
	TVTriggerSignalFormat() (TVTriggerSignalFormat, error)
	SetTVTriggerSignalFormat(format TVTriggerSignalFormat) error
	ConfigureTVTrigger(
		source TriggerSource,
		format TVTriggerSignalFormat,
		event TVTriggerEvent,
		polarity TVTriggerPolarity,
	) error
}
