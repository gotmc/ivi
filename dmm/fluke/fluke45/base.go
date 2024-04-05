// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fluke45

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi/dmm"
)

// MeasurementFunction needs a better comment.
func (d *DMM) MeasurementFunction() (dmm.MeasurementFunction, error) {
	fcn, err := d.QueryString("FUNC1?\n")
	if err != nil {
		return 0, err
	}
	switch strings.TrimSpace(fcn) {
	case "VDC":
		return dmm.DCVolts, nil
	case "VAC":
		return dmm.ACVolts, nil
	case "ADC":
		return dmm.DCCurrent, nil
	case "AAC":
		return dmm.ACCurrent, nil
	case "OHMS":
		return dmm.TwoWireResistance, nil
	case "VACDC":
		return dmm.ACPlusDCVolts, nil
	case "AACDC":
		return dmm.ACPlusDCCurrent, nil
	default:
		return 0, fmt.Errorf("%s is not a valid measurement function", fcn)
	}
}
