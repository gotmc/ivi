package dmm

// MeasurementFunction provides the defined values for the Measurement Function defined in
// Section 4.2.1 of IVI-4.2: IviDmm Class Specification.
type MeasurementFunction int

// The MeasurementFunction defined values are the available measurement functions.
const (
	DCVolts MeasurementFunction = iota
	ACVolts
	DCCurrent
	ACCurrent
	TwoWireResistance
	FourWireResistance
	ACPlusDCVolts
	ACPlusDCCurrent
	Frequency
	Period
	Temperature
)

var measurementFunctions = map[MeasurementFunction]string{
	DCVolts:            "DC Volts",
	ACVolts:            "AC Volts",
	DCCurrent:          "DC Current",
	ACCurrent:          "AC Current",
	TwoWireResistance:  "2-wire Resistance",
	FourWireResistance: "4-wire Resistance",
	ACPlusDCVolts:      "AC Plus DC Volts",
	ACPlusDCCurrent:    "AC Plus DC Current",
	Frequency:          "Frequency",
	Period:             "Period",
	Temperature:        "Temperature",
}

func (f MeasurementFunction) String() string {
	return measurementFunctions[f]
}
