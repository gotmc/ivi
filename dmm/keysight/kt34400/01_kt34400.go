// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package kt34400 implements the IVI Instrument driver for the Keysight
// Digital Multimeter Family of instruments.
//
// The Keysight 34400 family of DMMs use LAN port 5025 for SCPI Telnet sessions
// and port 5025 for SCPI Socket sessions (confirmed for the Keysight 34461A
// and assumed for the others). The default GPIB address for the 34461A is 22
// (per p. 475 of the manual).
//
// State Caching: Not implemented
package kt34400

import (
	"context"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
)

const (
	specMajorVersion = 4
	specMinorVersion = 2
	specRevision     = "4.1"
)

// Confirm the interfaces implemented by the driver.
var (
	_ dmm.Base                            = (*Driver)(nil)
	_ dmm.ACMeasurementExtension          = (*Driver)(nil)
	_ dmm.FrequencyMeasurementExtension   = (*Driver)(nil)
	_ dmm.TemperatureMeasurementExtension = (*Driver)(nil)
)

// Driver provides the IVI driver for the Keysight 3446x family of DMMs.
type Driver struct {
	inst    ivi.Transport
	timeout time.Duration
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 3446x series of DMMs. By
// default the constructor queries *IDN? and verifies the model against the
// supported list; pass [ivi.WithoutIDQuery] to skip that check. Use
// [ivi.WithReset] to reset on creation and [ivi.WithTimeout] to override the
// default I/O timeout.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	s, err := ivi.NewDriverSetup(inst, ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDmmBase",
			"IviDmmACMeasurement",
			"IviDmmFrequencyMeasurement",
			"IviDmmTemperatureMeasurement",
			// "IviDmmResistanceTemperatureDevice",
			// "IviDmmThermistor",
			// "IviDmmMultiPoint",
			// "IviDmmTriggerSlope",
			// "IviDmmSoftwareTrigger",
			// "IviDmmDeviceInfo",
			// "IviDmmAutoRangeValue",
			// "IviDmmAutoZero",
			// "IviDmmPowerLineFrequency",
		},
		SupportedInstrumentModels: []string{
			"34450A",
			"EDU34450A",
			"34460A",
			"34461A",
			"34465A",
			"34470A",
		},

		SupportedBusInterfaces: []string{"USB", "GPIB", "LAN"},
	}, opts)
	if err != nil {
		return nil, err
	}

	driver := Driver{
		inst:     inst,
		timeout:  s.Timeout,
		Inherent: s.Inherent,
	}

	if s.Config.Reset {
		if err := driver.Reset(); err != nil {
			return nil, err
		}
	}

	return &driver, nil
}

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// Close properly shuts down the DMM by returning it to local control.
// This ensures the instrument's front panel regains control after use.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
