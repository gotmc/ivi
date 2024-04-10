// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// Confirm the driver implements the IviFgenStdFunc extension group.
var _ fgen.StdFuncChannel = (*Channel)(nil)

// Amplitude reads the difference between the maximum and minimum waveform
// values, i.e., the peak-to-peak voltage value.
//
// Amplitude is the getter for the read-write IviFgenStdFunc Attribute
// Amplitude described in Section 5.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) Amplitude() (float64, error) {
	return query.Float64(ch.inst, "AMPL?")
}

// SetAmplitude specifies the difference between the maximum and minimum
// waveform values, i.e., the peak-to-peak voltage value.
//
// SetAmplitude is the setter for the read-write IviFgenStdFunc Attribute
// Amplitude described in Section 5.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetAmplitude(amp float64) error {
	return ch.inst.Command("AMPL %f VP", amp)
}

// DCOffset reads the difference between the average of the maximum and minimum
// waveform values and the x-axis (0 volts).
//
// DCOffset is the getter for the read-write IviFgenStdFunc Attribute DC Offset
// described in Section 5.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) DCOffset() (float64, error) {
	return query.Float64(ch.inst, "OFFS?")
}

// SetDCOffset sets the difference between the average of the maximum and
// minimum waveform values and the x-axis (0 volts).
//
// SetDCOffset is the setter for the read-write IviFgenStdFunc Attribute DC
// Offset described in Section 5.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetDCOffset(amp float64) error {
	return ch.inst.Command("OFFS %f", amp)
}

// DutyCycleHigh reads the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value.
//
// DutyCycle is the getter for the read-write IviFgenStdFunc Attribute Duty
// Cycle High described in Section 5.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) DutyCycleHigh() (float64, error) {
	return 0.0, errors.New("duty cycle not yet implemented")
}

// SetDutyCycleHigh sets the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value.
//
// SetDutyCycle is the setter for the read-write IviFgenStdFunc Attribute Duty
// Cycle High described in Section 5.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetDutyCycleHigh(duty float64) error {
	return errors.New("not yet implemented; difficult in ds345")
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
	return query.Float64(ch.inst, "PHSE?")
}

// SetStartPhase writes the start phase of the standard waveform the function
// generator produces. When the Waveform attribute is set to Waveform DC, this
// attribute does not affect signal output. The units are degrees.
//
// SetStartPhase is the setter for the read-write IviFgenStdFunc Attribute
// Start Phase described in Section 5.2.5 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetStartPhase(freq float64) error {
	return ch.inst.Command("PHSE %f", freq)
}

// StandardWaveform determines if one of the IVI Standard Waveforms is being
// output by the function generator. If not, an error is returned.
//
// StandardWaveform is the getter for the read-write IviFgenStdFunc Attribute
// Waveform described in Section 5.2.6 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) StandardWaveform() (fgen.StandardWaveform, error) {
	var wave fgen.StandardWaveform
	s, err := query.String(ch.inst, "FUNC?")
	if err != nil {
		return wave, err
	}
	s = strings.TrimSpace(s)
	switch s {
	case "0":
		return fgen.Sine, nil
	case "1":
		return fgen.Square, nil
	case "2":
		return fgen.Triangle, nil
	case "3":
		invrt, err := query.String(ch.inst, "INVT?")
		if err != nil {
			return wave, fmt.Errorf("unable to determine ramp up vs ramp down: %s", err)
		}
		switch invrt {
		case "0":
			return fgen.RampUp, nil
		case "1":
			return fgen.RampDown, nil
		}
	}
	return wave, fmt.Errorf("unable to determine standard waveform type: %s", s)
}

// SetStandardWaveform specifies which standard waveform the function
// generator produces.
//
// SetStandardWaveform is the setter for the read-write IviFgenStdFunc
// Attribute Waveform described in Section 5.2.6 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetStandardWaveform(wave fgen.StandardWaveform) error {
	// FIXME(mdr): May need to change the phase offset in order to match the
	// waveforms shown in Figure 5-1 of IVI-4.3: IviFgen Class Specification.
	if wave == fgen.DC {
		return errors.New("dc standard waveform not implemented")
	}
	return ch.inst.Command(waveformCommand[wave])
}

var waveformCommand = map[fgen.StandardWaveform]string{
	fgen.Sine:     "FUNC0",
	fgen.Square:   "FUNC1",
	fgen.Triangle: "FUNC2",
	fgen.RampUp:   "FUNC3; INVT0",
	fgen.RampDown: "FUNC3; INVT1",
}

// ConfigureStandardWaveform configures the attributes of the function
// generator that affect standard waveform generation.
//
// ConfigureStandardWaveform is the method that implements the Configure
// Standard Waveform function described in Section 5.3.1 of IVI-4.3: IviFgen
// Class Specification.
func (ch *Channel) ConfigureStandardWaveform(wave fgen.StandardWaveform, amp float64,
	offset float64, freq float64, phase float64) error {
	if wave == fgen.DC {
		return errors.New("dc standard waveform not implemented")
	}
	return ch.inst.Command(waveformApplyCommand[wave], freq, amp, offset, phase)
}

var waveformApplyCommand = map[fgen.StandardWaveform]string{
	fgen.Sine:     "FUNC 0; FREQ %.4f; AMPL %.4fVP; OFFS %.4f; PHSE %.3f",
	fgen.Square:   "FUNC 1; FREQ %.4f; AMPL %.4fVP; OFFS %.4f; PHSE %.3f",
	fgen.Triangle: "FUNC 2; FREQ %.4f; AMPL %.4fVP; OFFS %.4f; PHSE %.3f",
	fgen.RampUp:   "FUNC 3; INVT 0; FREQ %.4f; AMPL %.4fVP; OFFS %.4f; PHSE %.3f",
	fgen.RampDown: "FUNC 3; INVT 1; FREQ %.4f; AMPL %.4fVP; OFFS %.4f; PHSE %.3f",
}
