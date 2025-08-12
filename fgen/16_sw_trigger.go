// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 16 IviFgenSoftwareTrigger Extension Group

## Section 16.1 IviFgenSoftwareTrigger Overview

The IviFgenSoftwareTrigger Extension Group supports function generators that
can generate output based on a software trigger signal. The user can send a
software trigger to cause signal output to occur.

This extension affects instrument behavior when the Trigger Source attribute is
set to Software Trigger.


## Section 16.2 IviFgenSoftwareTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

16.2.1 void SendSoftwareTrigger ();

*/

// SoftwareTrigger provides the interface required for the
// IviFgenSoftwareTrigger extension group.
type SoftwareTrigger interface {
	SendStartSoftwareTrigger() error
}
