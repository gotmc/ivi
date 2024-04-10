// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// TriggerChannel provides the interface for the channel repeated capability for
// the IviFgenTrigger capability group.
type TriggerChannel interface {
	TriggerSource() (TriggerSource, error)
	SetTriggerSource(TriggerSource) error
}

// TriggerSource models the defined values for the Trigger Source defined in
// Section 9.2.1 of IVI-4.3: IviFgenClass Specification.
type TriggerSource int

// These are the available trigger sources.
const (
	InternalTrigger TriggerSource = iota
	ExternalTrigger
	SoftwareTrigger
)

func (ts TriggerSource) String() string {
	switch ts {
	case InternalTrigger:
		return "internal trigger"
	case ExternalTrigger:
		return "external trigger"
	case SoftwareTrigger:
		return "software trigger"
	}
	return ""
}
