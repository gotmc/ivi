// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package fluke45 implements the IVI driver for the Fluke 45 DMM.

State Caching: Not implemented
*/
package fluke45

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

// DMM provides the IVI driver for the Fluke 45 DMM.
type DMM struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*DMM, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 2,
		ClassSpecRevision:     "4.1",
		GroupCapabilities: []string{
			"DmmBase",
			"DmmACMeasurement",
			"DmmFrequencyMeasurement",
			"DmmDeviceInfo",
		},
		SupportedInstrumentModels: []string{
			"45",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := DMM{
		inst:     inst,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// QueryString queries the DMM and returns a string.
func (d *DMM) QueryString(cmd string) (string, error) {
	return query.String(d.inst, cmd)
}
