// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key33220 implements the IVI Instrument driver for the Agilent/Keysight
33220A function generator.

State Caching: Not implemented
*/
package key33220

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion = 4
	specMinorVersion = 3
	specRevision     = "5.2"
)

// Confirm that the device driver implements the IviFgenBase interface.
var _ fgen.Base = (*Device)(nil)

// Device provides the IVI driver for a Keysight/Agilent 33220A or 33210A
// function generator.
type Device struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new Key33220 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Device, error) {
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
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenStdfunc",
			"IviFgenTrigger",
			"IviFgenInternalTrigger",
			"IviFgenBurst",
		},
		SupportedInstrumentModels: []string{
			"33220A",
			"33210A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	device := Device{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := device.Reset()
		return &device, err
	}
	return &device, nil
}

// AvailableCOMPorts lists the avaialble COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "LAN", "USB"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return 10
}

// LANPorts returns a map of the different ports with the key being the type of
// port.
func LANPorts() map[string]int {
	return map[string]int{
		"telnet": 5024,
		"socket": 5025,
	}
}

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviFgenBase Attribute Output
// Count described in Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (a *Device) OutputCount() int {
	return len(a.Channels)
}
