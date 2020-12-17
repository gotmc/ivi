// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package e36xx implements the IVI driver for the Agilent/Keysight E3600 series
of power supplies.

State Caching: Not implemented
*/
package e36xx

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

// E36xx provides the IVI driver for the Agilent/Keysight E3600 series
// of power supplies.
type E36xx struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new AgilentE36xx IVI Instrument driver. Currently, only the
// E3631A model is supported, but in the future as other models are added, the
// New function will query the instrument to determine the model and ensure it
// is one of the supported models. If reset is true, then the instrument is
// reset.
func New(inst ivi.Instrument, reset bool) (*E36xx, error) {
	// FIXME(mdr): Need to query the instrument to determine the model and then
	// set any model specific attributes, such as quantity and names of channels.
	channelNames := []string{
		"P6V",
		"P25V",
		"N25V",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		baseChannel := dcpwr.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 4,
		ClassSpecRevision:     "3.0",
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
		},
		SupportedInstrumentModels: []string{
			"E3631A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := E36xx{
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
	return []string{"GPIB", "RS232"}
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
	return "DTE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{9600, 4800, 2400, 1200, 600, 300}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return 9600
}

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2", "7E2", "7O2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
