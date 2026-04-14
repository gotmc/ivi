// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import "github.com/gotmc/query"

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
func (d *Driver) InternalTriggerRate() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	// The Agilent 33220A stores burst internal period in seconds; the IVI API
	// expects triggers per second. Therefore, we need the inverse.
	per, err := query.Float64(ctx, d.inst, "BURS:INT:PER?")
	if err != nil {
		return 0.0, err
	}

	return 1 / per, nil
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
func (d *Driver) SetInternalTriggerRate(rate float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	// Convert rate (Hz) to period (seconds).
	return d.inst.Command(ctx, "BURS:INT:PER %v", 1/rate)
}
