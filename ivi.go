// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package ivi provides a Go-based implementation of the Interchangeable
// Virtual Instrument (IVI) standard. The IVI Specifications developed by the
// IVI Foundation provide standardized APIs for programming test instruments,
// such as oscilloscopes, power supplies, and function generators.
//
// The main advantage of the ivi package is not having to learn the Standard
// Commands for Programmable Instruments (SCPI) commands for each individual
// piece of test equipment. Contrary to the name, SCPI commands differ from one
// piece of test equipment to another. For instance, with ivi both the Agilent
// 33220A and the Stanford Research Systems DS345 function generators can be
// programmed using one API.
//
// Currently, ivi doesn't cache state. Every time an attribute is read directly
// from the instrument. Development focus is currently on fleshing out the APIs
// and creating a few IVI drivers for each instrument type.
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
