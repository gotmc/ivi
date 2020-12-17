// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

import "errors"

// Base provides the interface required for the IviDCPwrBase capability group.
type Base interface {
	// Channels() ([]*BaseChannel, error)
	// Channel(name string) (*BaseChannel, error)
	// ChannelByID(id int) (*BaseChannel, error)
	// ChannelCount() int
	OutputCount() int
}

// BaseChannel provides the interface for the channel repeated capability for
// the IviDCPwrBase capability group.
type BaseChannel interface {
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

// CurrentLimitBehavior provides the defined values for the Current Limit
// Behavior defined in Section 4.2.2 of IVI-4.4: IviDCPwr Class Specification.
type CurrentLimitBehavior int

// CurrentTrip and CurrentRegulate are the available Current Limit Behaviors.
// In CurrentTrip behavior, the power supply disables the output when the
// output current is equal to or greater than the value of the Current Limit
// attribute. In CurrentRegulate behavior, the power supply restricts the
// output voltage such that the output current is not greater than the value of
// the Current Limit attribute.
const (
	CurrentTrip CurrentLimitBehavior = iota
	CurrentRegulate
)

// RangeType provides the defined values for the Output Range Type defined in
// Section 4.3.3 of IVI-4.4: IviDCPwr Class Specification.
type RangeType int

// Available range types. In VoltageRange, the voltage range is set by the
// Range parameter. In CurrentRange, the current range is set by the Range
// parameter.
const (
	VoltageRange RangeType = iota
	CurrentRange
)

// OutputState provides the defined values for the state of the output defined
// in Section 4.3.9 of IVI-4.4: IviDCPwr Class Specification.
type OutputState int

// Available output states that can be queried.
const (
	ConstantVoltage OutputState = iota
	ConstantCurrent
	OverVoltage
	OverCurrent
	Unregulated
)

// Error codes related to the IviDCPwr Class Specification.
var (
	ErrNotImplemented error = errors.New("not implemented in ivi driver")
	ErrOVPUnsupported error = errors.New("ovp not supported")
)
