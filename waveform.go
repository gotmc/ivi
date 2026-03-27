// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

// Waveform represents acquired waveform data from an oscilloscope channel.
type Waveform struct {
	items []float64
}

// GetAllElements returns all waveform elements.
func (w *Waveform) GetAllElements() ([]float64, error) {
	return w.items, nil
}
