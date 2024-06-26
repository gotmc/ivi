// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
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

// ArbitrarySampleRate returns the sample rate of the arbitrary waveforms the
// function generator produces. The units are samples per second.
//
// ArbitrarySampleRate is the getter for the read-write IviFgenArbWfm Attribute
// Arbitrary Sample Rate described in Section 6.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (d *Driver) ArbitrarySampleRate() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

// SetArbitrarySampleRate specifies the sample rate of the arbitrary waveforms
// the function generator produces. The units are samples per second.
//
// SetArbitrarySampleRate is the setter for the read-write IviFgenArbWfm
// Attribute Arbitrary Sample Rate described in Section 6.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) SetArbitrarySampleRate(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

// ArbWfmNumberWaveformsMax returns the maximum number of arbitrary waveforms
// that the function generator allows. The 33220A includes five built-in
// arbitrary waveforms stord in non-volative memory, four user-defined
// wavesforms stored in non-volatile memory, and one user-defined waveform in
// volatile memory.
//
// ArbWfmNumberWaveformsMax is the getter for the read-only IviFgenArbWfm
// Attribute Number Waveforms Max described in Section 6.2.5 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) ArbWfmNumberWaveformsMax() int {
	return MaxArbitraryWaveforms
}

// ArbWfmMaxSize returns the maximum number of points the function generator
// allows in an arbitrary waveform.
//
// ArbWfmMaxSize is the getter for the read-only IviFgenArbWfm Waveform Size
// Max described in Section 6.2.6 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) ArbWfmMaxSize() int {
	return MaxPointsWaveform
}

// ArbWfmMinSize returns the minimum number of points the function generator
// allows in an arbitrary waveform.
//
// ArbWfmMinSize is the getter for the read-only IviFgenArbWfm Waveform Size
// Min described in Section 6.2.7 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) ArbWfmMinSize() int {
	return MinPointsWaveform
}

// ArbWfmQuantum returns the quantum value the function generator allows. The
// size of each arbitrary waveform shall be a multiple of a quantum value.
//
// ArbWfmQuantum is the getter for the read-only IviFgenArbWfm Waveform Quantum
// described in Section 6.2.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) ArbWfmQuantum() int {
	return WaveformQuantum
}

func (ch *Channel) ArbitraryGain() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetArbitraryGain(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) ArbitraryOffset() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetArbitraryOffset(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) ArbitraryWaveformHandle() (int, error) {
	return 0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetArbitraryWaveformHandle(_ int) error {
	return ivi.ErrFunctionNotSupported
}
