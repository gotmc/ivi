// Copyright (c) 2017-2021 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package pmx

import (
	"errors"

	"github.com/gotmc/ivi/dcpwr"
)

// Make sure the IviDCPwrBase capability group has been implemented.
var _ dcpwr.Base = (*PMX)(nil)
var _ dcpwr.BaseChannel = (*Channel)(nil)

// CurrentLimit determines the output current limit. The units are Amps.
//
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	return ch.QueryFloat64("CURR?\n")
}

// SetCurrentLimit specifies the output current limit in Amperes.
//
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	if ch.currentLimitBehavior == dcpwr.CurrentRegulate {
		return ch.Set("CURR %f;:CURR:PROT MAX\n", limit)
	} else if ch.currentLimitBehavior == dcpwr.CurrentTrip {
		return ch.Set("CURR %f;:CURR:PROT %f\n", limit, limit)
	}
	return errors.New("current limit behavior not set")
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (dcpwr.CurrentLimitBehavior, error) {
	return ch.currentLimitBehavior, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. Since the PMX series always has the OCP enabled, if the current
// limit behavior is set to regulate, the OCP is set to MAX (110% of the max
// current value of the device).  Whereas, setting the current limit behavior
// to trip, the OCP is set equal to the current limit.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetCurrentLimitBehavior(behavior dcpwr.CurrentLimitBehavior) error {
	if behavior == dcpwr.CurrentRegulate {
		ch.currentLimitBehavior = dcpwr.CurrentRegulate
		return ch.Set("CURR:PROT MAX\n")
	} else if behavior == dcpwr.CurrentTrip {
		ch.currentLimitBehavior = dcpwr.CurrentTrip
		limit, err := ch.QueryFloat64("CURR?\n")
		if err != nil {
			return err
		}
		return ch.Set("CURR:PROT %f\n", limit)
	}
	return errors.New("unknown current limit behavior")
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.
//
// OutputEnabled is the getter for the read-write IviDCPwrBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return ch.QueryBool("OUTP?\n")
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	if v {
		return ch.Set("OUTP 1\n")
	}
	return ch.Set("OUTP 0\n")
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

// OVPEnabled specifies whether the power supply provides Over-Voltage
// Protection (OVP).
//
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	max, err := ch.QueryFloat64("VOLT:PROT? MAX\n")
	if err != nil {
		return false, err
	}
	ovp, err := ch.QueryFloat64("VOLT:PROT?\n")
	if err != nil {
		return false, err
	}
	if ovp == max {
		return false, nil
	}
	return true, nil
}

// SetOVPEnabled enables or disables the Over-Voltage Protection (OVP). Since
// the OVP is always enabled on the PMX power supply, if false, the PMX's OVP
// is set to its maximum value, which is 110% of the maximum output voltage.
// voltage protection set.
//
// SetOVPEnabled is the setter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	if v == false {
		return ch.Set("VOLT:PROT MAX\n")
	}
	return nil
}

// DisableOVP is a convenience function for disabling Over-Voltage Protection
// (OVP).
func (ch *Channel) DisableOVP() error {
	return dcpwr.ErrNotImplemented
}

// EnableOVP is a convenience function for enabling Over-Voltage Protection
// (OVP).
func (ch *Channel) EnableOVP() error {
	return dcpwr.ErrNotImplemented
}

// OVPLimit returns the current Over-Voltage Protection (OVP) value.
//
// OPVLimit is the getter for the read-write IviDWPwrBase Attribute OVP Limit
// described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	return ch.QueryFloat64("VOLT:PROT?\n")
}

// SetOVPLimit specifies the voltage the power supply allows. The units are
// Volts. If the OVP Enabled attribute is set to True, the power supply
// disables the output when the output voltage is greater than or equal to the
// value of this attribute. If the OVP Enabled is set to False, this attribute
// does not affect the behavior of the instrument.
//
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	return ch.Set("VOLT:PROT %f\n", limit)
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate. The units are Volts.
//
// VoltageLevel is the getter for the read-write IviDCPwrBase Attribute Voltage
// Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	return ch.QueryFloat64("VOLT?\n")
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate. The units are Volts.
//
// SetVoltageLevel is the setter for the read-write IviDCPwrBase Attribute
// Voltage Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetVoltageLevel(level float64) error {
	return ch.Set("VOLT %f\n", level)
}

// OutputCount returns the number of available output channels.
//
// OutputCount is the getter for the read-only IviDCPwrBase Attribute Output
// Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (dev *PMX) OutputCount() int {
	return len(dev.Channels)
}

// ConfigureCurrentLimit specifies the output current limit value and the
// behavior of the power supply when the output current is greater than or
// equal to that value.
//
// ConfigureCurrentLimit implements the IviDCPwrBase Configure Current Limit
// function described in Section 4.3.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) ConfigureCurrentLimit(behavior dcpwr.CurrentLimitBehavior, limit float64) error {
	if behavior == dcpwr.CurrentRegulate {
		ch.currentLimitBehavior = dcpwr.CurrentRegulate
		return ch.Set("CURR %f;:CURR:PROT MAX\n", limit)
	} else if behavior == dcpwr.CurrentTrip {
		ch.currentLimitBehavior = dcpwr.CurrentTrip
		return ch.Set("CURR %f;:CURR:PROT %f\n", limit, limit)
	}
	return errors.New("unknown current limit behavior")

}

// ConfigureOutputRange configures either the power supply’s output voltage or
// current range on an output. Setting a voltage range can invalidate a
// previously configured current range. Setting a current range can invalidate
// a previously configured voltage range.
//
// ConfigureOutputRange implements the IviDCPwrBase function described in
// Section 4.3.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOutputRange(rt dcpwr.RangeType, rng float64) error {
	return dcpwr.ErrNotImplemented
}

// ConfigureOVP specifies the over-voltage limit and the behavior of the power
// supply when the output voltage is greater than or equal to that value.
//
// ConfigureOVP implements the IviDCPwrBase Configure OVP function described in
// Section 4.3.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOVP(enabled bool, limit float64) error {
	if enabled == false {
		return ch.Set("VOLT:PROT MAX\n")
	}
	return ch.Set("VOLT:PROT %f\n", limit)
}

// QueryCurrentLimitMax returns the maximum programmable current limit that the
// power supply accepts for a particular voltage level on an output.
//
// QueryCurrentLimitMax implements the IviDCPwrBase function described in
// Section 4.3.7 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryCurrentLimitMax(voltage float64) (float64, error) {
	return 0.0, dcpwr.ErrNotImplemented
}

// QueryVoltageLevelMax returns the maximum programmable voltage level that the
// power supply accepts for a particular current limit on an output.
//
// QueryVoltageLevelMax implements the IviDCPwrBase function described in
// Section 4.3.8 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryVoltageLevelMax(currentLimit float64) (float64, error) {
	return 0.0, dcpwr.ErrNotImplemented
}

// QueryOutputState returns whether the power supply is in a particular output
// state.
//
// QueryOutputState implements the IviDCPwrBase function described in Section
// 4.3.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryOutputState(os dcpwr.OutputState) (bool, error) {
	return false, dcpwr.ErrNotImplemented
}

// ResetOutputProtection resets the power supply output protection after an
// over-voltage or over-current condition occurs.
//
// ResetOutputProtection implements the IviDCPwrBase function described in
// Section 4.3.10 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ResetOutputProtection() error {
	return dcpwr.ErrNotImplemented
}
