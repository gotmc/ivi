// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dp800

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

func (d *Driver) OutputChannelCount() int {
	return len(d.channels)
}

func (ch *Channel) Name() string {
	return ch.name
}

// CurrentLimit determines the output current limit. The units are Amps.
//
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, ":SOUR%d:CURR?", ch.idx)
}

// SetCurrentLimit specifies the output current limit in Amperes.
//
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ":SOUR%d:CURR %f", ch.idx, limit)
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The DP800 supports both CurrentRegulate and CurrentTrip via the
// OCP feature.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior(
) (dcpwr.CurrentLimitBehavior, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	ocpEnabled, err := query.Boolf(ctx, ch.inst, ":OUTP:OCP? %s", ch.name)
	if err != nil {
		return 0, fmt.Errorf("CurrentLimitBehavior: %w", err)
	}

	if ocpEnabled {
		return dcpwr.CurrentTrip, nil
	}

	return dcpwr.CurrentRegulate, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the current limit
// attribute. When set to CurrentTrip, the DP800 enables OCP on the channel;
// when set to CurrentRegulate, OCP is disabled.
//
// SetCurrentLimitBehavior implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetCurrentLimitBehavior(
	behavior dcpwr.CurrentLimitBehavior,
) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	switch behavior {
	case dcpwr.CurrentRegulate:
		return ch.inst.Command(ctx, ":OUTP:OCP %s,OFF", ch.name)
	case dcpwr.CurrentTrip:
		return ch.inst.Command(ctx, ":OUTP:OCP %s,ON", ch.name)
	default:
		return fmt.Errorf(
			"SetCurrentLimitBehavior: %w", ivi.ErrValueNotSupported,
		)
	}
}

// OutputEnabled determines if the given output channel is enabled or disabled.
//
// OutputEnabled is the getter for the read-write IviDCPwrBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Boolf(ctx, ch.inst, ":OUTP? %s", ch.name)
}

// SetOutputEnabled sets the specified output channel to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	if v {
		return ivi.Set(ctx, ch.inst, ":OUTP %s,ON", ch.name)
	}

	return ivi.Set(ctx, ch.inst, ":OUTP %s,OFF", ch.name)
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

// OVPEnabled determines whether Over-Voltage Protection (OVP) is enabled on
// the specified channel.
//
// OVPEnabled is the getter for the read-write IviDCPwrBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Boolf(ctx, ch.inst, ":OUTP:OVP? %s", ch.name)
}

// SetOVPEnabled enables or disables Over-Voltage Protection (OVP) on the
// specified channel.
//
// SetOVPEnabled is the setter for the read-write IviDCPwrBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	if v {
		return ch.inst.Command(ctx, ":OUTP:OVP %s,ON", ch.name)
	}

	return ch.inst.Command(ctx, ":OUTP:OVP %s,OFF", ch.name)
}

// DisableOVP is a convenience function for disabling Over-Voltage Protection
// (OVP).
func (ch *Channel) DisableOVP() error {
	return ch.SetOVPEnabled(false)
}

// EnableOVP is a convenience function for enabling Over-Voltage Protection
// (OVP).
func (ch *Channel) EnableOVP() error {
	return ch.SetOVPEnabled(true)
}

// OVPLimit returns the Over-Voltage Protection (OVP) value in Volts.
//
// OVPLimit is the getter for the read-write IviDCPwrBase Attribute OVP Limit
// described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, ":OUTP:OVP:VAL? %s", ch.name)
}

// SetOVPLimit specifies the Over-Voltage Protection (OVP) value in Volts.
//
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ":OUTP:OVP:VAL %s,%f", ch.name, limit)
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate in Volts.
//
// VoltageLevel is the getter for the read-write IviDCPwrBase Attribute Voltage
// Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, ":SOUR%d:VOLT?", ch.idx)
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate in Volts.
//
// SetVoltageLevel is the setter for the read-write IviDCPwrBase Attribute
// Voltage Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetVoltageLevel(level float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ":SOUR%d:VOLT %f", ch.idx, level)
}

// ConfigureCurrentLimit configures the current limit. It specifies the output
// current limit value and the behavior of the power supply when the output
// current is greater than or equal to that value.
//
// ConfigureCurrentLimit implements the IviDCPwrBase function described in
// Section 4.3.1 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureCurrentLimit(
	behavior dcpwr.CurrentLimitBehavior,
	limit float64,
) error {
	if err := ch.SetCurrentLimit(limit); err != nil {
		return err
	}

	return ch.SetCurrentLimitBehavior(behavior)
}

