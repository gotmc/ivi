// Copyright (c) 2017-2023 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36xx

import (
	"errors"

	"github.com/gotmc/convert"
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

// Confirm that the output channel repeated capabilitiy implements the
// IviDCPwrBase and IviDCPwrMeasurement interfaces.
var _ dcpwr.BaseChannel = (*Channel)(nil)
var _ dcpwr.MeasurementChannel = (*Channel)(nil)

// Channel models the output channel repeated capabilitiy for the DC power
// supply output channel.
type Channel struct {
	name string
	inst ivi.Instrument
}

// CurrentLimit determines the output current limit. The units are Amps.
//
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	_, current, err := ch.readVoltageCurrent()
	return current, err
}

// SetCurrentLimit specifies the output current limit in Amperes.
//
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	return ivi.Set(ch.inst, "INST %s;:CURR %f\n", ch.name, limit)
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (dcpwr.CurrentLimitBehavior, error) {
	return dcpwr.CurrentRegulate, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the current limit
// attribute. The E3631A only supports the CurrentRegulate behavior, so
// attempting to set CurrentTrip will result in an error.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetCurrentLimitBehavior(behavior dcpwr.CurrentLimitBehavior) error {
	if behavior == dcpwr.CurrentTrip {
		return errors.New("current trip is not supported")
	}
	return nil
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.
//
// OutputEnabled is the getter for the read-write IviDCPwrBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return query.Bool(ch.inst, "OUTP?\n")
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	if v {
		return ivi.Set(ch.inst, "OUTP ON\n")
	}
	return ivi.Set(ch.inst, "OUTP OFF\n")
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
//
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	return false, nil
}

// SetOVPEnabled always returns an error for the E3631A since it doesn't have
// Over-Voltage Protection (OVP).
//
// SetOVPEnabled is the setter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	return dcpwr.ErrOVPUnsupported
}

// DisableOVP is a convenience function for disabling Over-Voltage Protection
// (OVP). DisableOVP always returns nil for the E3631A since support OVP.
func (ch *Channel) DisableOVP() error {
	return nil
}

// EnableOVP is a convenience function for enabling Over-Voltage Protection
// (OVP). EnableOVP always returns an error for the E3631A since support OVP.
func (ch *Channel) EnableOVP() error {
	return dcpwr.ErrOVPUnsupported
}

// OVPLimit returns an error, since the E3631A doesn't support Over-Voltage
// Protection (OVP).
//
// OVPLimit is the getter for the read-write IviDWPwrBase Attribute OVP Limit
// described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	return 0, dcpwr.ErrOVPUnsupported
}

// SetOVPLimit returns an error since the E3631A doesn't support Over-Voltage
// Protection (OVP).
//
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	return dcpwr.ErrOVPUnsupported
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate in Volts.
//
// VoltageLevel is the getter for the read-write IviDCPwrBase Attribute Voltage
// Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	voltage, _, err := ch.readVoltageCurrent()
	return voltage, err
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate in Volts.
//
// SetVoltageLevel is the setter for the read-write IviDCPwrBase Attribute
// Voltage Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetVoltageLevel(amp float64) error {
	return ivi.Set(ch.inst, "APPL %s, %f\n", ch.name, amp)
}

// ConfigureCurrentLimit configures the current limit. It specifies the output
// current limit value and the behavior of the power supply when the output
// current is greater than or equal to that value.
//
// ConfigureCurrentLimit implements the IviDCPwrBase function described in
// Section 4.3.1 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureCurrentLimit(behavior dcpwr.CurrentLimitBehavior, limit float64) error {
	return dcpwr.ErrNotImplemented
}

// ConfigureOutputRange configures either the power supply’s output voltage or
// current range on an output. Setting a voltage range can invalidate a
// previously configured current range. Setting a current range can invalidate
// a previously configured voltage range. The instrument driver coerces the
// range value to the closest value the instrument supports that is greater
// than or equal to the value specified.
//
// Some DC power supplies do not allow the user to explicitly specify an
// output’s range. Instead, they automatically change the range based on the
// values the user requests for the voltage level, OVP limit, and current
// limit. For instruments that automatically change the range, the
// ConfigureOutputRange function should perform range checking to verify that
// its input parameters are valid, but should not perform any communication
// with the instrument or set any attributes.
//
// ConfigureOutputRange implements the IviDCPwrBase function described in
// Section 4.3.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOutputRange(rt dcpwr.RangeType, rng float64) error {
	return dcpwr.ErrNotImplemented
}

// ConfigureOVP configures the Over-Voltage Protection (OVP). It specifies the
// over-voltage limit and the behavior of the power supply when the output
// voltage is greater than or equal to that value. When the Enabled parameter
// is False, the Limit parameter does not affect the instrument’s behavior, and
// the driver does not set the OVP Limit attribute.
//
// ConfigureOVP implements the IviDCPwrBase function described in Section 4.3.4
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOVP(b bool, limit float64) error {
	return dcpwr.ErrNotImplemented
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

// MeasureVoltage takes a measurement on the output signal and returns the
// measured voltage.
//
// MeasureVoltage implements the IviDCPwrMeasurement function Measure for the
// Voltage MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureVoltage() (float64, error) {
	return query.Float64f(ch.inst, "MEAS:VOLT? %s", ch.name)
}

// MeasureCurrent takes a measurement on the output signal and returns the
// measured current.
//
// MeasureCurrent implements the IviDCPwrMeasurement function Measure for the
// Current MeasurementType parameter described in Section 7.2.1 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) MeasureCurrent() (float64, error) {
	return query.Float64f(ch.inst, "MEAS:CURR? %s", ch.name)
}

// readVoltageCurrent queries the power supply's present voltage and current
// values for each output and returns a quoted string. The voltage and current
// are returned in sequence as shown in the sample string below (the quotation
// marks are returned as part of the string). If any output identifier is not
// specified, the voltage and the current of the currently selected output are
// returned. "5.000000,1.000000" In the above string, the first number 5.000000
// is the voltage limit value and the second number 1.000000 is the current
// limit value for the specified output.
func (ch *Channel) readVoltageCurrent() (float64, float64, error) {
	s, err := query.Stringf(ch.inst, "APPL? %s", ch)
	if err != nil {
		return 0.0, 0.0, err
	}
	floats, err := convert.StringToNFloats(s, ",", 2)
	if err != nil {
		return 0.0, 0.0, err
	}
	return floats[0], floats[1], nil
}
