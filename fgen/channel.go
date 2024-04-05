// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

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
func (ch *Channel) QueryBool(cmd string) (bool, error) {
	return query.Bool(ch.inst, cmd)
}

// QueryFloat64 queries the channel and returns a float64.
func (ch *Channel) QueryFloat64(cmd string) (float64, error) {
	return query.Float64(ch.inst, cmd)
}

// QueryInt queries the channel and returns an int.
func (ch *Channel) QueryInt(cmd string) (int, error) {
	return query.Int(ch.inst, cmd)
}

// QueryString queries the channel and returns a string.
func (ch *Channel) QueryString(cmd string) (string, error) {
	return query.String(ch.inst, cmd)
}
