// Copyright (c) 2017-2021 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package com provides information about communication ports (Ethernet, serial,
etc.) for various IVI devices.
*/
package com

// SerialModes returns the allowed serial communication modes for a device.
type SerialModes struct {
	BaudRates []int
	DataBits  int
	Parity    Parity
	StopBits  int
}

// Parity defines the serial port parity setting.
type Parity int

// Enum of available parity modes.
const (
	NoParity Parity = iota
	OddParity
	EvenParity
)
