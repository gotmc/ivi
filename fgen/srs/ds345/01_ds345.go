// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package ds345 implements the IVI driver for the Stanford Research Systems
// DS345 function generator.
//
// State Caching: Not implemented
//
// The serial port is a DB25 DCE. The [Green-utech USB RS-232 to DB25 serial
// cable](https://www.amazon.com/dp/B08J2VMNFY) does work.
package ds345

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion      = 4
	specMinorVersion      = 3
	specRevision          = "5.2"
	defaultGPIBAddress    = 19
	defaultSerialBuadRate = 9600
	defaultResetDelay     = 500 * time.Millisecond
	defaultClearDelay     = 500 * time.Millisecond
)

// Confirm the implemented interfaces by the driver.
var _ fgen.Base = (*Driver)(nil)
var _ fgen.BaseChannel = (*Channel)(nil)
var _ fgen.BurstChannel = (*Channel)(nil)
var _ fgen.IntTriggerChannel = (*Channel)(nil)
var _ fgen.StdFuncChannel = (*Channel)(nil)
var _ fgen.TriggerChannel = (*Channel)(nil)

// Driver provides the IVI driver for a SRS DS345 function generator.
type Driver struct {
	inst     ivi.Transport
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// Channel models the output channel repeated capability for the function
// generator output channel.
type Channel struct {
	inst    ivi.Transport
	name    string
	timeout time.Duration
}

// New creates a new DS345 IVI Instrument. By default the constructor queries
// *IDN? and verifies the model against the supported list; pass
// [ivi.WithoutIDQuery] to skip that check. Use [ivi.WithReset] to reset on
// creation and [ivi.WithTimeout] to override the default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

	channelNames := []string{
		"Output",
	}
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)

	for i, channelName := range channelNames {
		ch := Channel{
			name:    channelName,
			inst:    inst,
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
		// Commented out GroupCapabilities still need to be added.
		GroupCapabilities: []string{
			// "IviFgenArbFrequency",
			// "IviFgenArbSeq",
			// "IviFgenArbWaveform",
			"IviFgenBase",
			"IviFgenBurst",
			// "IviFgenInternalTrigger",
			// "IviFgenModulateFM",
			// "IviFgenModulateAM",
			// "IviFgenSoftwareTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: []string{
			"DS345",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"RS232",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if _, err := inherent.CheckID(); err != nil && !cfg.SkipIDQuery {
		return nil, err
	}

	driver := Driver{
		inst:     inst,
		channels: channels,
		timeout:  timeout,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(); err != nil {
			return &driver, err
		}
		// Default to internal trigger instead of single trigger when reset.
		ctx, cancel := driver.newContext()
		defer cancel()

		if err := driver.inst.Command(ctx, "TSRC1"); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
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

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}

// SerialConfig lists whether the RS-232 serial port is configured as a DCE
// (Data Circuit-Terminating Equipment) or a DTE (Data Terminal Equipment). Computers
// running the IVI program are DTEs; therefore, use a straight through serial
// cable when connecting to DCEs and a null modem cable when connecting to DTEs.
func SerialConfig() string {
	return "DCE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{19200, 9600, 4800, 2400, 1200, 600, 300}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return defaultSerialBuadRate
}

// SerialDataFrames lists the available RS-232 data frame formats. The DS345
// "always sends two stop bits, 8 data bits, and no parity, and will correctly
// receive data sent with eitehr one or two stop bits." per the User's Manual.
func SerialDataFrames() []string {
	return []string{"8N2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}
