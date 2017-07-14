// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilente36xx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// String implements the stringer interface for channel.
func (ch *Channel) String() string {
	return ch.name
}

func (ch *Channel) setFloat64(cmd string, value float64) error {
	return ivi.SetFloat64(ch.inst, cmd, value)
}

func (ch *Channel) queryBool(query string) (bool, error) {
	return ivi.QueryBool(ch.inst, query)
}

func (ch *Channel) queryFloat64(query string) (float64, error) {
	return ivi.QueryFloat64(ch.inst, query)
}

func (ch *Channel) queryString(query string) (string, error) {
	return ivi.QueryString(ch.inst, query)
}

func (ch *Channel) queryLimit(query dcpwr.VoltageCurrent) (float64, error) {
	cmd := fmt.Sprintf("APPL? %s", ch.name)
	s, err := ch.inst.Query(cmd)
	if err != nil {
		return 0.0, err
	}
	ret := strings.Split(s, ",")
	return strconv.ParseFloat(ret[query], 64)
}
