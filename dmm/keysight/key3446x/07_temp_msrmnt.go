// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

// TemperatureTransducerType returns the transducer probe type currently
// selected for temperature measurements.
//
// TemperatureTransducerType is the getter for the read-write
// IviDmmTemperatureMeasurement Attribute Temp Transducer Type described in
// Section 7.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) TemperatureTransducerType() (dmm.TempTransducerType, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	response, err := query.String(ctx, d.inst, "TEMP:TRAN:TYPE?")
	if err != nil {
		return 0, fmt.Errorf("TemperatureTransducerType: %w", err)
	}

	t, err := ivi.ReverseLookup(scpiToTransducerType, strings.TrimSpace(response))
	if err != nil {
		return 0, fmt.Errorf(
			"TemperatureTransducerType: invalid response %q: %w",
			response, err,
		)
	}

	return t, nil
}

// SetTemperatureTransducerType selects the transducer probe type for
// temperature measurements. The 34460A and 34461A support 2-wire and 4-wire
// RTDs and 2-wire and 4-wire thermistors; the 34465A and 34470A additionally
// support thermocouples.
//
// The IVI Thermistor enum maps to the 2-wire SCPI value (THERmistor) because
// the IVI-4.2 specification does not distinguish between 2-wire and 4-wire
// thermistors. Callers that need to select the 4-wire form (FTHermistor) must
// issue the SCPI command directly.
//
// Requesting [dmm.Thermocouple] on a model that lacks thermocouple hardware
// (34460A, 34461A) returns [ivi.ErrUnsupportedModel] without touching the
// instrument.
//
// SetTemperatureTransducerType is the setter for the read-write
// IviDmmTemperatureMeasurement Attribute Temp Transducer Type described in
// Section 7.2.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) SetTemperatureTransducerType(t dmm.TempTransducerType) error {
	if t == dmm.Thermocouple {
		if err := d.requireThermocoupleCapableModel(); err != nil {
			return err
		}
	}

	scpi, err := ivi.LookupSCPI(transducerTypeToSCPI, t)
	if err != nil {
		return fmt.Errorf(
			"SetTemperatureTransducerType: %v not supported: %w",
			t, err,
		)
	}

	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TEMP:TRAN:TYPE %s", scpi)
}

// thermocoupleCapableModels lists the Truevolt models that include
// thermocouple measurement hardware. Per the Operating and Service Guide, the
// TCouple transducer type is available only on the 34465A and 34470A.
var thermocoupleCapableModels = []string{"34465A", "34470A"}

// requireThermocoupleCapableModel returns [ivi.ErrUnsupportedModel] if the
// connected instrument's model cannot measure thermocouples. The model is
// normally cached on the Driver at construction time; when the caller
// suppressed the ID query with [ivi.WithoutIDQuery] and *IDN? failed, the
// cache is empty and we fall back to a live query here.
func (d *Driver) requireThermocoupleCapableModel() error {
	model := d.model
	if model == "" {
		fresh, err := d.InstrumentModel()
		if err != nil {
			return fmt.Errorf(
				"SetTemperatureTransducerType: cannot determine model: %w", err,
			)
		}

		model = strings.TrimSpace(fresh)
	}

	if !slices.Contains(thermocoupleCapableModels, model) {
		return fmt.Errorf(
			"SetTemperatureTransducerType: thermocouple not supported on %q: %w",
			model, ivi.ErrUnsupportedModel,
		)
	}

	return nil
}

// transducerTypeToSCPI maps IVI TempTransducerType values to the SCPI form
// accepted by TEMP:TRAN:TYPE. The 4-wire thermistor (FTHermistor) has no IVI
// equivalent; Thermistor maps to the 2-wire (THERmistor) short form.
var transducerTypeToSCPI = map[dmm.TempTransducerType]string{
	dmm.Thermocouple: "TC",
	dmm.Thermistor:   "THER", //nolint:misspell // SCPI keyword for THERmistor
	dmm.TwoWireRTD:   "RTD",
	dmm.FourWireRTD:  "FRTD",
}

// scpiToTransducerType maps SCPI responses from TEMP:TRAN:TYPE? back to the
// corresponding IVI TempTransducerType. FTH (4-wire thermistor) is mapped to
// the IVI Thermistor enum since the spec has no separate 4-wire variant.
var scpiToTransducerType = map[string]dmm.TempTransducerType{
	"TC":   dmm.Thermocouple,
	"THER": dmm.Thermistor, //nolint:misspell // SCPI response for THERmistor
	"FTH":  dmm.Thermistor,
	"RTD":  dmm.TwoWireRTD,
	"FRTD": dmm.FourWireRTD,
}
