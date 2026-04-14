// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package key33000 implements the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators, including the 33200A, 33500A,
// 33500B, and 33600A families. This driver corresponds to the Keysight
// IVI-C Kt33000 driver.
//
// The 33500B and 33600A models use LAN port 5025 for SCPI Socket sessions.
// The default GPIB address is 10.
//
// State Caching: Not implemented
package key33000

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion   = 4
	specMinorVersion   = 3
	specRevision       = "5.2"
	defaultGPIBAddress = 10
	telnetPort         = 5024
	socketPort         = 5025
)

// Confirm the interfaces implemented by the driver.
var _ fgen.Base = (*Driver)(nil)
var _ fgen.BaseChannel = (*Channel)(nil)
var _ fgen.StdFuncChannel = (*Channel)(nil)
var _ fgen.StartTriggerChannel = (*Channel)(nil)
var _ fgen.TriggerChannel = (*Channel)(nil)
var _ fgen.IntTrigger = (*Driver)(nil)
var _ fgen.BurstChannel = (*Channel)(nil)
var _ fgen.ArbWfm = (*Driver)(nil)
var _ fgen.ArbWfmChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators. The context is used for any I/O
// performed during construction (e.g., ID query, reset). Use [ivi.WithIDQuery]
// to verify the instrument model and [ivi.WithReset] to reset on creation.
func New(ctx context.Context, inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)
	channelNames := []string{
		"Output1",
		"Output2",
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
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenBurst",
			"IviFgenInternalTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{
			"33210A",
			"33220A",
			"33502A",
			"33509B",
			"33510B",
			"33511B",
			"33512B",
			"33519B",
			"33520B",
			"33521A",
			"33521B",
			"33522A",
			"33522B",
			"33609A",
			"33610A",
			"33611A",
			"33612A",
			"33619A",
			"33620A",
			"33621A",
			"33622A",
			"EDU33211A",
			"EDU33212A",
			"FG33531A",
			"FG33532A",
		},
		SupportedBusInterfaces: []string{
			"TCPIP",
			"GPIB",
			"USB",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(ctx); err != nil {
			return nil, err
		}
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(ctx); err != nil {
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

// Close properly shuts down the function generator by returning it to local
// control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// AvailableCOMPorts lists the available COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "LAN", "USB"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}

// LANPorts returns a map of the different ports with the key being the type of
// port.
func LANPorts() map[string]int {
	return map[string]int{
		"telnet": telnetPort,
		"socket": socketPort,
	}
}

// Channel models the output channel repeated capability for the function
// generator output channel.
type Channel struct {
	inst ivi.Transport
	name string
	num  int // 1-based channel number for SOURce[1|2]: prefix
}

// srcPrefix returns the SCPI source prefix for this channel (e.g., "SOUR1:").
func (ch *Channel) srcPrefix() string {
	return fmt.Sprintf("SOUR%d:", ch.num)
}
