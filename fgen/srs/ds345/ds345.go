// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package ds345 implements the IVI driver for the Stanford Research System
DS345 function generator. The serial port mode is 8N2.

State Caching: Not implemented
*/
package ds345

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/com"
	"github.com/gotmc/ivi/fgen"
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
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 3,
		ClassSpecRevision:     "5.2",
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

// SerialConfig returns the allowed configuration for the DS345's serial port.
// Baud rates are listed in order from fastest to slowest.
func SerialConfig() com.SerialModes {
	return com.SerialModes{
		BaudRates: []int{19200, 9600, 4800, 2400, 1200, 600, 300},
		DataBits:  8,
		Parity:    com.NoParity,
		StopBits:  2,
	}
}
