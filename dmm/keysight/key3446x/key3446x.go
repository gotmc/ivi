// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key3446x implements the IVI Instrument driver for the Keysight 3446x
family of DMM.

State Caching: Not implemented
*/
package key3446x

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

// Ag3446x provides the IVI Instrument driver for the Keysight 3446x family of
// DMM.
type Ag3446x struct {
	inst        ivi.Instrument
	outputCount int
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Ag3446x, error) {
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
			"34460A",
			"34461A",
			"34465A",
			"34470A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Ag3446x{
		inst:        inst,
		outputCount: 1,
		Inherent:    inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// QueryString queries the DMM and returns a string.
func (d *Ag3446x) QueryString(cmd string) (string, error) {
	return query.String(d.inst, cmd)
}
