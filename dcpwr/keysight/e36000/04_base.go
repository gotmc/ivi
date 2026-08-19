// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36000

import (
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

// OutputChannelCount returns the number of available output channels.
//
// OutputChannelCount is the getter for the read-only IviDCPwrBase Attribute
// Output Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (d *Driver) OutputChannelCount() int {
	return len(d.channels)
}

// Name returns the channel's symbolic name (e.g., "P6V", "Output").
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

	return query.Float64(ctx, ch.inst, ch.getCmd("CURR?"))
}

// SetCurrentLimit specifies the output current limit in Amperes.
//
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ch.setCmd("CURR %.4f"), limit)
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. On models with over-current protection the behavior is read from
// CURRent:PROTection:STATe, since enabling that protection is what makes the
// supply trip its output off at the limit instead of regulating there. Models
// without it, such as the E3631A, can only regulate.
//
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (dcpwr.CurrentLimitBehavior, error) {
	if !ch.protection.ocp {
		return dcpwr.CurrentRegulate, nil
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	tripping, err := query.Bool(ctx, ch.inst, ch.getCmd("CURR:PROT:STAT?"))
	if err != nil {
		return 0, fmt.Errorf("CurrentLimitBehavior: %w", err)
	}

	if tripping {
		return dcpwr.CurrentTrip, nil
	}

	return dcpwr.CurrentRegulate, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the current limit
// attribute. On models with over-current protection both behaviors are
// available and are selected through CURRent:PROTection:STATe. Models without
// it, such as the E3631A, only support CurrentRegulate, so asking them for
// CurrentTrip returns [ivi.ErrValueNotSupported].
//
// SetCurrentLimitBehavior implements the setter for the read-write
// IviDCPwrBase Attribute Current Limit Behavior described in Section 4.2.2 of
// IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetCurrentLimitBehavior(
	behavior dcpwr.CurrentLimitBehavior,
) error {
	if !ch.protection.ocp {
		if behavior == dcpwr.CurrentTrip {
			return fmt.Errorf(
				"SetCurrentLimitBehavior: CurrentTrip not supported. %w",
				ivi.ErrValueNotSupported,
			)
		}

		return nil
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	state := "OFF"
	if behavior == dcpwr.CurrentTrip {
		state = "ON"
	}

	return ch.inst.Command(ctx, ch.setCmd("CURR:PROT:STAT %s"), state)
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.
//
// OutputEnabled is the getter for the read-write IviDCPwrBase Attribute Output
// Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Bool(ctx, ch.inst, ch.outputGetCmd("OUTP?"))
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
//
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	if v {
		return ch.inst.Command(ctx, ch.outputSetCmd("OUTP ON"))
	}

	return ch.inst.Command(ctx, ch.outputSetCmd("OUTP OFF"))
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

// OVPEnabled determines whether Over-Voltage Protection (OVP) is enabled. It
// always returns false on models without OVP, such as the E3631A.
//
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	if !ch.protection.ovp {
		return false, nil
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	enabled, err := query.Bool(ctx, ch.inst, ch.getCmd("VOLT:PROT:STAT?"))
	if err != nil {
		return false, fmt.Errorf("OVPEnabled: %w", err)
	}

	return enabled, nil
}

// SetOVPEnabled enables or disables Over-Voltage Protection (OVP). It returns
// [dcpwr.ErrOVPUnsupported] on models without OVP, such as the E3631A.
//
// SetOVPEnabled is the setter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	if !ch.protection.ovp {
		return fmt.Errorf("SetOVPEnabled: %w", dcpwr.ErrOVPUnsupported)
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	state := "OFF"
	if v {
		state = "ON"
	}

	return ch.inst.Command(ctx, ch.setCmd("VOLT:PROT:STAT %s"), state)
}

// DisableOVP is a convenience function for setting the OVP Enabled attribute
// to false. It returns nil on models without Over-Voltage Protection (OVP),
// such as the E3631A, since their protection is already off.
func (ch *Channel) DisableOVP() error {
	if !ch.protection.ovp {
		return nil
	}

	return ch.SetOVPEnabled(false)
}

// EnableOVP is a convenience function for setting the OVP Enabled attribute to
// true. It returns [dcpwr.ErrOVPUnsupported] on models without Over-Voltage
// Protection (OVP), such as the E3631A.
func (ch *Channel) EnableOVP() error {
	return ch.SetOVPEnabled(true)
}

// OVPLimit returns the voltage, in Volts, at which Over-Voltage Protection
// (OVP) trips. It returns [dcpwr.ErrOVPUnsupported] on models without OVP,
// such as the E3631A.
//
// OVPLimit is the getter for the read-write IviDWPwrBase Attribute OVP Limit
// described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	if !ch.protection.ovp {
		return 0, fmt.Errorf("OVPLimit: %w", dcpwr.ErrOVPUnsupported)
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, ch.getCmd("VOLT:PROT?"))
}

