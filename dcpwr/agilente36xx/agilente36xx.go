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

// Required to implement the Inherent and DCPwr attributes and capabilities
const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 4
	classSpecRevision     = "3.0"
	groupCapabilities     = "DCPwrBase,DCPwrMeasurement"
)

var supportedInstrumentModels = []string{
	"E3631A",
}

var channelNames = []string{
	"P6V",
	"P25V",
	"N25V",
}

// AgilentE36xx provides the IVI driver for the Agilent/Keysight E3600 series
// of power supplies.
type AgilentE36xx struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new AgilentE36xx IVI Instrument.
func New(inst ivi.Instrument) (*AgilentE36xx, error) {
	outputCount := len(channelNames)
	channels := make([]Channel, outputCount)
	for i, ch := range channelNames {
		channels[i] = Channel{
			id:   i,
			name: ch,
			inst: inst,
		}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	dcpwr := AgilentE36xx{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	return &dcpwr, nil
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviDCPwrBase Attribute Output Channel Count
// described in Section 4.2.7 of IVI-4.4: IviDCPwr Class Specification.
func (dcpwr *AgilentE36xx) OutputCount() int {
	return len(dcpwr.Channels)
}
