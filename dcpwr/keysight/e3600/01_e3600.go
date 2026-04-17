// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package e3600 implements the IVI driver for the Keysight/Agilent E3600
// series of power supplies.
//
// State Caching: Not implemented
package e3600

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

// Confirm the interfaces implemented by the driver.
var (
	_ dcpwr.Base               = (*Driver)(nil)
	_ dcpwr.BaseChannel        = (*Channel)(nil)
	_ dcpwr.MeasurementChannel = (*Channel)(nil)
	_ dcpwr.Trigger            = (*Driver)(nil)
	_ dcpwr.TriggerChannel     = (*Channel)(nil)
	_ dcpwr.SoftwareTrigger    = (*Driver)(nil)
)

// Driver provides the IVI driver for the Agilent/Keysight E3600 series of DC
// power supplies.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	inst    ivi.Transport
	name    string
	timeout time.Duration
}

// New creates a new IVI driver for the Keysight/Agilent E3600 series of DC
// power supplies. The constructor always queries *IDN? since channel
// configuration depends on the model; by default it also validates the model
// against the supported list. Pass [ivi.WithoutIDQuery] to skip validation
// (the model is still queried). Use [ivi.WithReset] to reset on creation and
// [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	s, err := ivi.NewDriverSetup(inst, ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            700 * time.Millisecond,
		ClearDelay:            700 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
			"IviDCPwrTrigger",
			"IviDCPwrSoftwareTrigger",
		},
		SupportedInstrumentModels: []string{
			"E3631A",
			"E3632A",
			"E3633A",
			"E3634A",
			"E3640A",
			"E3641A",
			"E3642A",
			"E3643A",
			"E3644A",
			"E3645A",
			"E3646A",
			"E3647A",
			"E3648A",
			"E3649A",
			"E36102A",
			"E36103A",
			"E36104A",
			"E36105A",
			"E36106A",
			"E36102B",
			"E36103B",
			"E36104B",
			"E36105B",
			"E36106B",
			"E36231A",
			"E36232A",
			"E36154A",
			"E36155A",
			"E36233A",
			"E36234A",
			"E36311A",
			"E36312A",
			"E36313A",
			"E36441A",
			"E36731A",
			"EDU36311A",
		},

		SupportedBusInterfaces: []string{"GPIB", "SERIAL"},
	}, opts)
	if err != nil {
		return nil, err
	}

	// Channel configuration depends on the queried model.
	model, err := s.Inherent.InstrumentModel()
	if err != nil {
		return nil, fmt.Errorf("error determining instrument model: %w", err)
	}

	channelNames := availableChannels[model]
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

// availableChannels maps instrument models to their output channel names.
var availableChannels = map[string][]string{
	"E3631A":    {"P6V", "P25V", "N25V"},
	"E3632A":    {"Output"},
	"E3633A":    {"Output"},
	"E3634A":    {"Output"},
	"E3640A":    {"Output"},
	"E3641A":    {"Output"},
	"E3642A":    {"Output"},
	"E3643A":    {"Output"},
	"E3644A":    {"Output"},
	"E3645A":    {"Output"},
	"E3646A":    {"Output 1", "Output 2"},
	"E3647A":    {"Output 1", "Output 2"},
	"E3648A":    {"Output 1", "Output 2"},
	"E3649A":    {"Output 1", "Output 2"},
	"E36102A":   {"Output"},
	"E36103A":   {"Output"},
	"E36104A":   {"Output"},
	"E36105A":   {"Output"},
	"E36106A":   {"Output"},
	"E36102B":   {"Output"},
	"E36103B":   {"Output"},
	"E36104B":   {"Output"},
	"E36105B":   {"Output"},
	"E36106B":   {"Output"},
	"E36231A":   {"Output"},
	"E36232A":   {"Output"},
	"E36154A":   {"Output"},
	"E36155A":   {"Output"},
	"E36233A":   {"Output"},
	"E36234A":   {"Output"},
	"E36311A":   {"Output 1", "Output 2", "Output 3"},
	"E36312A":   {"Output 1", "Output 2", "Output 3"},
	"E36313A":   {"Output 1", "Output 2", "Output 3"},
	"E36441A":   {"Output 1", "Output 2", "Output 3", "Output 4"},
	"E36731A":   {"Output"},
	"EDU36311A": {"Output 1", "Output 2", "Output 3"},
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
