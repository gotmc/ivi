// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "context"

// Instrument provides the interface required for all IVI Instruments.
type Instrument interface {
	ReadContext(ctx context.Context, p []byte) (n int, err error)
	WriteContext(ctx context.Context, p []byte) (n int, err error)
	Command(ctx context.Context, cmd string, a ...any) error
	Query(ctx context.Context, s string) (value string, err error)
	Close() error
}

// Commander provides the interface to send a command to an instrument that is
// optionally formatted according to a format specifier.
type Commander interface {
	Command(ctx context.Context, cmd string, a ...any) error
}

// Querier provides the interface to query an instrument using a given command
// and then provide the resultant string.
type Querier interface {
	Query(ctx context.Context, s string) (value string, err error)
}
