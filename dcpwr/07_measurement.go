// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

/*

# Section 7 IviDCPwrMeasurement Extension Group

## Section 7.1 IviDCPwrMeasurement Overview

The IviDCPwrMeasurement extension group defines extensions for DC power
supplies capable of calculating various measurements such as voltage and
current from the output signal.

## Section 7.2 IviDCPwrMeasurement Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

7.2.1 Double Outputs[].Measure(Ivi.DCPwr.MeasurementType measurementType);

*/

// MeasurementChannel provides the interface for the channel repeated
// capability for the IviDCPwrMeasurement capability group.
type MeasurementChannel interface {
	Measure(msrType MeasurementType) (float64, error)
	MeasureVoltage() (float64, error)
	MeasureCurrent() (float64, error)
}
