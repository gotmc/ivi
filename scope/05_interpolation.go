// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

/*

# Section 5 IviScopeInterpolation Extension Group

## Section 5.1 IviScopeInterpolation Overview

The IviScopeInterpolation extension group defines extensions for oscilloscopes
capable of interpolating values in the waveform record that the oscilloscope’s
acquisition sub-system was unable to digitize.


## Section 5.2 IviScopeInterpolation Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   5.2.1 | Interpolation          | Int32    | R/W    | N/A         |

## Section 5.3 IviScopeInterpolation Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

None.

*/

type Interpolation interface {
	AcquisitionInterpolation(interp InterpolationMethod) error
}
