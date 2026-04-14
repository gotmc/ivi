// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package pmx implements the IVI driver for the KIKUSUI PMX series of
// regulated DC power supplies.
//
// State Caching: Not implemented
package pmx

import (
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

const (
	specMajorVersion = 4
	specMinorVersion = 4
	specRevision     = "3.0"
)

// Confirm implemented interfaces by the pmx driver.
var _ dcpwr.Base = (*Driver)(nil)
var _ dcpwr.BaseChannel = (*Channel)(nil)
var _ dcpwr.MeasurementChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Kikusui PMX series of DC power
// supplies.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	ivi.Inherent
}

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	name                 string
	inst                 ivi.Transport
	currentLimitBehavior dcpwr.CurrentLimitBehavior
}

// New creates a new PMX IVI Instrument. Use [ivi.WithIDQuery]
// to verify the instrument model, [ivi.WithReset] to reset on creation, and
// [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

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
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
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
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(); err != nil {
			return nil, err
		}
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(); err != nil {
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

// Close properly shuts down the power supply by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
