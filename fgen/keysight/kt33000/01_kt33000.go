// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

// Package kt33000 implements the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators, covering the 33200, 33500A,
// 33500B, 33600A, and EDU33210 families. This driver corresponds to the
// Keysight IVI.NET Kt33000 driver.
//
// The supported models span two generations of SCPI command set, described
// by [scpiFamily]. The 33500 and later models address channel-specific
// subsystems with a suffix (SOUR1:, OUTP2, TRIG1:), while the single-channel
// 33200 series has no such nodes and takes the unsuffixed form (FREQ, OUTP,
// TRIG:SLOP). New selects the right form from the connected model, so callers
// use the same API for either generation.
//
// The 33500B and 33600A models use LAN port 5025 for SCPI Socket sessions.
// The default GPIB address is 10.
//
// State Caching: Not implemented
package kt33000

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
)

const (
	specMajorVersion   = 4
	specMinorVersion   = 3
	specRevision       = "5.2"
	defaultGPIBAddress = 10
	telnetPort         = 5024
	socketPort         = 5025
)

// Confirm the interfaces implemented by the driver.
var _ fgen.Base = (*Driver)(nil)
var _ fgen.BaseChannel = (*Channel)(nil)
var _ fgen.StdFuncChannel = (*Channel)(nil)
var _ fgen.StartTriggerChannel = (*Channel)(nil)
var _ fgen.TriggerChannel = (*Channel)(nil)
var _ fgen.IntTriggerChannel = (*Channel)(nil)
var _ fgen.BurstChannel = (*Channel)(nil)
var _ fgen.ArbWfm = (*Driver)(nil)
var _ fgen.ArbWfmChannel = (*Channel)(nil)

// Driver provides the IVI driver for the Keysight 33000 series
// function/arbitrary waveform generators.
type Driver struct {
	inst     ivi.Transport
	gen      functionGenerator
	channels []Channel
	timeout  time.Duration
	ivi.Inherent
}

