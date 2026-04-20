// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcload

// Base is the interface required of every electronic-load driver; it
// provides discovery of the load's channel count.
type Base interface {
	OutputCount() int
}

// BaseChannel is the per-channel interface required of every electronic-load
// driver.
type BaseChannel interface {
	SetMode(mode string) error
}
