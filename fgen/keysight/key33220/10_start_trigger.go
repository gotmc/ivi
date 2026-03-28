// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"context"
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
}

// StartTriggerDelay returns the delay from the start trigger to the first
// point in the waveform generation.
//
// StartTriggerDelay is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Delay described in Section 10.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerDelay(_ context.Context) (time.Duration, error) {
	return time.Duration(0), ivi.ErrFunctionNotSupported
}

// SetStartTriggerDelay sets the delay from the start trigger to the first
// point in the waveform generation.
//
// SetStartTriggerDelay is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Delay described in Section 10.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerDelay(_ context.Context, _ time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

// StartTriggerSlope returns the slope of the trigger that starts the
// generator.
//
// StartTriggerSlope is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSlope(ctx context.Context) (fgen.TriggerSlope, error) {
	var slope fgen.TriggerSlope

	s, err := query.String(ctx, ch.inst, "TRIG:SLOP?")
	if err != nil {
		return slope, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))

	slope, err = ivi.ReverseLookup(scpiToTriggerSlope, s)
	if err != nil {
		return slope, fmt.Errorf("error determining start trigger slope: %w", err)
	}

	return slope, nil
}

// StartTriggerSlope sets the slope of the trigger that starts the generator.
//
// StartTriggerSlope is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSlope(ctx context.Context, slope fgen.TriggerSlope) error {
	triggerSlope, err := ivi.LookupSCPI(triggerSlopeToSCPI, slope)
	if err != nil {
		return fmt.Errorf("trigger slope %v not supported: %w", slope, err)
	}

	return ch.inst.Command(ctx, "TRIG:SLOP %s", triggerSlope)
}

// StartTriggerSource returns the source of the start trigger.
//
// StartTriggerSource is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSource(ctx context.Context) (fgen.TriggerSource, error) {
	var src fgen.TriggerSource

	s, err := query.String(ctx, ch.inst, "TRIG:SOUR?")
	if err != nil {
		return src, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))

	src, err = ivi.ReverseLookup(scpiToTriggerSource, s)
	if err != nil {
		return src, fmt.Errorf("error determining trigger source: %w", err)
	}

	return src, nil
}

// SetStartTriggerSource specifies the source of the start trigger.
//
// SetStartTriggerSource is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSource(ctx context.Context, src fgen.TriggerSource) error {
	triggerSource, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("trigger source %v not supported: %w", src, err)
	}

	return ch.inst.Command(ctx, "TRIG:SOUR %s", triggerSource)
}

func (ch *Channel) StartTriggerThreshold(_ context.Context) (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerThreshold(_ context.Context, _ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerConfigure(
	_ context.Context,
	_ fgen.TriggerSource,
	_ fgen.TriggerSlope,
) error {
	return ivi.ErrFunctionNotSupported
}
