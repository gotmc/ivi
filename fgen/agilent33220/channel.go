// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotmc/ivi"
)

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
	_, err := ch.inst.WriteString("OUTP OFF\n")
	return err
}

func (ch *Channel) EnableOutput() error {
	_, err := ch.inst.WriteString("OUTP ON\n")
	return err
}

// Amplitude provides the getter for the IviFgenStdFunc Attribute Amplitude
// specified in IVI-4.3: IviFgen Class Specification, Section 5.2.1 Amplitude.
func (ch *Channel) Amplitude() (float64, error) {
	return ch.queryFloat64("VOLT?\n")
}

// SetAmplitude provides the setter for the IviFgenStdFunc Attribute Amplitude
// specified in IVI-4.3: IviFgen Class Specification, Section 5.2.1 Amplitude.
func (ch *Channel) SetAmplitude(amp float64) error {
	return ch.setFloat64("VOLT %f VPP\n", amp)
}

// DCOffset provides the getter for the IviFgenStdFunc Attribute DC Offset
// specified in IVI-4.3: IviFgen Class Specification, Section 5.2.2 DC Offset.
func (ch *Channel) DCOffset() (float64, error) {
	return ch.queryFloat64("VOLT:OFFS?\n")
}

// SetDCOffset provides the setter for the IviFgenStdFunc Attribute DC Offset
// specified in IVI-4.3: IviFgen Class Specification, Section 5.2.2 DC Offset.
func (ch *Channel) SetDCOffset(amp float64) error {
	return ch.setFloat64("VOLT:OFFS %f\n", amp)
}

// DutyCycle provides the getter for the IviFgenStdFunc Attribute Duty Cycle
// High specified in IVI-4.3: IviFgen Class Specification, Section 5.2.3 Duty
// Cycle High.
func (ch *Channel) DutyCycle() (float64, error) {
	return ch.queryFloat64("FUNC:SQU:DCYC?\n")
}

// SetDutyCycle provides the setter for the IviFgenStdFunc Attribute Duty Cycle
// High specified in IVI-4.3: IviFgen Class Specification, Section 5.2.3 Duty
// Cycle High.
func (ch *Channel) SetDutyCycle(duty float64) error {
	return ch.setFloat64("FUNC:SQU:DCYC %f\n", duty)
}

func (ch *Channel) Frequency() (float64, error) {
	return ch.queryFloat64("FREQ?\n")
}

func (ch *Channel) SetFrequency(freq float64) error {
	return ch.setFloat64("FREQ %f\n", freq)
}

func (ch *Channel) StandardWaveform() (ivi.StandardWaveform, error) {
	var wave ivi.StandardWaveform
	s, err := ch.queryString("FUNC?\n")
	if err != nil {
		return wave, err
	}
	s = strings.TrimSpace(s)
	switch s {
	case "SIN":
		return ivi.Sine, nil
	case "SQU":
		return ivi.Square, nil
	case "RAMP":
		symm, err := ch.queryFloat64("FUNC:RAMP:SYMM?\n")
		if err != nil {
			return wave, fmt.Errorf("unable to get symmetry to determine standard waveform: %s", err)
		}
		switch symm {
		case 0.0:
			return ivi.RampDown, nil
		case 100.0:
			return ivi.RampUp, nil
		case 50.0:
			return ivi.Triangle, nil
		default:
			return wave, fmt.Errorf("unable to determine waveform type RAMP with SYMM %s", symm)
		}
	}
	return wave, fmt.Errorf("unable to determine standard waveform type: %s", s)
}

func (ch *Channel) SetStandardWaveform(wave ivi.StandardWaveform) error {
	var send string
	// FIXME(mdr): May need to change the phase offset in order to match the
	// waveforms shown in Figure 5-1 of IVI-4.3: IviFgen Class Specification.
	switch wave {
	case ivi.Sine:
		send = "FUNC SIN\n"
	case ivi.Square:
		send = "FUNC SQU\n"
	case ivi.Triangle:
		send = "FUNC RAMP; FUNC:RAMP:SYMM 50\n"
	case ivi.RampUp:
		send = "FUNC RAMP; FUNC:RAMP:SYMM 100\n"
	case ivi.RampDown:
		send = "FUNC RAMP; FUNC:RAMP:SYMM 0\n"
	case ivi.DC:
		send = "FUNC DC\n"
	}
	_, err := ch.inst.WriteString(send)
	return err
}

func (ch *Channel) SetOperationMode(mode OperationMode) {
	ch.operationMode = mode
}

func (ch *Channel) SetOutputMode(mode OutputMode) {
	ch.outputMode = mode
}

func (ch *Channel) queryFloat64(query string) (float64, error) {
	return queryFloat64(ch.inst, query)
}

func queryFloat64(q ivi.Querier, query string) (float64, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func (ch *Channel) setFloat64(cmd string, value float64) error {
	return setFloat64(ch.inst, cmd, value)
}

func setFloat64(sw ivi.StringWriter, cmd string, value float64) error {
	send := fmt.Sprintf(cmd, value)
	_, err := sw.WriteString(send)
	return err
}

func (ch *Channel) queryString(query string) (string, error) {
	return queryString(ch.inst, query)
}

func queryString(q ivi.Querier, query string) (string, error) {
	return q.Query(query)
}
