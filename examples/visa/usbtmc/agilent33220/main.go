// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"fmt"
	"log"

	"github.com/gotmc/ivi/fgen/agilent33220"
	_ "github.com/gotmc/usbtmc/driver/truveris"
	"github.com/gotmc/visa"
	_ "github.com/gotmc/visa/drivers/usbtmc"
)

func main() {

	address := "usb0::2391::1031::MY44035849::INSTR"
	res, err := visa.NewResource(address)
	if err != nil {
		log.Fatalf("VISA resource %s: %s", address, err)
	}
	defer res.Close()
	fgen, err := agilent33220.New(res)
	if err != nil {
		log.Fatalf("IVI instrument error: %s", err)
	}
	defer fgen.Close()
	amp, err := fgen.GetAmplitude(0)
	if err != nil {
		log.Fatalf("Problem reading amplitude: %s", err)
	}
	fmt.Printf("Amplitude = %.2f", amp)
	err = fgen.Amplitude(0, 0.24)
	if err != nil {
		log.Fatalf("Problem setting the amplitude: %s", err)
	}
	amp, err = fgen.GetAmplitude(0)
	if err != nil {
		log.Fatalf("Problem reading amplitude: %s", err)
	}
	fmt.Printf("Amplitude = %.2f", amp)
}
