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

// OperationMode provides the defined values for the Operation Mode defined in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
type OperationMode int

// Continuous and Burst are the available Operation Modes. "A burst consists of
// a discrete number of waveform cycles. the user uses the attribute of the
// IviFgenTrigger Extension Group to configure the trigger, and the attributes
// of the IviFgenBurst extension group to configure how the function generator
// produces bursts.
const (
	Continuous OperationMode = iota
	Burst
)

// OutputMode determines how the function generator produces waveforms. This
// attribute determines which extension group's functions and attributes are
// used to configure the waveform the function generator produces. OutputMode
// implements the defined values for the Output Mode read-write attribute
// defined in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
type OutputMode int

// Function indicates the IVI driver uses the attributes and functions of the
// IviFgenStdFunc extension group. Arbitrary indicates the IVI driver uses the
// attributes and functions of the IviFgenArbWfm extension group.
const (
	Function OutputMode = iota
	Arbitrary
)

// Channel represents a repeated capability of an output channel for the
// function generator.
type Channel struct {
	id            int
	inst          ivi.Instrument
	operationMode OperationMode
	outputMode    OutputMode
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel. OperationMode implements the
// getter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OperationMode() (OperationMode, error) {
	var mode OperationMode
	s, err := ch.queryString("BURS:STAT?\n")
	if err != nil {
		return mode, fmt.Errorf("error getting operation mode: %s", err)
	}
	switch s {
	case "OFF":
		return Continuous, nil
	case "ON":
		return Burst, nil
	default:
		return mode, fmt.Errorf("error determining operation mode from fgen: %s", s)
	}
}

// SetOperationMode specifies whether the function generator should produce a
// continuous or burst output on the channel. SetOperationMode implements the
// setter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetOperationMode(mode OperationMode) {
	ch.operationMode = mode
}

// OutputEnabled determines if the output channel is enabled or disabled.
// OutputEnabled is the getter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return ch.queryBool("OUTP?\n")
}

// SetOutputEnabled sets the output channel to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	var send string
	if v {
		send = "OUTP ON\n"
	} else {
		send = "OUTP OFF\n"
	}
	_, err := ch.inst.WriteString(send)
	return err
}

// DisableOutput is a convenience function for setting the Output Enabled
// attribute to false.
func (ch *Channel) DisableOutput() error {
	return ch.SetOutputEnabled(false)
}

// EnableOutput is a convenience function for setting the Output Enabled
// attribute to true.
func (ch *Channel) EnableOutput() error {
	return ch.SetOutputEnabled(true)
}

// OutputImpedance return the output channel's impedance in ohms.
// OutputImpedance is the getter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.4 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputImpedance() (float64, error) {
	return ch.queryFloat64("OUTP:LOAD?\n")
}

// SetOutputImpedance sets the output channel's impedance in ohms.
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(impedance float64) error {
	return ch.setFloat64("OUTP:LOAD %f\n", impedance)
}

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

func (ch *Channel) setFloat64(cmd string, value float64) error {
	return ivi.SetFloat64(ch.inst, cmd, value)
}

func (ch *Channel) queryBool(query string) (bool, error) {
	return ivi.QueryBool(ch.inst, query)
}

func (ch *Channel) queryFloat64(query string) (float64, error) {
	return ivi.QueryFloat64(ch.inst, query)
}

func (ch *Channel) queryString(query string) (string, error) {
	return ivi.QueryString(ch.inst, query)
}
