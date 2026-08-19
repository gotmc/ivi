// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivitest

import (
	"context"
	"fmt"
)

// Strict wraps [Mock] and runs [CheckSCPI] over every command and query the
// driver sends, collecting the malformed ones rather than failing at the call
// site. This lets a test drive a driver through many methods and then report
// every violation at once.
//
// Strict records commands in the embedded Mock's CommandsSent and queries in
// QueriesSent, so a test can assert on exact spelling as well as validity.
type Strict struct {
	Mock
	// QueriesSent captures every query string, in call order.
	QueriesSent []string
	// Violations holds one message per malformed command or query, naming the
	// offending string and why it cannot parse.
	Violations []string
}

// errorf is the subset of [testing.T] that Check needs. Taking an interface
// keeps the testing package out of this one's imports.
type errorf interface {
	Errorf(format string, args ...any)
	Helper()
}

// Check reports every collected violation to t. It is a no-op when the driver
// sent nothing malformed.
func (s *Strict) Check(t errorf) {
	t.Helper()

	for _, v := range s.Violations {
		t.Errorf("%s", v)
	}
}

// Command validates the formatted command before handing it to the embedded
// Mock, which records it in CommandsSent.
func (s *Strict) Command(ctx context.Context, format string, a ...any) error {
	cmd := fmt.Sprintf(format, a...)
	s.check("command", cmd)

	// Pass the already formatted string so Mock records exactly what was
	// validated. It holds no verbs, so it is safe as a format string.
	return s.Mock.Command(ctx, "%s", cmd)
}

// Query validates the query before handing it to the embedded Mock.
func (s *Strict) Query(ctx context.Context, cmd string) (string, error) {
	s.check("query", cmd)
	s.QueriesSent = append(s.QueriesSent, cmd)

	return s.Mock.Query(ctx, cmd)
}

// check records a violation when cmd is not well-formed SCPI.
func (s *Strict) check(kind, cmd string) {
	if err := CheckSCPI(cmd); err != nil {
		s.Violations = append(
			s.Violations, fmt.Sprintf("%s %q: %v", kind, cmd, err),
		)
	}
}
