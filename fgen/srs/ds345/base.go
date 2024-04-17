// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package ds345 implements the IVI driver for the Stanford Research System
DS345 function generator.

State Caching: Not implemented
*/
package ds345

import (
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion = 4
	specMinorVersion = 3
	specRevision     = "5.2"
)

// Confirm the driver implements the interface for the IviFgenBase capability
// group.
var _ fgen.Base = (*Driver)(nil)

// Driver provides the IVI driver for a SRS DS345 function generator.
type Driver struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new DS345 IVI Instrument.
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
		}
		channels[i] = ch
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
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
	return 19
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
	return 9600
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

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviFgenBase Attribute Output
// Count described in Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputCount() int {
	return len(d.Channels)
}

// AbortGeneration aborts a previously initiated signal generation. Since the
// DS345 constantly generates a signal and the output cannot be aborted, this
// function does nothing and returns no error.
//
// AbortGeneration implements the IviFgenBase function described in Section 4.3.1
// of IVI-4.3: IviFgen Class Specification.
func (d *Driver) AbortGeneration() error {
	return nil
}
