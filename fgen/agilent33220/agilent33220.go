// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"fmt"
	"io"
	"strconv"

	"github.com/gotmc/ivi"
)

type StandardWaveform int

const (
	Sine StandardWaveform = iota
	Square
	Triangle
	RampUp
	RampDown
	DC
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
	inst     ivi.Instrument
	Channels *[]Channel
}

type Channel struct {
}

// OutputCount returns the number of outputs for the function generator.
func (fgen *Agilent33220) OutputCount() int {
	return 1
}

// GetAmplitude returns the amplitude in volts for the given channel index.
func (fgen *Agilent33220) GetAmplitude(ch int) (float64, error) {
	return getAmplitude(fgen.inst, ch)
}

func (fgen *Agilent33220) Amplitude(ch int, a float64) error {
	return nil
}

func getAmplitude(q ivi.Querier, ch int) (float64, error) {
	if ch != 0 {
		return 0, fmt.Errorf("Channel doesn't exist: %d", ch)
	}
	value, err := q.Query("VOLT?")
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

// Close closes the IVI Instrument.
func (fgen *Agilent33220) Close() error {
	return fgen.inst.Close()
}

// New creates a new Agilent33220 IVI Instrument.
func New(inst ivi.Instrument) (*Agilent33220, error) {
	var fgen Agilent33220
	fgen.inst = inst
	return &fgen, nil
}
