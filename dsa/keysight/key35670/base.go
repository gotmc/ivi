// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key35670

import (
	"context"
	"fmt"

	"github.com/gotmc/query"
)

// SetStartFreq sets the start frequency in Hertz.
func (dev *Key35670) SetStartFreq(ctx context.Context, freq float64) error {
	if freq < 0.0 || freq > 114999.9023 {
		return fmt.Errorf("start frequency of %f out of range", freq)
	}
	_, err := fmt.Fprintf(dev.inst, "sens:freq:star %f", freq)
	return err
}

// StartFreq queries the start frequency in Hertz.
func (dev *Key35670) StartFreq(ctx context.Context) (float64, error) {
	return query.Float64(ctx, dev.inst, "sens:freq:star?")
}

// SetStartStopFreq sets the start and stop frequency in Hertz.