// New creates a new IVI driver for a Keysight 33000 series function/arbitrary
// waveform generator using the given transport layer. By default the
// instrument is queried to determine if it is one of the supported instrument
// models. Optional driver options can be provided to set the timeout, not
// query the instrument, etc.
func New(inst ivi.Transport, opts ...ivi.DriverOption) (*Driver, error) {
	s, err := ivi.NewDriverSetup(inst, ivi.InherentBase{
		ClassSpecMajorVersion: specMajorVersion,
		ClassSpecMinorVersion: specMinorVersion,
		ClassSpecRevision:     specRevision,
		ResetDelay:            500 * time.Millisecond,
		ClearDelay:            500 * time.Millisecond,
		ReturnToLocal:         true,
		GroupCapabilities: []string{
			"IviFgenBase",
			"IviFgenBurst",
			"IviFgenInternalTrigger",
			"IviFgenStdfunc",
			"IviFgenTrigger",
		},
		SupportedInstrumentModels: supportedModels(),
		SupportedBusInterfaces:    []string{"TCPIP", "GPIB", "USB"},
	}, opts)
	if err != nil {
		return nil, err
	}

	// Channel configuration depends on the instrument model.
	model, err := s.Inherent.InstrumentModel()
	if err != nil {
		return nil, fmt.Errorf("error determining instrument model: %w", err)
	}

	gen, err := generatorForModel(model)
	if err != nil {
		return nil, err
	}

	// The command used to return the instrument to local control is family
	// specific, and the family is not known until the model is queried.
	s.Inherent.LocalControlCommand = gen.family.localControlCommand()

	channels := make([]Channel, len(gen.channels))
	for i, name := range gen.channels {
		channels[i] = Channel{
			name: name, inst: inst, num: i,
			family: gen.family, timeout: s.Timeout,
		}
	}

	driver := Driver{
		inst:     inst,
		gen:      gen,
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

// Channel name sets shared by the entries in supportedGenerators. The slices
// are read-only; New copies each name into the Channel it creates.
var (
	oneOutput  = []string{"Output"}
	twoOutputs = []string{"Output 1", "Output 2"}
)

// supportedGenerators describes the model-specific configuration of every
// instrument this driver supports. generatorForModel selects the entry
// matching the model reported by *IDN?, and supportedModels derives the
// InherentBase model list from it, so this table is the single source of
// truth for which models the driver accepts.
var supportedGenerators = []functionGenerator{
	// 33200 series. These are the only models using the legacy33200 command
	// set. The 33210A gains arbitrary waveform capability only with Option
	// 002, which the driver cannot detect from *IDN?.
	{
		model: "33210A", channels: oneOutput, bandwidthHz: 10_000_000,
		family: legacy33200, arbWaveforms: false,
	},
	{
		model: "33220A", channels: oneOutput, bandwidthHz: 20_000_000,
		family: legacy33200, arbWaveforms: true,
	},

	// 33500A/33500B series. Within the series, models ending in 09/10 and
	// 19/20 are function generators without arbitrary waveform capability,
	// while those ending in 11/12 and 21/22 add it. The even-numbered model
	// of each pair is the two-channel version.
	{model: "33509B", channels: oneOutput, bandwidthHz: 20_000_000, arbWaveforms: false},
	{model: "33510B", channels: twoOutputs, bandwidthHz: 20_000_000, arbWaveforms: false},
	{model: "33511B", channels: oneOutput, bandwidthHz: 20_000_000, arbWaveforms: true},
	{model: "33512B", channels: twoOutputs, bandwidthHz: 20_000_000, arbWaveforms: true},
	{model: "33519B", channels: oneOutput, bandwidthHz: 30_000_000, arbWaveforms: false},
	{model: "33520B", channels: twoOutputs, bandwidthHz: 30_000_000, arbWaveforms: false},
	{model: "33521A", channels: oneOutput, bandwidthHz: 30_000_000, arbWaveforms: true},
	{model: "33521B", channels: oneOutput, bandwidthHz: 30_000_000, arbWaveforms: true},
	{model: "33522A", channels: twoOutputs, bandwidthHz: 30_000_000, arbWaveforms: true},
	{model: "33522B", channels: twoOutputs, bandwidthHz: 30_000_000, arbWaveforms: true},

	// 33600A series, following the same numbering convention as the 33500
	// series.
	{model: "33609A", channels: oneOutput, bandwidthHz: 80_000_000, arbWaveforms: false},
	{model: "33610A", channels: twoOutputs, bandwidthHz: 80_000_000, arbWaveforms: false},
	{model: "33611A", channels: oneOutput, bandwidthHz: 80_000_000, arbWaveforms: true},
	{model: "33612A", channels: twoOutputs, bandwidthHz: 80_000_000, arbWaveforms: true},
	{model: "33619A", channels: oneOutput, bandwidthHz: 120_000_000, arbWaveforms: false},
	{model: "33620A", channels: twoOutputs, bandwidthHz: 120_000_000, arbWaveforms: false},
	{model: "33621A", channels: oneOutput, bandwidthHz: 120_000_000, arbWaveforms: true},
	{model: "33622A", channels: twoOutputs, bandwidthHz: 120_000_000, arbWaveforms: true},

	// EDU33210 series.
	{model: "EDU33211A", channels: oneOutput, bandwidthHz: 20_000_000, arbWaveforms: true},
	{model: "EDU33212A", channels: twoOutputs, bandwidthHz: 20_000_000, arbWaveforms: true},

	// The 33502A is a two-channel isolated amplifier accessory rather than a
	// generator, so its channels are amplifier outputs and it has no
	// waveform generation of its own.
	{model: "33502A", channels: []string{"Channel 1", "Channel 2"}},

	// The FG33531A and FG33532A specifications have not been verified
	// against a datasheet, so their bandwidth is left unspecified.
	{model: "FG33531A", channels: oneOutput},
	{model: "FG33532A", channels: twoOutputs},
}

// functionGenerator describes the model-specific configuration and
// capabilities of one supported instrument.
type functionGenerator struct {
	// model is the model number as reported by the *IDN? query.
	model string
	// channels names the instrument's output channels, in channel order.
	channels []string
	// bandwidthHz is the maximum sine wave output frequency in hertz. A zero
	// value means the maximum frequency is not recorded for the model.
	bandwidthHz int
	// family selects the SCPI command set the model implements. The zero
	// value is trueform, which covers every model except the 33200 series.
	family scpiFamily
	// arbWaveforms reports whether the model provides arbitrary waveform
	// capability as standard equipment.
	arbWaveforms bool
}

// scpiFamily identifies the generation of SCPI command set a model
// implements. The two generations differ in whether subsystem keywords carry
// a channel suffix, so the family determines how commands are spelled.
type scpiFamily int

const (
	// trueform is the command set used by the 33500, 33600, and EDU33210
	// series. Channel-specific subsystems carry a suffix: the source
	// subsystem is addressed as SOUR1:/SOUR2:, outputs as OUTP1/OUTP2, and
	// triggers as TRIG1:/TRIG2:.
	trueform scpiFamily = iota
	// legacy33200 is the command set used by the single-channel 33200
	// series. Its command tree has no channel-suffixed nodes, so commands
	// are unprefixed and unsuffixed: FREQ, OUTP, and TRIG:SLOP. Sending the
	// suffixed form to a 33220A produces an undefined header error.
	legacy33200
)

// localControlCommand returns the SCPI command that returns the instrument to
// local front-panel control.
func (family scpiFamily) localControlCommand() string {
	if family == legacy33200 {
		return "SYST:LOC"
	}

	return "SYST:COMM:RLST LOC"
}

// supportedModels returns the model numbers described by supportedGenerators,
// in table order.
func supportedModels() []string {
	models := make([]string, len(supportedGenerators))
	for i, gen := range supportedGenerators {
		models[i] = gen.model
	}

	return models
}

// generatorForModel returns the supportedGenerators entry for the given model
// number, or [ivi.ErrUnsupportedModel] if the model is not in the table.
func generatorForModel(model string) (functionGenerator, error) {
	for _, gen := range supportedGenerators {
		if gen.model == model {
			return gen, nil
		}
	}

	return functionGenerator{}, fmt.Errorf("%q: %w", model, ivi.ErrUnsupportedModel)
}

// MaxFrequency returns the maximum sine wave output frequency in hertz for
// the connected instrument model. It returns 0 when the maximum frequency is
// not recorded for the model.
func (d *Driver) MaxFrequency() float64 {
	return float64(d.gen.bandwidthHz)
}

// SupportsArbWaveform reports whether the connected instrument model provides
// arbitrary waveform capability as standard equipment. Models that offer it
// only as an ordering option (such as the 33210A with Option 002) report
// false, since the option cannot be detected from the *IDN? response.
func (d *Driver) SupportsArbWaveform() bool {
	return d.gen.arbWaveforms
}

// newContext creates a context with the driver's configured timeout.
func (d *Driver) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d.timeout)
}

