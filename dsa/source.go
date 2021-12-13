// Copyright (c) 2017-2022 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

// SourceShape models the defined values for the Source Shape.
type SourceShape string

// These are the available source waveforms.
const (
	Sine          SourceShape = "SIN"  // Sinusoidal Waveform
	Random                    = "RAND" // Random Noise
	BurstRandom               = "BRAN" // Burst Random Noise
	PeriodicChirp             = "PCH"  // Periodic Chirp
	BurstChirp                = "BCH"  // Burst Chirp
	Pink                      = "PINK" // Pink Noise
)

func (shape SourceShape) String() string {
	return string(shape)
}

// SourceShapes maps the string value to the SourceShape.
var SourceShapes = map[string]SourceShape{
	"SIN":  Sine,
	"RAND": Random,
	"BRAN": BurstRandom,
	"PCH":  PeriodicChirp,
	"BCH":  BurstChirp,
	"PINK": Pink,
}
