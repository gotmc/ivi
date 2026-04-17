// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e3600

import (
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

// Measure takes a measurement of the requested type (voltage or current) on
// this channel.
//
// Measure implements the IviDCPwrMeasurement function described in Section
// 7.3.1 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) Measure(msrType dcpwr.MeasurementType) (float64, error) {
	switch msrType {
	case dcpwr.CurrentMeasurement:
		return ch.MeasureCurrent()
	case dcpwr.VoltageMeasurement:
		return ch.MeasureVoltage()
	}

	return 0.0, fmt.Errorf("Measure %v: %w", msrType, ivi.ErrValueNotSupported)
}

// MeasureVoltage takes a measurement on the output signal and returns the
// measured voltage.
//
// MeasureVoltage implements the IviDCPwrMeasurement function Measure for the
// Voltage MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureVoltage() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, "MEAS:VOLT? %s", ch.name)
}

// MeasureCurrent takes a measurement on the output signal and returns the
// measured current.
//
// MeasureCurrent implements the IviDCPwrMeasurement function Measure for the
// Current MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureCurrent() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, "MEAS:CURR? %s", ch.name)
}
