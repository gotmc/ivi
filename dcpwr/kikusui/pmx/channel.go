// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package pmx

import "github.com/gotmc/ivi/dcpwr"

type Channel struct {
	dcpwr.Channel
	currentLimitBehavior dcpwr.CurrentLimitBehavior
}
