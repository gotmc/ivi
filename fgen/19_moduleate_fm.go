// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 19 IviFgenModulateFM Extension Group

## Section 19.1 IviFgenModulateFM Overview

The IviFgenModulateFM Extension Group supports function generators that can
apply frequency modulation to an output signal. The user can enable or disable
frequency modulation, and specify the source of the modulating waveform. If the
function generator supports an internal modulating waveform source, the user
can specify the waveform type, frequency, and peak frequency deviation.


## Section 19.2 IviFgenModulateFM Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute         | Type   | Access | AppliesTo |
| ------- | ----------------  | ------ | ------ | --------- |
|  19.2.1 | FM Enabled        | Bool   | R/W    | Channel   |
|  19.2.2 | FM Internal Dev   | Real64 | R/W    | N/A       |
|  19.2.3 | FM Internal Freq  | Real64 | R/W    | N/A       |
|  19.2.4 | FM Int Waveform   | Int32  | R/W    | N/A       |
|  19.2.5 | FM Source         | Int32  | R/W    | Channel   |


## Section 19.3 IviFgenModulateFM Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

19.3.2 void FM.ConfigureInternal (Double deviation,
                                  StandardWaveform waveformFunction,
                                  Double frequency);

*/
