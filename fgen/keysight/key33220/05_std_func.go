// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

func (ch *Channel) Amplitude() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "VOLT?")
}

func (ch *Channel) SetAmplitude(amp float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "VOLT %f VPP", amp)
}

func (ch *Channel) DCOffset() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "VOLT:OFFS?")
}

func (ch *Channel) SetDCOffset(offset float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "VOLT:OFFS %f", offset)
}

func (ch *Channel) DutyCycleHigh() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "FUNC:SQU:DCYC?")
}

func (ch *Channel) SetDutyCycleHigh(duty float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "FUNC:SQU:DCYC %f", duty)
}

func (ch *Channel) Frequency() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "FREQ?")
}

func (ch *Channel) SetFrequency(freq float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "FREQ %f", freq)
}

func (ch *Channel) StartPhase() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, "PHAS?")
}

func (ch *Channel) SetStartPhase(phase float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "PHAS %f", phase)
}

// StandardWaveform determines if one of the IVI Standard Waveforms is being
// output by the function generator.
func (ch *Channel) StandardWaveform() (fgen.StandardWaveform, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	var wave fgen.StandardWaveform

	s, err := query.String(ctx, ch.inst, "FUNC?")
	if err != nil {
		return wave, err
	}

	s = strings.TrimSpace(s)
	switch s {
	case "SIN":
		return fgen.Sine, nil
	case "SQU":
		return fgen.Square, nil
	case "DC":
		return fgen.DC, nil
	case "RAMP":
		symm, err := query.Float64(ctx, ch.inst, "FUNC:RAMP:SYMM?")
		if err != nil {
			return wave, fmt.Errorf(
				"unable to get symmetry to determine standard waveform: %w",
				err,
			)
		}

		switch symm {
		case 0.0:
			return fgen.RampDown, nil
		case 100.0:
			return fgen.RampUp, nil
		case 50.0:
			return fgen.Triangle, nil
		default:
			return wave, fmt.Errorf(
				"unable to determine waveform type RAMP with SYMM %f", symm,
			)
		}
	}

	return wave, fmt.Errorf("unable to determine standard waveform type: %s", s)
}

// SetStandardWaveform specifies which standard waveform the function generator
// produces.
func (ch *Channel) SetStandardWaveform(wave fgen.StandardWaveform) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(waveformCommand, wave)
	if err != nil {
		return fmt.Errorf("SetStandardWaveform: %w", err)
	}

	return ch.inst.Command(ctx, cmd)
}

// ConfigureStandardWaveform configures the attributes of the function
// generator that affect standard waveform generation.
func (ch *Channel) ConfigureStandardWaveform(
	wave fgen.StandardWaveform,
	amp float64,
	offset float64,
	freq float64,
	phase float64,
) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	format, err := ivi.LookupSCPI(waveformApplyCommand, wave)
	if err != nil {
		return fmt.Errorf("ConfigureStandardWaveform: %w", err)
	}

	if err := ch.inst.Command(ctx, format, freq, amp, offset); err != nil {
		return err
	}

	return ch.inst.Command(ctx, "PHAS %.4f", phase)
}

var waveformCommand = map[fgen.StandardWaveform]string{
	fgen.Sine:     "FUNC SIN",
	fgen.Square:   "FUNC SQU",
	fgen.Triangle: "FUNC RAMP; RAMP:SYMM 50",
	fgen.RampUp:   "FUNC RAMP; RAMP:SYMM 100",
	fgen.RampDown: "FUNC RAMP; RAMP:SYMM 0",
	fgen.DC:       "FUNC DC",
}

var waveformApplyCommand = map[fgen.StandardWaveform]string{
	fgen.Sine:     "APPL:SIN %.4f, %.4f, %.4f",
	fgen.Square:   "APPL:SQU %.4f, %.4f, %.4f",
	fgen.Triangle: "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 50",
	fgen.RampUp:   "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 100",
	fgen.RampDown: "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 0",
	fgen.DC:       "APPL:DC %.4f, %.4f, %.4f",
}
