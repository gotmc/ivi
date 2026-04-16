// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package sr630 implements the IVI driver for the Stanford Research Systems
// SR630 16-channel thermocouple monitor.
//
// The SR630 uses flat 4-letter commands (not SCPI subsystem hierarchy) and
// supports thermocouple types B, E, J, K, R, S, and T. Each channel can be
// independently configured for sensor type and measurement units.
//
// State Caching: Not implemented
package sr630

import (
	"context"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/tempmon"
)

const (
	numChannels     = 16
	defaultGPIBAddr = 19
	defaultBaudRate = 9600
)

// Confirm the interfaces implemented by the driver.
var _ tempmon.Base = (*Driver)(nil)
var _ tempmon.Scanner = (*Driver)(nil)
var _ tempmon.RelativeTemperature = (*Driver)(nil)

// Driver provides the IVI driver for the SRS SR630 thermocouple monitor.
type Driver struct {
	inst    ivi.Transport
	timeout time.Duration
	model   string
	ivi.Inherent
}

// New creates a new IVI driver for the SRS SR630 thermocouple monitor. By
// default the constructor queries *IDN? and verifies the model against the
// supported list; pass [ivi.WithoutIDQuery] to skip that check. Use
// [ivi.WithReset] to reset on creation and [ivi.WithTimeout] to override the
// default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)

	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = ivi.DefaultTimeout
	}

	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 0,
		ClassSpecMinorVersion: 0,
		ClassSpecRevision:     "N/A",
		ResetDelay:            2 * time.Second,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         false, // SR630 has no SYST:LOC equivalent
		GroupCapabilities: []string{
			"TempMonBase",
			"TempMonScanner",
			"TempMonRelativeTemperature",
		},
		SupportedInstrumentModels: []string{
			"SR630",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"RS232",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	model, err := inherent.CheckID()
	if err != nil && !cfg.SkipIDQuery {
		return nil, err
	}

	driver := Driver{
		inst:     inst,
		timeout:  timeout,
		model:    model,
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
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// Close returns the instrument to its default state. The SR630 has no
// remote-to-local control command.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// DefaultGPIBAddress returns the default GPIB address for the SR630.
func DefaultGPIBAddress() int {
	return defaultGPIBAddr
}

// DefaultBaudRate returns the default RS-232 baud rate for the SR630.
func DefaultBaudRate() int {
	return defaultBaudRate
}
