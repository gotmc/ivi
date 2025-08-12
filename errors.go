// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "errors"

var (
	ErrNotImplemented       = errors.New("not implemented in ivi driver")
	ErrFunctionNotSupported = errors.New("function not supported")
	ErrValueNotSupported    = errors.New("value not supported")
	ErrUnexpectedResponse   = errors.New("unexpected response from instrument")
)
