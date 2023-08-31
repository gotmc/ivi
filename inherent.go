// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"errors"
	"fmt"
	"strings"
)

type idPart int

const (
	mfrID idPart = iota
	modelID
	snID
	fwrID
)

// Inherent provides the inherent capabilities for all IVI instruments.
type Inherent struct {
	inst Instrument
	InherentBase
}

// InherentBase provides the exported properties for the inherent capabilities
// of all IVI instruments.
type InherentBase struct {
	ClassSpecMajorVersion     int
	ClassSpecMinorVersion     int
	ClassSpecRevision         string
	GroupCapabilities         []string
	SupportedInstrumentModels []string
	IDNString                 string
}

// NewInherent creates a new Inherent struct using the given Instrument
// interface and the InherentBase struct.
func NewInherent(inst Instrument, base InherentBase) Inherent {
	return Inherent{
		inst:         inst,
		InherentBase: base,
	}
}

// FirmwareRevision queries the instrument and returns the firmware revision of
// the instrument. FirmwareRevision is the getter for the read-only inherent
// attribute Instrument Firmware Revision described in Section 5.18 of IVI-3.2:
// Inherent Capabilities Specification.
func (inherent *Inherent) FirmwareRevision() (string, error) {
	return inherent.queryIdentification(fwrID)
}

// InstrumentManufacturer queries the instrument and returns the manufacturer
// of the instrument. InstrumentManufacturer is the getter for the read-only
// inherent attribute Instrument Manufacturer described in Section 5.19 of
// IVI-3.2: Inherent Capabilities Specification.
func (inherent *Inherent) InstrumentManufacturer() (string, error) {
	return inherent.queryIdentification(mfrID)
}

// InstrumentModel queries the instrument and returns the model of the
// instrument.  InstrumentModel is the getter for the read-only inherent
// attribute Instrument Model described in Section 5.20 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) InstrumentModel() (string, error) {
	return inherent.queryIdentification(modelID)
}

// InstrumentSerialNumber queries the instrument and returns the S/N of the
// instrument.
func (inherent *Inherent) InstrumentSerialNumber() (string, error) {
	return inherent.queryIdentification(snID)
}

// Reset resets the instrument.
func (inherent *Inherent) Reset() error {
	_, err := inherent.inst.WriteString("*RST\n")
	return err
}

// Clear clears the instrument.
func (inherent *Inherent) Clear() error {
	_, err := inherent.inst.WriteString("*CLS\n")
	return err
}

// Disable places the instrument in a quiescent state as quickly as possible.
// Disable provides the method described in Section 6.4 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) Disable() error {
	// FIXME(mdr): Implement!!!!
	return errors.New("disable not implemented")
}

func (inherent *Inherent) queryIdentification(part idPart) (string, error) {
	s, err := inherent.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	return parseIdentification(s, part)
}

func parseIdentification(idn string, part idPart) (string, error) {
	parts := strings.Split(idn, ",")
	if len(parts) != 4 {
		return "", fmt.Errorf("idn string (`%s`) could not be split in four", idn)
	}
	return parts[part], nil
}
