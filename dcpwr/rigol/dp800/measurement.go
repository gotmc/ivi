// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dp800

import (
	"fmt"

	"github.com/gotmc/ivi/dcpwr"
)

// Make sure the IviDCPwrMeasurement capability group has been implemented.
var _ dcpwr.Measurement = (*DP800)(nil)
var _ dcpwr.MeasurementChannel = (*Channel)(nil)

// MeasureVoltage takes a measurement on the output signal and returns the
// measured voltage.
//
// MeasureVoltage implements the IviDCPwrMeasurement function Measure for the
// Voltage MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureVoltage() (float64, error) {
	cmd := fmt.Sprintf("MEAS:CURR? %s", ch.Name())
	return ch.QueryFloat64(cmd)
}

// MeasureCurrent takes a measurement on the output signal and returns the
// measured current.
//
// MeasureCurrent implements the IviDCPwrMeasurement function Measure for the
// Current MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureCurrent() (float64, error) {
	cmd := fmt.Sprintf("MEAS? %s", ch.Name())
	return ch.QueryFloat64(cmd)
}
