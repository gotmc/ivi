// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// Base provides the interface required for the IviFgenBase capability group.
type Base interface {
	// Channels() ([]*BaseChannel, error)
	// Channel(name string) (*BaseChannel, error)
	// ChannelByID(id int) (*BaseChannel, error)
	// ChannelCount() int
	OutputCount() int
}

// BaseChannel provides the interface for the channel repeated capability for
// the IviFgenBase capability group.
type BaseChannel interface {
	// Name() string
	// VirtualName() string
	OperationMode() (OperationMode, error)
	SetOperationMode(mode OperationMode) error
	OutputEnabled() (bool, error)
	SetOutputEnabled(b bool) error
	DisableOutput() error
	EnableOutput() error
	OutputImpedance() (float64, error)
	SetOutputImpedance(impedance float64) error
	// OutputMode() (OutputMode, error)
	// SetOutputMode(mode OutputMode) error
	// ReferenceClockSource() (ClockSource, error)
	// SetReferenceClockSource(src ClockSource) error
	// AbortGeneration() error
	// InitiateGeneration() error
}

// OperationMode provides the defined values for the Operation Mode defined in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
type OperationMode int

// ContinuousMode and BurstMode are the available Operation Modes. "A burst
// consists of a discrete number of waveform cycles. the user uses the
// attribute of the IviFgenTrigger Extension Group to configure the trigger,
// and the attributes of the IviFgenBurst extension group to configure how the
// function generator produces bursts.
const (
	ContinuousMode OperationMode = iota
	BurstMode
)

func (om OperationMode) String() string {
	switch om {
	case ContinuousMode:
		return "continuous mode"
	case BurstMode:
		return "burst mode"
	}
	return ""
}

// OutputMode determines how the function generator produces waveforms. This
// attribute determines which extension group's functions and attributes are
// used to configure the waveform the function generator produces. OutputMode
// implements the defined values for the Output Mode read-write attribute
// defined in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
type OutputMode int

// Function indicates the IVI driver uses the attributes and functions of the
// IviFgenStdFunc extension group. Arbitrary indicates the IVI driver uses the
// attributes and functions of the IviFgenArbWfm extension group.
const (
	Function OutputMode = iota
	Arbitrary
	Sequence
)

// ClockSource models the defined values for the Reference Clock Source defined
// in Section 4.2.7 of IVI-4.3 IviFgenClass Specification.
type ClockSource int

// Available reference clock source.
const (
	RefClockInternal ClockSource = iota
	RefClockExternal
	RefClockRTSIClock
)
