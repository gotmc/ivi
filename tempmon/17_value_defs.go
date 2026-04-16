// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tempmon

// TemperatureUnits specifies the measurement units for a channel.
type TemperatureUnits int

const (
	Celsius TemperatureUnits = iota
	Fahrenheit
	Kelvin
)

var temperatureUnitsDesc = map[TemperatureUnits]string{
	Celsius:    "Celsius",
	Fahrenheit: "Fahrenheit",
	Kelvin:     "Kelvin",
}

func (tu TemperatureUnits) String() string { return temperatureUnitsDesc[tu] }

// ThermocoupleType specifies the thermocouple type for a channel.
type ThermocoupleType int

const (
	TypeB ThermocoupleType = iota
	TypeE
	TypeJ
	TypeK
	TypeN
	TypeR
	TypeS
	TypeT
)

var thermocoupleTypeDesc = map[ThermocoupleType]string{
	TypeB: "Type B",
	TypeE: "Type E",
	TypeJ: "Type J",
	TypeK: "Type K",
	TypeN: "Type N",
	TypeR: "Type R",
	TypeS: "Type S",
	TypeT: "Type T",
}

func (tt ThermocoupleType) String() string { return thermocoupleTypeDesc[tt] }
