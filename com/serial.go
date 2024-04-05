// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
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
	Equipment  EquipClass
	BaudRates  []int
	DataFrames []DataFrame
}

// DataFrame models an RS-232 data frame format.
type DataFrame struct {
	DataBits int
	Parity   Parity
	StopBits int
}

// Parity defines the serial port parity setting.
type Parity int

// Enum of available parity modes.
const (
	NoParity Parity = iota
	OddParity
	EvenParity
)

// EquipClass defines whether the device is a DCE (Data Circuit-Terminating
// Equipment) or a DTE (Data Terminal Equipment). Note, the computer running
// the IVI software is a DTE. DTE-DTE and DCE-DCE connections require a null
// modem cable.
type EquipClass int

// Enum of available equipment classes.
const (
	DCE EquipClass = iota
	DTE
)
