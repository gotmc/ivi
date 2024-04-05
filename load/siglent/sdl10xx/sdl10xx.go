// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package sdl10xx implements the IVI driver for the Signlent SDL1000X and
SDL1030X DC electronic loads.

State Caching: Not implemented
*/
package sdl10xx

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/load"
)

// SDL10xx provides the IVI driver for the Siglent SDL1000X and SDL1030X DC
// Electronic Loads.
type SDL10xx struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new Siglent SDL10xx IVI Instrument driver.
func New(inst ivi.Instrument, reset bool) (*SDL10xx, error) {
	// FIXME(mdr): Need to query the instrument to determine the model and then
	// set any model specific attributes, such as quantity and names of channels.
	channelNames := []string{
		"Input",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		baseChannel := load.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 0,
		ClassSpecMinorVersion: 0,
		ClassSpecRevision:     "N/A",
		GroupCapabilities:     []string{},
		SupportedInstrumentModels: []string{
			"SDL1000X",
			"SDL1030X",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := SDL10xx{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// AvailableCOMPorts lists the avaialble COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"RS232", "USB", "LAN", "GPIB"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return 5
}

// SerialConfig lists whether the RS-232 serial port is configured as a DCE
// (Data Circuit-Terminating Equipment) or a DTE (Data Terminal Equipment). Computers
// running the IVI program are DTEs; therefore, use a straight through serial
// cable when connecting to DCEs and a null modem cable when connecting to DTEs.
func SerialConfig() string {
	// FIXME: What's the right answer for the SDL10x0X?:w
	return "DTE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{4800, 9600, 19200, 38400, 57600, 115200}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return 115200
}

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2", "7E2", "7O2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
