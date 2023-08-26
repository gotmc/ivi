// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package dp800 implements the IVI driver for the Rigol DP800 series
of programmable linear DC power supplies.

State Caching: Not implemented
*/
package dp800

import (
	"fmt"
	"slices"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

// Confirm that the device driver implements the IviDCPwrBase interface.
var _ dcpwr.Base = (*Device)(nil)

// Device provides the IVI driver for the Rigol DP800 series of DC power
// supplies.
type Device struct {
	inst     ivi.Instrument
	model    string
	Channels []Channel
	ivi.Inherent
}

// New creates a new Rigol DP800 IVI Instrument driver. The New function will
// query the instrument to determine the model and ensure it is one of the
// supported models. If reset is true, then the instrument is reset.
func New(inst ivi.Instrument, reset bool) (*Device, error) {
	supportedModels := []string{
		"DP831A",
		"DP832A",
		"DP821A",
		"DP811A",
		"DP831",
		"DP832",
		"DP821",
		"DP811",
	}
	inherentBase := ivi.InherentBase{
		ClassSpecMajorVersion: 4,
		ClassSpecMinorVersion: 4,
		ClassSpecRevision:     "3.0",
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
			"IviDCPwrTrigger",
		},
		SupportedInstrumentModels: supportedModels,
	}
	inherent := ivi.NewInherent(inst, inherentBase)

	model, err := inherent.InstrumentModel()
	if err != nil {
		return nil, fmt.Errorf("error determinig instrument model: %s", err)
	}

	if !slices.Contains(supportedModels, model) {
		return nil, fmt.Errorf("model %s not supported by this driver", model)
	}

	type genericChannel struct {
		name       string
		minVoltage float64
		maxVoltage float64
		minCurrent float64
		maxCurrent float64
	}

	availableChannels := map[string][]genericChannel{
		"DP831A": []genericChannel{
			{"CH1", 0.0, 8.4, 0.0, 5.3},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -32.0, 0.0, 0.0, 2.1},
		},
		"DP832A": []genericChannel{
			{"CH1", 0.0, 32.0, 0.0, 3.2},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -5.3, 0.0, 0.0, 3.2},
		},
		"DP821A": []genericChannel{
			{"CH1", 0.0, 63.0, 0.0, 1.05},
			{"CH2", 0.0, 8.4, 0.0, 10.5},
		},
		"DP811A": []genericChannel{
			{"Range1", 0.0, 21.0, 0.0, 10.5},
			{"Range2", 0.0, 42.0, 0.0, 5.3},
		},
		"DP831": []genericChannel{
			{"CH1", 0.0, 8.4, 0.0, 5.3},
			{"CH2", 0.0, 32.0, 0.0, 2.1},
			{"CH3", -32.0, 0.0, 0.0, 2.1},
		},
		"DP832": []genericChannel{
			{"CH1", 0.0, 32.0, 0.0, 3.2},
			{"CH2", 0.0, 32.0, 0.0, 3.2},
			{"CH3", -5.3, 0.0, 0.0, 3.2},
		},
		"DP821": []genericChannel{
			{"CH1", 0.0, 63.0, 0.0, 1.05},
			{"CH2", 0.0, 8.4, 0.0, 10.5},
		},
		"DP811": []genericChannel{
			{"Range1", 0.0, 63.0, 0.0, 1.05},
			{"Range2", 0.0, 8.4, 0.0, 10.5},
		},
	}

	genericChannels := availableChannels[model]

	outputCount := len(genericChannels)
	channels := make([]Channel, outputCount)
	for i, genericChannel := range genericChannels {
		ch := Channel{
			name:       genericChannel.name,
			inst:       inst,
			minVoltage: genericChannel.minVoltage,
			maxVoltage: genericChannel.maxVoltage,
			minCurrent: genericChannel.minCurrent,
			maxCurrent: genericChannel.maxCurrent,
		}
		channels[i] = ch
	}
	driver := Device{
		inst:     inst,
		Channels: channels,
		Inherent: inherent,
	}
	if reset {
		err := driver.Reset()
		return &driver, err
	}
	return &driver, nil
}

// AvailableCOMPorts lists the avaialble COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "RS232"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return 5
}

// SerialConfig lists whether the RS-232 serial port is configured as a DCE
// (Data Circuit-Terminating Equipment) or a DTE (Data Terminal Equipment). Computers
// running the IVI program are DTEs; therefore, use a straight through serial
// cable when connecting to DCEs and a null modem cable when connecting to DTEs.
func SerialConfig() string {
	return "DTE"
}

// SerialBaudRates lists the available baud rates for the RS-232 serial port
// from the fastest to the slowest.
func SerialBaudRates() []int {
	return []int{9600, 4800, 2400, 1200, 600, 300}
}

// DefaultSerialBaudRate returns the default baud rate for the RS-232 serial
// port.
func DefaultSerialBaudRate() int {
	return 9600
}

// SerialDataFrames lists the available RS-232 data frame formats.
func SerialDataFrames() []string {
	return []string{"8N2", "7E2", "7O2"}
}

// DefaultSerialDataFrame returns the default RS-232 data frame format.
func DefaultSerialDataFrame() string {
	return "8N2"
}

// ChannelCount returns the number of available output channels.
//
// ChannelCount is the getter for the read-only IviDCPwrBase Attribute Output
// Channel Count described in Section 4.2.7 of IVI-4.4: IviDCPwr Class
// Specification.
func (dev *Device) ChannelCount() int {
	return len(dev.Channels)
}
