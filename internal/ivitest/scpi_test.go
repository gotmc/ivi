// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivitest

import (
	"errors"
	"testing"
)

func TestCheckSCPI_Valid(t *testing.T) {
	// Every string here is sent by a driver in this repository, so a failure
	// means the validator rejects real traffic rather than that a driver is
	// wrong.
	valid := []string{
		"*RST",
		"*CLS",
		"*IDN?",
		"*TRG",
		"OUTP ON",
		"OUTP?",
		"OUTP1 OFF",
		"OUTP1:LOAD 50.000000",
		"OUTP:PROT:CLE",
		"VOLT 4.1000",
		"VOLT?",
		"VOLT 2.500000 VPP",
		"VOLT:OFFS 0.500000",
		"INST P6V; VOLT 4.1000",
		"INST P25V; CURR?",
		"MEAS:VOLT?",
		"MEAS:VOLT? P6V",
		"SOUR1:FREQ 1000.000000",
		"SOUR2:FUNC:RAMP:SYMM 50",
		"SOUR1:BURS:INT:PER 0.001",
		"TRIG1:SLOP POS",
		"TRIG:SOUR EXT",
		"SYST:COMM:RLST LOC",
		"SYST:LOC",
		"APPL:SIN 1.0000, 2.0000, 3.0000",
		"FUNC:ARB:SRAT 1000.000000",
		"SOUR1:BURS:MODE TRIG",
		"CURR:PROT:STAT ON",
		"MEAS:CURR? (@1)",
	}

	for _, cmd := range valid {
		t.Run(cmd, func(t *testing.T) {
			if err := CheckSCPI(cmd); err != nil {
				t.Errorf("CheckSCPI(%q) = %v, want nil", cmd, err)
			}
		})
	}
}

func TestCheckSCPI_Malformed(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
	}{
		{
			// The bug this validator exists to catch.
			name: "space in channel name",
			cmd:  "INST Output 1; VOLT?",
		},
		{
			name: "space in channel name on measure",
			cmd:  "MEAS:VOLT? Output 1",
		},
		{
			name: "three word parameter",
			cmd:  "INST Output 1 extra",
		},
		{
			name: "empty",
			cmd:  "",
		},
		{
			name: "empty message unit",
			cmd:  "INST P6V;; VOLT?",
		},
		{
			name: "formatting error reached the wire",
			cmd:  "%!(EXTRA string=P6V)VOLT?",
		},
		{
			name: "unformatted verb",
			cmd:  "VOLT %f",
		},
		{
			name: "double space before parameters",
			cmd:  "VOLT  4.1000",
		},
		{
			name: "illegal mnemonic",
			cmd:  "1VOLT 4.1",
		},
		{
			name: "empty parameter",
			cmd:  "APPL:SIN 1.0,,3.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckSCPI(tt.cmd)
			if err == nil {
				t.Fatalf("CheckSCPI(%q) = nil, want an error", tt.cmd)
			}
			if !errors.Is(err, ErrMalformedSCPI) {
				t.Errorf("CheckSCPI(%q) error = %v, want ErrMalformedSCPI",
					tt.cmd, err)
			}
		})
	}
}
