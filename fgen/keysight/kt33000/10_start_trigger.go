// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33000

import (
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

var triggerSlopeToSCPI = map[fgen.TriggerSlope]string{
	fgen.TriggerSlopePositive: "POS",
	fgen.TriggerSlopeNegative: "NEG",
}

var scpiToTriggerSlope = map[string]fgen.TriggerSlope{
	"POS": fgen.TriggerSlopePositive,
	"NEG": fgen.TriggerSlopeNegative,
}

var triggerSourceToSCPI = map[fgen.TriggerSource]string{
	fgen.TriggerSourceInternal: "IMM",
	fgen.TriggerSourceExternal: "EXT",
	fgen.TriggerSourceSoftware: "BUS",
}

var scpiToTriggerSource = map[string]fgen.TriggerSource{
	"IMM": fgen.TriggerSourceInternal,
	"EXT": fgen.TriggerSourceExternal,
	"BUS": fgen.TriggerSourceSoftware,
	"TIM": fgen.TriggerSourceInternal,
}

// StartTriggerDelay returns the delay from the start trigger to the first
// point in the waveform generation.
//
// StartTriggerDelay is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Delay described in Section 10.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerDelay() (time.Duration, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	sec, err := query.Float64f(ctx, ch.inst, "TRIG%d:DEL?", ch.num+1)
	if err != nil {
		return 0, fmt.Errorf("StartTriggerDelay: %w", err)
	}

	return time.Duration(sec * float64(time.Second)), nil
}

// SetStartTriggerDelay sets the delay from the start trigger to the first
// point in the waveform generation.
//
// SetStartTriggerDelay is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Delay described in Section 10.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerDelay(delay time.Duration) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	return ch.inst.Command(ctx, "TRIG%d:DEL %f", ch.num+1, delay.Seconds())
}

// StartTriggerSlope returns the slope of the trigger that starts the
// generator.
//
// StartTriggerSlope is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSlope() (fgen.TriggerSlope, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	var slope fgen.TriggerSlope

	s, err := query.Stringf(ctx, ch.inst, "TRIG%d:SLOP?", ch.num+1)
	if err != nil {
		return slope, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))

	slope, err = ivi.ReverseLookup(scpiToTriggerSlope, s)
	if err != nil {
		return slope, fmt.Errorf("StartTriggerSlope: %w", err)
	}

	return slope, nil
}

// SetStartTriggerSlope sets the slope of the trigger that starts the
// generator.
//
// SetStartTriggerSlope is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSlope(slope fgen.TriggerSlope) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	triggerSlope, err := ivi.LookupSCPI(triggerSlopeToSCPI, slope)
	if err != nil {
		return fmt.Errorf("SetStartTriggerSlope %v: %w", slope, err)
	}

	return ch.inst.Command(ctx, "TRIG%d:SLOP %s", ch.num+1, triggerSlope)
}

// StartTriggerSource returns the source of the start trigger.
//
// StartTriggerSource is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSource() (fgen.TriggerSource, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	var src fgen.TriggerSource

	s, err := query.Stringf(ctx, ch.inst, "TRIG%d:SOUR?", ch.num+1)
	if err != nil {
		return src, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))

	src, err = ivi.ReverseLookup(scpiToTriggerSource, s)
	if err != nil {
		return src, fmt.Errorf("StartTriggerSource: %w", err)
	}

	return src, nil
}

// SetStartTriggerSource specifies the source of the start trigger.
//
// SetStartTriggerSource is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSource(src fgen.TriggerSource) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	triggerSource, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("SetStartTriggerSource %v: %w", src, err)
	}

	return ch.inst.Command(ctx, "TRIG%d:SOUR %s", ch.num+1, triggerSource)
}

func (ch *Channel) StartTriggerThreshold() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerThreshold(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerConfigure(
	_ fgen.TriggerSource,
	_ fgen.TriggerSlope,
) error {
	return ivi.ErrFunctionNotSupported
}
