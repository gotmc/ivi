// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package ivitest provides shared test fixtures used by driver packages
// inside this repository. It is not part of the public ivi API.
package ivitest

import (
	"context"
	"errors"
	"fmt"
)

// Mock satisfies the [github.com/gotmc/ivi.Transport] interface for unit
// testing. Command appends the formatted SCPI string to CommandsSent, Query
// returns QueryResp. Setting ShouldError causes Command and Query to return
// an error without touching CommandsSent or QueryResp.
//
// The binary I/O methods are minimal no-ops suitable for tests that never
// exercise binary transfers. Tests that need richer behavior can embed Mock
// and override the methods they care about.
type Mock struct {
	// CommandsSent captures every formatted SCPI command passed to Command,
	// in call order.
	CommandsSent []string
	// QueryResp is the string returned by Query.
	QueryResp string
	// ShouldError, when true, makes Command and Query return a generic error.
	ShouldError bool
}

// ReadBinary returns (0, nil) unconditionally.
func (m *Mock) ReadBinary(_ context.Context, _ []byte) (int, error) {
	return 0, nil
}

// WriteBinary returns (len(p), nil) unconditionally.
func (m *Mock) WriteBinary(_ context.Context, p []byte) (int, error) {
	return len(p), nil
}

// Close returns nil unconditionally.
func (m *Mock) Close() error { return nil }

// Command formats the SCPI command and appends it to CommandsSent, or
// returns an error when ShouldError is set.
func (m *Mock) Command(_ context.Context, format string, a ...any) error {
	if m.ShouldError {
		return errors.New("mock command error")
	}

	m.CommandsSent = append(m.CommandsSent, fmt.Sprintf(format, a...))

	return nil
}

// Query returns QueryResp, or an error when ShouldError is set. The query
// string is ignored.
func (m *Mock) Query(_ context.Context, _ string) (string, error) {
	if m.ShouldError {
		return "", errors.New("mock query error")
	}

	return m.QueryResp, nil
}
