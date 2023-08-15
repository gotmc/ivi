// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

/*
Package dsa provides the Defined Values and other structs, methods, etc.
that are common among all Dynamic Signal Analyzers (DSA). Note, DSAs are not
included in one of the IVI specifications.
*/
package dsa

// AmpUnits models the defined values for amplitude units.
type AmpUnits int

// These are the available amplitude units.
const (
	DBVRMS AmpUnits = iota
	VPK
	DBVPK
	V
	DBV
	VRMS
)

var ampUnits = map[AmpUnits]string{
	DBVRMS: "dBVrms",
	VPK:    "Vpeak",
	DBVPK:  "dBVpeak",
	V:      "V",
	DBV:    "dBV",
	VRMS:   "Vrms",
}

func (a AmpUnits) String() string {
	return ampUnits[a]
}
