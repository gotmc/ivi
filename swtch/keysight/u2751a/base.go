// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package u2751a

import (
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/swtch"
)

// ACCurrentCarryMax returns the maximum AC current the channel can carry in
// amperes RMS. ACCurrentCarryMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.ACCurrentCarryMax described in Section 4.2.1 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) ACCurrentCarryMax() float64 {
	return ch.acCurrentCarryMax
}

// ACCurrentSwitchingMax returns the maximum AC current the channel can switch in
// amperes RMS. ACCurrentSwitchingMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.ACCurrentSwitchingMax described in Section 4.2.2 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) ACCurrentSwitchingMax() float64 {
	return ch.acCurrentSwitchingMax
}

// ACPowerCarryMax returns the maximum AC power the channel can handle in
// volt-amperes. ACPowerCarryMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.ACPowerCarryMax described in Section 4.2.3 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) ACPowerCarryMax() float64 {
	return ch.acPowerCarryMax
}

// ACPowerSwitchingMax returns the maximum AC power the channel can switch in
// volt-amperes. ACPowerSwitchingMax is the getter for the read-only
// IviSwtchBase attribute Characteristics.ACPowerSwitchingMax described in
// Section 4.2.4 of IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) ACPowerSwitchingMax() float64 {
	return ch.acPowerSwitchingMax
}

// ACVoltageMax returns the maximum AC voltage the channel can handle in volts
// RMS. ACVoltageMax is the getter for the read-only IviSwtchBase attribute
// Characteristics.ACVoltageMax described in Section 4.2.5 of IVI-4.6:IviSwtch
// Class Specification.
func (ch *Channel) ACVoltageMax() float64 {
	return ch.acVoltageMax
}

// Bandwidth returns the maximum frequency signal, in Hertz, that can pass
// through the channel without attenuating it by more than 3dB.  Bandwidth is
// the getter for the read-only IviSwtchBase attribute
// Characteristics.Bandwidth described in Section 4.2.6 of IVI-4.6:IviSwtch
// Class Specification.
func (ch *Channel) Bandwidth() float64 {
	return ch.bw
}

// ChannelCount returns the number of channels. ChannelCount is the getter for
// the read-only IviSwtchBase attribute Channels.Count described in Section
// 4.2.7 of IVI-4.6:IviSwtch Class Specification.
func (d *U2751A) ChannelCount() int {
	return len(d.channels)
}

// Impedance returns the characteristic impedance of the channel in ohms.
// Impedance is the getter for the read-only IviSwtchBase attribute
// Characteristics.Impedance described in Section 4.2.10 of IVI-4.6:IviSwtch
// Class Specification.
func (ch *Channel) Impedance() float64 {
	return ch.impedance
}

// DCCurrentCarryMax returns the maximum DC current the channel can carry in
// amperes. DCCurrentCarryMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.DCCurrentCarryMax described in Section 4.2.11 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) DCCurrentCarryMax() float64 {
	return ch.dcCurrentCarryMax
}

// DCCurrentSwitchingMax returns the maximum DC current the channel can switch
// in amperes. DCCurrentSwitchingMax is the getter for the read-only
// IviSwtchBase attribute Characteristics.DCCurrentSwitchingMax described in
// Section 4.2.12 of IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) DCCurrentSwitchingMax() float64 {
	return ch.dcCurrentSwitchingMax
}

// DCPowerCarryMax returns the maximum DC power the channel can handle in
// watts. DCPowerCarryMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.DCPowerCarryMax described in Section 4.2.13 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) DCPowerCarryMax() float64 {
	return ch.dcPowerCarryMax
}

// DCPowerSwitchingMax returns the maximum DC power the channel can switch in
// watts. DCPowerSwitchingMax is the getter for the read-only IviSwtchBase
// attribute Characteristics.DCPowerSwitchingMax described in Section 4.2.14 of
// IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) DCPowerSwitchingMax() float64 {
	return ch.dcPowerSwitchingMax
}

// DCVoltageMax returns the maximum DC voltage the channel can handle in volts.
// DCVoltageMax is the getter for the read-only IviSwtchBase attribute
// DCVoltageMax described in Section 4.2.15 of IVI-4.6:IviSwtch Class
// Specification.
func (ch *Channel) DCVoltageMax() float64 {
	return ch.dcVoltageMax
}

// IsConfigChannel returns whether the specific driver uses the channel for
// internal path creation. If set to True, the channel is no longer accessible
// to the user and can be used by the specific driver for path creation. If set
// to False, the channel is considered a standard channel and can be explicitly
// connected to another channel. IsConfigChannel is the getter for the
// read-write IviSwtchBase Attribute IsConfigurationChannel described in
// Section 4.2.16 of IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) IsConfigChannel() bool {
	return ch.isConfigChannel
}

// SetConfigChannel specifies whether the specific driver uses the channel for
// internal path creation. If set to True, the channel is no longer accessible
// to the user and can be used by the specific driver for path creation. If set
// to False, the channel is considered a standard channel and can be explicitly
// connected to another channel. IsConfigChannel is the setter for the
// read-write IviSwtchBase Attribute IsConfigurationChannel described in
// Section 4.2.16 of IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) SetConfigChannel(b bool) {
	ch.isConfigChannel = b
}

// IsDebounced indicates whether the switch module has settled from the
// switching commands and completed the debounce. If True, the switch module
// has settled from the switching commands and completed the debounce.It
// indicates that the signal going through the switch module is valid, assuming
// that the switches in the path have the correct characteristics. If False,
// the switch module has not settled. IsDebounced is the getter for the
// read-write IviSwtchBase Attribute Is Source Channel described in Section
// 4.2.17 of IVI-4.6: IviSwtch Class Specification.
func (ch *Channel) IsDebounced() bool {
	return ch.isDebounced
}

