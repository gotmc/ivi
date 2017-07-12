// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilente36xx

import "strings"

func (dcpwr *AgilentE36xx) InstrumentManufacturer() (string, error) {
	s, err := dcpwr.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	ret := strings.Split(s, ",")
	return ret[0], nil
}
