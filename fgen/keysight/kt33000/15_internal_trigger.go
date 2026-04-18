// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33000

import (
	"github.com/gotmc/query"
)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
//
// Deviation from IVI-4.3: The IVI specification (Section 15.2.1) defines
// InternalTriggerRate as a driver-level attribute. We implement it on the
// channel because the 33500B/33600A series have independent burst periods per
// channel via the SOURce[1|2]:BURSt:INTernal:PERiod command.
func (ch *Channel) InternalTriggerRate() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	// The 33500B stores the burst internal period in seconds; the IVI API
	// expects the number of triggers per second. Therefore, we need the inverse.
	per, err := query.Float64(ctx, ch.inst, ch.srcPrefix()+"BURS:INT:PER?")
	if err != nil {
		return 0.0, err
	}

	return 1 / per, nil
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
//
// Deviation from IVI-4.3: See InternalTriggerRate.
func (ch *Channel) SetInternalTriggerRate(rate float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	// Convert rate (Hz) to period (seconds).
	return ch.inst.Command(ctx, ch.srcPrefix()+"BURS:INT:PER %v", 1/rate)
}
