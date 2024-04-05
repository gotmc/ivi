// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package load

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

// Channel represents the repeated capability of an output channel for a DC
// power supply.
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// NewChannel returns a Channel for an electronic load.
func NewChannel(id int, name string, inst ivi.Instrument) Channel {
	return Channel{id, name, inst}
}

// Name returns the name of the output channel. Name is the getter for the
// read-only IviDCPwrBase Attribute Output Channel Name described in Section
// 4.2.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) Name() string {
	// TODO(mdr): Should I get rid of the Name getter and instead use the more Go
	// idiomatic Stringer interface?
	return ch.name
}

// String implements the stringer interface for channel.
func (ch *Channel) String() string {
	return ch.name
}

// Set takes the same inputs as fmt.Sprintf() and writes the resultant command
// to the IVI device.
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

// QueryString queries the channel and returns a string.
// FIXME(mdr): Change to take a format string and ...interface{}
func (ch *Channel) QueryString(cmd string) (string, error) {
	return query.String(ch.inst, cmd)
}
