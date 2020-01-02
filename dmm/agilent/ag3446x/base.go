// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ag3446x

import "github.com/gotmc/ivi/dmm"

// MeasurementFunction needs a better comment.
func (d *Ag3446x) MeasurementFunction() (dmm.MeasurementFunction, error) {
	f, err := d.QueryString("CONF?\n")
	return dmm.DCVolts, nil
}
