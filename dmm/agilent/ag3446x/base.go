// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ag3446x

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi/dmm"
)

// MeasurementFunction needs a better comment.
func (d *Ag3446x) MeasurementFunction() (dmm.MeasurementFunction, error) {
	conf, err := d.QueryString("CONF?\n")
	if err != nil {
		return 0, err
	}
	conf = strings.TrimSpace(conf)
	fcnString := strings.Split(conf, " ")[0]
	fcn, ok := dmm.MeasurementFunctionMap[fcnString]
	if !ok {
		return 0, fmt.Errorf("%s is not a valid measurement function", fcnString)
	}
	return fcn, nil
}
