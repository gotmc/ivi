// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"testing"
)

func TestParseIdentification(t *testing.T) {
	testCases := map[string]struct {
		given    string
		part     idPart
		expected string
	}{
		"model": {
			"HEWLETT-PACKARD,E3631A,0,1.4-5.0-1.0",
			modelID,
			"E3631A",
		},
		"firmware": {
			"HEWLETT-PACKARD,E3631A,0,1.4-5.0-1.0",
			fwrID,
			"1.4-5.0-1.0",
		},
		"manufacturer": {
			"HEWLETT-PACKARD,E3631A,0,1.4-5.0-1.0",
			mfrID,
			"HEWLETT-PACKARD",
		},
		"serial_number": {
			"HEWLETT-PACKARD,E3631A,0,1.4-5.0-1.0",
			snID,
			"0",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := parseIdentification(testCase.given, testCase.part)
			if err != nil {
				t.Errorf("got error %s / expected nil", err)
			}
			if got != testCase.expected {
				t.Errorf("got %s / expected %s", got, testCase.expected)
			}
		})
	}
}

func TestParseIdentificationErrors(t *testing.T) {
	testCases := map[string]struct {
		given       string
		expectedMsg string
	}{
		"empty_idn": {
			"",
			"idn string (``) could not be split in four",
		},
		"partial_idn": {
			"HEWLETT-PACKARD,E3631A",
			"idn string (`HEWLETT-PACKARD,E3631A`) could not be split in four",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := parseIdentification(testCase.given, fwrID)
			if err == nil {
				t.Errorf("expected error, got nil")
			} else if err.Error() != testCase.expectedMsg {
				t.Errorf("got error %q / expected %q", err.Error(), testCase.expectedMsg)
			}
		})
	}
}
