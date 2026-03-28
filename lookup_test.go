// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"errors"
	"testing"
)

type testEnum int

const (
	enumA testEnum = iota
	enumB
	enumC
)

var forwardMap = map[testEnum]string{
	enumA: "SCPI_A",
	enumB: "SCPI_B",
}

var reverseMap = map[string]testEnum{
	"SCPI_A": enumA,
	"SCPI_B": enumB,
}

func TestLookupSCPI(t *testing.T) {
	tests := []struct {
		name    string
		val     testEnum
		want    string
		wantErr error
	}{
		{"found A", enumA, "SCPI_A", nil},
		{"found B", enumB, "SCPI_B", nil},
		{"missing C", enumC, "", ErrValueNotSupported},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LookupSCPI(forwardMap, tt.val)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("LookupSCPI() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if err != nil {
				t.Errorf("LookupSCPI() unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("LookupSCPI() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestReverseLookup(t *testing.T) {
	tests := []struct {
		name    string
		scpi    string
		want    testEnum
		wantErr error
	}{
		{"found A", "SCPI_A", enumA, nil},
		{"found B", "SCPI_B", enumB, nil},
		{"missing", "SCPI_X", 0, ErrUnexpectedResponse},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReverseLookup(reverseMap, tt.scpi)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("ReverseLookup() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if err != nil {
				t.Errorf("ReverseLookup() unexpected error: %v", err)
			}

			if got != tt.want {
				t.Errorf("ReverseLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
