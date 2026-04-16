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
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	model    string
	ivi.Inherent
}

// Channel models the output channel repeated capability for the function
// generator output channel.
type Channel struct {
	inst    ivi.Transport
	name    string
	num     int
	timeout time.Duration
}

// New creates a new InfiniiVision IVI Instrument. By default the constructor
// queries *IDN? and verifies the model against the supported list; pass
// [ivi.WithoutIDQuery] to skip that check. Use [ivi.WithReset] to reset on
// creation and [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

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
			name:    channelName,
			inst:    inst,
			num:     i + 1,
			timeout: timeout,
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
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	model, err := inherent.CheckID()
	if err != nil && !cfg.SkipIDQuery {
		return nil, err
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		timeout:  timeout,
		model:    model,
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

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// Close properly shuts down the oscilloscope by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}
