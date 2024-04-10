// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// Confirm the driver implements the IviFgenInternalTrigger interface.
var _ fgen.IntTriggerChannel = (*Channel)(nil)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
//
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) InternalTriggerRate() (float64, error) {
	return query.Float64(ch.inst, "TRAT?")
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
//
// SetInternalTriggerRate is the setter for the read-write
// IviFgenInternalTrigger Attribute Internal Trigger Rate described in Section
// 15.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetInternalTriggerRate(rate float64) error {
	return ch.inst.Command("TRAT %.3f", rate)
}
