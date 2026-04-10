// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package fgen

import "context"

/*

The following information is from the IVI-4.3: IviFgen Class Specification date
October 14, 2016, revision 5.2.

# Section 21 IviFgenTerminalConfiguration Extension Group

## Section 21.1 IviFgenTerminalConfiguration Overview

The IviFgenTerminalConfiguration extension group supports function generators
with the ability to specify whether the output terminals are single-ended or
differential.


## Section 21.2 IviFgenTerminalConfiguration Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute         | Type   | Access | AppliesTo |
| ------- | ----------------  | ------ | ------ | --------- |
|  21.2.1 | Terminal Config   | Int32  | R/W    | Channel   |


## Section 21.3 IviFgenTerminalConfiguration Functions

None.

*/

// TerminalConfiguration provides the interface required for the
// IviFgenTerminalConfiguration extension group.
type TerminalConfigurator interface {
	TerminalConfiguration(ctx context.Context) (TerminalConfigurationType, error)
	SetTerminalConfiguration(ctx context.Context, t TerminalConfigurationType) error
}
