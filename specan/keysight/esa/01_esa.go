// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package esa implements the IVI driver for the Agilent/Keysight ESA series
// spectrum analyzers. The driver also supports the PSA, EMC, and X-Series
// analyzers that share compatible SCPI command sets.
//
// State Caching: Not implemented
package esa

import (
	"context"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/specan"
)

const (
	specMajorVersion = 4
	specMinorVersion = 8
	specRevision     = "2.0"
)

// Confirm the interfaces implemented by the driver.
var _ specan.Base = (*Driver)(nil)

// Driver provides the IVI driver for Keysight/Agilent ESA, PSA, EMC, and
// X-Series spectrum analyzers.
type Driver struct {
	inst ivi.Transport
	ivi.Inherent
}

// New creates a new IVI driver for Keysight/Agilent spectrum analyzers. The
// context is used for any I/O performed during construction (e.g., ID query,
// reset). Use [ivi.WithIDQuery] to verify the instrument model and
// [ivi.WithReset] to reset on creation.
func New(ctx context.Context, inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	cfg := ivi.ApplyOptions(opts)
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviSpecAnBase",
		},
		SupportedInstrumentModels: []string{
			// ESA-L Series
			"E4411B",
			// ESA-E Series
			"E4401B", "E4402B", "E4403B", "E4404B",
			"E4405B", "E4407B", "E4408B",
			// EMC Series
			"E7401A", "E7402A", "E7403A", "E7404A", "E7405A",
			// PSA Series
			"E4440A", "E4443A", "E4445A", "E4446A",
			"E4447A", "E4448A", "N8201A",
			// X-Series
			"N9030A", "N9020A", "N9010A", "N9000A",
		},
		SupportedBusInterfaces: []string{
			"GPIB",
			"TCPIP",
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
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(ctx); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// Close properly shuts down the spectrum analyzer by returning it to local
// control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
