// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

/*
# Section 4 IviDCPwrBase Capability Group

## Section 4.1 IviDCPwrBase Overview

The IviDCPwrBase capability group supports the most basic DC power supply
capabilities. The user can enable or disable outputs, specify the DC voltage to
generate, specify output limits, and control the behavior of the power supply
when the output is greater than or equal to one of the limits.

## Section 4.2 IviDCPwrBase Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type        | Access | AppliesTo   |
| ------- | ---------------------- | ----------- | ------ | ----------- |
|   4.2.1 | Current Limit          | Real64      | R/W    | Channel     |
|   4.2.2 | Current Limit Behavior | Int32       | R/W    | Channel     |
|   4.2.3 | Output Enabled         | Bool        | R/W    | Channel     |
|   4.2.4 | OVP Enabled            | Bool        | R/W    | Channel     |
|   4.2.5 | OVP Limit              | Real64      | R/W    | Channel     |
|   4.2.6 | Voltage Level          | Real64      | R/W    | Channel     |
|   4.2.7 | Output Channel Count   | Int32       | RO     | DCPwrOutput |
|   4.2.8 | Output Channel Item    | DCPwrOutput | RO     | DCPwrOutput |
|   4.2.9 | Output Channel Name    | String      | RO     | DCPwrOutput |

## Section 4.3 IviDCPwrBase Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

4.3.1 void Outputs[].ConfigureCurrentLimit(CurrentLimitBehavior behavior,
                                           Double limit);
4.3.3 void Outputs[].ConfigureRange(Ivi.DCPwr.RangeType rangeType, Double range);
4.3.4 void Outputs[].ConfigureOVP(Boolean enabled, Double limit);
4.3.7 Double Outputs[].QueryCurrentLimitMax(Double voltageLevel);
4.3.8 Double Outputs[].QueryVoltageLevelMax(Double currentLimit);
4.3.9 Boolean Outputs[].QueryState(OutputState outputState);
4.3.10 void Outputs[].ResetOutputProtection();

*/

// Base provides the interface for the IviDCPwrBase capability group.
type Base interface {
	OutputChannelCount() int
	// OutputChannelItem(name string) (BaseChannel, error)
}

// BaseChannel provides the interface for the channel repeated capability for
// the IviDCPwrBase capability group.
type BaseChannel interface {
	Name() string
	CurrentLimit() (float64, error)
	SetCurrentLimit(limit float64) error
	CurrentLimitBehavior() (CurrentLimitBehavior, error)
	SetCurrentLimitBehavior(behavior CurrentLimitBehavior) error
	OutputEnabled() (bool, error)
	SetOutputEnabled(b bool) error
	DisableOutput() error
	EnableOutput() error
	OVPEnabled() (bool, error)
	SetOVPEnabled(b bool) error
	DisableOVP() error
	EnableOVP() error
	VoltageLevel() (float64, error)
	SetVoltageLevel(level float64) error
	ConfigureCurrentLimit(behavior CurrentLimitBehavior, limit float64) error
	ConfigureOutputRange(rt RangeType, rng float64) error
	ConfigureOVP(b bool, limit float64) error
	QueryCurrentLimitMax(voltage float64) (float64, error)
	QueryVoltageLevelMax(currentLimit float64) (float64, error)
	QueryOutputState(os OutputState) (bool, error)
	ResetOutputProtection() error
}
