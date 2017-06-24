// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"github.com/gotmc/ivi/fgen/agilent33220"
	"github.com/gotmc/visa"
	_ "github.com/gotmc/visa/drivers/usbtmc"
)

func main() {

	address := "usb0::2391::1031::MY44035849::INSTR"
	res, err := visa.NewResource(address)
	fgen, err := agilent33220.New(res)

}
