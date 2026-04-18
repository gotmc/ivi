// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33220

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
}

func (ch *Channel) StartTriggerDelay() (time.Duration, error) {
	return time.Duration(0), ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerDelay(_ time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerSlope() (fgen.TriggerSlope, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	var slope fgen.TriggerSlope

	s, err := query.String(ctx, ch.inst, "TRIG:SLOP?")
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

func (ch *Channel) SetStartTriggerSlope(slope fgen.TriggerSlope) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	triggerSlope, err := ivi.LookupSCPI(triggerSlopeToSCPI, slope)
	if err != nil {
		return fmt.Errorf("SetStartTriggerSlope %v: %w", slope, err)
	}

	return ch.inst.Command(ctx, "TRIG:SLOP %s", triggerSlope)
}

func (ch *Channel) StartTriggerSource() (fgen.TriggerSource, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	var src fgen.TriggerSource

	s, err := query.String(ctx, ch.inst, "TRIG:SOUR?")
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

func (ch *Channel) SetStartTriggerSource(src fgen.TriggerSource) error {
	ctx, cancel := ch.newContext()
	defer cancel()

	triggerSource, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("SetStartTriggerSource %v: %w", src, err)
	}

	return ch.inst.Command(ctx, "TRIG:SOUR %s", triggerSource)
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
