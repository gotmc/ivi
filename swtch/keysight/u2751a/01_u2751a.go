// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package u2751a implements the IVI driver for the Keysight U2751A 4x8 2-wire
// switch matrix.
//
// State Caching: Not implemented
package u2751a

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/swtch"
)

var (
	_ swtch.Base        = (*U2751A)(nil)
	_ swtch.BaseChannel = (*Channel)(nil)
)

const (
	specMajorVersion = 4
	specMinorVersion = 6
	specRevision     = "4.0"
)

// ChannelType is used to determine if the channel is a row or a column.
type ChannelType int

// Available channel types.
const (
	Row ChannelType = iota
	Col
)

// String implements the stringer interface for ChannelType.
func (ct ChannelType) String() string {
	if ct == Row {
		return "row"
	}

	return "column"
}

// U2751A provides the IVI driver for a Keysight U2751A 4x8 2-wire switch
// matrix.
type U2751A struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
	paths []path
}

type path []string

// New creates a new U2751A IVI Instrument. Use [ivi.WithIDQuery] to verify the
// instrument model, [ivi.WithReset] to reset on creation, [ivi.WithStandalone]
// to configure standalone voltage ratings, and [ivi.WithTimeout] to override
// the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (U2751A, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

	infoChannels := []struct {
		name     string
		chType   ChannelType
		switchID int
	}{
		{"Row1", Row, 1},
		{"Row2", Row, 2},
		{"Row3", Row, 3},
		{"Row4", Row, 4},
		{"Col1", Col, 1},
		{"Col2", Col, 2},
		{"Col3", Col, 3},
		{"Col4", Col, 4},
		{"Col5", Col, 5},
		{"Col6", Col, 6},
		{"Col7", Col, 7},
		{"Col8", Col, 8},
	}
	outputCount := len(infoChannels)

	channels := make([]Channel, outputCount)
	for i, ch := range infoChannels {
		channels[i] = newChannel(i, ch.name, ch.chType, ch.switchID, inst, timeout, cfg.Standalone)
	}

	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviSwtchBase",
			"IviSwtchScanner",
			"IviSwtchSoftware",
		},
		SupportedInstrumentModels: []string{
			"U2751A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(); err != nil {
			return U2751A{}, err
		}
	}

	driver := U2751A{
		inst:     inst,
		channels: channels,
		timeout:  timeout,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(); err != nil {
			return driver, err
		}
	}

	return driver, nil
}

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	id                    int
	name                  string
	virtualName           string
	switchID              int
	inst                  ivi.Transport
	timeout               time.Duration
	chType                ChannelType
	acCurrentCarryMax     float64
	acCurrentSwitchingMax float64
	acPowerCarryMax       float64
	acPowerSwitchingMax   float64
	acVoltageMax          float64
	bw                    float64
	impedance             float64
	dcCurrentCarryMax     float64
	dcCurrentSwitchingMax float64
	dcPowerCarryMax       float64
	dcPowerSwitchingMax   float64
	dcVoltageMax          float64
	isConfigChannel       bool
	isDebounced           bool
	isSourceChannel       bool
	settlingTime          time.Duration
	numWires              int
}

// newContext creates a context with the driver's configured timeout.
func (d *U2751A) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// Close properly shuts down the switch matrix by returning it to local
// control.
func (d *U2751A) Close() error {
	return d.Inherent.Close()
}

// Disable causes the switch to disconnect all paths.
func (d *U2751A) Disable() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return ivi.Set(ctx, d.inst, "rout:open (@101:408)\n")
}

// channel returns the concrete channel based on either the virtual name or the
// physical name. Virtual names are checked first.
func (d *U2751A) channel(name string) (*Channel, error) {
	// See if the given name matches one of the virtual channel names.
	for _, ch := range d.channels {
		if name == ch.virtualName {
			return &d.channels[ch.id], nil
		}
	}
	// See if the given name matches one of the physical channel names.
	for _, ch := range d.channels {
		if name == ch.name {
			return &d.channels[ch.id], nil
		}
	}
	return nil, fmt.Errorf("channel %s not found", name)
}

// Channel returns the channel based on either the virtual name or the physical
// name. Virtual names are checked first.
func (d *U2751A) Channel(name string) (swtch.BaseChannel, error) {
	return d.channel(name)
}

// ChannelByID returns the channel based on the ID (0-based).
func (d *U2751A) ChannelByID(id int) (swtch.BaseChannel, error) {
	if id < 0 || id >= len(d.channels) {
		return nil, fmt.Errorf("channel %d not found", id)
	}

	return &d.channels[id], nil
}

// Channels returns all channels.
func (d *U2751A) Channels() ([]swtch.BaseChannel, error) {
	channels := make([]swtch.BaseChannel, len(d.channels))
	for i := range d.channels {
		channels[i] = &d.channels[i]
	}
	return channels, nil
}

func newChannel(
	id int,
	name string,
	chType ChannelType,
	switchID int,
	inst ivi.Transport,
	timeout time.Duration,
	standalone bool,
) Channel {
	dcVoltageMax := 42.0
	acVoltageMax := 35.0

	if !standalone {
		dcVoltageMax = 180.0
		acVoltageMax = 180.0
	}

	return Channel{
		id:                    id,
		name:                  name,
		virtualName:           name,
		switchID:              switchID,
		inst:                  inst,
		timeout:               timeout,
		chType:                chType,
		acCurrentCarryMax:     2.0,
		acCurrentSwitchingMax: 2.0,
		acPowerCarryMax:       62.5,
		acVoltageMax:          acVoltageMax,
		acPowerSwitchingMax:   62.5,
		bw:                    30e6,
		dcCurrentCarryMax:     2.0,
		dcCurrentSwitchingMax: 2.0,
		dcPowerCarryMax:       60,
		dcPowerSwitchingMax:   60,
		dcVoltageMax:          dcVoltageMax,
		impedance:             50,
		isSourceChannel:       false,
		isConfigChannel:       false,
		isDebounced:           false,
		numWires:              2,
		settlingTime:          4 * time.Millisecond,
	}
}

// TODO(mdr): Instead of having a SetVirtualNames method, should the virtual
// names be set at creation and not allowed to be changed?

// SetVirtualNames sets the virtual name for the channels given as a map with
// the physical name provided as the key. Each virtual name must be unique and
// the number of virtual names provided must match the numder of channels
// otherwise an error is returned.
func (d *U2751A) SetVirtualNames(names map[string]string) error {
	for physicalName, virtualName := range names {
		for i, ch := range d.channels {
			if physicalName == ch.name {
				d.channels[i].virtualName = virtualName
			}
		}
	}

	return nil
}
