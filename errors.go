// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "errors"

// Sentinel errors returned by the ivi package and by driver implementations.
// Callers should use [errors.Is] to check for these rather than comparing
// error values directly, since drivers typically wrap them with additional
// context.
var (
	// ErrNotImplemented indicates the requested IVI capability has not been
	// implemented by the driver.
	ErrNotImplemented = errors.New("not implemented in ivi driver")
	// ErrFunctionNotSupported indicates the requested measurement function or
	// operation is not supported by the connected instrument model.
	ErrFunctionNotSupported = errors.New("function not supported")
	// ErrValueNotSupported indicates the requested value (e.g., a range or
	// setting) is outside what the instrument accepts.
	ErrValueNotSupported = errors.New("value not supported")
	// ErrUnexpectedResponse indicates the instrument returned a response the
	// driver could not parse or did not expect.
	ErrUnexpectedResponse = errors.New("unexpected response from instrument")
	// ErrChannelNotFound indicates the caller requested a channel index or
	// name that the driver does not know about.
	ErrChannelNotFound = errors.New("channel index out of range")
	// ErrUnsupportedModel indicates the connected instrument's model is not in
	// the driver's SupportedInstrumentModels list.
	ErrUnsupportedModel = errors.New("unsupported instrument model")
)
