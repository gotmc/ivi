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
	"context"
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
	timeout  time.Duration
	ivi.Inherent
}

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	name                 string
	inst                 ivi.Transport
	timeout              time.Duration
	currentLimitBehavior dcpwr.CurrentLimitBehavior
}

// New creates a new PMX IVI Instrument. By default the constructor queries
// *IDN? and verifies the model against the supported list; pass
// [ivi.WithoutIDQuery] to skip that check. Use [ivi.WithReset] to reset on
// creation and [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	s, err := ivi.NewDriverSetup(inst, ivi.InherentBase{
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
			"PMX18-2A", "PMX18-5A", "PMX35-1A", "PMX35-3A", "PMX70-1A",
			"PMX110-0.6A", "PMX250-0.25A", "PMX350-0.2A", "PMX500-0.2A",
		},
	}, opts)
	if err != nil {
		return nil, err
	}

	channelNames := []string{"DCOutput"}
	channels := make([]Channel, len(channelNames))
	for i, name := range channelNames {
		channels[i] = Channel{name: name, inst: inst, timeout: s.Timeout}
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		timeout:  s.Timeout,
		Inherent: s.Inherent,
	}

	if s.Config.Reset {
		if err := driver.Reset(); err != nil {
			return nil, err
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

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// Close properly shuts down the power supply by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
