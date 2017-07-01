// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"io"

	"github.com/gotmc/ivi"
)

type StandardFunction struct {
	amplitude     float64
	dCOffset      float64
	dutyCycleHigh float64
	frequency     float64
	startPhase    float64
	waveform      StandardWaveform
}

func (stdFunc *StandardFunction) ConfigureWaveform(w io.Writer) {

}

// Agilent33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Agilent33220 struct {
	inst        ivi.Instrument
	outputCount int
	Ch          []Channel
}

// OutputCount returns the number of outputs for the function generator.
func (fgen *Agilent33220) OutputCount() int {
	return fgen.outputCount
}

// Close closes the IVI Instrument.
func (fgen *Agilent33220) Close() error {
	return fgen.inst.Close()
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
	fgen := Agilent33220{
		inst:        inst,
		outputCount: outputCount,
		Ch:          channels,
	}
	return &fgen, nil
}
