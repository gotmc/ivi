// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

import "github.com/gotmc/ivi"

// Channel models a generic DSA input channel.
type Channel struct {
	id   int
	name string
	inst ivi.Transport
}

// NewChannel returns a Channel for a DSA.
func NewChannel(id int, name string, inst ivi.Transport) Channel {
	return Channel{id, name, inst}
}

// ID returns the channel's numeric ID.
func (ch *Channel) ID() int { return ch.id }

// Name returns the channel's name.
func (ch *Channel) Name() string { return ch.name }
