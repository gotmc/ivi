// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "github.com/gotmc/ivi"

// Channel models a generic FGen channel
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// NewChannel returns a Channel for a function generator.
func NewChannel(id int, name string, inst ivi.Instrument) Channel {
	return Channel{id, name, inst}
}

// Set writes the format string, using the given paarameters to the channel.
func (ch *Channel) Set(format string, a ...interface{}) error {
	return ivi.Set(ch.inst, format, a...)
}

// QueryBool queries the channel and returns a bool.
func (ch *Channel) QueryBool(query string) (bool, error) {
	return ivi.QueryBool(ch.inst, query)
}

// QueryFloat64 queries the channel and returns a float64.
func (ch *Channel) QueryFloat64(query string) (float64, error) {
	return ivi.QueryFloat64(ch.inst, query)
}

// QueryInt queries the channel and returns an int.
func (ch *Channel) QueryInt(query string) (int, error) {
	return ivi.QueryInt(ch.inst, query)
}

// QueryString queries the channel and returns a string.
func (ch *Channel) QueryString(query string) (string, error) {
	return ivi.QueryString(ch.inst, query)
}
