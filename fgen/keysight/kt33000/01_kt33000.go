// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package kt33000 implements the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators, including the 33200A, 33500A,
// 33500B, and 33600A families. This driver corresponds to the Keysight
// IVI-C Kt33000 driver.
//
// The 33500B and 33600A models use LAN port 5025 for SCPI Socket sessions.
// The default GPIB address is 10.
//
// State Caching: Not implemented
package kt33000

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
var _ fgen.IntTriggerChannel = (*Channel)(nil)
var _ fgen.BurstChannel = (*Channel)(nil)
var _ fgen.ArbWfm = (*Driver)(nil)
var _ fgen.ArbWfmChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators. By default the constructor queries
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
		LocalControlCommand:   "SYST:COMM:RLST LOC",
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenBurst",
			"IviFgenInternalTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{
			"33210A", "33220A", "33502A", "33509B", "33510B", "33511B",
			"33512B", "33519B", "33520B", "33521A", "33521B", "33522A",
			"33522B", "33609A", "33610A", "33611A", "33612A", "33619A",
			"33620A", "33621A", "33622A", "EDU33211A", "EDU33212A",
			"FG33531A", "FG33532A",
		},
		SupportedBusInterfaces: []string{"TCPIP", "GPIB", "USB"},
	}, opts)
	if err != nil {
		return nil, err
	}

	channelNames := []string{"Output1", "Output2"}
	channels := make([]Channel, len(channelNames))
	for i, name := range channelNames {
		channels[i] = Channel{name: name, inst: inst, num: i, timeout: s.Timeout}
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

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
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
	inst    ivi.Transport
	name    string
	num     int // 0-based channel index
	timeout time.Duration
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// srcPrefix returns the SCPI source prefix for this channel (e.g., "SOUR1:").
func (ch *Channel) srcPrefix() string {
	return fmt.Sprintf("SOUR%d:", ch.num+1)
}
