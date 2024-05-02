// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 5 IviDmmACMeasurement Extension Group

## Section 5.1 IviDmmACMeasurement Overview

The IviDmmACMeasurement extension group supports DMMs that take AC voltage or
AC current measurements. It defines attributes that configure additional
settings for AC measurements. These attributes are the minimum and maximum
frequency components of the input signal. This extension group also defines
functions that configure these attributes.


## Section 5.2 IviDmmACMeasurement Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute           | Type     | Access | AppliesTo |
| ------- | ----------------    | ------   | ------ | --------- |
|   5.2.1 | AC Max Freq         | Real64   | R/W    | N/A       |
|   5.2.2 | AC Min Freq         | Real64   | R/W    | N/A       |

## Section 5.3 IviDmmACMeasurement Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

5.3.1 void AC.ConfigureBandwidth (Double MinFreq, Double MaxFreq);

*/

// ACMeasurementExtension provides the interface required for the
// IviDmmACMeasurement extension group described in Section 5 of IVI-4.2 IviDmm
// Class Specification.
type ACMeasurementExtension interface {
	MaxACFrequency() (float64, error)
	SetMaxACFrequency(maxFreq float64) error
	MinACFrequency() (float64, error)
	SetMinACFrequency(minFreq float64) error
	ConfigureACBandwidth(minFreq, maxFreq float64) error
}
