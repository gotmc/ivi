// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
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

// PMX provides the IVI driver for the Agilent/Keysight E3600 series of power
// supplies.
type PMX struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new PMX IVI Instrument. Currently, only the E3631A
// model is supported, but in the future as other models are added, the New
// function will query the instrument to determine the model and ensure it is
// one of the supported models. If reset is true, then the instrument is reset.
func New(inst ivi.Instrument, reset bool) (*PMX, error) {
	// FIXME(mdr): Need to query the instrument to determine the model and then
	// set any model specific attributes, such as quantity and names of channels.
	channelNames := []string{
		"DCOutput",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		baseChannel := dcpwr.NewChannel(i, ch, inst)
		channels[i] = Channel{
			Channel:              baseChannel,
			currentLimitBehavior: dcpwr.Regulate,
		}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 4,
		ClassSpecRevision:     "3.0",
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
	driver := PMX{
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
