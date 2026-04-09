// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "context"

// Set formats according to a format specifier and then sends the resulting
// command to the given Commander interface.
func Set(ctx context.Context, cmdr Commander, format string, a ...any) error {
	return cmdr.Command(ctx, format, a...)
}

// QueryID queries the identity of the instrument.
func QueryID(ctx context.Context, q Querier) (string, error) {
	return q.Query(ctx, "*IDN?\n")
}
