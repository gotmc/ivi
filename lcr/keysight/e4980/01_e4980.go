// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package e4980 implements the IVI driver for the Keysight/Agilent E4980A and
// E4980AL precision LCR meters.
//
// State Caching: Not implemented
package e4980

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/lcr"
)

// Confirm the interfaces implemented by the driver.
var _ lcr.Base = (*Driver)(nil)
var _ lcr.DCBias = (*Driver)(nil)
var _ lcr.Compensation = (*Driver)(nil)

// Driver provides the IVI driver for the Keysight E4980A and E4980AL
// precision LCR meters.
type Driver struct {
	inst    ivi.Transport
	timeout time.Duration
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight E4980A/AL LCR meter. By
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
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviLCRBase",
			"IviLCRDCBias",
			"IviLCRCompensation",
		},
		SupportedInstrumentModels: []string{
			"E4980A",
			"E4980AL",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"USB",
			"TCPIP",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase, timeout)

	if _, err := inherent.CheckID(); err != nil && !cfg.SkipIDQuery {
		return nil, err
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

// Close properly shuts down the LCR meter by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// DefaultGPIBAddress returns the default GPIB interface address.
func DefaultGPIBAddress() int {
	return 17
}

// LANPort returns the default SCPI socket port.
func LANPort() int {
	return 5025
}

// MeasurementResult holds the primary and secondary measurement values along
// with the measurement status.
type MeasurementResult struct {
	Primary   float64
	Secondary float64
	Status    lcr.MeasurementStatus
	Bin       int // Comparator bin number; 0 if comparator is off
}

// String implements the fmt.Stringer interface for MeasurementResult.
func (r MeasurementResult) String() string {
	return fmt.Sprintf(
		"primary=%g, secondary=%g, status=%s",
		r.Primary, r.Secondary, r.Status,
	)
}
