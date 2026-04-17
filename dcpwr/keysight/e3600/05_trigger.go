// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e3600

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/query"
)

// The E3631A supports only the Immediate ("IMM") and Software/Bus ("BUS")
// trigger sources; the E36300 series supports the same set. External and
// hardware trigger sources defined by IVI-4.4 are not supported by either
// family, so SetTriggerSource returns [ivi.ErrValueNotSupported] for them.
var triggerSourceToSCPI = map[dcpwr.TriggerSource]string{
	dcpwr.TriggerSourceImmediate: "IMM",
	dcpwr.TriggerSourceSoftware:  "BUS",
}

var scpiToTriggerSource = map[string]dcpwr.TriggerSource{
	"IMM": dcpwr.TriggerSourceImmediate,
	"BUS": dcpwr.TriggerSourceSoftware,
}

// AbortTrigger sends the SCPI root-level ABORt command to return the
// instrument to the trigger-idle state.
//
// AbortTrigger implements the IviDCPwrTrigger function described in Section
// 5.3.1 of IVI-4.4: IviDCPwr Class Specification.
func (d *Driver) AbortTrigger() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "ABOR")
}

// InitiateTrigger initiates the trigger system so that the next trigger
// event (immediate or software-triggered via [Driver.SendSoftwareTrigger])
// will transfer the triggered voltage and current limit values to the output.
//
// InitiateTrigger implements the IviDCPwrTrigger function described in
// Section 5.3.5 of IVI-4.4: IviDCPwr Class Specification.
func (d *Driver) InitiateTrigger() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INIT")
}

// TriggerSource returns the source the instrument will accept a trigger
// from. The only values returned are [dcpwr.TriggerSourceImmediate] and
// [dcpwr.TriggerSourceSoftware].
//
// TriggerSource is the getter for the read-write IviDCPwrTrigger Attribute
// Trigger Source described in Section 5.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) TriggerSource() (dcpwr.TriggerSource, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	// Trigger source is an instrument-wide setting on the E3631A and applies
	// to the currently selected output; select this channel first so the
	// query result reflects the channel the caller asked about.
	s, err := query.Stringf(ctx, ch.inst, "INST %s; TRIG:SOUR?", ch.name)
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	src, err := ivi.ReverseLookup(scpiToTriggerSource, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	return src, nil
}

// SetTriggerSource specifies the source from which the instrument will accept
// a trigger. Only [dcpwr.TriggerSourceImmediate] and
// [dcpwr.TriggerSourceSoftware] are supported; any other value returns
// [ivi.ErrValueNotSupported].
//
// SetTriggerSource is the setter for the read-write IviDCPwrTrigger Attribute
// Trigger Source described in Section 5.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetTriggerSource(source dcpwr.TriggerSource) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	scpi, err := ivi.LookupSCPI(triggerSourceToSCPI, source)
	if err != nil {
		return fmt.Errorf(
			"SetTriggerSource: trigger source %v not supported: %w",
			source, err,
		)
	}

	return ch.inst.Command(ctx, "INST %s; TRIG:SOUR %s", ch.name, scpi)
}

// TriggeredCurrentLimit returns the current limit (in Amps) that the
// instrument will switch to when the next trigger event occurs on this
// channel after [Driver.InitiateTrigger] has been called.
//
// TriggeredCurrentLimit is the getter for the read-write IviDCPwrTrigger
// Attribute Triggered Current Limit described in Section 5.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) TriggeredCurrentLimit() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, "INST %s; CURR:TRIG?", ch.name)
}

// SetTriggeredCurrentLimit specifies the current limit (in Amps) that the
// instrument will switch to on the next trigger event on this channel.
//
// SetTriggeredCurrentLimit is the setter for the read-write IviDCPwrTrigger
// Attribute Triggered Current Limit described in Section 5.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetTriggeredCurrentLimit(limit float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "INST %s; CURR:TRIG %.4f", ch.name, limit)
}

// TriggeredVoltageLevel returns the voltage level (in Volts) that the
// instrument will switch to when the next trigger event occurs on this
// channel after [Driver.InitiateTrigger] has been called.
//
// TriggeredVoltageLevel is the getter for the read-write IviDCPwrTrigger
// Attribute Triggered Voltage Level described in Section 5.2.3 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) TriggeredVoltageLevel() (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	return query.Float64f(ctx, ch.inst, "INST %s; VOLT:TRIG?", ch.name)
}

// SetTriggeredVoltageLevel specifies the voltage level (in Volts) that the
// instrument will switch to on the next trigger event on this channel.
//
// SetTriggeredVoltageLevel is the setter for the read-write IviDCPwrTrigger
// Attribute Triggered Voltage Level described in Section 5.2.3 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetTriggeredVoltageLevel(level float64) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "INST %s; VOLT:TRIG %.4f", ch.name, level)
}
