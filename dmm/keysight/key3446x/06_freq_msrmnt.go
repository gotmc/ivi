// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"fmt"

	"github.com/gotmc/query"
)

// FrequencyVoltageRange returns the voltage range used to detect the input
// signal during frequency and period measurements, along with a flag
// indicating whether the instrument is autoranging the voltage input.
//
// FrequencyVoltageRange is the getter for the read-write
// IviDmmFrequencyMeasurement Attribute Frequency Voltage Range described in
// Section 6.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) FrequencyVoltageRange() (bool, float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	autoRange, err := query.Bool(ctx, d.inst, "FREQ:VOLT:RANG:AUTO?")
	if err != nil {
		return false, 0, fmt.Errorf("FrequencyVoltageRange: %w", err)
	}

	rng, err := query.Float64(ctx, d.inst, "FREQ:VOLT:RANG?")
	if err != nil {
		return false, 0, fmt.Errorf("FrequencyVoltageRange: %w", err)
	}

	return autoRange, rng, nil
}

// SetFrequencyVoltageRange configures the voltage range used to detect the
// input signal during frequency and period measurements. When autoRange is
// true, the instrument selects the range automatically and rangeValue is
// ignored; otherwise the driver selects the smallest supported fixed range
// (100 mV, 1 V, 10 V, 100 V, or 750 V) that accommodates rangeValue.
//
// SetFrequencyVoltageRange is the setter for the read-write
// IviDmmFrequencyMeasurement Attribute Frequency Voltage Range described in
// Section 6.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetFrequencyVoltageRange(autoRange bool, rangeValue float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if autoRange {
		return d.inst.Command(ctx, "FREQ:VOLT:RANG:AUTO ON")
	}

	rng, err := determineManualFrequencyVoltageRange(rangeValue)
	if err != nil {
		return fmt.Errorf("SetFrequencyVoltageRange: %w", err)
	}

	// Setting a fixed range implicitly disables autoranging on the
	// instrument, so no separate AUTO OFF command is required.
	return d.inst.Command(ctx, "FREQ:VOLT:RANG %s", rng)
}
