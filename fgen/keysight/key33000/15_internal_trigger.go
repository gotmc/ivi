// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33000

import (
	"context"

	"github.com/gotmc/query"
)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
//
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) InternalTriggerRate(ctx context.Context) (float64, error) {
	// The 33500B stores the burst internal period in seconds; the IVI API
	// expects the number of triggers per second. Therefore, we need the inverse.
	per, err := query.Float64(ctx, d.inst, d.channels[0].srcPrefix()+"BURS:INT:PER?")
	if err != nil {
		return 0.0, err
	}

	return 1 / per, nil
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
//
// SetInternalTriggerRate is the setter for the read-write
// IviFgenInternalTrigger Attribute Internal Trigger Rate described in Section
// 15.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) SetInternalTriggerRate(ctx context.Context, rate float64) error {
	// Convert rate (Hz) to period (seconds).
	return d.inst.Command(ctx, d.channels[0].srcPrefix()+"BURS:INT:PER %v", 1/rate)
}
