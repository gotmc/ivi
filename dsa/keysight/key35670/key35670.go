// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key35670

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dsa"
)

/*
Package key35670 implements the IVI driver for the Keysight 35670 Dynamic
Signal Analyzer (DSA).
*/

// Required to implement the Inherent Capabilities & Attributes
const (
	classSpecMajorVersion = 100
	classSpecMinorVersion = 0
	classSpecRevision     = "100.0"
	groupCapabilities     = "IviDSABase"
)

var supportedInstrumentModels = []string{
	"35670A",
}

var channelNames = []string{
	"CH1",
	"CH2",
	"CH3",
	"CH4",
}

// Key35670 provides the IVI driver for a Keysight 35670A Dynamic Signal
// Analyzer.
type Key35670 struct {
	inst     ivi.Instrument
	Channels []Channel
	ivi.Inherent
}

// New creates a new Key35670 IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Key35670, error) {
	inputCount := len(channelNames)
	channels := make([]Channel, inputCount)
	for i, ch := range channelNames {
		baseChannel := dsa.NewChannel(i, ch, inst)
		channels[i] = Channel{baseChannel}
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion:     classSpecMajorVersion,
		ClassSpecMinorVersion:     classSpecMinorVersion,
		ClassSpecRevision:         classSpecRevision,
		GroupCapabilities:         groupCapabilities,
		SupportedInstrumentModels: supportedInstrumentModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	dev := Key35670{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := dev.Reset()
		return &dev, err
	}
	return &dev, nil
}

// Channel represents a repeated capability of an input channel for the Dynamic
// Signal Analyzer (DSA).
type Channel struct {
	dsa.Channel
}
