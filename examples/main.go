// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package main

import (
	"log"
	"time"

	_ "github.com/gotmc/ivi/driver/tcpip"
	_ "github.com/gotmc/ivi/driver/usbtmc"
	"github.com/gotmc/ivi/fgen/agilent33220"
)

// Can use either a USBTMC or TCP/IP socker to communicate with the function
// generator. Below are two different VISA address strings.
const (
	usbAddress string = "USB0::2391::1031::MY44035349::INSTR"
	tcpAddress string = "TCPIP::10.12.100.242::2055::SOCKET"
)

func main() {

	start := time.Now()
	fgen, err := agilent33220.New(visa.NewInstrument(usbAddress))
	if err != nil {
		log.Fatal("Couldn't open the Agilent 33220A function generator")
	}
	defer fgen.Close()

	log.Printf("%.2fs to setup instrument\n", time.Since(start).Seconds())

	channelNum := 1
	ch, err := fgen.Channel(channelNum)
	if err != nil {
		log.Fatal("Couldn't access channel %d\n", channelNum)
	}

	ch.Waveform(Sine)
	ch.Amplitude(0.1)
	ch.Offset(0.0)
	ch.Frequency(2340)
	ch.Enabled()

}
