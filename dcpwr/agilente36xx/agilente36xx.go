// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilente36xx implements the IVI driver for the Agilent/Keysight E3600
series of power supplies.

State Caching: Not implemented
*/
package agilente36xx

import "github.com/gotmc/ivi"

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 4
	classSpecRevision     = "3.0"
	groupCapabilities     = "DCPwrBase,DCPwrMeasurement"
	idnString             = `^(?P<mfr>[^,]+),` +
		`(?P<model>[^,]+),0,` +
		`(?P<fwr>\d{1}.\d{1})-` +
		`(?P<boot>\d{1}.\d{1})-` +
		`(?P<asic>\d{1})$`
)

var supportedInstrumentModels = []string{
	"E3631A",
}

// AgilentE36xx provides the IVI driver for the Agilent/Keysight E3600 series
// of power supplies.
type AgilentE36xx struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
	ivi.Inherent
}

// New creates a new AgilentE36xx IVI Instrument.
func New(inst ivi.Instrument) (*AgilentE36xx, error) {
	outputCount := 3
	p6v := Channel{
		id:   0,
		name: "P6V",
		inst: inst,
	}
	p25v := Channel{
		id:   1,
		name: "P25V",
		inst: inst,
	}
	n25v := Channel{
		id:   2,
		name: "N25V",
		inst: inst,
	}
	channels := make([]Channel, outputCount)
	channels[0] = p6v
	channels[1] = p25v
	channels[2] = n25v
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	dcpwr := AgilentE36xx{
		inst:        inst,
		outputCount: outputCount,
		Channels:    channels,
		Inherent:    inherent,
	}
	return &dcpwr, nil
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviDCPwrBase Attribute Output Channel Count
// described in Section 4.2.7 of IVI-4.4: IviDCPwr Class Specification.
func (dcpwr *AgilentE36xx) OutputCount() int {
	return dcpwr.outputCount
}
