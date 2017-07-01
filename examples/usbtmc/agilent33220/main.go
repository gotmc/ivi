// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"

	"github.com/gotmc/ivi/fgen/agilent33220"
	"github.com/gotmc/usbtmc"
	_ "github.com/gotmc/usbtmc/driver/truveris"
)

func main() {

	ctx, err := usbtmc.NewContext()
	if err != nil {
		log.Fatalf("Error creating new USB context: %s", err)
	}
	defer ctx.Close()

	res, err := ctx.NewDevice("USB0::2391::1031::MY44035849::INSTR")
	if err != nil {
		log.Fatalf("NewDevice error: %s", err)
	}
	defer res.Close()
	fgen, err := agilent33220.New(res)
	if err != nil {
		log.Fatalf("IVI instrument error: %s", err)
	}
	defer fgen.Close()
	// You can access the channel from the fgen instrument.
	fgen.Ch[0].DisableOutput()
	fgen.Ch[0].SetAmplitude(0.4)
	// Or you can assign the channel to a variabl.
	ch := fgen.Ch[0]
	ch.SetStandardWaveform(agilent33220.Sine)
	ch.SetDCOffset(0.1)
	ch.SetFrequency(2340)
	f, err := ch.Frequency()
	log.Printf("Frequency is %.0f Hz", f)
	ch.EnableOutput()
}
