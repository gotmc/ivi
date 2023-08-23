// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36xx

import (
	"fmt"

	"github.com/gotmc/convert"
	"github.com/gotmc/ivi/dcpwr"
)

type voltageCurrent int

const (
	voltageQuery voltageCurrent = iota
	currentQuery
)

// Channel models an E36xx DC power supply output channel.
type Channel struct {
	dcpwr.Channel
}

func (ch *Channel) queryLimit(query voltageCurrent) (float64, error) {
	cmd := fmt.Sprintf("APPL? %s", ch)
	s, err := ch.QueryString(cmd)
	if err != nil {
		return 0.0, err
	}
	floats, err := convert.StringToNFloats(s, ",", 2)
	if err != nil {
		return 0.0, err
	}
	return floats[query], nil
}
