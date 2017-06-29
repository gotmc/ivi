// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "io"

// Instrument provides the interface required for all IVI Instruments.
type Instrument interface {
	io.ReadWriteCloser
	StringWriter
	Querier
}

type StringWriter interface {
	WriteString(s string) (n int, err error)
}

type Querier interface {
	Query(s string) (value string, err error)
}
