// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 14 IviDmmDeviceInfo Extension Group

## Section 14.1 IviDmmDeviceInfo Overview

The IviDmmDeviceInfo extension group defines a set of read-only attributes for
DMMs that can be queried to determine how they are presently configured. These
attributes are the aperture time and the aperture time units.

## Section 14.2 IviDmmDeviceInfo Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|  14.2.1 | Aperture Time           | Real64   | RO     | N/A       |
|  14.2.2 | Aperture Time Units     | Int32    | RO     | N/A       |

## Section 14.3 IviDmmDeviceInfo Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// DeviceInfoExtension provides the interface required for the IviDmmDeviceInfo
// extension group described in Section 14 of IVI-4.2 IviDmm Class
// Specification.
type DeviceInfoExtension interface {
	ApertureTime() (float64, error)
	ApertureTimeUnits() (ApertureTimeUnits, error)
}
