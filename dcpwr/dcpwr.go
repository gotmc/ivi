// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package dcpwr provides the Defined Values and other structs, methods, etc.
that are common among all intstruments meeting the IVI-4.4: IviDCPwr Class
Specification.

Files are split based on the class capability groups.
*/
package dcpwr

import "errors"

// Base provides the interface required for the IviDCPwrBase capability group.
type Base interface {
	ChannelCount() int
}

// Error codes related to the IviDCPwr Class Specification.
var (
	ErrNotImplemented     error = errors.New("not implemented in ivi driver")
	ErrOVPUnsupported     error = errors.New("ovp not supported")
	ErrTriggerNotSoftware error = errors.New("trigger source is not set to software trigger.")
)

// CommType defines the available types of communication for a DC power supply.
type CommType int

// Available communiction interfaces for remote communction of a DC power
// supply.
const (
	GPIB CommType = iota
	RS232
	USB
	LAN
)
