// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package e36000 implements the IVI driver for the Keysight/Agilent E3600
// series of power supplies.
//
// The supported models do not share one command set. Multi-output supplies
// such as the E3631A select the output to program with the INSTrument
// subsystem, while single-output supplies such as the E36100B series have no
// INSTrument subsystem at all. The difference is captured by [scpiFamily],
// which determines how every channel-scoped command is spelled, and the model
// reported by *IDN? selects the family. Protection capability varies the same
// way and is recorded per model in [supportedSupplies].
//
// State Caching: Not implemented
package e36000

import (
	"context"
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
)

const (
	specMajorVersion = 4
	specMinorVersion = 4
	specRevision     = "3.0"
)

// Confirm the interfaces implemented by the driver.
var (
	_ dcpwr.Base               = (*Driver)(nil)
	_ dcpwr.BaseChannel        = (*Channel)(nil)
	_ dcpwr.MeasurementChannel = (*Channel)(nil)
	_ dcpwr.Trigger            = (*Driver)(nil)
	_ dcpwr.TriggerChannel     = (*Channel)(nil)
	_ dcpwr.SoftwareTrigger    = (*Driver)(nil)
)

// Driver provides the IVI driver for the Agilent/Keysight E3600 series of DC
// power supplies.
type Driver struct {
	inst     ivi.Transport
	supply   powerSupply
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// Channel models the output channel repeated capability for the DC power
// supply output channel.
type Channel struct {
	inst       ivi.Transport
	name       string
	family     scpiFamily
	protection protectionSupport
	timeout    time.Duration
}

// New creates a new IVI driver for the Keysight/Agilent E3600 series of DC
// power supplies. By default the instrument is queried to determine the model
// and ensure it is supported. Optional driver options can be provided to set
// the timeout, not query the instrument, etc.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	s, err := ivi.NewDriverSetup(inst, ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            700 * time.Millisecond,
		ClearDelay:            700 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviDCPwrBase",
			"IviDCPwrMeasurement",
			"IviDCPwrTrigger",
			"IviDCPwrSoftwareTrigger",
		},
		SupportedInstrumentModels: supportedModels(),
		SupportedBusInterfaces:    []string{"TCPIP", "USB", "GPIB", "SERIAL"},
	}, opts)
	if err != nil {
		return nil, err
	}

	// Channel configuration and command spelling depend on the model.
	model, err := s.Inherent.InstrumentModel()
	if err != nil {
		return nil, fmt.Errorf("error determining instrument model: %w", err)
	}

	supply, err := supplyForModel(model)
	if err != nil {
		return nil, err
	}

	channels := make([]Channel, len(supply.channels))

	for i, name := range supply.channels {
		channels[i] = Channel{
			name:       name,
			inst:       inst,
			family:     supply.family,
			protection: supply.protection,
			timeout:    s.Timeout,
		}
	}

	driver := Driver{
		inst:     inst,
		supply:   supply,
		channels: channels,
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

// scpiFamily identifies how a model names the output a command applies to.
// The families differ in whether the command tree has an INSTrument
// subsystem, so the family determines how every channel-scoped command is
// spelled.
type scpiFamily int

const (
	// instSelect is the command set used by supplies whose outputs are
	// chosen with INSTrument[:SELect], such as the E3631A. The
	// [SOURce:]VOLTage, [SOURce:]CURRent, and MEASure subsystems all act on
	// the selected output, so channel-scoped commands carry an
	// "INST <name>; " prefix and MEASure names the output it reads.
	instSelect scpiFamily = iota
	// singleOutput is the command set used by supplies with exactly one
	// output, such as the E36100B series. Their command tree has no
	// INSTrument subsystem, so commands are unprefixed and MEASure takes no
	// parameter. Sending "INST Output; VOLT?" to an E36102B produces error
	// -113 (Undefined header), and "MEAS:VOLT? Output" produces a parameter
	// error rather than a reading.
	singleOutput
)

// protectionSupport records which output protection subsystems a model
// implements. The zero value claims none, which is correct for the E3631A.
type protectionSupport struct {
	// ovp reports whether the model implements the
	// [SOURce:]VOLTage:PROTection subsystem, which provides the trip level,
	// an enable state, a tripped query, and a clear command.
	ovp bool
	// ocp reports whether the model implements the
	// [SOURce:]CURRent:PROTection subsystem. Enabling it is what makes the
	// supply trip its output off at the current limit rather than regulate
	// there, so it is what backs [dcpwr.CurrentTrip].
	ocp bool
	// outputClear reports whether the model implements OUTPut:PROTection:CLEar,
	// which clears a latched protection trip on the output.
	outputClear bool
	// statusRegisters reports whether the model implements the
	// STATus:OPERation and STATus:QUEStionable condition registers that
	// [Channel.QueryOutputState] reads.
	statusRegisters bool
}

// powerSupply describes the model-specific configuration and capabilities of
// one supported instrument.
type powerSupply struct {
	// model is the model number as reported by the *IDN? query.
	model string
	// channels names the instrument's output channels, in channel order. For
	// the instSelect family these are the identifiers INSTrument[:SELect]
	// accepts; for the singleOutput family the name is descriptive only,
	// since that command set has no way to name an output.
	channels []string
	// family selects the SCPI command set the model implements. The zero
	// value is instSelect.
	family scpiFamily
	// protection records the output protection subsystems the model
	// implements.
	protection protectionSupport
}

// Channel name sets shared by the entries in supportedSupplies. The slices are
// read-only; New copies each name into the Channel it creates.
var (
	oneOutput  = []string{"Output"}
	twoOutputs = []string{"Output 1", "Output 2"}
)

// e36100Protection is the protection capability of the E36100 series, which
// implements over-voltage protection, over-current protection,
// OUTPut:PROTection:CLEar, and the operation and questionable status
// registers.
var e36100Protection = protectionSupport{
	ovp:             true,
	ocp:             true,
	outputClear:     true,
	statusRegisters: true,
}

// supportedSupplies describes the model-specific configuration of every
// instrument this driver supports. supplyForModel selects the entry matching
// the model reported by *IDN?, and supportedModels derives the InherentBase
// model list from it, so this table is the single source of truth for which
// models the driver accepts.
var supportedSupplies = []powerSupply{
	// The E3631A selects its outputs with INSTrument[:SELect] P6V | P25V |
	// N25V and has no protection subsystem of its own.
	{model: "E3631A", channels: []string{"P6V", "P25V", "N25V"}},

	// E36100B series. Verified against the Keysight E36100B Series Operating
	// and Service Guide: the command tree has no INSTrument subsystem, so
	// every output-scoped command is unprefixed, and the series implements
	// VOLTage:PROTection, CURRent:PROTection, OUTPut:PROTection:CLEar, and
	// the STATus condition registers. The E36100A models share the
	// single-output topology, so they are given the same command set; their
	// protection capability has not been verified against a programming
	// guide and is left unclaimed.
	{model: "E36102A", channels: oneOutput, family: singleOutput},
	{model: "E36103A", channels: oneOutput, family: singleOutput},
	{model: "E36104A", channels: oneOutput, family: singleOutput},
	{model: "E36105A", channels: oneOutput, family: singleOutput},
	{model: "E36106A", channels: oneOutput, family: singleOutput},
	{model: "E36102B", channels: oneOutput, family: singleOutput, protection: e36100Protection},
	{model: "E36103B", channels: oneOutput, family: singleOutput, protection: e36100Protection},
	{model: "E36104B", channels: oneOutput, family: singleOutput, protection: e36100Protection},
	{model: "E36105B", channels: oneOutput, family: singleOutput, protection: e36100Protection},
	{model: "E36106B", channels: oneOutput, family: singleOutput, protection: e36100Protection},

	// The models below have not been verified against a programming guide
	// and are left on the instSelect command set they have always used. The
	// single-output models among them almost certainly belong to the
	// singleOutput family for the same reason the E36100 series does, and
	// the multi-output models need the identifiers their INSTrument
	// subsystem actually accepts rather than the descriptive names below
	// (the E3646A through E3649A use OUT1 and OUT2; the E36300 series uses
	// CH1, CH2, and CH3). Both are corrections worth making once a guide is
	// on hand to confirm them.
	{model: "E3632A", channels: oneOutput},
	{model: "E3633A", channels: oneOutput},
	{model: "E3634A", channels: oneOutput},
	{model: "E3640A", channels: oneOutput},
	{model: "E3641A", channels: oneOutput},
	{model: "E3642A", channels: oneOutput},
	{model: "E3643A", channels: oneOutput},
	{model: "E3644A", channels: oneOutput},
	{model: "E3645A", channels: oneOutput},
	{model: "E3646A", channels: twoOutputs},
	{model: "E3647A", channels: twoOutputs},
	{model: "E3648A", channels: twoOutputs},
	{model: "E3649A", channels: twoOutputs},
	{model: "E36154A", channels: oneOutput},
	{model: "E36155A", channels: oneOutput},
	{model: "E36231A", channels: oneOutput},
	{model: "E36232A", channels: oneOutput},
	{model: "E36233A", channels: oneOutput},
	{model: "E36234A", channels: oneOutput},
	{model: "E36311A", channels: []string{"Output 1", "Output 2", "Output 3"}},
	{model: "E36312A", channels: []string{"Output 1", "Output 2", "Output 3"}},
	{model: "E36313A", channels: []string{"Output 1", "Output 2", "Output 3"}},
	{model: "E36441A", channels: []string{"Output 1", "Output 2", "Output 3", "Output 4"}},
	{model: "E36731A", channels: oneOutput},
	{model: "EDU36311A", channels: []string{"Output 1", "Output 2", "Output 3"}},
}

// supportedModels returns the model numbers described by supportedSupplies,
// in table order.
func supportedModels() []string {
	models := make([]string, len(supportedSupplies))
	for i, supply := range supportedSupplies {
		models[i] = supply.model
	}

	return models
}

// supplyForModel returns the supportedSupplies entry for the given model
// number, or [ivi.ErrUnsupportedModel] if the model is not in the table.
func supplyForModel(model string) (powerSupply, error) {
	for _, supply := range supportedSupplies {
		if supply.model == model {
			return supply, nil
		}
	}

	return powerSupply{}, fmt.Errorf("%q: %w", model, ivi.ErrUnsupportedModel)
}

// Channel returns the Channel at the given index, with bounds checking.
func (d *Driver) Channel(index int) (*Channel, error) {
	if index < 0 || index >= len(d.channels) {
		return nil, fmt.Errorf("channel %d: %w", index, ivi.ErrChannelNotFound)
	}

	return &d.channels[index], nil
}

// prefix returns the command prefix that selects this channel for the
// subsystems INSTrument[:SELect] applies to. It is empty for models whose
// command set has no way to select an output.
func (ch *Channel) prefix() string {
	if ch.family == singleOutput {
		return ""
	}

	return "INST " + ch.name + "; "
}

// measureParameter returns the output identifier that the MEASure subsystem
// takes on this model, including the leading space that separates it from the
// query. It is empty for models whose MEASure commands take no parameter.
func (ch *Channel) measureParameter() string {
	if ch.family == singleOutput {
		return ""
	}

	return " " + ch.name
}

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// Close properly shuts down the power supply by returning it to local control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// AvailableCOMPorts lists the available COM ports, including optional ports.
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
