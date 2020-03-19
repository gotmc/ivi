// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*Package fgen provides the Defined Values and other structs, methods, etc.
that are common among all intstruments meeting the IVI-4.3: IviFgen Class
Specification.

Files are split based on the class capability groups.
*/
package fgen

// TriggerSource models the defined values for the Trigger Source defined in
// Section 9.2.1 of IVI-4.3: IviFgenClass Specification.
type TriggerSource int

// These are the available trigger sources.
const (
	InternalTrigger TriggerSource = iota
	ExternalTrigger
	SoftwareTrigger
)
