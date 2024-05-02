// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 6 IviDmmFrequencyMeasurement Extension Group

## Section 6.1 IviDmmFrequencyMeasurement Overview

The IviDmmFrequencyMeasurement extension group supports DMMs that take
frequency measurements. It defines attributes that are required to configure
additional parameters needed for frequency measurements.

## Section 6.2 IviDmmFrequencyMeasurement Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute               | Type     | Access | AppliesTo |
| ------- | ----------------------- | ------   | ------ | --------- |
|   6.2.1 | Frequency Voltage Range | Real64   | R/W    | N/A       |
|   6.2.2 | Freq Voltage Range Auto | Bool     | R/W    | N/A       |

## Section 6.3 IviDmmFrequencyMeasurement Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

// FrequencyMeasurementExtension provides the interface required for the
// IviDmmFrequencyMeasurement extension group described in Section 6 of IVI-4.2
// IviDmm Class Specification.
type FrequencyMeasurementExtension interface {
	FrequencyVoltageRange() (autoRange bool, rangeValue float64, err error)
	SetFrequencyVoltageRange(autoRange bool, rangeValue float64) error
}
