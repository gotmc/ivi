// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"github.com/gotmc/ivi/fgen"
)

// Make sure that the ds345 driver implements the IviFgenBurst capability
// group.
var _ fgen.Burst = (*DS345)(nil)
var _ fgen.BurstChannel = (*Channel)(nil)

// BurstCount returns the number of waveform cycles that the function generator
// produces after it receives a trigger.  BurstCount is the getter for the
// read-write IviFgenBurst Attribute Burst Count described in Section 17.2.1 of
// IVI-4.3: IviFgen Class Specification.
func (ch *Channel) BurstCount() (int, error) {
	return ch.QueryInt("BCNT?\n")
}

// SetBurstCount sets the number of waveform cycles that the function generator
// produces after it receives a trigger.  SetBurstCount is the setter for the
// read-write IviFgenBurst Attribute Burst Count described in Section 17.2.1 of
// IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetBurstCount(count int) error {
	return ch.Set("BCNT %d\n", count)
}
