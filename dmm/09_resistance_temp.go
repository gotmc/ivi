// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 9 IviDmmResistanceTemperatureDevice Extension Group

## Section 9.1 IviDmmResistanceTemperatureDevice Overview

The IviDmmResistanceTemperatureDevice extension group supports DMMs that take
temperature measurements using a resistance temperature device (RTD) transducer
type.

The IviDmm class assumes that you are using a Platinum Resistance Temperature
Device.

## Section 9.2 IviDmmResistanceTemperatureDevice Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|   9.2.1 | RTD Alpha               | Real64   | R/W    | N/A       |
|   9.2.2 | RTD Resistance          | Real64   | R/W    | N/A       |

## Section 9.3 IviDmmResistanceTemperatureDevice Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

9.3.1 void Temperature.Rtd.Configure(Double alpha, Double resistance);

*/

// RTDExtension provides the interface required for the
// IviDmmResistanceTemperatureDevice extension group described in Section 9 of
// IVI-4.2 IviDmm Class Specification.
type RTDExtension interface {
	RTDAlpha() (float64, error)
	SetRTDAlpha(alpha float64) error
	RTDResistance() (float64, error)
	SetRTDResistance(resistance float64) error
	ConfigureRTD(alpha, resistance float64) error
}
