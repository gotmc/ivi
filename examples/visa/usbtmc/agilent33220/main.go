// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/fgen/agilent33220"
	_ "github.com/gotmc/usbtmc/driver/google"
	"github.com/gotmc/visa"
	_ "github.com/gotmc/visa/driver/usbtmc"
)

func main() {

	address := "usb0::2391::1031::MY44035849::INSTR"
	res, err := visa.NewResource(address)
	if err != nil {
		log.Fatalf("VISA resource %s: %s", address, err)
	}
	fgen, err := agilent33220.New(res)
	if err != nil {
		log.Fatalf("IVI instrument error: %s", err)
	}
	ch := fgen.Channels[0]
	ch.DisableOutput()
	ch.SetStandardWaveform(ivi.Triangle)
	ch.SetFrequency(1000)
	ch.SetAmplitude(0.25)
	ch.SetDCOffset(0.1)
	ch.EnableOutput()
	wave, err := ch.StandardWaveform()
	if err != nil {
		log.Printf("Error determining standard waveform: %s", err)
	} else {
		log.Printf("Waveform = %s", wave)
	}
	amp, err := ch.Amplitude()
	if err != nil {
		log.Printf("Error determining the amplitude: %s", err)
	} else {
		log.Printf("Amplitude = %.2f Vpp", amp)
	}
	freq, err := ch.Frequency()
	if err != nil {
		log.Printf("Error determining the frequency: %s", err)
	} else {
		log.Printf("Frequency = %.2f Hz", freq)
	}
	// FIXME(mdr): I need to fix the ivi/visa/usbtmc APIs, so that I don't get a
	// double close. Am I going to require that a separate USB context be created
	// separately, or will that be part of opening the device?
	// log.Printf("About to close fgen")
	// err = fgen.Close()
	// if err != nil {
	// log.Printf("Error closing fgen: %s", err)
	// }
	// log.Printf("Closed fgen.")
	err = res.Close()
	if err != nil {
		log.Printf("Error closing VISA resource: %s", err)
	}
	log.Printf("Here")
	// FIXME(mdr): By changing from the truveris USBTMC driver to the google
	// driver, there's no a hang in closing the USB context.
}
