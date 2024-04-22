// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

// ACMeasurement provides the interface required for the IviDmmACMeasurement
// extension group described in Section 5 of IVI-4.2 IviDmm Class
// Specification.
type ACMeasurement interface {
	MaxACFrequency() (float64, error)
	SetMaxACFrequency(maxFreq float64) error
	MinACFrequency() (float64, error)
	SetMinACFrequency(minFreq float64) error
	ConfigureACBandwidth(minFreq, maxFreq float64) error
}
