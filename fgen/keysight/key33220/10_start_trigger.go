// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

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
	switch s {
	case "POS":
		slope = fgen.TriggerSlopePositive
	case "NEG":
		slope = fgen.TriggerSlopeNegative
	default:
		return slope, errors.New("error determining start trigger slope")
	}

	return slope, nil
}

// StartTriggerSlope sets the slope of the trigger that starts the generator.
//
// StartTriggerSlope is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSlope(ctx context.Context, slope fgen.TriggerSlope) error {
	slopes := map[fgen.TriggerSlope]string{
		fgen.TriggerSlopePositive: "POS",
		fgen.TriggerSlopeNegative: "NEG",
	}

	triggerSlope, ok := slopes[slope]
	if !ok {
		return fmt.Errorf("trigger slope %v not supported", slope)
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
	switch s {
	case "IMM":
		src = fgen.TriggerSourceInternal
	case "EXT":
		src = fgen.TriggerSourceExternal
	case "BUS":
		src = fgen.TriggerSourceSoftware
	default:
		return src, errors.New("error determining trigger source")
	}

	return src, nil
}

// SetStartTriggerSource specifies the source of the start trigger.
//
// SetStartTriggerSource is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerSource(ctx context.Context, src fgen.TriggerSource) error {
	triggers := map[fgen.TriggerSource]string{
		fgen.TriggerSourceInternal: "IMM",
		fgen.TriggerSourceExternal: "EXT",
		fgen.TriggerSourceSoftware: "BUS",
	}

	triggerSource, ok := triggers[src]
	if !ok {
		return fmt.Errorf("trigger source %v not supported", src)
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
