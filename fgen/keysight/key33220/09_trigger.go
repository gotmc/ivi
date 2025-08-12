// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33220

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// TriggerSource determines the trigger srouce.
//
// TriggerSource is the getter for the read-write IviFgenTrigger Attribute
// Trigger Source described in Section 9.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) TriggerSource() (fgen.OldTriggerSource, error) {
	var src fgen.OldTriggerSource

	s, err := query.String(ch.inst, "TRIG:SOUR?")
	if err != nil {
		return src, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))
	switch s {
	case "IMM":
		src = fgen.OldTriggerSourceInternal
	case "EXT":
		src = fgen.OldTriggerSourceExternal
	case "BUS":
		src = fgen.OldTriggerSourceSoftware
	default:
		return src, errors.New("error determining trigger source")
	}

	return src, nil
}

// SetTriggerSource specifies the trigger srouce.
//
// SetTriggerSource is the setter for the read-write IviFgenTrigger Attribute
// Trigger Source described in Section 9.2.1 of IVI-4.3: IviFgen Class
// Specification.
func (ch *Channel) SetTriggerSource(src fgen.OldTriggerSource) error {
	triggers := map[fgen.OldTriggerSource]string{
		fgen.OldTriggerSourceInternal: "IMM",
		fgen.OldTriggerSourceExternal: "EXT",
		fgen.OldTriggerSourceSoftware: "BUS",
	}

	triggerSource, ok := triggers[src]
	if !ok {
		return fmt.Errorf("trigger source %s not supported", src)
	}

	return ch.inst.Command("TRIG:SOUR %s", triggerSource)
}
