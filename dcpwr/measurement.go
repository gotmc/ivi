// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

// Measurement provides the interface required for the IviDCPwrMeasurement
// capability group.
type Measurement interface {
	// Channels() ([]*MeasurementChannel, error)
	// Channel(name string) (*MeasurementChannel, error)
	// ChannelByID(id int) (*MeasurementChannel, error)
	// ChannelCount() int
}

// MeasurementChannel provides the interface for the channel repeated
// capability for the IviDCPwrMeasurement capability group.
type MeasurementChannel interface {
	MeasureVoltage() (float64, error)
	MeasureCurrent() (float64, error)
}
