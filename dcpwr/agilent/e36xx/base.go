// Copyright (c) 2017-2019 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36xx

import (
	"errors"

	"github.com/gotmc/ivi/dcpwr"
)

// CurrentLimit determines the output current limit. The units are Amps.
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	return ch.queryLimit(currentQuery)
}

// SetCurrentLimit specifies the output current limit. The units are Amp.s
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	return ch.Set("INST %s;:CURR %f\n", ch.Name(), limit)
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior.
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (dcpwr.CurrentLimitBehavior, error) {
	return dcpwr.Regulate, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior, so
// attempting to set CurrentTrip will result in an error.  CurrentLimitBehavior
// implements the getter for the read-write IviDCPwrBase Attribute Current
// Limit Behavior described in Section 4.2.2 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetCurrentLimitBehavior(behavior dcpwr.CurrentLimitBehavior) error {
	if behavior == dcpwr.Trip {
		return errors.New("current trip is not supported")
	}
	return nil
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.  OutputEnabled is the getter for the read-write IviDCPwrBase
// Attribute Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return ch.QueryBool("OUTP?\n")
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	if v {
		return ch.Set("OUTP ON\n")
	}
	return ch.Set("OUTP OFF\n")
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

// OVPEnabled always returns false for the E3631A since it doesn't have OVP.
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	return false, nil
}

// SetOVPEnabled always returns an error for the E3631A since it doesn't have
// OVP.  SetOVPEnabled is the setter for the read-write IviFgenBase Attribute
// OVP Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	return errors.New("OVP not supported")
}

// OVPLimit returns an error, since the E3631A doesn't support OVP. OVPLimit is
// the getter for the read-write IviDWPwrBase Attribute OVP Limit described in
// Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	return 0, errors.New("OVP not supported")
}

// SetOVPLimit returns an error since the E3631A doesn't support OVP.
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	return errors.New("OVP not supported")
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate. The units are Volts.  VoltageLevel is the getter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	return ch.queryLimit(voltageQuery)
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate. The units are Volts.  SetVoltageLevel is the setter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetVoltageLevel(amp float64) error {
	return ch.Set("APPL %s, %f\n", ch.Name(), amp)
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviDCPwrBase Attribute Output Channel Count
// described in Section 4.2.7 of IVI-4.4: IviDCPwr Class Specification.
func (dev *E36xx) OutputCount() int {
	return len(dev.Channels)
}
