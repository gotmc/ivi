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
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Driver, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
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
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
		inst:     inst,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}
