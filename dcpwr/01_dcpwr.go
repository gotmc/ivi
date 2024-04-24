// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package dcpwr provides the defined values, interfaces, structs, methods, and
enums that are common among all intstruments meeting the IVI-4.4: IviDCPwr
Class Specification.

Files are split based on the class capability groups listed in Table 2-1
IviDCPwr Group Names in the IVI-4.4 IviDCPwr Class Specification.
*/
package dcpwr

import "errors"

// Error codes related to the IviDCPwr Class Specification.
var (
	ErrNotImplemented     = errors.New("not implemented in ivi driver")
	ErrOVPUnsupported     = errors.New("ovp not supported")
	ErrTriggerNotSoftware = errors.New("trigger source is not set to software trigger")
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
