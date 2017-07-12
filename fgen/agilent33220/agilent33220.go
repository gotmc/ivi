// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilent33220 implements the IVI driver for the Agilent 33220A function
generator.

State Caching: Not implemented
*/
package agilent33220

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 3
	classSpecRevision     = "5.2"
	groupCapabilities     = "FgenBase,FgenStdfunc"
	idnString             = `^(?P<mfr>[^,]+),` +
		`(?P<model>[^,]+),0,` +
		`(?P<fwr>\d{1}.\d{2})-` +
		`(?P<boot>\d{1}.\d{2})-` +
		`(?P<asic>\d{2})-` +
		`(?P<pcb>\d{1}$`
)

var supportedInstrumentModels = []string{
	"33220A",
	"33210A",
}

// Agilent33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Agilent33220 struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
	ivi.Inherent
}

// New creates a new Agilent33220 IVI Instrument.
func New(inst ivi.Instrument) (*Agilent33220, error) {
	outputCount := 1
	ch := Channel{
		id:   0,
		inst: inst,
	}
	channels := make([]Channel, outputCount)
	channels[0] = ch
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
		IDNString:                 idnString,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	fgen := Agilent33220{
		inst:        inst,
		outputCount: outputCount,
		Channels:    channels,
		Inherent:    inherent,
	}
	return &fgen, nil
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviFgenBase Attribute Output Count described in
// Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (fgen *Agilent33220) OutputCount() int {
	return fgen.outputCount
}
