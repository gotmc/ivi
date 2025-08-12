// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
	"github.com/gotmc/query"
)

func (ch *Channel) FetchWaveformMeasurement(msrmnt scope.WaveformMeasurement) (float64, error) {
	switch msrmnt {
	case scope.RiseTime:
		return query.Float64f(ch.inst, ":MEAS:RIS? %s", ch.name)
	case scope.FallTime:
		return query.Float64f(ch.inst, ":MEAS:FALL? %s", ch.name)
	case scope.Frequency:
		return query.Float64f(ch.inst, ":MEAS:FREQ? %s", ch.name)
	case scope.Period:
		return query.Float64f(ch.inst, ":MEAS:PER? %s", ch.name)
	case scope.VoltageRMS:
		return query.Float64f(ch.inst, ":MEAS:VRMS? %s", ch.name)
	case scope.VoltageCycleRMS:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltageMax:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltageMin:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltagePeakToPeak:
		return query.Float64f(ch.inst, ":MEAS:VPP? %s", ch.name)
	case scope.VoltageHigh:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltageLow:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltageAverage:
		return 0.0, ivi.ErrValueNotSupported
	case scope.VoltageCycleAverage:
		return 0.0, ivi.ErrValueNotSupported
	case scope.WidthNegative:
		return 0.0, ivi.ErrValueNotSupported
	case scope.WidthPositive:
		return 0.0, ivi.ErrValueNotSupported
	case scope.DutyCycleNegative:
		return 0.0, ivi.ErrValueNotSupported
	case scope.DutyCyclePositive:
		return 0.0, ivi.ErrValueNotSupported
	case scope.Amplitude:
		return 0.0, ivi.ErrValueNotSupported
	case scope.Overshoot:
		return 0.0, ivi.ErrValueNotSupported
	case scope.Preshoot:
		return 0.0, ivi.ErrValueNotSupported
	}

	return 0.0, ivi.ErrValueNotSupported
}
