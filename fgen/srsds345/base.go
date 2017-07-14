// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package srsds345

import (
	"errors"
	"fmt"

	"github.com/gotmc/ivi/fgen"
)

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviFgenBase Attribute Output Count described in
// Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (a *SRSDS345) OutputCount() int {
	return a.outputCount
}

// OperationMode determines whether the function generator should produce a
// continuous or burst output on the channel. OperationMode implements the
// getter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) OperationMode() (fgen.OperationMode, error) {
	var mode fgen.OperationMode
	s, err := ch.queryString("MENA?\n")
	if err != nil {
		return mode, fmt.Errorf("error getting operation mode: %s", err)
	}
	switch s {
	case "0":
		return fgen.Continuous, nil
	case "1":
		mod, err := ch.queryString("MTYP?\n")
		if err != nil {
			return mode, fmt.Errorf("error determining modulation type: %s", err)
		}
		switch mod {
		case "5":
			return fgen.Burst, nil
		default:
			return mode, fmt.Errorf("error determining operation mode, mtyp = %s", mod)
		}
	default:
		return mode, fmt.Errorf("error determining operation mode from fgen: %s", s)
	}
}

// SetOperationMode specifies whether the function generator should produce a
// continuous or burst output on the channel. SetOperationMode implements the
// setter for the read-write IviFgenBase Attribute Operation Mode described in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetOperationMode(mode fgen.OperationMode) error {
	var cmd string
	switch mode {
	case fgen.Burst:
		cmd = "MENA1;MTYP5\n"
	case fgen.Continuous:
		cmd = "MENA0\n"
	}
	_, err := ch.inst.WriteString(cmd)
	return err
}

// OutputEnabled determines if the output channel is enabled or disabled.
// OutputEnabled is the getter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return false, errors.New("output enabled not implemented")
}

// SetOutputEnabled sets the output channel to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	return errors.New("set output enabled not implemented")
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
	return 50.0, nil
}

// SetOutputImpedance sets the output channel's impedance in ohms.
// SetOutputImpedance is the setter for the read-write IviFgenBase Attribute
// Output Impedance described in Section 4.2.3 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetOutputImpedance(impedance float64) error {
	if impedance != 50 {
		return errors.New("output impedance must be 50 ohms")
	}
	return nil
}
