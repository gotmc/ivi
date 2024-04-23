// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key33220 implements the IVI driver for the Keysight/Agilent 33220A and
33210A function/arbitrary waveform generators.

State Caching: Not implemented
*/
package key33220

import (
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

const (
	specMajorVersion   = 4
	specMinorVersion   = 3
	specRevision       = "5.2"
	defaultGPIBAddress = 10
	telnetPort         = 5024
	socketPort         = 5025
)

// Confirm the driver implements the interface for the IviFgenBase capability
// group.
var _ fgen.Base = (*Driver)(nil)

// Driver provides the IVI driver for a Keysight/Agilent 33220A or 33210A
// function generator.
type Driver struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 33210A and 33220A
// function/arbitrary waveform generators.
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
		GroupCapabilities: []string{
			"IviFgenBase",
			// "IviFgenArbFrequency",
			// "IviFgenArbWfm",
			"IviFgenBurst",
			"IviFgenInternalTrigger",
			// "IviFgenModulateAM",
			// "IviFgenModulateFM",
			// "IviFgenSoftwareTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{
			"33220A",
			"33210A",
		},
		SupportedBusInterfaces: []string{
			"TCPIP",
			"GPIB",
			"USB",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
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

// AvailableCOMPorts lists the available COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "LAN", "USB"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}

// LANPorts returns a map of the different ports with the key being the type of
// port.
func LANPorts() map[string]int {
	return map[string]int{
		"telnet": telnetPort,
		"socket": socketPort,
	}
}

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviFgenBase Attribute Output
// Count described in Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputCount() int {
	return len(d.Channels)
}

// OutputMode returns the determines how the function generator produces
// waveforms. This attribute determines which extension group’s functions and
// attributes are used to configure the waveform the function generator
// produces.
//
// OutputMode is the getter for the read-only IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) OutputMode() (fgen.OutputMode, error) {
	var outputMode fgen.OutputMode

	funcType, err := query.String(d.inst, "FUNC?")
	if err != nil {
		return outputMode, fmt.Errorf("error determining the output function type: %w", err)
	}

	switch funcType {
	case "SIN", "SQU", "RAMP":
		return fgen.OutputModeFunction, nil
	case "NOIS":
		return fgen.OutputModeNoise, nil
	case "USER":
		return fgen.OutputModeArbitrary, nil
	}

	return 0, fmt.Errorf("unknown output mode type")
}

// SetOutputMode sets how the function generator produces waveforms. This
// attribute determines which extension group’s functions and attributes are
// used to configure the waveform the function generator produces.
//
// OutputMode is the setter for the read-only IviFgenBase Attribute Output
// Mode described in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) SetOutputMode(outputMode fgen.OutputMode) error {
	switch outputMode {
	case fgen.OutputModeFunction:
		return d.inst.Command("FUNC SIN")
	case fgen.OutputModeArbitrary:
		return d.inst.Command("FUNC USER")
	case fgen.OutputModeSequence:
		return fmt.Errorf("function generator does not support output mode sequency")
	case fgen.OutputModeNoise:
		return d.inst.Command("FUNC NOIS")
	}

	return fmt.Errorf("error setting output mode")
}

// InitiateGeneration initiates signal generation by enabling all outputs.
// Instead of calling this function, the user can simply enable outputs.
//
// InitiateGeneration implements the IviFgenBase function described in Section
// 4.3.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) InitiateGeneration() error {
	for _, channel := range d.Channels {
		if err := channel.EnableOutput(); err != nil {
			return err
		}
	}

	return nil
}

// AbortGeneration aborts a previously initiated signal generation by disabling
// all outputs.
//
// AbortGeneration implements the IviFgenBase function described in Section 4.3.1
// of IVI-4.3: IviFgen Class Specification.
func (d *Driver) AbortGeneration() error {
	for _, channel := range d.Channels {
		if err := channel.DisableOutput(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Driver) ReferenceClockSource() (fgen.ClockSource, error) {
	return fgen.RefClockInternal, nil
}

func (d *Driver) SetReferenceClockSource(_ fgen.ClockSource) error {
	return nil
}

func (ch *Channel) Name() string {
	return "output"
}
