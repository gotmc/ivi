// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

/*

# Section 4 IviFgenBase Capability Group

## Section 4.1 IviFgenBase Overview

The IviFgenBase capability group supports the most basic function generator
capabilities. The user can configure the output impedance and reference clock
source, and enable or disable the function generator’s output channels.

## Section 4.2 IviFgenBase Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute        | Type   | Access | AppliesTo |
| ------- | ---------------- | ------ | ------ | --------- |
|   4.2.1 | Output Count     | Int32  | RO     | N/A       |
|   4.2.2 | Operation Mode   | Int32  | R/W    | Channel   |
|   4.2.3 | Output Enabled   | Bool   | R/W    | Channel   |
|   4.2.4 | Output Impedance | Real64 | R/W    | Channel   |
|   4.2.5 | Output Mode      | Int32  | R/W    | N/A       |
|   4.2.7 | Ref Clock Source | String | R/W    | N/A       |

## Section 4.3 IviFgenBase Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

4.3.1 void AbortGeneration();
4.3.7 String Output.GetChannelNamne (Int32 index);
4.3.8 void InitiateGeneration()

*/

// Base provides the interface required for the IviFgenBase capability group.
type Base interface {
	OutputCount() int
	OutputMode() (OutputMode, error)
	SetOutputMode(mode OutputMode) error
	ReferenceClockSource() (ClockSource, error)
	SetReferenceClockSource(src ClockSource) error
	AbortGeneration() error
	InitiateGeneration() error
}

// BaseChannel provides the interface required for the channel repeated
// capability for the IviFgenBase capability group.
type BaseChannel interface {
	Name() string
	OperationMode() (OperationMode, error)
	SetOperationMode(mode OperationMode) error
	OutputEnabled() (bool, error)
	SetOutputEnabled(b bool) error
	OutputImpedance() (float64, error)
	SetOutputImpedance(impedance float64) error
}
