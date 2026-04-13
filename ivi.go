// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "context"

// Transport provides the interface required for all IVI Instruments to
// communicate with test equipment regardless of the transport layer driver
// used (e.g., usbtmc, asrl, lxi, prologix).
type Transport interface {
	Command(ctx context.Context, cmd string, a ...any) error
	Query(ctx context.Context, cmd string) (value string, err error)
	ReadBinary(ctx context.Context, p []byte) (n int, err error)
	WriteBinary(ctx context.Context, p []byte) (n int, err error)
	Close() error
}

// Commander provides the interface to send a SCPI/ASCII command with a
// terminator to an instrument that is optionally formatted according to a
// format specifier.
type Commander interface {
	Command(ctx context.Context, cmd string, a ...any) error
}

// Querier provides the interface to query an instrument using a given
// SCPI/ASCII command that reads until the appropriate terminator and then
// provides the resultant string.
type Querier interface {
	Query(ctx context.Context, cmd string) (value string, err error)
}

// BinaryReader provides the interface to read binary data without
// terminator interpretation.
type BinaryReader interface {
	ReadBinary(ctx context.Context, p []byte) (n int, err error)
}

// BinaryWriter provides the interface to write binary data without adding a
// terminator.
type BinaryWriter interface {
	WriteBinary(ctx context.Context, p []byte) (n int, err error)
}
