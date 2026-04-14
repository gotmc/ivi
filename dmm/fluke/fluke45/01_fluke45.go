// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package fluke45 implements the IVI driver for the Fluke 45 DMM.
//
// State Caching: Not implemented
package fluke45

import (
	"context"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
)

const (
	specMajorVersion = 4
	specMinorVersion = 2
	specRevision     = "4.1"
)

// Confirm the interfaces implemented by the driver.
var _ dmm.Base = (*Driver)(nil)

// Driver provides the IVI driver for the Fluke 45 DMM.
type Driver struct {
	inst    ivi.Transport
	timeout time.Duration
	ivi.Inherent
}

// New creates a new Fluke 45 IVI Instrument. Use [ivi.WithIDQuery] to verify
// the instrument model, [ivi.WithReset] to reset on creation, and
// [ivi.WithTimeout] to override the default I/O timeout.
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
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDmmBase",
			"IviDmmACMeasurement",
			"IviDmmFrequencyMeasurement",
			"IviDmmDeviceInfo",
		},
		SupportedInstrumentModels: []string{
			"45",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"Serial",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(); err != nil {
			return nil, err
		}
	}

	driver := Driver{
		inst:     inst,
		timeout:  timeout,
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

// Close properly shuts down the DMM by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
