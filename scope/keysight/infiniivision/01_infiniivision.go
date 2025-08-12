// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
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
var _ scope.BaseChannel = (*Channel)(nil)

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
	// FIXME: Need to query the instrument for the model and then determine the
	// number of channels based on the model returned.
	channelNames := []string{
		"CHAN1",
		"CHAN2",
		"CHAN3",
		"CHAN4",
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
	}

	return &driver, nil
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}
