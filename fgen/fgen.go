// Copyright (c) 2017-2019 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*Package fgen provides the Defined Values and other structs, methods, etc.
that are common among all intstruments meeting the IVI-4.3: IviFgen Class
Specification.

Files are split based on the class capaiblity groups.
*/
package fgen

// OperationMode provides the defined values for the Operation Mode defined in
// Section 4.2.2 of IVI-4.3: IviFgen Class Specification.
type OperationMode int

// Continuous and Burst are the available Operation Modes. "A burst consists of
// a discrete number of waveform cycles. the user uses the attribute of the
// IviFgenTrigger Extension Group to configure the trigger, and the attributes
// of the IviFgenBurst extension group to configure how the function generator
// produces bursts.
const (
	Continuous OperationMode = iota
	Burst
)

// OutputMode determines how the function generator produces waveforms. This
// attribute determines which extension group's functions and attributes are
// used to configure the waveform the function generator produces. OutputMode
// implements the defined values for the Output Mode read-write attribute
// defined in Section 4.2.5 of IVI-4.3: IviFgen Class Specification.
type OutputMode int

// Function indicates the IVI driver uses the attributes and functions of the
// IviFgenStdFunc extension group. Arbitrary indicates the IVI driver uses the
// attributes and functions of the IviFgenArbWfm extension group.
const (
	Function OutputMode = iota
	Arbitrary
	Sequence
)

// TriggerSource models the defined values for the Trigger Source defined in
// Section 9.2.1 of IVI-4.3: IviFgenClass Specification.
type TriggerSource int

// These are the available trigger sources.
const (
	InternalTrigger TriggerSource = iota
	ExternalTrigger
	SoftwareTrigger
)
