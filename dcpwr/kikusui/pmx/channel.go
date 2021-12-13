// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package pmx

import "github.com/gotmc/ivi/dcpwr"

// Channel represents the repeated capability of an output channel for the PMX
// DC power supply.
type Channel struct {
	dcpwr.Channel
	currentLimitBehavior dcpwr.CurrentLimitBehavior
}
