// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package template provides a template to create additional IVI drivers for DC
power supplies.

State Caching: Not implemented
*/
package template

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

// Confirm that the device driver implements the IviDCPwrBase interface.
var _ dcpwr.Base = (*Device)(nil)

const (
	specMajorVersion = 4
	specMinorVersion = 4
	specRevision     = "3.0"
)

// Device provides the IVI driver for the Rigol DP800 series of DC power
// supplies.
type Device struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new DC Power Supply IVI Instrument driver. The New function will
// query the instrument to determine the model and ensure it is one of the
// supported models. If reset is true, then the instrument is reset.
func New(inst ivi.Instrument, reset bool) (*Device, error) {
	channelNames := []string{
		"CH1",
		"CH2",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, channelName := range channelNames {
		ch := Channel{
			name: channelName,
			inst: inst,
		}
		channels[i] = ch
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
			"IviDCPwrTrigger",
		},
		SupportedInstrumentModels: []string{
			"TBD",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Device{
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
	return []string{"GPIB", "RS232", "Ethernet", "USB"}
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

// ChannelCount returns the number of available output channels.
//
// ChannelCount is the getter for the read-only IviDCPwrBase Attribute Output
// Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (dev *Device) ChannelCount() int {
	return len(dev.Channels)
}
