// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package key3446x implements the IVI Instrument driver for the Keysight 3446x
family of DMM.

State Caching: Not implemented
*/
package key3446x

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

const (
	specMajorVersion = 4
	specMinorVersion = 2
	specRevision     = "4.1"
)

// Ag3446x provides the IVI Instrument driver for the Keysight 3446x family of
// DMM.
type Ag3446x struct {
	inst        ivi.Instrument
	outputCount int
	ivi.Inherent
}

// New creates a new Agilent3446x IVI Instrument.
func New(inst ivi.Instrument, reset bool) (*Ag3446x, error) {
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		GroupCapabilities: []string{
			"DmmBase",
			"DmmACMeasurement",
			"DmmFrequencyMeasurement",
			"DmmDeviceInfo",
		},
		SupportedInstrumentModels: []string{
			"34460A",
			"34461A",
			"34465A",
			"34470A",
		},
	}
	inherent := ivi.NewInherent(inst, inherentBase)
	driver := Ag3446x{
		inst:        inst,
		outputCount: 1,
		Inherent:    inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// QueryString queries the DMM and returns a string.
func (d *Ag3446x) QueryString(cmd string) (string, error) {
	return query.String(d.inst, cmd)
}

// MeasurementFunction needs a better comment.
func (d *Ag3446x) MeasurementFunction() (dmm.MeasurementFunction, error) {
	conf, err := d.QueryString("CONF?")
	if err != nil {
		return 0, err
	}
	conf = strings.TrimSpace(conf)
	fcnString := strings.Split(conf, " ")[0]
	fcn, ok := fcnMap[fcnString]
	if !ok {
		return 0, fmt.Errorf("%s is not a valid measurement function", fcnString)
	}
	return fcn, nil
}

// MeasurementFunctionMap maps the string name of a measurement function to the
// MeasurementFunction.
var fcnMap = map[string]dmm.MeasurementFunction{
	"VOLT":    dmm.DCVolts,
	"VOLT:DC": dmm.DCVolts,
	"VOLT:AC": dmm.ACVolts,
	"CURR":    dmm.DCCurrent,
	"CURR:DC": dmm.DCCurrent,
	"CURR:AC": dmm.ACCurrent,
	"RES":     dmm.TwoWireResistance,
	"FRES":    dmm.FourWireResistance,
	"FREQ":    dmm.Frequency,
	"TEMP":    dmm.Temperature,
}
