// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package u2751a implements the IVI driver for the Keysight U2751A 4x8 2-wire
switch matrix.

State Caching: Not implemented
*/
package u2751a

import (
	"fmt"
	"time"

	"github.com/gotmc/ivi"
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
	inst     ivi.Instrument
	channels []Channel
	ivi.Inherent
	paths []path
}

type path []string

// New creates a new U2751A IVI Instrument.
func New(inst ivi.Instrument, reset, standalone bool) (U2751A, error) {
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
		channels[i] = newChannel(i, ch.name, ch.chType, ch.switchID, inst, standalone)
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		GroupCapabilities: []string{
			"IviSwtchBase",
			"IviSwtchScanner",
			"IviSwtchSoftware",
		},
		SupportedInstrumentModels: []string{
			"U2751A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := U2751A{
		inst:     inst,
		channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return driver, err
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
	inst                  ivi.Instrument
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

// Disable causes the switch to disconnect all paths.
func (d *U2751A) Disable() error {
	return ivi.Set(d.inst, "rout:open (@101:408)\n")
}

// Channel returns the channel based on either the virtual name or the physical
// name. Virtual names are checked first.
func (d *U2751A) Channel(name string) (*Channel, error) {
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
	return &Channel{}, fmt.Errorf("channel %s not found", name)
}

// ChannelByID returns the channel based on the ID (0-based).
func (d *U2751A) ChannelByID(id int) (*Channel, error) {
	if id < 0 || id > len(d.channels) {
		return &Channel{}, fmt.Errorf("channel %d not found", id)
	}
	return &d.channels[id], nil
}

// Channels returns all channels.
func (d *U2751A) Channels() ([]Channel, error) {
	return d.channels, nil
}

func newChannel(id int, name string, chType ChannelType, switchID int, inst ivi.Instrument, standalone bool) Channel {
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
