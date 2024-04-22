// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

// FrequencyMeasurement provides the interface required for the
// IviDmmFrequencyMeasurement extension group described in Section 6 of IVI-4.2
// IviDmm Class Specification.
type FrequencyMeasurement interface {
	FrequencyVoltageRange() (FreqAutoRange, float64, error)
	SetFrequencyVoltageRange(autoRange FreqAutoRange, rangeValue float64) error
}

type FreqAutoRange int

// The AutoRange defined values are the available auto range settings.
const (
	FreqAutoRangeOn FreqAutoRange = iota
	FreqAutoRangeOff
)
