// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilent33220 implements the IVI driver for the Agilent 33220A function
generator.

IVI Instrument Class: IviFgen
IVI Class Specification: IVI-4.3
IVI Specification Revision: 5.2
IVI Specification Edition: October 14, 2016
Capability Groups Supported (specification section):
   4. IviFgenBase               Partially (missing 4.2.5, 4.2.7)
   5. IviFgenStdFunc            Yes
   6. IviFgenArbWfm             Not Yet
   7. IviFgenArbFrequency       Not Yet
   8. IviFgenArbSeq             No
   9. IviFgenTrigger            Not Yet
  10. IviFgenStartTrigger       Not Yet
  11. IviFgenStopTrigger        Not Yet
  12. IviFgenHoldTrigger        Not Yet
  16. IviFgenSoftwareTrigger    Not Yet
  17. IviFgenBurst              Not Yet (next to work on)

Hardware Information:
  Instrument Manufacturer:          Keysight Technologies
	Supported Instrument Models:      33210A, 33220A

State Caching: Not implemented
*/
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
	waveform      ivi.StandardWaveform
}

func (stdFunc *StandardFunction) ConfigureWaveform(w io.Writer) {

}

// Agilent33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Agilent33220 struct {
	inst        ivi.Instrument
	outputCount int
	Channels    []Channel
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviFgenBase Attribute Output Count described in
// Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
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
		Channels:    channels,
	}
	return &fgen, nil
}
