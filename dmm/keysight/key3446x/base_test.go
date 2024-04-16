// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"testing"

	"github.com/gotmc/ivi/dmm"
)

func TestDetermineVoltageRange(t *testing.T) {
	testCases := []struct {
		autoRange   dmm.AutoRange
		rangeValue  float64
		expectedErr error
		expected    string
	}{
		{
			dmm.AutoOn,
			0.0,
			nil,
			"AUTO",
		},
	}
	for _, tc := range testCases {
		got, err := determineVoltageRange(tc.autoRange, tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf("wanted err %s / got err %s", tc.expectedErr, err)
		}
		if got != tc.expected {
			t.Errorf("wanted %v / got %v", tc.expected, got)
		}
	}
}
