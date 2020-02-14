// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
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

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 2
	classSpecRevision     = "4.1"
	groupCapabilities     = "DmmBase,DmmACMeasurement,DmmFrequencyMeasurement,DmmDeviceInfo"
)

var supportedInstrumentModels = []string{
	"45",
}

// DMM provides the IVI driver for the Fluke 45 DMM.
type DMM struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*DMM, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	dmm := DMM{
		inst:     inst,
		Inherent: inherent,
	}
	if reset {
		err := dmm.Reset()
		return &dmm, err
	}
	return &dmm, nil
}

// QueryString queries the DMM and returns a string.
func (d *DMM) QueryString(cmd string) (string, error) {
	return query.String(d.inst, cmd)
}
