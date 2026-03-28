// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

var triggerSourceToSCPI = map[fgen.TriggerSource]string{
	fgen.TriggerSourceInternal: "1",
	fgen.TriggerSourceExternal: "2",
}

var scpiToTriggerSource = map[string]fgen.TriggerSource{
	"0": fgen.TriggerSourceSoftware,
	"1": fgen.TriggerSourceInternal,
	"2": fgen.TriggerSourceExternal,
	"3": fgen.TriggerSourceExternal,
}

func (ch *Channel) StartTriggerDelay(_ context.Context) (time.Duration, error) {
	return time.Duration(0), ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerDelay(_ context.Context, _ time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerSlope(_ context.Context) (fgen.TriggerSlope, error) {
	return 0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerSlope(_ context.Context, slope fgen.TriggerSlope) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerSource(ctx context.Context) (fgen.TriggerSource, error) {
	var src fgen.TriggerSource

	s, err := query.String(ctx, ch.inst, "TSRC?")
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

func (ch *Channel) SetStartTriggerSource(ctx context.Context, src fgen.TriggerSource) error {
	triggerSource, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("trigger source %s not supported: %w", src, err)
	}

	return ch.inst.Command(ctx, "TSRC%s", triggerSource)
}

func (ch *Channel) StartTriggerThreshold(_ context.Context) (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerThreshold(_ context.Context, threshold float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerConfigure(
	_ context.Context,
	source fgen.TriggerSource,
	slope fgen.TriggerSlope,
) error {
	return ivi.ErrFunctionNotSupported
}
