// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tempmon

import "time"

/*

The temperature monitor class is not part of the IVI Foundation
specifications. This interface is designed based on the common capabilities
of multi-channel temperature monitoring instruments.

A temperature monitor measures temperature from multiple input channels, each
of which can be independently configured with a sensor type (thermocouple,
RTD, thermistor, etc.) and measurement units.

*/

// Base provides the interface for core temperature monitor capabilities:
// channel configuration, temperature measurement, and alarm limits.
type Base interface {
	// Channel configuration
	ChannelCount() int
	ThermocoupleType(channel int) (ThermocoupleType, error)
	SetThermocoupleType(channel int, tcType ThermocoupleType) error
	TemperatureUnits(channel int) (TemperatureUnits, error)
	SetTemperatureUnits(channel int, units TemperatureUnits) error

	// Measurement
	MeasureTemperature(channel int) (float64, error)

	// Alarm limits
	AlarmEnabled(channel int) (bool, error)
	SetAlarmEnabled(channel int, enabled bool) error
	AlarmHighLimit(channel int) (float64, error)
	SetAlarmHighLimit(channel int, limit float64) error
	AlarmLowLimit(channel int) (float64, error)
	SetAlarmLowLimit(channel int, limit float64) error
}

// Scanner provides the interface for automated multi-channel scanning.
type Scanner interface {
	ScanEnabled() (bool, error)
	SetScanEnabled(enabled bool) error
	ChannelScanEnabled(channel int) (bool, error)
	SetChannelScanEnabled(channel int, enabled bool) error
	DwellTime() (time.Duration, error)
	SetDwellTime(dwell time.Duration) error
}

// RelativeTemperature provides the interface for nominal/delta-T measurements.
type RelativeTemperature interface {
	NominalTemperature(channel int) (float64, error)
	SetNominalTemperature(channel int, nominal float64) error
	DeltaTemperature(channel int) (float64, error)
}
