// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package infiniivision implements the IVI driver for various Keysight
oscilloscopes.

State Caching: Not implemented
*/
package infiniivision

import (
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
)

const (
	specMajorVersion   = 4
	specMinorVersion   = 1
	specRevision       = "4.1"
	defaultGPIBAddress = 0
	defaultResetDelay  = 500 * time.Millisecond
	defaultClearDelay  = 500 * time.Millisecond
)

// Confirm the implemented interfaces by the driver.
var _ scope.Base = (*Driver)(nil)
var _ scop3.BaseChannel = (*Channel)(nil)

// Driver provides the IVI driver for a Keysigh InfiniiVision family of
// oscilloscopes.
type Driver struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// Channel models the output channel repeated capability for the function
// generator output channel.
type Channel struct {
	inst ivi.Instrument
	name string
	num  int
}

// New creates a new InfiniiVision IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Driver, error) {
	channelNames := []string{
		"Output",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)

	for i, channelName := range channelNames {
		ch := Channel{
			name: channelName,
			inst: inst,
			num:  i + 1,
		}
		channels[i] = ch
	}

	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            defaultResetDelay,
		ClearDelay:            defaultClearDelay,
		// Commented out GroupCapabilities still need to be added.
		GroupCapabilities: []string{
			// "IviFgenArbFrequency",
			// "IviFgenArbSeq",
			// "IviFgenArbWaveform",
			"IviFgenBase",
			"IviFgenBurst",
			// "IviFgenInternalTrigger",
			// "IviFgenModulateFM",
			// "IviFgenModulateAM",
			// "IviFgenSoftwareTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{
			"DS345",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"RS232",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}

	if reset {
		if err := driver.Reset(); err != nil {
			return &driver, err
		}
		// Default to internal trigger instead of single trigger when reset.
		if err := driver.inst.Command("TSRC1"); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}

// SerialConfig lists whether the RS-232 serial port is configured as a DCE
// (Data Circuit-Terminating Equipment) or a DTE (Data Terminal Equipment). Computers
// running the IVI program are DTEs; therefore, use a straight through serial
// cable when connecting to DCEs and a null modem cable when connecting to DTEs.
func SerialConfig() string {
	return "DCE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{19200, 9600, 4800, 2400, 1200, 600, 300}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return defaultSerialBuadRate
}

// SerialDataFrames lists the available RS-232 data frame formats. The DS345
// "always sends two stop bits, 8 data bits, and no parity, and will correctly
// receive data sent with eitehr one or two stop bits." per the User's Manual.
func SerialDataFrames() []string {
	return []string{"8N2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
