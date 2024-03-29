// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"github.com/gotmc/ivi/fgen"
)

// Make sure that the key33220 driver implements the IviFgenInternalTrigger
// capability group.
var _ fgen.IntTriggerChannel = (*Channel)(nil)

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) InternalTriggerRate() (float64, error) {
	// The Agilent 33220A needs to know the burst period in seconds; however, the
	// IVI API expects the number of triggers per second. Therefore, we need the
	// inverse.
	per, err := ch.QueryFloat64("BURS:INT:PER?\n")
	if err != nil {
		return 0.0, err
	}
	return 1 / per, nil
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
// SetInternalTriggerRate is the setter for the read-write
// IviFgenInternalTrigger Attribute Internal Trigger Rate described in Section
// 15.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetInternalTriggerRate(rate float64) error {
	// The Agilent 33220A needs to know the burst period in seconds; however, the
	// IVI API expects the number of triggers per second. Therefore, we need the
	// inverse.
	return ch.Set("BURS:INT:PER %v\n", 1/rate)
}

// InternalTriggerPeriod determines the internal trigger period in seconds.
// InternalTriggerPeriod is not part of the IVI API, which only provides
// InternalTriggerRate, but this is a convenience function.
func (ch *Channel) InternalTriggerPeriod() (float64, error) {
	return ch.QueryFloat64("BURS:INT:PER?\n")
}

// SetInternalTriggerPeriod specifies the internal trigger period in seconds.
// SetInternalTriggerPeriod is not part of the IVI API, which only provides
// SetInternalTriggerRate, but this is a convenience function.
func (ch *Channel) SetInternalTriggerPeriod(period float64) error {
	return ch.Set("BURS:INT:PER %v\n", period)
}
