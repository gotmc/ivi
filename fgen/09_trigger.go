// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

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
	TriggerSource() (OldTriggerSource, error)
	SetTriggerSource(OldTriggerSource) error
}
