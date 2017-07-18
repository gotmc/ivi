// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kikusuipmx

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

type Channel struct {
	id                   int
	name                 string
	inst                 ivi.Instrument
	currentLimitBehavior dcpwr.CurrentLimitBehavior
}

// String implements the stringer interface for channel.
func (ch *Channel) String() string {
	return ch.name
}

func (ch *Channel) Set(format string, a ...interface{}) error {
	return ivi.Set(ch.inst, format, a...)
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
