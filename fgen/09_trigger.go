// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "context"

/*

# Section 9 IviFgenTrigger Extension Group

***DEPRECATED***

## Section 9.1 IviFgenTrigger Overview

The IviFgenTrigger Extension Group supports function generators capable of
configuring a trigger. This trigger source is used by other extension groups
like IviFgenBurst to determine when to produce output generation. This
extension group has been deprecated by the IviFgenStartTrigger Extension group.
Drivers that support the IviFgenTrigger Extension group shall also support the
IviFgenStartTrigger Extension group in order to be compliant with version 5.0
or later of the IviFgen class specification.

This extension affects instrument behavior when the Operation Mode attribute is
set to Operate Burst.


## Section 9.2 IviFgenTrigger Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

For .NET, see section 10.2.3 Start Trigger Source.


## Section 9.3 IviFgenTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

There are no .NET functions for Section 9.3.


*/

// TriggerChannel provides the interface for the channel repeated capability for
// the IviFgenTrigger extension group.
type TriggerChannel interface {
	TriggerSource(ctx context.Context) (OldTriggerSource, error)
	SetTriggerSource(ctx context.Context, src OldTriggerSource) error
}

// oldToNewTriggerSource maps deprecated OldTriggerSource values to their
// TriggerSource equivalents.
var oldToNewTriggerSource = map[OldTriggerSource]TriggerSource{
	OldTriggerSourceInternal: TriggerSourceInternal,
	OldTriggerSourceExternal: TriggerSourceExternal,
	OldTriggerSourceSoftware: TriggerSourceSoftware,
	OldTriggerSourceTTL0:     TriggerSourceTTL0,
	OldTriggerSourceTTL1:     TriggerSourceTTL1,
	OldTriggerSourceTTL2:     TriggerSourceTTL2,
	OldTriggerSourceTTL3:     TriggerSourceTTL3,
	OldTriggerSourceTTL4:     TriggerSourceTTL4,
	OldTriggerSourceTTL5:     TriggerSourceTTL5,
	OldTriggerSourceTTL6:     TriggerSourceTTL6,
	OldTriggerSourceTTL7:     TriggerSourceTTL7,
	OldTriggerSourceECL0:     TriggerSourcePXITrig0,
	OldTriggerSourceECL1:     TriggerSourcePXITrig1,
	OldTriggerSourcePXIStar:  TriggerSourcePXIStar,
	OldTriggerSourceRTSI0:    TriggerSourceRTSI0,
	OldTriggerSourceRTSI1:    TriggerSourceRTSI1,
	OldTriggerSourceRTSI2:    TriggerSourceRTSI2,
	OldTriggerSourceRTSI3:    TriggerSourceRTSI3,
	OldTriggerSourceRTSI4:    TriggerSourceRTSI4,
	OldTriggerSourceRTSI5:    TriggerSourceRTSI5,
	OldTriggerSourceRTSI6:    TriggerSourceRTSI6,
}

// newToOldTriggerSource maps TriggerSource values back to their deprecated
// OldTriggerSource equivalents.
var newToOldTriggerSource = map[TriggerSource]OldTriggerSource{
	TriggerSourceInternal: OldTriggerSourceInternal,
	TriggerSourceExternal: OldTriggerSourceExternal,
	TriggerSourceSoftware: OldTriggerSourceSoftware,
	TriggerSourceTTL0:     OldTriggerSourceTTL0,
	TriggerSourceTTL1:     OldTriggerSourceTTL1,
	TriggerSourceTTL2:     OldTriggerSourceTTL2,
	TriggerSourceTTL3:     OldTriggerSourceTTL3,
	TriggerSourceTTL4:     OldTriggerSourceTTL4,
	TriggerSourceTTL5:     OldTriggerSourceTTL5,
	TriggerSourceTTL6:     OldTriggerSourceTTL6,
	TriggerSourceTTL7:     OldTriggerSourceTTL7,
	TriggerSourcePXITrig0: OldTriggerSourceECL0,
	TriggerSourcePXITrig1: OldTriggerSourceECL1,
	TriggerSourcePXIStar:  OldTriggerSourcePXIStar,
	TriggerSourceRTSI0:    OldTriggerSourceRTSI0,
	TriggerSourceRTSI1:    OldTriggerSourceRTSI1,
	TriggerSourceRTSI2:    OldTriggerSourceRTSI2,
	TriggerSourceRTSI3:    OldTriggerSourceRTSI3,
	TriggerSourceRTSI4:    OldTriggerSourceRTSI4,
	TriggerSourceRTSI5:    OldTriggerSourceRTSI5,
	TriggerSourceRTSI6:    OldTriggerSourceRTSI6,
}

// OldToNewTriggerSource converts a deprecated OldTriggerSource to the
// equivalent TriggerSource. Returns the converted value and true if the
// mapping exists, or the zero value and false otherwise.
func OldToNewTriggerSource(old OldTriggerSource) (TriggerSource, bool) {
	ts, ok := oldToNewTriggerSource[old]
	return ts, ok
}

// NewToOldTriggerSource converts a TriggerSource to the equivalent deprecated
// OldTriggerSource. Returns the converted value and true if the mapping
// exists, or the zero value and false otherwise.
func NewToOldTriggerSource(ts TriggerSource) (OldTriggerSource, bool) {
	old, ok := newToOldTriggerSource[ts]
	return old, ok
}
