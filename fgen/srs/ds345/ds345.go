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
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion = 4
	specMinorVersion = 3
	specRevision     = "5.2"
)

// DS345 provides the IVI driver for a SRS DS345 function generator.
type DS345 struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new DS345 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*DS345, error) {
	channelNames := []string{
		"Output",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		baseChannel := fgen.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenStdfunc",
			"IviFgenTrigger",
			"IviFgenInternalTrigger",
			"IviFgenBurst",
		},
		SupportedInstrumentModels: []string{
			"DS345",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := DS345{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		if err := driver.Reset(); err != nil {
			return &driver, err
		}
		// Default to internal trigger instead of single trigger.
		if _, err := driver.inst.WriteString("TSRC1\n"); err != nil {
			return &driver, err
		}
		return &driver, nil
	}
	return &driver, nil
}

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	fgen.Channel
}

// AvailableCOMPorts lists the avaialble COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "RS232"}
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

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
