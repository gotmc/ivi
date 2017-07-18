// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import "github.com/gotmc/ivi"

func (ch *Channel) Set(cmd string, value float64) error {
	return ivi.SetFloat64(ch.inst, cmd, value)
}

func (ch *Channel) setInt(cmd string, value int) error {
	return ivi.SetInt(ch.inst, cmd, value)
}

func (ch *Channel) queryBool(query string) (bool, error) {
	return ivi.QueryBool(ch.inst, query)
}

func (ch *Channel) queryFloat64(query string) (float64, error) {
	return ivi.QueryFloat64(ch.inst, query)
}

func (ch *Channel) queryInt(query string) (int, error) {
	return ivi.QueryInt(ch.inst, query)
}

func (ch *Channel) queryString(query string) (string, error) {
	return ivi.QueryString(ch.inst, query)
}
