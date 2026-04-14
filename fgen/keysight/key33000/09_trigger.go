// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key33000

import (
	"fmt"

	"github.com/gotmc/ivi/fgen"
)

// TriggerSource determines the trigger source.
//
// TriggerSource is the getter for the read-write IviFgenTrigger Attribute
// Trigger Source described in Section 9.2.1 of IVI-4.3: IviFgen Class
// Specification.
//
// Deprecated: Use StartTriggerSource instead (Section 10).
func (ch *Channel) TriggerSource() (fgen.OldTriggerSource, error) {
	src, err := ch.StartTriggerSource()
	if err != nil {
		return 0, err
	}

	old, ok := fgen.NewToOldTriggerSource(src)
	if !ok {
		return 0, fmt.Errorf("trigger source %s has no deprecated equivalent", src)
	}

	return old, nil
}

// SetTriggerSource specifies the trigger source.
//
// SetTriggerSource is the setter for the read-write IviFgenTrigger Attribute
// Trigger Source described in Section 9.2.1 of IVI-4.3: IviFgen Class
// Specification.
//
// Deprecated: Use SetStartTriggerSource instead (Section 10).
func (ch *Channel) SetTriggerSource(src fgen.OldTriggerSource) error {
	ts, ok := fgen.OldToNewTriggerSource(src)
	if !ok {
		return fmt.Errorf("trigger source %s not supported", src)
	}

	return ch.SetStartTriggerSource(ts)
}
