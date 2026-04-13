// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package dp800 implements the IVI driver for the Rigol DP800 series of
// programmable linear DC power supplies.
//
// State Caching: Not implemented
package dp800

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

// Confirm that the device driver implements the IviDCPwrBase interface.
var _ dcpwr.Base = (*Driver)(nil)
var _ dcpwr.BaseChannel = (*Channel)(nil)
var _ dcpwr.MeasurementChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Rigol DP800 series of DC power
// supplies.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	ivi.Inherent
}

// New creates a new Rigol DP800 IVI Instrument driver. The New function always
// queries the instrument to determine the model for channel configuration. Use
// [ivi.WithIDQuery] to also validate the model against the supported models
// list. Use [ivi.WithReset] to reset the instrument on creation.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)
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
			"DP831A",
			"DP832A",
			"DP821A",
			"DP811A",
			"DP831",
			"DP832",
			"DP821",
			"DP811",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)

	// Always query the model since channel configuration depends on it.
	model, err := inherent.CheckID(context.Background())
	if err != nil && cfg.IDQuery {
		return nil, err
	} else if err != nil {
		// Without idQuery, still need the model for channel config.
		// Try to get it directly.
		model, err = inherent.InstrumentModel(context.Background())
		if err != nil {
			return nil, fmt.Errorf("error determining instrument model: %w", err)
		}
	}

	type genericChannel struct {
		name       string
		minVoltage float64
		maxVoltage float64
		minCurrent float64
		maxCurrent float64
	}

	availableChannels := map[string][]genericChannel{
		"DP831A": {
			{"CH1", 0.0, 8.4, 0.0, 5.3},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -32.0, 0.0, 0.0, 2.1},
		},
		"DP832A": {
			{"CH1", 0.0, 32.0, 0.0, 3.2},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -5.3, 0.0, 0.0, 3.2},
		},
		"DP821A": {
			{"CH1", 0.0, 63.0, 0.0, 1.05},
			{"CH2", 0.0, 8.4, 0.0, 10.5},
		},
		"DP811A": {
			{"Range1", 0.0, 21.0, 0.0, 10.5},
			{"Range2", 0.0, 42.0, 0.0, 5.3},
		},
		"DP831": {
			{"CH1", 0.0, 8.4, 0.0, 5.3},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -32.0, 0.0, 0.0, 2.1},
		},
		"DP832": {
			{"CH1", 0.0, 32.0, 0.0, 3.2},
			{"CH2", 0.0, 32.0, 0.0, 3.2},
			{"CH3", -5.3, 0.0, 0.0, 3.2},
		},
		"DP821": {
			{"CH1", 0.0, 63.0, 0.0, 1.05},
			{"CH2", 0.0, 8.4, 0.0, 10.5},
		},
		"DP811": {
			{"Range1", 0.0, 63.0, 0.0, 1.05},
			{"Range2", 0.0, 8.4, 0.0, 10.5},
		},
	}

	genericChannels := availableChannels[model]

	outputCount := len(genericChannels)
	channels := make([]Channel, outputCount)
	for i, genericChannel := range genericChannels {
		ch := Channel{
			name:       genericChannel.name,
			idx:        i + 1, // 1-based channel index
			inst:       inst,
			minVoltage: genericChannel.minVoltage,
			maxVoltage: genericChannel.maxVoltage,
			minCurrent: genericChannel.minCurrent,
			maxCurrent: genericChannel.maxCurrent,
		}
		channels[i] = ch
	}
	driver := Driver{
		inst:     inst,
		channels: channels,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(context.Background()); err != nil {
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

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	name       string
	idx        int // 1-based channel index for :SOURce[<n>] commands
	inst       ivi.Transport
	minVoltage float64
	maxVoltage float64
	minCurrent float64
	maxCurrent float64
}

// AvailableCOMPorts lists the available COM ports, including optional ports.
func AvailableCOMPorts() []string {
	// FIXME: Is this accurate for all supported models? What about USB?
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
	return []int{128000, 115200, 57600, 38400, 19200, 14400, 9600, 7200, 4800}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return 128000
}

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2", "7E2", "7O2"}
}

// SerialEndMark lists the end mark for commands sent over the RS-232 serial port.
func SerialEndMark() string {
	return "\r\n"
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}

// ChannelCount returns the number of available output channels.
//
// ChannelCount is the getter for the read-only IviDCPwrBase Attribute Output
// Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (dev *Driver) ChannelCount() int {
	return len(dev.channels)
}
