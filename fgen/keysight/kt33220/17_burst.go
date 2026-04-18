// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33220

import "github.com/gotmc/query"

// BurstCount returns the number of waveform cycles that the function generator
// produces after it receives a trigger.
func (ch *Channel) BurstCount() (int, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Int(ctx, ch.inst, "BURS:NCYC?")
}

// SetBurstCount sets the number of waveform cycles that the function generator
// produces after it receives a trigger.
func (ch *Channel) SetBurstCount(count int) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "BURS:NCYC %d", count)
}
