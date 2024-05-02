// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
)

func TestCreateConfigureVoltageACCommand(t *testing.T) {
	testCases := []struct {
		autoRange   dmm.AutoRange
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{dmm.AutoOn, 0.0, "CONF:VOLT:AC AUTO", nil},
		{dmm.AutoOff, 0.09, "CONF:VOLT:AC 0.1", nil},
		{dmm.AutoOn, 10.0, "CONF:VOLT:AC AUTO", nil},
		{dmm.AutoOff, 9.0, "CONF:VOLT:AC 10", nil},
		{dmm.AutoOff, 10.0, "CONF:VOLT:AC 10", nil},
		{dmm.AutoOff, 10.001, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 99.99, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 100.00, "CONF:VOLT:AC 100", nil},
		{dmm.AutoOff, 100.001, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 999.999, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 1000.0, "CONF:VOLT:AC 1000", nil},
		{dmm.AutoOff, 1000.01, "", ivi.ErrValueNotSupported},
	}
	for _, tc := range testCases {
		got, err := createConfigureVoltageACCommand(tc.autoRange, tc.rangeValue)
		if err != tc.expectedErr {
			t.Errorf("wanted err %s / got err %s", tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Errorf("wanted %v / got %v", tc.expected, got)
		}
	}
}

func TestDetermineVoltageRange(t *testing.T) {
	testCases := []struct {
		autoRange   dmm.AutoRange
		rangeValue  float64
		expected    string
		expectedErr error
	}{
		{dmm.AutoOn, 0.0, "AUTO", nil},
		{dmm.AutoOff, 0.09, "0.1", nil},
		{dmm.AutoOn, 10.0, "AUTO", nil},
		{dmm.AutoOff, 9.0, "10", nil},
		{dmm.AutoOff, 10.0, "10", nil},
		{dmm.AutoOff, 10.001, "100", nil},
		{dmm.AutoOff, 99.99, "100", nil},
		{dmm.AutoOff, 100.00, "100", nil},
		{dmm.AutoOff, 100.001, "1000", nil},
		{dmm.AutoOff, 999.999, "1000", nil},
		{dmm.AutoOff, 1000.0, "1000", nil},
		{dmm.AutoOff, 1000.01, "", ivi.ErrValueNotSupported},
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
