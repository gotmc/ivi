// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

import "context"

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
	CurrentLimit(ctx context.Context) (float64, error)
	SetCurrentLimit(ctx context.Context, limit float64) error
	CurrentLimitBehavior(ctx context.Context) (CurrentLimitBehavior, error)
	SetCurrentLimitBehavior(ctx context.Context, behavior CurrentLimitBehavior) error
	OutputEnabled(ctx context.Context) (bool, error)
	SetOutputEnabled(ctx context.Context, b bool) error
	DisableOutput(ctx context.Context) error
	EnableOutput(ctx context.Context) error
	OVPEnabled(ctx context.Context) (bool, error)
	SetOVPEnabled(ctx context.Context, b bool) error
	DisableOVP(ctx context.Context) error
	EnableOVP(ctx context.Context) error
	OVPLimit(ctx context.Context) (float64, error)
	SetOVPLimit(ctx context.Context, limit float64) error
	VoltageLevel(ctx context.Context) (float64, error)
	SetVoltageLevel(ctx context.Context, level float64) error
	ConfigureCurrentLimit(ctx context.Context, behavior CurrentLimitBehavior, limit float64) error
	ConfigureOutputRange(ctx context.Context, rt RangeType, rng float64) error
	ConfigureOVP(ctx context.Context, b bool, limit float64) error
	QueryCurrentLimitMax(ctx context.Context, voltage float64) (float64, error)
	QueryVoltageLevelMax(ctx context.Context, currentLimit float64) (float64, error)
	QueryOutputState(ctx context.Context, os OutputState) (bool, error)
	ResetOutputProtection(ctx context.Context) error
}
