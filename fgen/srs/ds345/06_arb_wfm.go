// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

const (
	MaxPointsWaveform     = 16300
	MinPointsWaveform     = 8
	MaxArbitraryWaveforms = 10
	WaveformQuantum       = 8
)

func (d *Driver) ArbWfmNumberWaveformsMax() int {
	return MaxArbitraryWaveforms
}

func (d *Driver) ArbWfmMaxSize() int {
	return MaxPointsWaveform
}

func (d *Driver) ArbWfmMinSize() int {
	return MinPointsWaveform
}

func (d *Driver) ArbWfmQuantum() int {
	return WaveformQuantum
}
