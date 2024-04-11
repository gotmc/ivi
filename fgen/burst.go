// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

// BurstChannel provides the interface for the channel repeated capability for
// the IviFgenBurst capability group.
type BurstChannel interface {
	BurstCount() (int, error)
	SetBurstCount(count int) error
}
