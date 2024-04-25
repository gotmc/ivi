// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 12 IviDmmTriggerSlope Capability Group

## Section 12.1 IviDmmTriggerSlope Overview

The IviDmmTriggerSlope extension group supports DMMs that can specify the
polarity of the external trigger signal. It defines an attribute and a function
to configure this polarity.

Typically the specific driver disables extension groups that the application
program does not explicitly use or enable. The IviDmmTriggerSlope extension
capability group affects the behavior of the instrument regardless of the value
of the Trigger Slope attribute. It is not possible for the specific driver to
disable this extension capability group.

Therefore, it is the responsibility of the user to ensure that the slope of the
trigger signal is correct for their application. Most DMMs do not have a
programmable trigger slope and do not implement the IviDmmTriggerSlope
extension capability group. Users should avoid using the IviDmmTriggerSlope
extension capability in their test program source code so that they can
maximize the set of instruments that they can use interchangeably.

For instrument drivers that implement the IviDmmTriggerSlope extension, the
user can set the value of the Trigger Slope attribute in the IVI configuration
file. For instruments that do not implement the IviDmmTriggerSlope extension
group, the user must ensure that trigger signal that their instrument receives
has the correct polarity. This can be done with external circuitry.

## Section 12.2 IviDmmTriggerSlope Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|  12.2.1 | Trigger Slope           | Int32    | R/W    | N/A       |

## Section 12.3 IviDmmTriggerSlope Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/
