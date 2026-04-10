// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package esa implements the IVI driver for the Agilent ESA series spectrum
// analyzers.
//
// State Caching: Not implemented
package esa

import (
	"context"

	"github.com/gotmc/ivi"
)

const (
	specMajorVersion = 4
	specMinorVersion = 8
	specRevision     = "2.0"
)

// E4411B provides the IVI driver for an Agilent E4411B ESA spectrum
// analyzer.
type E4411B struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new E4411B IVI Instrument. Use [ivi.WithIDQuery] to verify the
// instrument model and [ivi.WithReset] to reset on creation.
func New(inst ivi.Instrument, opts ...ivi.DriverOption) (*E4411B, error) {
	cfg := ivi.ApplyOptions(opts)
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviSpecAnBase",
		},
		SupportedInstrumentModels: []string{
			"E4411B",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)

	if cfg.IDQuery {
		if _, err := inherent.CheckID(context.Background()); err != nil {
			return nil, err
		}
	}

	driver := E4411B{
		inst:     inst,
		Inherent: inherent,
	}

	if cfg.Reset {
		if err := driver.Reset(context.Background()); err != nil {
			return &driver, err
		}
	}

	return &driver, nil
}

// Close properly shuts down the spectrum analyzer by returning it to local
// control.
func (d *E4411B) Close() error {
	return d.Inherent.Close()
}
