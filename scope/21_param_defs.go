// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

/*

# Section 21 IviScope Function Parameter Value Definitions

*/

type AcquisitionStatus int

const (
	AcquisitionComplete AcquisitionStatus = iota
	AcquisitionInprogress
	AcquisitionStatusUnknown
)

// String implements the stringer interface for AcquisitionStatus.
func (as AcquisitionStatus) String() string {
	return map[AcquisitionStatus]string{
		AcquisitionComplete:      "acquisition complete",
		AcquisitionInprogress:    "acquisition in progress",
		AcquisitionStatusUnknown: "acquisition status unknown",
	}[as]
}