// ConfigureOutputRange configures either the power supply's output voltage or
// current range on an output. The DP800 series automatically changes the range
// based on the values the user requests, so this function performs no
// communication with the instrument.
//
// ConfigureOutputRange implements the IviDCPwrBase function described in
// Section 4.3.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOutputRange(
	rt dcpwr.RangeType,
	rng float64,
) error {
	return nil
}

// ConfigureOVP configures the Over-Voltage Protection (OVP). It specifies the
// over-voltage limit and the behavior of the power supply when the output
// voltage is greater than or equal to that value. When the Enabled parameter
// is False, the Limit parameter does not affect the instrument's behavior, and
// the driver does not set the OVP Limit attribute.
//
// ConfigureOVP implements the IviDCPwrBase function described in Section 4.3.4
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOVP(enabled bool, limit float64) error {
	if err := ch.SetOVPEnabled(enabled); err != nil {
		return err
	}

	if enabled {
		return ch.SetOVPLimit(limit)
	}

	return nil
}

// QueryCurrentLimitMax returns the maximum programmable current limit that the
// power supply accepts for a particular voltage level on an output.
//
// QueryCurrentLimitMax implements the IviDCPwrBase function described in
// Section 4.3.7 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryCurrentLimitMax(
	voltage float64,
) (float64, error) {
	return ch.maxCurrent, nil
}

// QueryVoltageLevelMax returns the maximum programmable voltage level that the
// power supply accepts for a particular current limit on an output.
//
// QueryVoltageLevelMax implements the IviDCPwrBase function described in
// Section 4.3.8 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryVoltageLevelMax(
	currentLimit float64,
) (float64, error) {
	return ch.maxVoltage, nil
}

// QueryOutputState returns whether the power supply is in a particular output
// state. Uses the :OUTPut:CVCC? command which returns CV, CC, or UR.
//
// QueryOutputState implements the IviDCPwrBase function described in Section
// 4.3.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryOutputState(
	os dcpwr.OutputState,
) (bool, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	switch os {
	case dcpwr.ConstantVoltage:
		mode, err := query.Stringf(ctx, ch.inst, ":OUTP:CVCC? %s", ch.name)
		if err != nil {
			return false, fmt.Errorf("QueryOutputState: %w", err)
		}

		return strings.TrimSpace(mode) == "CV", nil
	case dcpwr.ConstantCurrent:
		mode, err := query.Stringf(ctx, ch.inst, ":OUTP:CVCC? %s", ch.name)
		if err != nil {
			return false, fmt.Errorf("QueryOutputState: %w", err)
		}

		return strings.TrimSpace(mode) == "CC", nil
	case dcpwr.Unregulated:
		mode, err := query.Stringf(ctx, ch.inst, ":OUTP:CVCC? %s", ch.name)
		if err != nil {
			return false, fmt.Errorf("QueryOutputState: %w", err)
		}

		return strings.TrimSpace(mode) == "UR", nil
	case dcpwr.OverVoltage:
		resp, err := query.Stringf(
			ctx, ch.inst, ":OUTP:OVP:QUES? %s", ch.name,
		)
		if err != nil {
			return false, fmt.Errorf("QueryOutputState: %w", err)
		}

		return strings.TrimSpace(resp) == "YES", nil
	case dcpwr.OverCurrent:
		resp, err := query.Stringf(
			ctx, ch.inst, ":OUTP:OCP:QUES? %s", ch.name,
		)
		if err != nil {
			return false, fmt.Errorf("QueryOutputState: %w", err)
		}

		return strings.TrimSpace(resp) == "YES", nil
	default:
		return false, fmt.Errorf(
			"QueryOutputState %v: %w", os, ivi.ErrValueNotSupported,
		)
	}
}

// ResetOutputProtection resets the power supply output protection after an
// over-voltage or over-current condition occurs.
//
// ResetOutputProtection implements the IviDCPwrBase function described in
// Section 4.3.10 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ResetOutputProtection() error {
	ctx, cancel := ch.newContext()
	defer cancel()

	if err := ch.inst.Command(ctx, ":OUTP:OVP:CLEAR %s", ch.name); err != nil {
		return fmt.Errorf("ResetOutputProtection (OVP): %w", err)
	}

	if err := ch.inst.Command(ctx, ":OUTP:OCP:CLEAR %s", ch.name); err != nil {
		return fmt.Errorf("ResetOutputProtection (OCP): %w", err)
	}

	return nil
}
