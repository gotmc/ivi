// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"fmt"
	"strconv"

	"github.com/gotmc/ivi"
)

type StandardWaveform int

const (
	Sine StandardWaveform = iota
	Square
	Ramp
	Pulse
	Noise
	DC
	User
)

var standardWaveforms = map[StandardWaveform]string{
	Sine:   "SIN",
	Square: "SQU",
	Ramp:   "RAMP",
	Pulse:  "PULS",
	Noise:  "NOIS",
	DC:     "DC",
	User:   "USER",
}

func (wave StandardWaveform) String() string {
	return standardWaveforms[wave]
}

type OperationMode int

const (
	Continuous OperationMode = iota
	Burst
)

type OutputMode int

const (
	Function OutputMode = iota
	Arbitrary
)

type Channel struct {
	id            int
	inst          ivi.Instrument
	operationMode OperationMode
	outputMode    OutputMode
}

func (ch *Channel) DisableOutput() error {
	_, err := ch.inst.WriteString("OUTP OFF")
	return err
}

func (ch *Channel) EnableOutput() error {
	_, err := ch.inst.WriteString("OUTP ON")
	return err
}

func (ch *Channel) SetAmplitude(amp float64) error {
	send := fmt.Sprintf("VOLT %f", amp)
	_, err := ch.inst.WriteString(send)
	return err
}

func (ch *Channel) SetDCOffset(amp float64) error {
	send := fmt.Sprintf("VOLT:OFFS %f", amp)
	_, err := ch.inst.WriteString(send)
	return err
}

func (ch *Channel) SetFrequency(freq float64) error {
	send := fmt.Sprintf("FREQ %f", freq)
	_, err := ch.inst.WriteString(send)
	return err
}

func (ch *Channel) Frequency() (float64, error) {
	s, err := ch.inst.Query("FREQ?\n")
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(s, 64)
}

func (ch *Channel) SetStandardWaveform(wave StandardWaveform) error {
	send := fmt.Sprintf("FUNC %s", wave)
	_, err := ch.inst.WriteString(send)
	return err
}

func (ch *Channel) SetOperationMode(mode OperationMode) {
	ch.operationMode = mode
}

func (ch *Channel) SetOutputMode(mode OutputMode) {
	ch.outputMode = mode
}
