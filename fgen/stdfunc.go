// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// StandardWaveform models the defined values for the Standard Waveform defined
// in Section 5.2.6 of IVI-4.3: IviFgen Class Specification.
type StandardWaveform int

// These are the available standard waveforms.
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
