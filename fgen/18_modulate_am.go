// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 18 IviFgenModulateAM Extension Group

## Section 18.1 IviFgenModulateAM Overview

The IviFgenModulateAM Extension Group supports function generators that can
apply amplitude modulation to an output signal. The user can enable or disable
amplitude modulation, and specify the source of the modulating waveform. If the
function generator supports an internal modulating waveform source, the user
can specify the waveform, frequency, and modulation depth.


## Section 18.2 IviFgenModulateAM Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute         | Type   | Access | AppliesTo |
| ------- | ----------------  | ------ | ------ | --------- |
|  18.2.1 | AM Enabled        | Bool   | R/W    | Channel   |
|  18.2.2 | AM Internal Depth | Real64 | R/W    | N/A       |
|  18.2.3 | AM Internal Freq  | Real64 | R/W    | N/A       |
|  18.2.4 | AM Int Waveform   | Int32  | R/W    | N/A       |
|  18.2.5 | AM Source         | Int32  | R/W    | Channel   |


## Section 18.3 IviFgenModulateAM Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

18.3.2 void AM.ConfigureInternal (Double depth,
                                  StandardWaveform waveformFunction,
                                  Double frequency);

*/

// ModulateAM provides the interface required for the IviFgenModulateAM
// extension group.
type ModulateAM interface {
	AMModulationInternalDepth() (float64, error)
	SetAMModulationInternalDepth(depth float64) error
	AMModulationInternalFrequency() (float64, error)
	SetAMModulationInternalFrequency(freq float64) error
	AMModulationInternalWaveform() (StandardWaveform, error)
	SetAMModulationInternalWaveform(StandardWaveform) error
	ConfigureInternalAM(depth float64, wave StandardWaveform, freq float64) error
}

// ModulateAMChannel provides the interface for the channel repeated
// capability for the IviFgenModulateAM extension group.
type ModulateAMChannel interface {
	AMModulationEnabled() (bool, error)
	SetAMModulationEnabled(b bool) error
	AMModulationSource() (AMSource, error)
	SetAMModulationSource(AMSource) error
}
