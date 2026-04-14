// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key35670

import (
	"fmt"

	"github.com/gotmc/query"
)

// SetStartFreq sets the start frequency in Hertz.
func (dev *Key35670) SetStartFreq(freq float64) error {
	ctx, cancel := dev.newContext()
	defer cancel()

	if freq < 0.0 || freq > 114999.9023 {
		return fmt.Errorf("start frequency of %f out of range", freq)
	}
	return dev.inst.Command(ctx, "sens:freq:star %f", freq)
}

// StartFreq queries the start frequency in Hertz.
func (dev *Key35670) StartFreq() (float64, error) {
	ctx, cancel := dev.newContext()
	defer cancel()

	return query.Float64(ctx, dev.inst, "sens:freq:star?")
}

// SetStartStopFreq sets the start and stop frequency in Hertz.
