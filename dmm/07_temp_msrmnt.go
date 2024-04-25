// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 7 IviDmmTemperatureMeasurement Capability Group

## Section 7.1 IviDmmTemperatureMeasurement Overview

The IviDmmTemperatureMeasurement extension group supports DMMs that take
temperature measurements with a thermocouple, an RTD, or a thermistor
transducer type. This extension group selects the transducer type. Other
capability groups further configure temperature settings based on the
transducer type.

## Section 7.2 IviDmmTemperatureMeasurement Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | -------- | ------ | --------- |
|   7.2.1 | Temp Transducer Type    | Int32    | R/W    | N/A       |

## Section 7.3 IviDmmTemperatureMeasurement Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// TemperatureMeasurement provides the interface required for the
// IviDmmTemperatureMeasurement extension group described in Section 7 of IVI-4.2
// IviDmm Class Specification.
type TemperatureMeasurement interface {
	TemperatureTransducerType() (TempTransducerType, error)
	SetTemperatureTransducerType(t TempTransducerType) error
}
