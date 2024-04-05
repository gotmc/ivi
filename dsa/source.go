// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dsa

// SourceShape models the defined values for the Source Shape.
type SourceShape string

// These are the available source waveforms.
const (
	Sine          SourceShape = "SIN"  // Sinusoidal Waveform
	Random        SourceShape = "RAND" // Random Noise
	BurstRandom   SourceShape = "BRAN" // Burst Random Noise
	PeriodicChirp SourceShape = "PCH"  // Periodic Chirp
	BurstChirp    SourceShape = "BCH"  // Burst Chirp
	Pink          SourceShape = "PINK" // Pink Noise
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
