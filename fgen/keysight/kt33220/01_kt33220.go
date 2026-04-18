// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package kt33220 implements the IVI driver for the Keysight/Agilent 33220A
// and 33210A function/arbitrary waveform generators.
//
// State Caching: Not implemented
package kt33220

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

// Driver provides the IVI driver for a Keysight/Agilent 33220A or 33210A
// function generator.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 33210A and 33220A
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
		GroupCapabilities: []string{
			"IviFgenBase",
			// "IviFgenArbFrequency",
			// "IviFgenArbWfm",
			"IviFgenBurst",
			"IviFgenInternalTrigger",
			// "IviFgenModulateAM",
			// "IviFgenModulateFM",
			// "IviFgenSoftwareTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{"33220A", "33210A"},
		SupportedBusInterfaces:    []string{"TCPIP", "GPIB", "USB"},
	}, opts)
	if err != nil {
		return nil, err
	}

	channelNames := []string{"Output"}
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

// Close properly shuts down the function generator by returning it to local control.
// This ensures the instrument's front panel regains control after use.
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
	timeout time.Duration
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}
