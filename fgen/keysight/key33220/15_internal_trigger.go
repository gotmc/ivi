// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"github.com/gotmc/query"
)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second (Hz).
//
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) InternalTriggerRate() (float64, error) {
	// The Agilent 33220A needs to know the burst period in seconds; however, the
	// IVI API expects the number of triggers per second. Therefore, we need the
	// inverse.
	per, err := query.Float64(d.inst, "BURS:INT:PER?")
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
func (d *Driver) SetInternalTriggerRate(rate float64) error {
	// The Agilent 33220A needs to know the burst period in seconds; however, the
	// IVI API expects the number of triggers per second. Therefore, we need the
	// inverse.
	return d.inst.Command("BURS:INT:PER %v", 1/rate)
}
