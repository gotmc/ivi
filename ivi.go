// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

// Instrument provides the interface required for all IVI Instruments.
type Instrument interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	StringWriter
	Querier
}

type StringWriter interface {
	WriteString(s string) (n int, err error)
}

type Querier interface {
	Query(s string) (value string, err error)
}

type StandardWaveform int

const (
	Sine StandardWaveform = iota
	Square
	Triangle
	RampUp
	RampDown
	DC
)

var standardWaveforms = map[StandardWaveform]string{
	Sine:     "Sine",
	Square:   "Square",
	Triangle: "Triangle",
	RampUp:   "Ramp Up",
	RampDown: "Ramp Down",
	DC:       "DC",
}

func (wave StandardWaveform) String() string {
	return standardWaveforms[wave]
}
