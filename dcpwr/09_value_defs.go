// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dcpwr

// CurrentLimitBehavior provides the defined values for the Current Limit
// Behavior defined in Section 4.2.2 and 9 of IVI-4.4: IviDCPwr Class
// Specification.
type CurrentLimitBehavior int

// CurrentTrip and CurrentRegulate are the available Current Limit Behaviors.
// In CurrentTrip behavior, the power supply disables the output when the
// output current is equal to or greater than the value of the Current Limit
// attribute. In CurrentRegulate behavior, the power supply restricts the
// output voltage such that the output current is not greater than the value of
// the Current Limit attribute.
const (
	CurrentTrip CurrentLimitBehavior = iota
	CurrentRegulate
)

// TriggerSource models the defined values for the Trigger Source defined in
// Section 9 IviDCPwr Attribute Value Definitions of IVI-4.4: IviDCPwrClass
// Specification.
type TriggerSource int

// The TriggerSource defined values are the available trigger sources.
const (
	TriggerSourceImmediate TriggerSource = iota
	TriggerSourceExternal
	TriggerSourceSoftware
	TriggerSourceTTL0
	TriggerSourceTTL1
	TriggerSourceTTL2
	TriggerSourceTTL3
	TriggerSourceTTL4
	TriggerSourceTTL5
	TriggerSourceTTL6
	TriggerSourceTTL7
	TriggerSourceECL0
	TriggerSourceECL1
	TriggerSourcePXIStar
	TriggerSourceRTSI0
	TriggerSourceRTSI1
	TriggerSourceRTSI2
	TriggerSourceRTSI3
	TriggerSourceRTSI4
	TriggerSourceRTSI5
	TriggerSourceRTSI6
)

// String implements the Stringer interface for TriggerSource.
func (ts TriggerSource) String() string {
	triggerSources := map[TriggerSource]string{
		TriggerSourceImmediate: "immediate",
		TriggerSourceExternal:  "external",
		TriggerSourceSoftware:  "software",
		TriggerSourceTTL0:      "ttl0",
		TriggerSourceTTL1:      "ttl1",
		TriggerSourceTTL2:      "ttl2",
		TriggerSourceTTL3:      "ttl3",
		TriggerSourceTTL4:      "ttl4",
		TriggerSourceTTL5:      "ttl5",
		TriggerSourceTTL6:      "ttl6",
		TriggerSourceTTL7:      "ttl7",
		TriggerSourceECL0:      "ecl0",
		TriggerSourceECL1:      "ecl1",
		TriggerSourcePXIStar:   "pxi star",
		TriggerSourceRTSI0:     "rtsi0",
		TriggerSourceRTSI1:     "rtsi1",
		TriggerSourceRTSI2:     "rtsi2",
		TriggerSourceRTSI3:     "rtsi3",
		TriggerSourceRTSI4:     "rtsi4",
		TriggerSourceRTSI5:     "rtsi5",
		TriggerSourceRTSI6:     "rtsi6",
	}

	return triggerSources[ts]
}
