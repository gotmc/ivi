// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"context"
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

func (ch *Channel) FetchWaveformMeasurement(
	ctx context.Context,
	msrmnt scope.WaveformMeasurement,
) (float64, error) {
	scpiCmd, err := ivi.LookupSCPI(waveformMeasurementToSCPI, msrmnt)
	if err != nil {
		return 0.0, fmt.Errorf("waveform measurement %v not supported: %w", msrmnt, err)
	}

	return query.Float64f(ctx, ch.inst, "%s %s", scpiCmd, ch.name)
}
