// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dp800

import (
	"context"
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

func (d *Driver) OutputChannelCount() int {
	return len(d.Channels)
}

func (ch *Channel) Name() string {
	return ch.name
}

// CurrentLimit determines the output current limit. The units are Amps.
//
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit(ctx context.Context) (float64, error) {
	return 0.0, fmt.Errorf("CurrentLimit: %w", ivi.ErrNotImplemented)
}

// SetCurrentLimit specifies the output current limit in Amperes.
//
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(ctx context.Context, limit float64) error {
	return fmt.Errorf("SetCurrentLimit: %w", ivi.ErrNotImplemented)
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior(ctx context.Context) (dcpwr.CurrentLimitBehavior, error) {
	return dcpwr.CurrentRegulate, fmt.Errorf("CurrentLimitBehavior: %w", ivi.ErrNotImplemented)
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the current limit
// attribute. The E3631A only supports the CurrentRegulate behavior, so
// attempting to set CurrentTrip will result in an error.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetCurrentLimitBehavior(
	ctx context.Context,
	behavior dcpwr.CurrentLimitBehavior,
) error {
	if behavior == dcpwr.CurrentTrip {
		return fmt.Errorf("SetCurrentLimitBehavior: CurrentTrip not supported. %w", ivi.ErrValueNotSupported)
	}
	return nil
}

// OutputEnabled determines if the given output channel is enabled or disabled.
//
// OutputEnabled is the getter for the read-write IviDCPwrBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OutputEnabled(ctx context.Context) (bool, error) {
	return query.Boolf(ctx, ch.inst, ":OUTP? %s", ch.name)
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(ctx context.Context, v bool) error {
	if v {
		return ivi.Set(ch.inst, ":OUTP %s,ON", ch.name)
	}
	return ivi.Set(ch.inst, ":OUTP %s,OFF", ch.name)
}

// DisableOutput is a convenience function for setting the Output Enabled
// attribute to false.
func (ch *Channel) DisableOutput(ctx context.Context) error {
	return ch.SetOutputEnabled(ctx, false)
}

// EnableOutput is a convenience function for setting the Output Enabled
// attribute to true.
func (ch *Channel) EnableOutput(ctx context.Context) error {
	return ch.SetOutputEnabled(ctx, true)
}

// OVPEnabled always returns false for the E3631A since it doesn't have OVP.
//
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled(ctx context.Context) (bool, error) {
	return false, nil
}

// SetOVPEnabled always returns an error for the E3631A since it doesn't have
// Over-Voltage Protection (OVP).
//
// SetOVPEnabled is the setter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(ctx context.Context, v bool) error {
	return fmt.Errorf("SetOVPEnabled: %w", dcpwr.ErrOVPUnsupported)
}

// DisableOVP is a convenience function for disabling Over-Voltage Protection
// (OVP). DisableOVP always returns nil for the E3631A since support OVP.
func (ch *Channel) DisableOVP(ctx context.Context) error {
	return nil
}

// EnableOVP is a convenience function for enabling Over-Voltage Protection
// (OVP). EnableOVP always returns an error for the E3631A since support OVP.
func (ch *Channel) EnableOVP(ctx context.Context) error {
	return fmt.Errorf("EnableOVP: %w", dcpwr.ErrOVPUnsupported)
}

// OVPLimit returns an error, since the E3631A doesn't support Over-Voltage
// Protection (OVP).
//
// OVPLimit is the getter for the read-write IviDWPwrBase Attribute OVP Limit
// described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit(ctx context.Context) (float64, error) {
	return 0, fmt.Errorf("OVPLimit: %w", dcpwr.ErrOVPUnsupported)
}

// SetOVPLimit returns an error since the E3631A doesn't support Over-Voltage
// Protection (OVP).
//
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(ctx context.Context, limit float64) error {
	return fmt.Errorf("SetOVPLimit: %w", dcpwr.ErrOVPUnsupported)
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate in Volts.
//
// VoltageLevel is the getter for the read-write IviDCPwrBase Attribute Voltage
// Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel(ctx context.Context) (float64, error) {
	return query.Float64f(ctx, ch.inst, ":APPL? %s,VOLT", ch.name)
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate in Volts.
//
// SetVoltageLevel is the setter for the read-write IviDCPwrBase Attribute
// Voltage Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetVoltageLevel(ctx context.Context, amp float64) error {
	return ch.inst.Command(ctx, ":APPL %s,%f", ch.name, amp)
}

// ConfigureCurrentLimit configures the current limit. It specifies the output
// current limit value and the behavior of the power supply when the output
// current is greater than or equal to that value.
//
// ConfigureCurrentLimit implements the IviDCPwrBase function described in
// Section 4.3.1 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureCurrentLimit(
	ctx context.Context,
	behavior dcpwr.CurrentLimitBehavior,
	limit float64,
) error {
	return fmt.Errorf("ConfigureCurrentLimit: %w", ivi.ErrNotImplemented)
}

// ConfigureOutputRange configures either the power supply’s output voltage or
// current range on an output. Setting a voltage range can invalidate a
// previously configured current range. Setting a current range can invalidate
// a previously configured voltage range.
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
func (ch *Channel) ConfigureOutputRange(
	ctx context.Context,
	rt dcpwr.RangeType,
	rng float64,
) error {
	return fmt.Errorf("ConfigureOutputRange: %w", ivi.ErrNotImplemented)
}

// ConfigureOVP configures the Over-Voltage Protection (OVP). It specifies the
// over-voltage limit and the behavior of the power supply when the output
// voltage is greater than or equal to that value. When the Enabled parameter
// is False, the Limit parameter does not affect the instrument’s behavior, and
// the driver does not set the OVP Limit attribute.
//
// ConfigureOVP implements the IviDCPwrBase function described in Section 4.3.4
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ConfigureOVP(ctx context.Context, b bool, limit float64) error {
	return fmt.Errorf("ConfigureOVP: %w", ivi.ErrNotImplemented)
}

// QueryCurrentLimitMax returns the maximum programmable current limit that the
// power supply accepts for a particular voltage level on an output.
//
// QueryCurrentLimitMax implements the IviDCPwrBase function described in
// Section 4.3.7 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryCurrentLimitMax(ctx context.Context, voltage float64) (float64, error) {
	return 0.0, fmt.Errorf("QueryCurrentLimitMax: %w", ivi.ErrNotImplemented)
}

// QueryVoltageLevelMax returns the maximum programmable voltage level that the
// power supply accepts for a particular current limit on an output.
//
// QueryVoltageLevelMax implements the IviDCPwrBase function described in
// Section 4.3.8 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryVoltageLevelMax(
	ctx context.Context,
	currentLimit float64,
) (float64, error) {
	return 0.0, fmt.Errorf("QueryVoltageLevelMax: %w", ivi.ErrNotImplemented)
}

// QueryOutputState returns whether the power supply is in a particular output
// state.
//
// QueryOutputState implements the IviDCPwrBase function described in Section
// 4.3.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryOutputState(ctx context.Context, os dcpwr.OutputState) (bool, error) {
	return false, fmt.Errorf("QueryOutputState: %w", ivi.ErrNotImplemented)
}

// ResetOutputProtection resets the power supply output protection after an
// over-voltage or over-current condition occurs.
//
// ResetOutputProtection implements the IviDCPwrBase function described in
// Section 4.3.10 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ResetOutputProtection(ctx context.Context) error {
	return fmt.Errorf("ResetOutputProtection: %w", ivi.ErrNotImplemented)
}
