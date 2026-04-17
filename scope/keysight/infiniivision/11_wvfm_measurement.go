// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
	"github.com/gotmc/query"
)

var waveformMeasurementToSCPI = map[scope.WaveformMeasurement]string{
	scope.RiseTime:          ":MEAS:RIS?",
	scope.FallTime:          ":MEAS:FALL?",
	scope.Frequency:         ":MEAS:FREQ?",
	scope.Period:            ":MEAS:PER?",
	scope.VoltageRMS:        ":MEAS:VRMS?",
	scope.VoltagePeakToPeak: ":MEAS:VPP?",
}

// FetchWaveformMeasurement fetches a specified waveform measurement from a
// previously acquired waveform on this channel.
//
// FetchWaveformMeasurement implements the IviScopeWaveformMeasurement
// function described in Section 11 of IVI-4.1: IviScope Class Specification.
func (ch *Channel) FetchWaveformMeasurement(
	msrmnt scope.WaveformMeasurement,
) (float64, error) {
	ctx, cancel := ch.newContext()
	defer cancel()

	scpiCmd, err := ivi.LookupSCPI(waveformMeasurementToSCPI, msrmnt)
	if err != nil {
		return 0.0, fmt.Errorf("waveform measurement %v not supported: %w", msrmnt, err)
	}

	return query.Float64f(ctx, ch.inst, "%s %s", scpiCmd, ch.name)
}
