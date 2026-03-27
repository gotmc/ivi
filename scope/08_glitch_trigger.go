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

# Section 8 IviScopeGlitchTrigger Extension Group

## Section 8.1 IviScopeGlitchTrigger Overview

In addition to the fundamental capabilities, the IviScopeGlitchTrigger
extension group defines extensions for oscilloscopes that can trigger on a
“glitch” pulses.

A glitch occurs when the oscilloscope detects a pulse width that is less than
or a greater than a specified glitch duration. The figure below shows both
positive and negative glitches for the “less than” condition as well as the
positive “greater than” glitch.


## Section 8.2 IviScopeGlitchTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   8.2.1 | Glitch Condition       | Int32    | R/W    | N/A         |
|   8.2.2 | Glitch Polarity        | Polarity | R/W    | N/A         |
|   8.2.3 | Runt Polarity          | Polarity | R/W    | N/A         |
|   8.2.4 | Glitch Width           | TimeSpan | R/W    | N/A         |

## Section 8.3 IviScopeGlitchTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

8.3.1 void Trigger.Glitch.Configure (String source,
                                     Double level,
                                     Ivi.Driver.PrecisionTimeSpan width,
                                     Polarity polarity,
                                     GlitchCondition condition)

*/

// GlitchTriggerer provides the interface required for the
// IviScopeGlitchTrigger extension group.
type GlitchTriggerer interface {
	GlitchCondition(ctx context.Context) (GlitchCondition, error)
	SetGlitchCondition(ctx context.Context, condition GlitchCondition) error
	GlitchPolarity(ctx context.Context) (Polarity, error)
	SetGlitchPolarity(ctx context.Context, polarity Polarity) error
	RuntPolarity(ctx context.Context) (Polarity, error)
	SetRuntPolarity(ctx context.Context, polarity Polarity) error
	GlitchWidth(ctx context.Context) (time.Duration, error)
	SetGlitchWidth(ctx context.Context, width time.Duration) error
	ConfigureGlitchTrigger(
		ctx context.Context,
		source TriggerSource,
		width time.Duration,
		polarity Polarity,
		condition GlitchCondition,
	) error
}