// IsSourceChannel returns true if this channel has been delcared as a source
// channel. If a user ever attempts to connect two channels that are either
// sources or have their own connections to sources, the path creation
// operation returns an error. IsSourceChannel is the getter for the read-write
// IviSwtchBase Attribute Is Source Channel described in Section 4.2.18 of
// IVI-4.6: IviSwtch Class Specification.
func (ch *Channel) IsSourceChannel() bool {
	return ch.isSourceChannel
}

// SetSourceChannel declares whether or not a particular channel is a source
// channel. By default, all channels are not source channels. IsConfigChannel
// is the setter for the read-write IviSwtchBase Attribute
// IsConfigurationChannel described in Section 4.2.18 of IVI-4.6:IviSwtch Class
// Specification.
func (ch *Channel) SetSourceChannel(b bool) {
	ch.isSourceChannel = b
}

// SettlingTime returns the maximum total settling time for the channel before
// the signal going through it is considered stable. This includes both the
// activation time for the channel as well as any debounce time. SettlingTime
// is the getter for the read-only IviSwtchBase Attribute Settling Time
// described in Section 4.2.19 of IVI-4.6:IviSwtch Class Specification.
func (ch *Channel) SettlingTime() time.Duration {
	return ch.settlingTime
}

// WireMode returns the number of conductors in the given channel. WireMode is
// the getter for the read-only IviSwtchBase Attribute Wire Mode described in
// Section 4.2.20 of IVI-4.6: IviSwtch Class Specification.
func (ch *Channel) WireMode() int {
	return ch.numWires
}

// CanConnect allows the user to verify whether the switch module can create a
// given path without the switch module actually creating the path. If the path
// can be created, true is returned. In addition, the operation indicates
// whether the switch module can create the path at the moment based on the
// current paths in existence. CanConnect implements the IviSwtch Base Function
// Can Connect described in Section 4.3.1 of IVI-4.6: IviSwtch Class
// Specification.
func (d *U2751A) CanConnect(ch1name, ch2name string) (bool, error) {
	ch1, err := d.Channel(ch1name)
	if err != nil {
		return false, err
	}
	ch2, err := d.Channel(ch2name)
	if err != nil {
		return false, err
	}

	smallID := ch1.id
	largeID := ch2.id
	if largeID < smallID {
		smallID = ch2.id
		largeID = ch1.id
	}

	// Check to see if one of the channels is a configuration channel.
	if d.channels[smallID].isConfigChannel || d.channels[largeID].isConfigChannel {
		return false, swtch.ErrChannelNotAvailable
	}

	// Check to see if we're trying to connect a row to a row or a column to a
	// column.
	if ch1.chType == ch2.chType {
		// FIXME(mdr): Add support for using a column as a configuration channel to
		// connect two rows, and add support for using a row as a configuration
		// channel to connect two columns.
		return false, swtch.ErrPathUnsupported
	}
	// Make sure that we aren't trying to connect two source channels.
	if ch1.isSourceChannel && ch2.isSourceChannel {
		return false, swtch.ErrSourceConflict
	}
	// See if the path already exists.
	newPath := []string{d.channels[smallID].name, d.channels[largeID].name}
	for _, path := range d.paths {
		if pathsEqual(newPath, path) {
			return false, swtch.ErrPathExists
		}
	}
	return true, nil
}

// Connect takes two channel names and, if possible, creates a path between the
// two channels.  Connect implements the IviSwtch Base Function Connect
// described in Section 4.3.2 of IVI-4.6: IviSwtch Class Specification.
func (d *U2751A) Connect(ch1name, ch2name string) error {
	ch1, err := d.Channel(ch1name)
	if err != nil {
		// Should I return an Unknown Channel Name per IVI-3.2 Table 9-2?
		return err
	}
	ch2, err := d.Channel(ch2name)
	if err != nil {
		return err
	}

	smallID := ch1.id
	largeID := ch2.id
	if largeID < smallID {
		smallID = ch2.id
		largeID = ch1.id
	}

	// Check to see if we're trying to connect a channel to itself.
	if smallID == largeID {
		return swtch.ErrCannotConnectToSelf
	}

	// Check to see if one of the channels is a configuration channel.
	if d.channels[smallID].isConfigChannel || d.channels[largeID].isConfigChannel {
		return swtch.ErrIsConfigChannel
	}

	// Check to see if we're trying to connect a row to a row or a column to a
	// column.
	if ch1.chType == ch2.chType {
		// FIXME(mdr): Add support for using a column as a configuration channel to
		// connect two rows, and add support for using a row as a configuration
		// channel to connect two columns.
		return swtch.ErrPathUnsupported
	}

	// Make sure that we aren't trying to connect two source channels.
	if ch1.isSourceChannel && ch2.isSourceChannel {
		return swtch.ErrAttemptToConnectSources
	}
	// See if the path already exists.
	newPath := []string{d.channels[smallID].name, d.channels[largeID].name}
	for _, path := range d.paths {
		if pathsEqual(newPath, path) {
			return swtch.ErrExplicitConnectionExists
		}
	}

	// Make the connection.
	row := ch1.switchID
	col := ch2.switchID
	if ch1.chType != Row || ch2.chType != Col {
		return fmt.Errorf("expected a row and a col got: %s and %s", ch1.chType, ch2.chType)
	}
	err = ivi.Set(d.inst, "ROUT:CLOS (@%1d%02d)\n", row, col)
	if err != nil {
		return err
	}
	d.paths = append(d.paths, newPath)
	return nil
}

func pathsEqual(pathA, pathB []string) bool {
	if len(pathA) != len(pathB) {
		return false
	}
	for i := range pathA {
		if pathA[i] != pathB[i] {
			return false
		}
	}
	return true
}
