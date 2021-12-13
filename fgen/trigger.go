// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// Trigger provides the interface required for the IviFgenTrigger extension group.
type Trigger interface {
	Channels() ([]*TriggerChannel, error)
	Channel(name string) (*TriggerChannel, error)
	ChannelByID(id int) (*TriggerChannel, error)
	ChannelCount() int
}

// TriggerChannel provides the interface for the channel repeated capability for
// the IviFgenTrigger capability group.
type TriggerChannel interface {
	TriggerSource() (TriggerSource, error)
	SetTriggerSource(TriggerSource) error
}

// TriggerSource models the defined values for the Trigger Source defined in
// Section 9.2.1 of IVI-4.3: IviFgenClass Specification.
type TriggerSource int

// These are the available trigger sources.
const (
	InternalTrigger TriggerSource = iota
	ExternalTrigger
	SoftwareTrigger
)
