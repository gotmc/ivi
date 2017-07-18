// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36xx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotmc/ivi/dcpwr"
)

type voltageCurrent int

const (
	voltageQuery voltageCurrent = iota
	currentQuery
)

type Channel struct {
	dcpwr.Channel
}

func (ch *Channel) queryLimit(query voltageCurrent) (float64, error) {
	cmd := fmt.Sprintf("APPL? %s", ch)
	s, err := ch.QueryString(cmd)
	if err != nil {
		return 0.0, err
	}
	ret := strings.Split(s, ",")
	return strconv.ParseFloat(ret[query], 64)
}