// SetOVPLimit specifies the voltage, in Volts, at which Over-Voltage
// Protection (OVP) trips. It returns [dcpwr.ErrOVPUnsupported] on models
// without OVP, such as the E3631A.
//
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	if !ch.protection.ovp {
		return fmt.Errorf("SetOVPLimit: %w", dcpwr.ErrOVPUnsupported)
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ch.setCmd("VOLT:PROT %.4f"), limit)
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate in Volts.
//
// VoltageLevel is the getter for the read-write IviDCPwrBase Attribute Voltage
// Level described in Section 4.2.6 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64(ctx, ch.inst, ch.getCmd("VOLT?"))
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

	return ch.inst.Command(ctx, ch.setCmd("VOLT %.4f"), level)
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
	if err := ch.SetCurrentLimitBehavior(behavior); err != nil {
		return fmt.Errorf("ConfigureCurrentLimit: %w", err)
	}

	if err := ch.SetCurrentLimit(limit); err != nil {
		return fmt.Errorf("ConfigureCurrentLimit: %w", err)
	}

	return nil
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
func (ch *Channel) ConfigureOutputRange(_ dcpwr.RangeType, _ float64) error {
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
func (ch *Channel) ConfigureOVP(enabled bool, limit float64) error {
	if !ch.protection.ovp {
		return fmt.Errorf("ConfigureOVP: %w", dcpwr.ErrOVPUnsupported)
	}

	// Per Section 4.3.4, the limit is only applied when OVP is being enabled.
	// Set it before enabling so the protection never arms at a stale level.
	if enabled {
		if err := ch.SetOVPLimit(limit); err != nil {
			return fmt.Errorf("ConfigureOVP: %w", err)
		}
	}

	if err := ch.SetOVPEnabled(enabled); err != nil {
		return fmt.Errorf("ConfigureOVP: %w", err)
	}

	return nil
}

// QueryCurrentLimitMax returns the maximum programmable current limit that the
// power supply accepts for a particular voltage level on an output.
//
// QueryCurrentLimitMax implements the IviDCPwrBase function described in
// Section 4.3.7 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryCurrentLimitMax(_ float64) (float64, error) {
	return 0.0, fmt.Errorf("QueryCurrentLimitMax: %w", ivi.ErrNotImplemented)
}

// QueryVoltageLevelMax returns the maximum programmable voltage level that the
// power supply accepts for a particular current limit on an output.
//
// QueryVoltageLevelMax implements the IviDCPwrBase function described in
// Section 4.3.8 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryVoltageLevelMax(_ float64) (float64, error) {
	return 0.0, fmt.Errorf("QueryVoltageLevelMax: %w", ivi.ErrNotImplemented)
}

// Bit masks for the STATus:OPERation and STATus:QUEStionable condition
// registers. The two registers report different things, so outputStateBit
// returns the register to read along with the mask to apply.
const (
	// operationCV is set while the output is in constant voltage mode.
	operationCV = 1 << 8
	// operationCC is set while the output is in constant current mode.
	operationCC = 1 << 10
	// questionableOV is set while the output is disabled by over-voltage
	// protection.
	questionableOV = 1 << 0
	// questionableOC is set while the output is disabled by over-current
	// protection.
	questionableOC = 1 << 1
	// questionableUNR is set while the output is unregulated.
	questionableUNR = 1 << 10
)

// outputStateBit maps an IVI output state onto the status register that
// reports it and the bit mask that selects it.
func outputStateBit(os dcpwr.OutputState) (register string, mask int, err error) {
	switch os {
	case dcpwr.ConstantVoltage:
		return "STAT:OPER:COND?", operationCV, nil
	case dcpwr.ConstantCurrent:
		return "STAT:OPER:COND?", operationCC, nil
	case dcpwr.OverVoltage:
		return "STAT:QUES:COND?", questionableOV, nil
	case dcpwr.OverCurrent:
		return "STAT:QUES:COND?", questionableOC, nil
	case dcpwr.Unregulated:
		return "STAT:QUES:COND?", questionableUNR, nil
	}

	return "", 0, fmt.Errorf("output state %v: %w", os, ivi.ErrValueNotSupported)
}

// QueryOutputState returns whether the power supply is in a particular output
// state. The state is read from the STATus:OPERation and STATus:QUEStionable
// condition registers, so models that do not implement them return
// [ivi.ErrNotImplemented].
//
// QueryOutputState implements the IviDCPwrBase function described in Section
// 4.3.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) QueryOutputState(os dcpwr.OutputState) (bool, error) {
	if !ch.protection.statusRegisters {
		return false, fmt.Errorf("QueryOutputState: %w", ivi.ErrNotImplemented)
	}

	register, mask, err := outputStateBit(os)
	if err != nil {
		return false, fmt.Errorf("QueryOutputState: %w", err)
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	condition, err := query.Int(ctx, ch.inst, ch.getCmd(register))
	if err != nil {
		return false, fmt.Errorf("QueryOutputState: %w", err)
	}

	return condition&mask != 0, nil
}

// ResetOutputProtection resets the power supply output protection after an
// over-voltage or over-current condition occurs. Models that do not implement
// OUTPut:PROTection:CLEar return [ivi.ErrNotImplemented].
//
// ResetOutputProtection implements the IviDCPwrBase function described in
// Section 4.3.10 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) ResetOutputProtection() error {
	if !ch.protection.outputClear {
		return fmt.Errorf("ResetOutputProtection: %w", ivi.ErrNotImplemented)
	}

	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, ch.setCmd("OUTP:PROT:CLE"))
}
