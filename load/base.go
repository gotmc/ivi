// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package load

type Base interface {
	OutputCount() int
}

type BaseChannel interface {
	SetMode(string) error
}
