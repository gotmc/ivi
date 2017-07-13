// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

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
