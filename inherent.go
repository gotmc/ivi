// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"errors"
	"regexp"
)

type Inherent struct {
	inst Instrument
	InherentBase
}

type InherentBase struct {
	ClassSpecMajorVersion     int
	ClassSpecMinorVersion     int
	ClassSpecRevision         string
	GroupCapabilities         string
	SupportedInstrumentModels []string
	IDNString                 string
}

func NewInherent(inst Instrument, base InherentBase) Inherent {
	return Inherent{
		inst:         inst,
		InherentBase: base,
	}
}

// FirmwardRevision queries the instrument and returns the firmware revision of
// the instrument. FirmwareRevision is the getter for the read-only inherent
// attribute Instrument Firmware Revision described in Section 5.18 of IVI-3.2:
// Inherent Capabilities Specification.
func (inherent *Inherent) FirmwareRevision() (string, error) {
	return inherent.parseIdentification("fwr")
}

// InstrumentManufacturer queries the instrument and returns the manufacturer
// of the instrument. InstrumentManufacturer is the getter for the read-only
// inherent attribute Instrument Manufacturer described in Section 5.19 of
// IVI-3.2: Inherent Capabilities Specification.
func (inherent *Inherent) InstrumentManufacturer() (string, error) {
	return inherent.parseIdentification("mfr")
}

// InstrumentModel queries the instrument and returns the model of the
// instrument.  InstrumentModel is the getter for the read-only inherent
// attribute Instrument Model described in Section 5.20 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) InstrumentModel() (string, error) {
	return inherent.parseIdentification("model")
}

// InstrumentSerialNumber queries the instrument and returns the S/N of the
// instrument.
func (inherent *Inherent) InstrumentSerialNumber() (string, error) {
	return inherent.parseIdentification("sn")
}

func (inherent *Inherent) Reset() error {
	_, err := inherent.inst.WriteString("*RST\n")
	return err
}

func (inherent *Inherent) Clear() error {
	_, err := inherent.inst.WriteString("*CLS\n")
	return err
}

// Disable places the instrument in a quiescent state as quickly as possible.
// Disable provides the method described in Section 6.4 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) Disable() error {
	// FIXME(mdr): Implement!!!!
	return errors.New("disable is not yet implemented.")
}

func (inherent *Inherent) parseIdentification(part string) (string, error) {
	s, err := inherent.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(inherent.IDNString)
	res := re.FindStringSubmatch(s)
	subexpNames := re.SubexpNames()
	matchMap := map[string]string{}
	for i, n := range res {
		matchMap[subexpNames[i]] = string(n)
	}
	return matchMap[part], nil
}
