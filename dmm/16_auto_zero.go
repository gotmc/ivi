// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 16 IviDmmAutoZero Extension Group

## Section 16.1 IviDmmAutoZero Overview

The IviDmmAutoZero extension supports DMMs with the capability to take an auto
zero reading. In general, the auto zero capability of a DMM normalizes all
measurements based on a Zero Reading.

## Section 16.2 IviDmmAutoZero Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|  16.2.1 | Auto Zero               | Int32    | R/W    | N/A       |

## Section 16.3 IviDmmAutoZero Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// AutoZeroExtension provides the interface required for the IviDmmAutoZero
// extension group described in Section 16 of IVI-4.2 IviDmm Class
// Specification.
type AutoZeroExtension interface {
	AutoZero() (AutoZero, error)
	SetAutoZero(autoZero AutoZero) error
}
