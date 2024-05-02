// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 13 IviDmmSoftwareTrigger Extension Group

## Section 13.1 IviDmmSoftwareTrigger Overview

The IviDmmSoftwareTrigger extension group supports DMMs that can initiate a
measurement based on a software trigger signal. The user can send a software
trigger to cause the DMM to initiate a measurement.

## Section 13.2 IviDmmSoftwareTrigger Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

13.2 void SendSoftwareTrigger();

*/

// SoftwareTriggerExtension provides the interface required for the
// IviDmmSoftwareTrigger extension group described in Section 13 of IVI-4.2
// IviDmm Class Specification.
type SoftwareTriggerExtension interface {
	SendSoftwareTrigger() error
}
