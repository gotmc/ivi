// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "fmt"

// LookupSCPI returns the SCPI string for the given enum value using the
// provided mapping.
func LookupSCPI[E comparable](m map[E]string, val E) (string, error) {
	s, ok := m[val]
	if !ok {
		return "", fmt.Errorf("%w: %v", ErrValueNotSupported, val)
	}

	return s, nil
}

// ReverseLookup returns the enum value for the given SCPI string using the
// provided mapping.
func ReverseLookup[E comparable](m map[string]E, scpi string) (E, error) {
	val, ok := m[scpi]
	if !ok {
		var zero E
		return zero, fmt.Errorf("%w: %q", ErrUnexpectedResponse, scpi)
	}

	return val, nil
}
