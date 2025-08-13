// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key3446x implements the IVI Instrument driver for the Keysight 3446x
family of DMM.

The Keysight 3446x family of DMMs use LAN port 5025 for SCPI Telnet sessions
and port 5025 for SCPI Socket sessions (confirmed for the Keysight 34461A and
assumed for the others). The default GPIB address for the 34461A is 22 (per p.
475 of the manual).

State Caching: Not implemented
*/
package key3446x

import (
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
var _ dmm.Base = (*Driver)(nil)

// Driver provides the IVI driver for the Keysight 3446x family of DMMs.
type Driver struct {
	inst ivi.Instrument
	ivi.Inherent
}

// New creates a new IVI driver for the Keysight 3446x series of DMMs.
func New(inst ivi.Instrument, reset bool) (*Driver, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true, // Default to returning to local control
		GroupCapabilities: []string{
			"IviDmmBase",
			"IviDmmACMeasurement",
			"IviDmmFrequencyMeasurement",
			// "IviDmmTemperatureMeasurement",
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
			"34460A",
			"34461A",
			"34465A",
			"34470A",
		},
		SupportedBusInterfaces: []string{
			"USB",
			"GPIB",
			"LAN",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Driver{
		inst:     inst,
		Inherent: inherent,
	}

	if reset {
		err := driver.Reset()
		return &driver, err
	}

	return &driver, nil
}

// Close properly shuts down the DMM by returning it to local control.
// This ensures the instrument's front panel regains control after use.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}
