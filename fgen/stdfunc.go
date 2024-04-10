// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// StdFuncChannel provides the interface for the channel repeated capability for
// the IviFgenStdFunc capability group.
type StdFuncChannel interface {
	Amplitude() (float64, error)
	SetAmplitude(amp float64) error
	DCOffset() (float64, error)
	SetDCOffset(offset float64) error
	DutyCycleHigh() (float64, error)
	SetDutyCycleHigh(duty float64) error
	Frequency() (float64, error)
	SetFrequency(freq float64) error
	StartPhase() (float64, error)
	SetStartPhase(start float64) error
	StandardWaveform() (StandardWaveform, error)
	SetStandardWaveform(StandardWaveform) error
	ConfigureStandardWaveform(wave StandardWaveform, amp, offset, freq, phase float64) error
}

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
