// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ds345

import (
	"errors"
	"log"
	"strings"

	"github.com/gotmc/ivi/fgen"
)

// TriggerSource determines the trigger srouce. TriggerSource is the getter for
// the read-write IviFgenTrigger Attribute Trigger Source described in Section
// 9.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) TriggerSource() (fgen.TriggerSource, error) {
	var src fgen.TriggerSource
	s, err := ch.QueryString("TSRC?\n")
	if err != nil {
		return src, err
	}
	s = strings.TrimSpace(strings.ToUpper(s))
	switch s {
	case "1":
		src = fgen.InternalTrigger
	case "2", "3":
		src = fgen.ExternalTrigger
	default:
		return src, errors.New("error determining trigger source")
	}
	return src, nil
}

// SetTriggerSource specifies the trigger srouce. SetTriggerSource is the
// setter for the read-write IviFgenTrigger Attribute Trigger Source described
// in Section 9.2.1 of IVI-4.3: IviFgen Class Specification.
func (ch *Channel) SetTriggerSource(src fgen.TriggerSource) error {
	if src == fgen.SoftwareTrigger {
		return errors.New("software trigger not supported")
	}
	triggers := map[fgen.TriggerSource]string{
		fgen.InternalTrigger: "1",
		fgen.ExternalTrigger: "2",
	}
	log.Printf("Sending command TSRC%s\n", triggers[src])
	return ch.Set("TSRC%s\n", triggers[src])
}
