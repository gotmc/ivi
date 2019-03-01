// Copyright (c) 2017-2019 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

// InternalTriggerRate determines the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
// InternalTriggerRate is the getter for the read-write IviFgenInternalTrigger
// Attribute Internal Trigger Rate described in Section 15.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) InternalTriggerRate() (float64, error) {
	return ch.QueryFloat64("TRAT?\n")
}

// SetInternalTriggerRate specifies the rate at which the function generator's
// internal trigger source produces a trigger in triggers per second.
// SetInternalTriggerRate is the setter for the read-write
// IviFgenInternalTrigger Attribute Internal Trigger Rate described in Section
// 15.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetInternalTriggerRate(rate float64) error {
	return ch.Set("TRAT %.3f\n", rate)
}
