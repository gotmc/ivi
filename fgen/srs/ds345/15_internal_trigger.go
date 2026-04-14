// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second. The DS345
// rounds the trigger rate to two significant digits and may range from 0.001
// Hz to 10 kHz.
//
// Deviation from IVI-4.3: The IVI specification (Section 15.2.1) defines
// InternalTriggerRate as a driver-level attribute. We implement it on the
// channel to support multi-channel function generators. The DS345 is
// single-channel, so the behavior is unchanged.
func (ch *Channel) InternalTriggerRate() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "TRAT?")
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second. The DS345
// rounds the trigger rate to two significant digits and may range from 0.001
// Hz to 10 kHz.
//
// Deviation from IVI-4.3: See InternalTriggerRate.
func (ch *Channel) SetInternalTriggerRate(rate float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	// The DS345 supports internal trigger rates from 0.001 Hz to 10 kHz.
	if rate < 0.001 || rate > 10000 {
		return ivi.ErrValueNotSupported
	}

	return ch.inst.Command(ctx, "TRAT %.3f", rate)
}