// Channel returns the Channel at the given index, with bounds checking.
func (d *Driver) Channel(index int) (*Channel, error) {
	if index < 0 || index >= len(d.channels) {
		return nil, fmt.Errorf("channel %d: %w", index, ivi.ErrChannelNotFound)
	}

	return &d.channels[index], nil
}

// Close properly shuts down the function generator by returning it to local
// control.
func (d *Driver) Close() error {
	return d.Inherent.Close()
}

// AvailableCOMPorts lists the available COM ports, including optional ports.
func AvailableCOMPorts() []string {
	return []string{"GPIB", "LAN", "USB"}
}

// DefaultGPIBAddress lists the default GPIB interface address.
func DefaultGPIBAddress() int {
	return defaultGPIBAddress
}

// LANPorts returns a map of the different ports with the key being the type of
// port.
func LANPorts() map[string]int {
	return map[string]int{
		"telnet": telnetPort,
		"socket": socketPort,
	}
}

// Channel models the output channel repeated capability for the function
// generator output channel.
type Channel struct {
	inst    ivi.Transport
	name    string
	num     int // 0-based channel index
	family  scpiFamily
	timeout time.Duration
}

// newContext creates a context with the channel's configured timeout.
func (ch *Channel) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), ch.timeout)
}

// srcPrefix returns the SCPI source subsystem prefix for this channel, such
// as "SOUR1:". It is empty on the 33200 series, whose commands hang directly
// off the root of the command tree.
func (ch *Channel) srcPrefix() string {
	if ch.family == legacy33200 {
		return ""
	}

	return fmt.Sprintf("SOUR%d:", ch.num+1)
}

// chanSuffix returns the channel suffix appended to subsystem keywords that
// carry one, such as the "1" in OUTP1 and TRIG1:. It is empty on the 33200
// series, which rejects the suffixed form.
func (ch *Channel) chanSuffix() string {
	if ch.family == legacy33200 {
		return ""
	}

	return strconv.Itoa(ch.num + 1)
}

// trigPrefix returns the SCPI trigger subsystem prefix for this channel, such
// as "TRIG1:". It is "TRIG:" on the 33200 series.
func (ch *Channel) trigPrefix() string {
	return "TRIG" + ch.chanSuffix() + ":"
}
