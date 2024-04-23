// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
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
func (ch *Channel) StartTriggerDelay() (time.Duration, error) {
	return time.Duration(0), ivi.ErrFunctionNotSupported
}

// SetStartTriggerDelay sets the delay from the start trigger to the first
// point in the waveform generation.
//
// SetStartTriggerDelay is the setter for the read-write IviFgenTrigger
// Attribute Start Trigger Delay described in Section 10.2.1 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) SetStartTriggerDelay(_ time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

// StartTriggerSlope returns the slope of the trigger that starts the
// generator.
//
// StartTriggerSlope is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Slope described in Section 10.2.2 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSlope() (fgen.TriggerSlope, error) {
	var slope fgen.TriggerSlope

	s, err := query.String(ch.inst, "TRIG:SLOP?")
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
func (ch *Channel) SetStartTriggerSlope(slope fgen.TriggerSlope) error {
	slopes := map[fgen.TriggerSlope]string{
		fgen.TriggerSlopePositive: "POS",
		fgen.TriggerSlopeNegative: "NEG",
	}

	triggerSlope, ok := slopes[slope]
	if !ok {
		return fmt.Errorf("trigger slope %v not supported", slope)
	}

	return ch.inst.Command("TRIG:SLOP %s", triggerSlope)
}

// StartTriggerSource returns the source of the start trigger.
//
// StartTriggerSource is the getter for the read-write IviFgenTrigger
// Attribute Start Trigger Source described in Section 10.2.3 of IVI-4.3:
// IviFgen Class Specification.
func (ch *Channel) StartTriggerSource() (fgen.TriggerSource, error) {
	var src fgen.TriggerSource

	s, err := query.String(ch.inst, "TRIG:SOUR?")
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
func (ch *Channel) SetStartTriggerSource(src fgen.TriggerSource) error {
	triggers := map[fgen.TriggerSource]string{
		fgen.TriggerSourceInternal: "IMM",
		fgen.TriggerSourceExternal: "EXT",
		fgen.TriggerSourceSoftware: "BUS",
	}

	triggerSource, ok := triggers[src]
	if !ok {
		return fmt.Errorf("trigger source %v not supported", src)
	}

	return ch.inst.Command("TRIG:SOUR %s", triggerSource)
}

func (ch *Channel) StartTriggerThreshold() (float64, error) {
	return 0.0, ivi.ErrFunctionNotSupported
}

func (ch *Channel) SetStartTriggerThreshold(_ float64) error {
	return ivi.ErrFunctionNotSupported
}

func (ch *Channel) StartTriggerConfigure(_ fgen.TriggerSource, _ fgen.TriggerSlope) error {
	return ivi.ErrFunctionNotSupported
}
