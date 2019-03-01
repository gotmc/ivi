// Copyright (c) 2017-2019 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

import "github.com/gotmc/ivi"

// Channel represents the repeated capability of an output channel for a DC
// power supply.
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// NewChannel returns a Channel for a DC power supply.
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
func (ch *Channel) QueryBool(query string) (bool, error) {
	return ivi.QueryBool(ch.inst, query)
}

// QueryFloat64 queries the channel and returns a float64.
func (ch *Channel) QueryFloat64(query string) (float64, error) {
	return ivi.QueryFloat64(ch.inst, query)
}

// QueryString queries the channel and returns a string.
// FIXME(mdr): Change to take a format string and ...interface{}
func (ch *Channel) QueryString(query string) (string, error) {
	return ivi.QueryString(ch.inst, query)
}
