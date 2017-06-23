// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilent34461a

import "github.com/gotmc/visa"

type Agilent34461A struct {
	Resource visa.Resource
}

func (dev *Agilent34461A) DCVolts() (v float64, err error) {

}
