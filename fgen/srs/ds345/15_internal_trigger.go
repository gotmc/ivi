// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
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
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) InternalTriggerRate() (float64, error) {
	return query.Float64(d.inst, "TRAT?")
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second. The DS345
// rounds the trigger rate to two significant digits and may range from 0.001
// Hz to 10 kHz.
//
// SetInternalTriggerRate is the setter for the read-write
// IviFgenInternalTrigger Attribute Internal Trigger Rate described in Section
// 15.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) SetInternalTriggerRate(rate float64) error {
	// The DS345 supports internal trigger rates from 0.001 Hz to 10 kHz.
	if rate < 0.001 || rate > 10000 {
		return ivi.ErrValueNotSupported
	}

	return d.inst.Command("TRAT %.3f", rate)
}
