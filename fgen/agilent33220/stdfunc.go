// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent33220

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
)

// Amplitude reads the difference between the maximum and minimum waveform
// values, i.e., the peak-to-peak voltage value. Amplitude is the getter for
// the read-write IviFgenStdFunc Attribute Amplitude described in Section 5.2.1
// of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) Amplitude() (float64, error) {
	return ch.queryFloat64("VOLT?\n")
}

// SetAmplitude specifies the difference between the maximum and minimum
// waveform values, i.e., the peak-to-peak voltage value. Amplitude is the
// setter for the read-write IviFgenStdFunc Attribute Amplitude described in
// Section 5.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetAmplitude(amp float64) error {
	return ch.setFloat64("VOLT %f VPP\n", amp)
}

// DCOffset reads the difference between the average of the maximum and minimum
// waveform values and the x-axis (0 volts). DCOffset is the getter for
// the read-write IviFgenStdFunc Attribute DC Offset described in Section 5.2.2
// of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) DCOffset() (float64, error) {
	return ch.queryFloat64("VOLT:OFFS?\n")
}

// SetDCOffset sets the difference between the average of the maximum and
// minimum waveform values and the x-axis (0 volts). SetDCOffset is the setter
// for the read-write IviFgenStdFunc Attribute DC Offset described in Section
// 5.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetDCOffset(amp float64) error {
	return ch.setFloat64("VOLT:OFFS %f\n", amp)
}

// DutyCycle reads the percentage of time, specified as 0-100, during one cycle
// for which the square wave is at its high value.  DutyCycle is the getter for
// the read-write IviFgenStdFunc Attribute Duty Cycle High described in Section
// 5.2.3 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) DutyCycle() (float64, error) {
	return ch.queryFloat64("FUNC:SQU:DCYC?\n")
}

// SetDutyCycle sets the percentage of time, specified as 0-100, during one
// cycle for which the square wave is at its high value. SetDutyCycle is the
// setter for the read-write IviFgenStdFunc Attribute Duty Cycle High described
// in Section 5.2.3 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetDutyCycle(duty float64) error {
	return ch.setFloat64("FUNC:SQU:DCYC %f\n", duty)
}

// Frequency reads the number of waveform cycles generated in one second (i.e.,
// Hz). Frequency is not applicable for a DC waveform.  Frequency is the getter
// for the read-write IviFgenStdFunc Attribute Frequency described in Section
// 5.2.4 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) Frequency() (float64, error) {
	return ch.queryFloat64("FREQ?\n")
}

// SetFrequency sets the number of waveform cycles generated in one second
// (i.e., Hz). Frequency is not applicable for a DC waveform. SetFrequency is
// the setter for the read-write IviFgenStdFunc Attribute Frequency described
// in Section 5.2.4 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetFrequency(freq float64) error {
	return ch.setFloat64("FREQ %f\n", freq)
}

// StandardWaveform determines if one of the IVI Standard Waveforms is being
// output by the function generator. If not, an error is returned.
// StandwardWaveform is the getter for the read-write IviFgenStdFunc Attribute
// Waveform described in Section 5.2.6 of IVI-4.3: IviFgen Class Specification.
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
	case "DC":
		return ivi.DC, nil
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
			return wave, fmt.Errorf("unable to determine waveform type RAMP with SYMM %f", symm)
		}
	}
	return wave, fmt.Errorf("unable to determine standard waveform type: %s", s)
}

// SetStandardWaveform specifices which standard waveform the function
// generator produces.  SetStandwardWaveform is the setter for the read-write
// IviFgenStdFunc Attribute Waveform described in Section 5.2.6 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStandardWaveform(wave ivi.StandardWaveform) error {
	// FIXME(mdr): May need to change the phase offset in order to match the
	// waveforms shown in Figure 5-1 of IVI-4.3: IviFgen Class Specification.
	_, err := ch.inst.WriteString(waveformCommand[wave])
	return err
}

var waveformCommand = map[ivi.StandardWaveform]string{
	ivi.Sine:     "FUNC SIN",
	ivi.Square:   "FUNC SQU",
	ivi.Triangle: "FUNC RAMP; RAMP:SYMM 50",
	ivi.RampUp:   "FUNC RAMP; RAMP:SYMM 100",
	ivi.RampDown: "FUNC RAMP; RAMP:SYMM 0",
	ivi.DC:       "FUNC DC",
}

var waveformApplyCommand = map[ivi.StandardWaveform]string{
	ivi.Sine:     "APPL:SIN %.4f, %.4f, %.4f\n",
	ivi.Square:   "APPL:SQU %.4f, %.4f, %.4f\n",
	ivi.Triangle: "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 50\n",
	ivi.RampUp:   "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 100\n",
	ivi.RampDown: "APPL:RAMP %.4f, %.4f, %.4f;:FUNC:RAMP:SYMM 0\n",
	ivi.DC:       "APPL:DC %.4f, %.4f, %.4f\n",
}

// ConfigureStandardWaveform configures the attributes of the function
// generator that affect standard waveform generation.
// ConfigureStandardWaveform is the method that implements the Configure
// Standard Waveform function described in Section 5.3.1 of IVI-4.3: IviFgen
// Class Specification.
func (ch *Channel) ConfigureStandardWaveform(wave ivi.StandardWaveform, amp float64,
	offset float64, freq float64, phase float64) error {
	cmd := fmt.Sprintf(waveformApplyCommand[wave], freq, amp, offset)
	_, err := ch.inst.WriteString(cmd)
	return err
}
