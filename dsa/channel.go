// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

import (
	"context"

	"github.com/gotmc/ivi"
	"github.com/gotmc/query"
)

// Channel models a generic DSA channel
type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// NewChannel returns a Channel for a DSA.
func NewChannel(id int, name string, inst ivi.Instrument) Channel {
	return Channel{id, name, inst}
}

// Set writes the format string, using the given parameters to the channel.
func (ch *Channel) Set(ctx context.Context, format string, a ...interface{}) error {
	return ivi.Set(ctx, ch.inst, format, a...)
}

// QueryBool queries the channel and returns a bool.
func (ch *Channel) QueryBool(ctx context.Context, cmd string) (bool, error) {
	return query.Bool(ctx, ch.inst, cmd)
}

// QueryFloat64 queries the channel and returns a float64.
func (ch *Channel) QueryFloat64(ctx context.Context, cmd string) (float64, error) {
	return query.Float64(ctx, ch.inst, cmd)
}

// QueryInt queries the channel and returns an int.
func (ch *Channel) QueryInt(ctx context.Context, cmd string) (int, error) {
	return query.Int(ctx, ch.inst, cmd)
}

// QueryString queries the channel and returns a string.
func (ch *Channel) QueryString(ctx context.Context, cmd string) (string, error) {
	return query.String(ctx, ch.inst, cmd)
}
