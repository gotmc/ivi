// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

/*

# Section 7 IviScopeRuntTrigger Extension Group

## Section 7.1 IviScopeRuntTrigger Overview

In addition to the fundamental capabilities, the IviScopeRuntTrigger extension
group defines extensions for oscilloscopes with the capability to trigger on
“runt” pulses.

A runt condition occurs when the oscilloscope detects a positive or negative
going pulse that crosses one voltage threshold but fails to cross a second
threshold before re-crossing the first. The figure below shows both positive
and negative runt polarities.


## Section 7.2 IviScopeRuntTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   7.2.1 | Runt High Threshold    | Real64   | R/W    | N/A         |
|   7.2.2 | Runt Low Threshold     | Real64   | R/W    | N/A         |
|   7.2.3 | Runt Polarity          | Polarity | R/W    | N/A         |

## Section 7.3 IviScopeRuntTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

7.3.1 void Trigger.Runt.Configure (String source,
                                   Double thresholdLow,
                                   Double thresholdHigh,
                                   Polarity polarity);

*/

// RuntTriggerer provides the interface required for the IviScopeRuntTrigger
// extension group.
type RuntTriggerer interface {
	RuntHighThreshold() (float64, error)
	SetRuntHighThreshold(float64) error
	RuntLowThreshold() (float64, error)
	SetRuntLowThreshold(float64) error
	RuntPolarity() (Polarity, error)
	SetRuntPolarity(polarity Polarity) error
	ConfigureRuntTrigger(
		source TriggerSource,
		lowThreshold, highThreshold float64,
		polarity Polarity,
	) error
}
