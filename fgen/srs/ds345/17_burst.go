// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"context"

	"github.com/gotmc/query"
)

// BurstCount returns the number of waveform cycles that the function generator
// produces after it receives a trigger.
//
// BurstCount is the getter for the read-write IviFgenBurst Attribute Burst
// Count described in Section 17.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) BurstCount(ctx context.Context) (int, error) {
	return query.Int(ctx, ch.inst, "BCNT?")
}

// SetBurstCount sets the number of waveform cycles that the function generator
// produces after it receives a trigger.
//
// SetBurstCount is the setter for the read-write IviFgenBurst Attribute Burst
// Count described in Section 17.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetBurstCount(ctx context.Context, count int) error {
	return ch.inst.Command(ctx, "BCNT %d", count)
}
