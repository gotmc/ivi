// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package e36xx implements the IVI driver for the Keysight/Agilent E3600
// series of power supplies.
//
// State Caching: Not implemented
package e36xx

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

// Confirm the interfaces implemented by the driver.
var _ dcpwr.Base = (*Driver)(nil)
var _ dcpwr.BaseChannel = (*Channel)(nil)
var _ dcpwr.MeasurementChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Agilent/Keysight E3600 series of DC
// power supplies.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	ivi.Inherent
}

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	inst ivi.Transport
	name string
}

// New creates a new IVI driver for the Keysight/Agilent E3600 series of DC
// power supplies. The New function always queries the instrument to determine
// the model for channel configuration. Use [ivi.WithIDQuery] to also validate
// the model against the supported models list, [ivi.WithReset] to reset
// the instrument on creation, and [ivi.WithTimeout] to override the default
// I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            700 * time.Millisecond,
		ClearDelay:            700 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
		},
		SupportedInstrumentModels: []string{
			"E3631A",
			"E3632A",
			"E3633A",
			"E3634A",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"SERIAL",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	// Always query the model since channel configuration depends on it.
	model, err := inherent.CheckID()
	if err != nil && cfg.IDQuery {
		return nil, err
	} else if err != nil {
		// Without idQuery, still need the model for channel config.
		model, err = inherent.InstrumentModel()
		if err != nil {
			return nil, fmt.Errorf("error determining instrument model: %w", err)
		}
	}

	channelNames := availableChannels[model]
	channels := make([]Channel, len(channelNames))

	for i, channelName := range channelNames {
		channels[i] = Channel{
			name: channelName,
			inst: inst,
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

// availableChannels maps instrument models to their output channel names.
var availableChannels = map[string][]string{
	"E3631A": {"P6V", "P25V", "N25V"},
	"E3632A": {"Output"},
	"E3633A": {"Output"},
	"E3634A": {"Output"},
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

// AvailableCOMPorts lists the available COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "RS232"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return 5
}

// SerialConfig lists whether the RS-232 serial port is configured as a DCE
// (Data Circuit-Terminating Equipment) or a DTE (Data Terminal Equipment). Computers
// running the IVI program are DTEs; therefore, use a straight through serial
// cable when connecting to DCEs and a null modem cable when connecting to DTEs.
func SerialConfig() string {
	return "DTE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{9600, 4800, 2400, 1200, 600, 300}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return 9600
}

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2", "7E2", "7O2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
