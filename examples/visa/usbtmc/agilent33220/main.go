// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen/agilent33220"
	_ "github.com/gotmc/usbtmc/driver/truveris"
	"github.com/gotmc/visa"
	_ "github.com/gotmc/visa/driver/usbtmc"
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
	ch := fgen.Channels[0]
	ch.SetStandardWaveform(ivi.Triangle)
	ch.SetFrequency(1000)
	ch.SetAmplitude(0.24)
	wave, err := ch.StandardWaveform()
	if err != nil {
		log.Fatalf("Error determining standard waveform: %s", err)
	}
	log.Printf("Waveform = %s", wave)
}
