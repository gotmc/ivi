// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

// CurrentLimitBehavior provides the defined values for the Current Limit
// Behavior defined in Section 4.2.2 of IVI-4.4: IviDCPwr Class Specification.
type CurrentLimitBehavior int

// These are the defined values for the Current Limit Behavior.
const (
	Trip CurrentLimitBehavior = iota
	Regulate
)
