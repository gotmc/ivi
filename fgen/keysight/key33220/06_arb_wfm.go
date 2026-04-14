// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import "github.com/gotmc/ivi"

const (
	MaxPointsWaveform     = 65536
	MinPointsWaveform     = 1
	MaxArbitraryWaveforms = 10
	WaveformQuantum       = 1
)

func (d *Driver) ArbitrarySampleRate() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (d *Driver) SetArbitrarySampleRate(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (d *Driver) ArbWfmNumberWaveformsMax() int { return MaxArbitraryWaveforms }
func (d *Driver) ArbWfmMaxSize() int            { return MaxPointsWaveform }
func (d *Driver) ArbWfmMinSize() int            { return MinPointsWaveform }
func (d *Driver) ArbWfmQuantum() int            { return WaveformQuantum }

func (ch *Channel) ArbitraryGain() (float64, error) {
	amp, err := ch.Amplitude()
	if err != nil {
		return 0.0, err
	}
	return amp / 2, nil
}

func (ch *Channel) SetArbitraryGain(gain float64) error {
	return ch.SetAmplitude(2 * gain)
}

func (ch *Channel) ArbitraryOffset() (float64, error) {
	return ch.DCOffset()
}

func (ch *Channel) SetArbitraryOffset(offset float64) error {
	return ch.SetDCOffset(offset)
}

func (ch *Channel) ArbitraryWaveformHandle() (int, error) {
	return 0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetArbitraryWaveformHandle(_ int) error {
	return ivi.ErrFunctionNotSupported
}
