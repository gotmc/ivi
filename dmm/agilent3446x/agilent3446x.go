// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilent3446x implements the IVI driver for the Keysight 3446x family of
DMM.

State Caching: Not implemented
*/
package agilent3446x

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 2
	classSpecRevision     = "4.1"
	groupCapabilities     = "DmmBase,DmmACMeasurement,DmmFrequencyMeasurement,DmmDeviceInfo"
	idnString             = `^(?P<mfr>[^,]+),` +
		`(?P<model>[^,]+),` +
		`(?P<sn>[^,]+),` +
		`(?P<hw>\d{1})-` +
		`(?P<fwr>\d{2}.\d{2})-` +
		`(?P<pp>\d{2}.\d{2})-` +
		`(?P<mm>\d{2}.\d{2})-` +
		`(?P<fpga>\d{2}.\d{2})-` +
		`(?P<bb>\d{2})-` +
		`(?P<fp>\d{1}$`
)

var supportedInstrumentModels = []string{
	"34460A",
	"34461A",
	"34465A",
	"34470A",
}

// Required for base DMM
const (
	outputCount = 1
)

// Agilent3446x provides the IVI driver for the Keysight 3446x family of DMM.
type Agilent3446x struct {
	inst        ivi.Instrument
	outputCount int
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument) (*Agilent3446x, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
		IDNString:                 idnString,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	dmm := Agilent3446x{
		inst:        inst,
		outputCount: outputCount,
		Inherent:    inherent,
	}
	return &dmm, nil
}
