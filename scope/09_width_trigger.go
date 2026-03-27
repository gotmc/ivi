// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

import (
	"context"
	"time"
)

/*

# Section 9 IviScopeWidthTrigger Extension Group

## Section 9.1 IviScopeWidthTrigger Overview

In addition to the fundamental capabilities, the IviScopeWidthTrigger extension
group defines extensions for oscilloscopes capable of triggering on
user-specified pulse widths.

Width triggering occurs when the oscilloscope detects a positive or negative
pulse with a width between, or optionally outside, the user-specified
thresholds. The figure below shows positive and negative pulses that fall
within the user-specified thresholds.

## Section 9.2 IviScopeWidthTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   9.2.1 | Width Condition        | Int32    | R/W    | N/A         |
|   9.2.2 | Width High Threshold   | TimeSpan | R/W    | N/A         |
|   9.2.3 | Width Low Threshold    | TimeSpan | R/W    | N/A         |
|   9.2.4 | Width Polarity         | Polarity | R/W    | N/A         |

## Section 9.3 IviScopeWidthTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

9.3.1 void Trigger.Width.Configure (String source,
                                    Double level,
                                    Ivi.Driver.PrecisionTimeSpan thresholdLow,
                                    Ivi.Driver.PrecisionTimeSpan thresholdHigh,
                                    Polarity polarity,
                                    WidthCondition condition);

*/

// WidthTriggerer provides the interface required for the
// IviScopeWidthTrigger extension group.
type WidthTriggerer interface {
	WidthCondition(ctx context.Context) (WidthCondition, error)
	SetWidthCondition(ctx context.Context, condition WidthCondition) error
	WidthHighThreshold(ctx context.Context) (time.Duration, error)
	SetWidthHighThreshold(ctx context.Context, highTime time.Duration) error
	WidthLowThreshold(ctx context.Context) (time.Duration, error)
	SetWidthLowThreshold(ctx context.Context, lowTime time.Duration) error
	WidthPolarity(ctx context.Context) (Polarity, error)
	SetWidthPolarity(ctx context.Context, polarity Polarity) error
	ConfigureWidthTrigger(
		ctx context.Context,
		source TriggerSource,
		lowTime, highTime time.Duration,
		polarity Polarity,
		condition WidthCondition,
	) error
}
