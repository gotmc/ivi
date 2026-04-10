// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package infiniivision implements the IVI driver for various Keysight
// oscilloscopes.
//
// State Caching: Not implemented
package infiniivision

import (
	"context"
	"fmt"
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
	channels []Channel
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
func New(inst ivi.Instrument, idQuery, reset bool) (*Driver, error) {
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
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviScopeBase",
			"IviScopeWaveformMeasurement",
		},
		SupportedInstrumentModels: []string{
			"DSOX3024A",
			"DSOX3034A",
			"MSOX3024A",
			"MSOX3034A",
		},
		SupportedBusInterfaces: []string{
			"USB",
			"GPIB",
			"LAN",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)

	if idQuery {
		if _, err := inherent.CheckID(context.Background()); err != nil {
			return nil, err
		}
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		Inherent: inherent,
	}

	if reset {
		if err := driver.Reset(context.Background()); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// Channel returns the Channel at the given index, with bounds checking.
func (d *Driver) Channel(index int) (*Channel, error) {
	if index < 0 || index >= len(d.channels) {
		return nil, fmt.Errorf("channel %d: %w", index, ivi.ErrChannelNotFound)
	}

	return &d.channels[index], nil
}

// Close properly shuts down the oscilloscope by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}
