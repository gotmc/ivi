// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package lcr

import "time"

/*

The LCR meter class is not part of the IVI Foundation specifications. This
interface is designed based on the common capabilities of precision LCR meters,
using the Keysight E4980A as the reference instrument.

An LCR meter measures impedance parameters of a device under test (DUT) by
applying a test signal (AC voltage or current) at a specified frequency. The
measurement function determines which impedance model and parameters are
reported (e.g., Cp-D for parallel capacitance and dissipation factor, Ls-Rs
for series inductance and series resistance).

*/

// Base provides the interface for core LCR meter capabilities: measurement
// function selection, test signal configuration, triggering, and result
// retrieval.
type Base interface {
	// Measurement function
	MeasurementFunction() (MeasurementFunction, error)
	SetMeasurementFunction(fcn MeasurementFunction) error

	// Test signal
	Frequency() (float64, error)
	SetFrequency(freq float64) error
	TestVoltageLevel() (float64, error)
	SetTestVoltageLevel(volts float64) error
	TestCurrentLevel() (float64, error)
	SetTestCurrentLevel(amps float64) error

	// Impedance range
	ImpedanceAutoRange() (bool, error)
	SetImpedanceAutoRange(auto bool) error
	ImpedanceRange() (float64, error)
	SetImpedanceRange(ohms float64) error

	// Measurement speed and averaging
	MeasurementSpeed() (MeasurementSpeed, error)
	SetMeasurementSpeed(speed MeasurementSpeed) error
	AveragingCount() (int, error)
	SetAveragingCount(count int) error

	// Trigger
	TriggerSource() (TriggerSource, error)
	SetTriggerSource(src TriggerSource) error
	TriggerDelay() (time.Duration, error)
	SetTriggerDelay(delay time.Duration) error

	// Measurement control
	Initiate() error
	Trigger() error
	Abort() error

	// Results — returns primary value, secondary value, and status
	FetchMeasurement() (float64, float64, MeasurementStatus, error)
}

// DCBias provides the interface for DC bias control on LCR meters that
// support it. Not all LCR meters have DC bias capability.
type DCBias interface {
	DCBiasEnabled() (bool, error)
	SetDCBiasEnabled(enabled bool) error
	DCBiasVoltageLevel() (float64, error)
	SetDCBiasVoltageLevel(volts float64) error
}

// Compensation provides the interface for open, short, and load correction
// on LCR meters.
type Compensation interface {
	OpenCorrectionEnabled() (bool, error)
	SetOpenCorrectionEnabled(enabled bool) error
	ExecuteOpenCorrection() error
	ShortCorrectionEnabled() (bool, error)
	SetShortCorrectionEnabled(enabled bool) error
	ExecuteShortCorrection() error
	LoadCorrectionEnabled() (bool, error)
	SetLoadCorrectionEnabled(enabled bool) error
}
