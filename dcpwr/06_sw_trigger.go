// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

/*

# Section 6 IviDCPwrSoftwareTrigger Extension Group

## Section 6.1 IviDCPwrSoftwareTrigger Overview

The IviDCPwrSoftwareTrigger extension group defines extensions for DC power
supplies capable of changing the output signal based on a software trigger
event.

## Section 6.2 IviDCPwrSoftwareTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

6.2.1 void SendSoftwareTrigger();

*/

// SoftwareTrigger provides the interface required for the
// IviDCPwrSoftwareTrigger extension group.
type SoftwareTrigger interface {
	SendSoftwareTrigger() error
}
