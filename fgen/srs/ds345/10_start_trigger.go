// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

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
	switch s {
	case "0":
		src = fgen.TriggerSourceSoftware
	case "1":
		src = fgen.TriggerSourceInternal
	case "2", "3":
		src = fgen.TriggerSourceExternal
	default:
		return src, errors.New("error determining trigger source")
	}

	return src, nil
}

func (ch *Channel) SetStartTriggerSource(ctx context.Context, src fgen.TriggerSource) error {
	if src == fgen.TriggerSourceSoftware {
		return errors.New("software trigger not supported")
	}

	triggers := map[fgen.TriggerSource]string{
		fgen.TriggerSourceInternal: "1",
		fgen.TriggerSourceExternal: "2",
	}

	return ch.inst.Command(ctx, "TSRC%s", triggers[src])
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
