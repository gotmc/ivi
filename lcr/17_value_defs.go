// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package lcr

// MeasurementFunction specifies what impedance parameters the LCR meter
// measures. Each function produces a primary and secondary result value.
type MeasurementFunction int

// Available MeasurementFunction values; each pairs a primary quantity with
// a secondary one.
const (
	CpD       MeasurementFunction = iota // Parallel capacitance, Dissipation factor
	CpQ                                  // Parallel capacitance, Quality factor
	CpG                                  // Parallel capacitance, Conductance
	CpRp                                 // Parallel capacitance, Parallel resistance
	CsD                                  // Series capacitance, Dissipation factor
	CsQ                                  // Series capacitance, Quality factor
	CsRs                                 // Series capacitance, Series resistance
	LpD                                  // Parallel inductance, Dissipation factor
	LpQ                                  // Parallel inductance, Quality factor
	LpG                                  // Parallel inductance, Conductance
	LpRp                                 // Parallel inductance, Parallel resistance
	LsD                                  // Series inductance, Dissipation factor
	LsQ                                  // Series inductance, Quality factor
	LsRs                                 // Series inductance, Series resistance
	RX                                   // Resistance, Reactance
	ZThetaDeg                            // |Z|, Phase angle (degrees)
	ZThetaRad                            // |Z|, Phase angle (radians)
	GB                                   // Conductance, Susceptance
	YThetaDeg                            // |Y|, Phase angle (degrees)
	YThetaRad                            // |Y|, Phase angle (radians)
)

var measurementFunctionDesc = map[MeasurementFunction]string{
	CpD:       "Cp-D",
	CpQ:       "Cp-Q",
	CpG:       "Cp-G",
	CpRp:      "Cp-Rp",
	CsD:       "Cs-D",
	CsQ:       "Cs-Q",
	CsRs:      "Cs-Rs",
	LpD:       "Lp-D",
	LpQ:       "Lp-Q",
	LpG:       "Lp-G",
	LpRp:      "Lp-Rp",
	LsD:       "Ls-D",
	LsQ:       "Ls-Q",
	LsRs:      "Ls-Rs",
	RX:        "R-X",
	ZThetaDeg: "Z-θ(deg)",
	ZThetaRad: "Z-θ(rad)",
	GB:        "G-B",
	YThetaDeg: "Y-θ(deg)",
	YThetaRad: "Y-θ(rad)",
}

func (mf MeasurementFunction) String() string { return measurementFunctionDesc[mf] }

// MeasurementSpeed specifies the integration time for a measurement. Longer
// integration times produce more accurate results at the cost of speed.
type MeasurementSpeed int

// Available MeasurementSpeed values.
const (
	MeasurementSpeedShort MeasurementSpeed = iota
	MeasurementSpeedMedium
	MeasurementSpeedLong
)

var measurementSpeedDesc = map[MeasurementSpeed]string{
	MeasurementSpeedShort:  "Short",
	MeasurementSpeedMedium: "Medium",
	MeasurementSpeedLong:   "Long",
}

func (ms MeasurementSpeed) String() string { return measurementSpeedDesc[ms] }

// TriggerSource specifies the source of measurement triggers.
type TriggerSource int

// Available TriggerSource values.
const (
	TriggerSourceInternal TriggerSource = iota
	TriggerSourceExternal
	TriggerSourceBus
	TriggerSourceHold
)

var triggerSourceDesc = map[TriggerSource]string{
	TriggerSourceInternal: "Internal",
	TriggerSourceExternal: "External",
	TriggerSourceBus:      "Bus",
	TriggerSourceHold:     "Hold",
}

func (ts TriggerSource) String() string { return triggerSourceDesc[ts] }

// MeasurementStatus indicates the status of a measurement result.
type MeasurementStatus int

// Available MeasurementStatus values, with the integer values matching the
// status codes returned by the instrument.
const (
	MeasurementStatusNormal          MeasurementStatus = 0
	MeasurementStatusOverload        MeasurementStatus = 1
	MeasurementStatusSourceOverload  MeasurementStatus = 3
	MeasurementStatusALCNotRegulated MeasurementStatus = 4
	MeasurementStatusNoData          MeasurementStatus = -1
)

var measurementStatusDesc = map[MeasurementStatus]string{
	MeasurementStatusNormal:          "Normal",
	MeasurementStatusOverload:        "Overload",
	MeasurementStatusSourceOverload:  "Signal Source Overloaded",
	MeasurementStatusALCNotRegulated: "ALC Unable to Regulate",
	MeasurementStatusNoData:          "No Data",
}

func (ms MeasurementStatus) String() string { return measurementStatusDesc[ms] }
