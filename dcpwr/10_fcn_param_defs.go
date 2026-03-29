// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

// RangeType provides the defined values for the Output Range Type defined in
// Section 4.3.3 and 10 of IVI-4.4: IviDCPwr Class Specification.
type RangeType int

// Available range types. In VoltageRange, the voltage range is set by the
// Range parameter. In CurrentRange, the current range is set by the Range
// parameter. The actual values of the range types are the same as those shown
// in Section 10 IviDCPwr Function Parameter Value Definitions of IVI-4.4:
// IviDCPwr Class Specification.
const (
	CurrentRange RangeType = iota
	VoltageRange
)

var rangeTypes = map[RangeType]string{
	CurrentRange: "current range",
	VoltageRange: "voltage range",
}

// String implements the Stringer interface for RangeType.
func (rt RangeType) String() string {
	return rangeTypes[rt]
}

// OutputState provides the defined values for the state of the output defined
// in Section 4.3.9 and 10 of IVI-4.4: IviDCPwr Class Specification.
type OutputState int

// Available output states that can be queried. The actual values of the output
// states are the same as those shown in Section 10 IviDCPwr Function Parameter
// Value Definitions of IVI-4.4: IviDCPwr Class Specification.
const (
	ConstantVoltage OutputState = iota
	ConstantCurrent
	OverVoltage
	OverCurrent
	Unregulated
)

var outputStates = map[OutputState]string{
	ConstantVoltage: "constant voltage",
	ConstantCurrent: "constant current",
	OverVoltage:     "over voltage",
	OverCurrent:     "over current",
	Unregulated:     "unregulated",
}

// String implements the Stringer interface for OutputState.
func (os OutputState) String() string {
	return outputStates[os]
}

// MeasurementType provides the defined values for the type of measurement
// defined in Section 7.2.1 and 10 of IVI-4.4: IviDCPwr Class Specification.
type MeasurementType int

// Available output states that can be queried. The actual values of the output
// states are the same as those shown in Section 10 IviDCPwr Function Parameter
// Value Definitions of IVI-4.4: IviDCPwr Class Specification.
const (
	CurrentMeasurement MeasurementType = iota
	VoltageMeasurement
)

var measurementTypes = map[MeasurementType]string{
	CurrentMeasurement: "current",
	VoltageMeasurement: "voltage",
}

// String implements the Stringer interface for MeasurementType.
func (mt MeasurementType) String() string {
	return measurementTypes[mt]
}
