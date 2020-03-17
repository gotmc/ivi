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
