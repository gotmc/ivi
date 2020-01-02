// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package esa implements the IVI driver for the Agilent ESA series spectrum
analyzers.

State Caching: Not implemented
*/
package esa

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 8
	classSpecRevision     = "2.0"
	groupCapabilities     = "IviSpecAnBase"
)

var supportedInstrumentModels = []string{
	"E4411B",
}

// E4411B provides the IVI driver for an Agilent E4411B ESA spectrum
// analyzer.
type E4411B struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new E4411B IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*E4411B, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	sa := E4411B{
		inst:     inst,
		Inherent: inherent,
	}
	if reset {
		err := sa.Reset()
		return &sa, err
	}
	return &sa, nil
}
