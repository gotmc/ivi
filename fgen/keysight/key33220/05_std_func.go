// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// Amplitude reads the difference between the maximum and minimum waveform
// values, i.e., the peak-to-peak voltage value.
//
// Amplitude is the getter for the read-write IviFgenStdFunc Attribute
// Amplitude described in Section 5.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) Amplitude() (float64, error) {
	return query.Float64(ch.inst, "VOLT?")
}

// SetAmplitude specifies the difference between the maximum and minimum
// waveform values, i.e., the peak-to-peak voltage value.
//
// SetAmplitude is the setter for the read-write IviFgenStdFunc Attribute
// Amplitude described in Section 5.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetAmplitude(amp float64) error {
	return ch.inst.Command("VOLT %f VPP", amp)
}

// DCOffset reads the difference between the average of the maximum and minimum
// waveform values and the x-axis (0 volts).
//
// DCOffset is the getter for the read-write IviFgenStdFunc Attribute DC Offset
// described in Section 5.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) DCOffset() (float64, error) {
	return query.Float64(ch.inst, "VOLT:OFFS?")
}

// SetDCOffset sets the difference between the average of the maximum and
// minimum waveform values and the x-axis (0 volts).
//
// SetDCOffset is the setter for the read-write IviFgenStdFunc Attribute DC
// Offset described in Section 5.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetDCOffset(amp float64) error {
	return ch.inst.Command("VOLT:OFFS %f", amp)
}

// DutyCycleHigh reads the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value.
//
// DutyCycle is the getter for the read-write IviFgenStdFunc Attribute Duty
// Cycle High described in Section 5.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) DutyCycleHigh() (float64, error) {
	return query.Float64(ch.inst, "FUNC:SQU:DCYC?")
}

// SetDutyCycleHigh sets the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value.
//
// SetDutyCycle is the setter for the read-write IviFgenStdFunc Attribute Duty
// Cycle High described in Section 5.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetDutyCycleHigh(duty float64) error {
	return ch.inst.Command("FUNC:SQU:DCYC %f", duty)
}

// Frequency reads the number of waveform cycles generated in one second (i.e.,
// Hz). Frequency is not applicable for a DC waveform.
//
// Frequency is the getter for the read-write IviFgenStdFunc Attribute
// Frequency described in Section 5.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) Frequency() (float64, error) {
	return query.Float64(ch.inst, "FREQ?")
}

// SetFrequency sets the number of waveform cycles generated in one second
// (i.e., Hz). Frequency is not applicable for a DC waveform.
//
// SetFrequency is the setter for the read-write IviFgenStdFunc Attribute
// Frequency described in Section 5.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetFrequency(freq float64) error {
	return ch.inst.Command("FREQ %f", freq)
}

// StartPhase reads the start phase of the standard waveform the function
// generator produces. When the Waveform attribute is set to Waveform DC, this
// attribute does not affect signal output. The units are degrees.
//
// StartPhase is the getter for the read-write IviFgenStdFunc Attribute Start
// Phase described in Section 5.2.5 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) StartPhase() (float64, error) {
	return query.Float64(ch.inst, "PHAS?")
}

// SetStartPhase writes the start phase of the standard waveform the function
// generator produces. When the Waveform attribute is set to Waveform DC, this
// attribute does not affect signal output. The units are degrees.
//
// SetStartPhase is the setter for the read-write IviFgenStdFunc Attribute
// Start Phase described in Section 5.2.5 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetStartPhase(freq float64) error {
	return ch.inst.Command("PHAS %f", freq)
}

// StandardWaveform determines if one of the IVI Standard Waveforms is being output by
// the function generator. If not, an error is returned.
//
// StandardWaveform is the getter for the read-write IviFgenStdFunc Attribute Waveform
// described in Section 5.2.6 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) StandardWaveform() (fgen.StandardWaveform, error) {
	var wave fgen.StandardWaveform

	s, err := query.String(ch.inst, "FUNC?")
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
		symm, err := query.Float64(ch.inst, "FUNC:RAMP:SYMM?")
		if err != nil {
			return wave, fmt.Errorf(
				"unable to get symmetry to determine standard waveform: %s",
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
			return wave, fmt.Errorf("unable to determine waveform type RAMP with SYMM %f", symm)
		}
	}

	return wave, fmt.Errorf("unable to determine standard waveform type: %s", s)
}

// SetStandardWaveform specifies which standard waveform the function generator
// produces.
//
// SetStandardWaveform is the setter for the read-write IviFgenStdFunc
// Attribute Waveform described in Section 5.2.6 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetStandardWaveform(wave fgen.StandardWaveform) error {
	// FIXME(mdr): May need to change the phase offset in order to match the
	// waveforms shown in Figure 5-1 of IVI-4.3: IviFgen Class Specification.
	return ch.inst.Command(waveformCommand[wave])
}

// ConfigureStandardWaveform configures the attributes of the function
// generator that affect standard waveform generation.
//
// ConfigureStandardWaveform is the method that implements the Configure
// Standard Waveform function described in Section 5.3.1 of IVI-4.3: IviFgen
// Class Specification.
func (ch *Channel) ConfigureStandardWaveform(
	wave fgen.StandardWaveform,
	amp float64,
	offset float64,
	freq float64,
	phase float64,
) error {
	if err := ch.inst.Command(waveformApplyCommand[wave], freq, amp, offset); err != nil {
		return err
	}

	return ch.inst.Command("PHAS %.4f", phase)
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
