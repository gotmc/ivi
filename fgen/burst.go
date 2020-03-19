// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// Burst provides the interface required for the IviFgenBurst capability group.
type Burst interface {
	Channels() ([]*BurstChannel, error)
	Channel(name string) (*BurstChannel, error)
	ChannelByID(id int) (*BurstChannel, error)
	ChannelCount() int
}

// BurstChannel provides the interface for the channel repeated capability for
// the IviFgenBurst capability group.
type BurstChannel interface {
	BurstCount() (int, error)
	SetBurstCount(count int) error
}
