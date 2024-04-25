// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 17 IviDmmPowerLineFrequency Capability Group

## Section 17.1 IviDmmPowerLineFrequency Overview

The IviDmmPowerLineFrequency extension supports DMMs with the capability to
specify the power line frequency.

Typically the specific driver disables extension groups that the application
program does not explicitly use or enable. The IviDmmPowerLineFrequency
extension capability group affects the behavior of the instrument regardless of
the value of the Powerline Freq attribute. It is not possible for the specific
driver to disable this extension capability group.

Therefore, it is the responsibility of the user to ensure that the power line
frequency is correct for their application. Most DMMs do not have a
programmable power line frequency and do not implement the
IviDmmPowerLineFrequency extension capability group. Users should avoid using
the IviDmmPowerLineFrequency extension group in their test program source code
so that they can maximize the set of instruments that they can use
interchangeably.

For instrument drivers that implement the IviDmmPowerLineFrequency extension,
the user can set the value of the Powerline Freq attribute in the IVI
configuration file. For instruments that do not implement the
IviDmmPowerLineFrequency extension group, the user must ensure that their
instrument is set to use the correct power line frequency. Users can manually
change the power line frequency setting on most DMMs by means of a switch on
the instrument’s back panel.

## Section 17.2 IviDmmPowerLineFrequency Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|  17.2.1 | Powerline Frequency     | Real64   | R/W    | N/A       |

## Section 17.3 IviDmmPowerLineFrequency Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/
