// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

/*

# Section 8 IviDmmThermocouple Capability Group

## Section 8.1 IviDmmThermocouple Overview

The IviDmmThermocouple extension group supports DMMs that take temperature
measurements using a thermocouple transducer type.

## Section 8.2 IviDmmThermocouple Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                 | Type     | Access | AppliesTo |
| ------- | ------------------------- | -------- | ------ | --------- |
|   8.2.1 | Thermo Fixed Ref Junction | Real64   | R/W    | N/A       |
|   8.2.2 | Thermo Ref Junction Type  | Int32    | R/W    | N/A       |
|   8.2.3 | Thermocouple Type         | Int32    | R/W    | N/A       |

## Section 8.3 IviDmmThermocouple Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

8.3.2 void Temperature.Thermocouple.Configure(
                            ThermocoupleType type,
                            ReferenceJunctionType referenceJunctionType);

*/
