// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// IntTrigger provides the interface required to support the
// IviFgenInternalTrigger extension group as described in Section 15 of the
// IVI-4.3: IviFgen Class Specification.
type IntTrigger interface {
	Channels() ([]*IntTriggerChannel, error)
	Channel(name string) (*IntTriggerChannel, error)
	ChannelByID(id int) (*IntTriggerChannel, error)
	ChannelCount() int
}

// IntTriggerChannel provides the interface required for the channel repeated
// capability to support the IviFgenInternalTrigger capability group as
// described in Section 15 of the IVI-4.3: IviFgen Class Specification.
type IntTriggerChannel interface {
	InternalTriggerRate() (float64, error)
	SetInternalTriggerRate(rate float64) error
}
