// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package sdl10xx

import (
	"github.com/gotmc/ivi/load"
)

type voltageCurrent int

const (
	voltageQuery voltageCurrent = iota
	currentQuery
)

// Channel models a SDL10xx load channel.
type Channel struct {
	load.Channel
}
