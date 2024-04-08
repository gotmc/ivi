// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package pmx implements the IVI driver for the KIKUSUI PMX series of
regulated DC power supplies.

State Caching: Not implemented
*/
package pmx

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

const (
	specMajorVersion = 4
	specMinorVersion = 4
	specRevision     = "3.0"
)

// Confirm that the device driver implements the IviDCPwrBase interface.
var _ dcpwr.Base = (*Device)(nil)

// Device provides the IVI driver for the Kikusui PMX series of DC power
// supplies.
type Device struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new PMX IVI Instrument. Currently, only the E3631A
// model is supported, but in the future as other models are added, the New
// function will query the instrument to determine the model and ensure it is
// one of the supported models. If reset is true, then the instrument is reset.
func New(inst ivi.Instrument, reset bool) (*Device, error) {
	channelNames := []string{
		"DCOutput",
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
			"PMX18-2A",
			"PMX18-5A",
			"PMX35-1A",
			"PMX35-3A",
			"PMX70-1A",
			"PMX110-0.6A",
			"PMX250-0.25A",
			"PMX350-0.2A",
			"PMX500-0.2A",
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

// ChannelCount returns the number of available output channels.
//
// ChannelCount is the getter for the read-only IviDCPwrBase Attribute Output
// Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (dev *Device) ChannelCount() int {
	return len(dev.Channels)
}
