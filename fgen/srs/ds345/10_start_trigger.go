// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"errors"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

func (ch *Channel) StartTriggerDelay() (time.Duration, error) {
	return time.Duration(0), ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerDelay(_ time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerSlope() (fgen.TriggerSlope, error) {
	return 0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerSlope(slope fgen.TriggerSlope) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerSource() (fgen.TriggerSource, error) {
	var src fgen.TriggerSource

	s, err := query.String(ch.inst, "TSRC?")
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

func (ch *Channel) SetStartTriggerSource(src fgen.TriggerSource) error {
	if src == fgen.TriggerSourceSoftware {
		return errors.New("software trigger not supported")
	}

	triggers := map[fgen.TriggerSource]string{
		fgen.TriggerSourceInternal: "1",
		fgen.TriggerSourceExternal: "2",
	}

	return ch.inst.Command("TSRC%s", triggers[src])
}

func (ch *Channel) StartTriggerThreshold() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerThreshold(threshold float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerConfigure(source fgen.TriggerSource, slope fgen.TriggerSlope) error {
	return ivi.ErrFunctionNotSupported
}
