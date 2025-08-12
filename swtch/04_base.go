// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package swtch

import (
	"errors"
	"time"

	"github.com/gotmc/ivi"
)

// Base provides the interface required for the IviSwtchBase capability group.
type Base interface {
	CanConnect(ch1, ch2 string) error
	Channel(name string) (BaseChannel, error)
	ChannelByID(id int) (BaseChannel, error)
	ChannelCount() int
	Channels() ([]BaseChannel, error)
	Connect(ch1, ch2 string) error
	Disconnect(ch1, ch2 string) error
	DisconnectAll() error
	GetPath(ch1, ch2 string) ([]string, error)
	SetPath(chs []string) error
	WaitForDebounce(maxTime time.Duration) error
	SetVirtualNames(names []string) error
}

// BaseChannel provides the interface for the channel repeated capability for
// the IviSwtchBase capability group.
type BaseChannel interface {
	ACCurrentCarryMax() float64
	ACCurrentSwitchingMax() float64
	ACPowerCarryMax() float64
	ACPowerSwitchingMax() float64
	ACVoltageMax() float64
	Bandwidth() float64
	DCCurrentCarryMax() float64
	DCCurrentSwitchingMax() float64
	DCPowerCarryMax() float64
	DCPowerSwitchingMax() float64
	DCVoltageMax() float64
	DisableConfigChannel() error
	DisableSourceChannel() error
	EnableConfigChannel() error
	EnableSourceChannel() error
	Impedance() float64
	IsConfigChannel() bool
	IsDebounced() bool
	IsSourceChannel() bool
	Name() string
	SetConfigChannel(b bool) error
	SetSourceChannel(b bool) error
	SettlingTime() time.Duration
	VirtualName() string
	WireMode() int
}

// Error values for the PathCapability Parameter used in the CanConnect method
// as defined in Section 4.3.1 of IVI-4.6: IviSwtch Class Specification.
var (
	ErrPathExists          = errors.New("path exists")
	ErrPathUnsupported     = errors.New("path unsupported")
	ErrResourceInUse       = errors.New("resource in use")
	ErrSourceConflict      = errors.New("source conflict")
	ErrChannelNotAvailable = errors.New("channel not available")
)

// Error values that can be returned by the Connect method as defined in
// Section 4.3.2 of IVI-6.4: IviSwtch Class Specification.
var (
	ErrExplicitConnectionExists = errors.New("explicit connection exists")
	ErrIsConfigChannel          = errors.New("is config channel")
	ErrAttemptToConnectSources  = errors.New("attempt to connect sources")
	ErrCannotConnectToSelf      = errors.New("cannot connect to self")
	ErrPathNotFound             = errors.New("path not found")
)

// Channel models the repeated capability of a generic channel.
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// NewChannel returns a Channel for a switch.
func NewChannel(id int, name string, inst ivi.Instrument) Channel {
	return Channel{id, name, inst}
}

// Set writes the format string, using the given paarameters to the channel.
func (ch *Channel) Set(format string, a ...interface{}) error {
	return ivi.Set(ch.inst, format, a...)
}
