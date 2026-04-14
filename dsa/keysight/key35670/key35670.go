// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package key35670 implements the IVI Instrument driver for the Keysight 35670
// dynamic signal analyzer.
//
// Note: Dynamic Signal Analyzers are not part of the IVI Specification.
//
// State Caching: Not implemented
package key35670

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dsa"
)

/*
Package key35670 implements the IVI Instrument driver for the Keysight 35670
Dynamic Signal Analyzer (DSA).
*/

const (
	specMajorVersion = 0
	specMinorVersion = 0
	specRevision     = "N/A"
)

// Key35670 provides the IVI driver for a Keysight 35670A Dynamic Signal
// Analyzer.
type Key35670 struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// New creates a new Key35670 IVI Instrument driver. Use [ivi.WithIDQuery] to
// verify the instrument model, [ivi.WithReset] to reset on creation, and
// [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Key35670, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

	channelNames := []string{
		"CH1",
		"CH2",
		"CH3",
		"CH4",
	}
	inputCount := len(channelNames)
	channels := make([]Channel, inputCount)
	for i, ch := range channelNames {
		baseChannel := dsa.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDSABase",
		},
		SupportedInstrumentModels: []string{
			"35670A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(); err != nil {
			return nil, err
		}
	}

	driver := Key35670{
		inst:     inst,
		channels: channels,
		timeout:  timeout,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// newContext creates a context with the driver's configured timeout.
func (d *Key35670) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// Channel returns the Channel at the given index, with bounds checking.
func (d *Key35670) Channel(index int) (*Channel, error) {
	if index < 0 || index >= len(d.channels) {
		return nil, fmt.Errorf("channel %d: %w", index, ivi.ErrChannelNotFound)
	}

	return &d.channels[index], nil
}

// InputChannel represents a repeated capability of an input channel for the
// Dynamic Signal Analyzer (DSA).
type Channel struct {
	dsa.Channel
}
