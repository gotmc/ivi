// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"errors"
	"strings"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/query"
)

// TriggerSource determines the trigger srouce. TriggerSource is the getter for
// the read-write IviFgenTrigger Attribute Trigger Source described in Section
// 9.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) TriggerSource() (fgen.OldTriggerSource, error) {
	var src fgen.OldTriggerSource

	s, err := query.String(ch.inst, "TSRC?")
	if err != nil {
		return src, err
	}

	s = strings.TrimSpace(strings.ToUpper(s))
	switch s {
	case "1":
		src = fgen.OldTriggerSourceInternal
	case "2", "3":
		src = fgen.OldTriggerSourceExternal
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
	if src == fgen.OldTriggerSourceSoftware {
		return errors.New("software trigger not supported")
	}

	triggers := map[fgen.OldTriggerSource]string{
		fgen.OldTriggerSourceInternal: "1",
		fgen.OldTriggerSourceExternal: "2",
	}

	return ch.inst.Command("TSRC%s", triggers[src])
}
