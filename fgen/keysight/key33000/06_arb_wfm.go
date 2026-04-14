// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33000

import (
	"context"

	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

const (
	MaxPointsWaveform     = 16777216 // 16 MSa
	MinPointsWaveform     = 8
	MaxArbitraryWaveforms = 512
	WaveformQuantum       = 1
)

// ArbitrarySampleRate returns the sample rate of the arbitrary waveforms the
// function generator produces. The units are samples per second.
//
// ArbitrarySampleRate is the getter for the read-write IviFgenArbWfm Attribute
// Arbitrary Sample Rate described in Section 6.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (d *Driver) ArbitrarySampleRate(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, d.channels[0].srcPrefix()+"FUNC:ARB:SRAT?")
}

// SetArbitrarySampleRate specifies the sample rate of the arbitrary waveforms
// the function generator produces. The units are samples per second.
//
// SetArbitrarySampleRate is the setter for the read-write IviFgenArbWfm
// Attribute Arbitrary Sample Rate described in Section 6.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (d *Driver) SetArbitrarySampleRate(ctx context.Context, rate float64) error {
	return d.inst.Command(ctx, d.channels[0].srcPrefix()+"FUNC:ARB:SRAT %f", rate)
}

// ArbWfmNumberWaveformsMax returns the maximum number of arbitrary waveforms
// that the function generator allows.
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

// ArbWfmQuantum returns the quantum value the function generator allows.
//
// ArbWfmQuantum is the getter for the read-only IviFgenArbWfm Waveform
// Quantum described in Section 6.2.8 of IVI-4.3: IviFgen Class Specification.
func (d *Driver) ArbWfmQuantum() int {
	return WaveformQuantum
}

// ArbitraryGain returns the gain of the arbitrary waveform the function
// generator produces. This value is unitless.
//
// ArbitraryGain is the getter for the read-write IviFgenArbWfm attribute
// Arbitrary Gain described in Section 6.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) ArbitraryGain(ctx context.Context) (float64, error) {
	amp, err := ch.Amplitude(ctx)
	if err != nil {
		return 0.0, err
	}
	return amp / 2, nil
}

// SetArbitraryGain sets the gain of the arbitrary waveform the function
// generator produces. This value is unitless.
//
// SetArbitraryGain is the setter for the read-write IviFgenArbWfm attribute
// Arbitrary Gain described in Section 6.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetArbitraryGain(ctx context.Context, gain float64) error {
	return ch.SetAmplitude(ctx, 2*gain)
}

// ArbitraryOffset returns the offset of the arbitrary waveform the function
// generator produces. The units are volts.
//
// ArbitraryOffset is the getter for the read-write IviFgenArbWfm attribute
// Arbitrary Offset described in Section 6.2.2 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) ArbitraryOffset(ctx context.Context) (float64, error) {
	return ch.DCOffset(ctx)
}

// SetArbitraryOffset sets the offset of the arbitrary waveform the function
// generator produces. The units are volts.
//
// SetArbitraryOffset is the setter for the read-write IviFgenArbWfm attribute
// Arbitrary Offset described in Section 6.2.2 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetArbitraryOffset(ctx context.Context, offset float64) error {
	return ch.SetDCOffset(ctx, offset)
}

func (ch *Channel) ArbitraryWaveformHandle(_ context.Context) (int, error) {
	return 0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetArbitraryWaveformHandle(_ context.Context, _ int) error {
	return ivi.ErrFunctionNotSupported
}
