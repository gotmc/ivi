// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package agilent33220 implements the IVI driver for the Agilent 33220A function
generator.

IVI Instrument Class: IviFgen
Capability Groups Supported (specification section):
   4. IviFgenBase               Partially (missing 4.2.5, 4.2.7)
   5. IviFgenStdFunc            Yes
   6. IviFgenArbWfm             Not Yet
   7. IviFgenArbFrequency       Not Yet
   8. IviFgenArbSeq             No
   9. IviFgenTrigger            Not Yet
  10. IviFgenStartTrigger       Not Yet
  11. IviFgenStopTrigger        Not Yet
  12. IviFgenHoldTrigger        Not Yet
  16. IviFgenSoftwareTrigger    Not Yet
  17. IviFgenBurst              Not Yet (next to work on)

Hardware Information:

State Caching: Not implemented
*/
package agilent33220

import (
	"errors"
	"strings"

	"github.com/gotmc/ivi"
)

const (
	classSpecMajorVersion = 4
	classSpecMinorVersion = 3
	classSpecRevision     = "5.2"
)

var supportedInstrumentModels = []string{
	"33220A",
	"33210A",
}

// Agilent33220 provides the IVI driver for an Agilent 33220A or 33210A
// function generator.
type Agilent33220 struct {
	inst                       ivi.Instrument
	outputCount                int
	Channels                   []Channel
	ClassSpecMajorVersion      int
	ClassSpecMinorVersion      int
	ClassSpecRevision          string
	SupportedInstrumentsModels []string
}

// New creates a new Agilent33220 IVI Instrument.
func New(inst ivi.Instrument) (*Agilent33220, error) {
	outputCount := 1
	ch := Channel{
		id:   0,
		inst: inst,
	}
	channels := make([]Channel, outputCount)
	channels[0] = ch
	fgen := Agilent33220{
		inst:                       inst,
		outputCount:                outputCount,
		Channels:                   channels,
		ClassSpecMajorVersion:      classSpecMajorVersion,
		ClassSpecMinorVersion:      classSpecMinorVersion,
		ClassSpecRevision:          classSpecRevision,
		SupportedInstrumentsModels: supportedInstrumentModels,
	}
	return &fgen, nil
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviFgenBase Attribute Output Count described in
// Section 4.2.1 of IVI-4.3: IviFgen Class Specification.
func (fgen *Agilent33220) OutputCount() int {
	return fgen.outputCount
}

// FirmwardRevision queries the instrument and returns the firmware revision of
// the instrument. FirmwareRevision is the getter for the read-only inherent
// attribute Instrument Firmware Revision described in Section 5.18 of IVI-3.2:
// Inherent Capabilities Specification.
func (fgen *Agilent33220) FirmwareRevision() (string, error) {
	s, err := fgen.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	ret := strings.Split(s, ",")
	ret = strings.Split(ret[3], "-")
	return ret[0], nil
}

// InstrumentManufacturer queries the instrument and returns the manufacturer
// of the instrument. InstrumentManufacturer is the getter for the read-only
// inherent attribute Instrument Manufacturer described in Section 5.19 of
// IVI-3.2: Inherent Capabilities Specification.
func (fgen *Agilent33220) InstrumentManufacturer() (string, error) {
	s, err := fgen.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	ret := strings.Split(s, ",")
	return ret[0], nil
}

// InstrumentModel queries the instrument and returns the model of the
// instrument.  InstrumentModel is the getter for the read-only inherent
// attribute Instrument Model described in Section 5.20 of IVI-3.2: Inherent
// Capabilities Specification.
func (fgen *Agilent33220) InstrumentModel() (string, error) {
	s, err := fgen.inst.Query("*IDN?\n")
	if err != nil {
		return "", err
	}
	ret := strings.Split(s, ",")
	return ret[1], nil
}

// Disable places the instrument in a quiescent state as quickly as possible.
// Disable provides the method described in Section 6.4 of IVI-3.2: Inherent
// Capabilities Specification.
func (fgen *Agilent33220) Disable() error {
	// FIXME(mdr): Implement!!!!
	return errors.New("disable is not yet implemented.")
}

func (fgen *Agilent33220) Reset() error {
	_, err := fgen.inst.WriteString("*RST\n")
	return err
}

func (fgen *Agilent33220) Clear() error {
	_, err := fgen.inst.WriteString("*CLS\n")
	return err
}
