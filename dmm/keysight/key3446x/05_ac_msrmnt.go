// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dmm"
	"github.com/gotmc/query"
)

// The Truevolt DMMs use three AC input filters (3 Hz, 20 Hz, 200 Hz) selected
// by specifying the lowest frequency expected in the input signal. The usable
// input bandwidth extends to 300 kHz regardless of which filter is selected.
const (
	maxACInputFrequency = 300e3
	minACInputFrequency = 3.0
)

// MaxACFrequency returns the maximum AC input frequency component. The Truevolt
// family does not expose a settable upper cutoff; the usable AC bandwidth ends
// at 300 kHz for all filter settings.
//
// MaxACFrequency is the getter for the read-write IviDmmACMeasurement
// Attribute AC Max Freq described in Section 5.2.1 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) MaxACFrequency() (float64, error) {
	return maxACInputFrequency, nil
}

// SetMaxACFrequency configures the maximum AC input frequency component. The
// Truevolt family has a fixed 300 kHz upper bandwidth, so this method accepts
// any value ≥ 300 kHz (no SCPI command is issued) and rejects smaller values.
//
// SetMaxACFrequency is the setter for the read-write IviDmmACMeasurement
// Attribute AC Max Freq described in Section 5.2.1 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetMaxACFrequency(maxFreq float64) error {
	if maxFreq < maxACInputFrequency {
		return fmt.Errorf(
			"SetMaxACFrequency: %g Hz below fixed instrument bandwidth %g Hz: %w",
			maxFreq, maxACInputFrequency, ivi.ErrValueNotSupported,
		)
	}

	return nil
}

// MinACFrequency returns the minimum AC input frequency component, which
// corresponds to the AC filter cutoff configured on the instrument. The filter
// applies to whichever AC measurement function is currently selected.
//
// MinACFrequency is the getter for the read-write IviDmmACMeasurement
// Attribute AC Min Freq described in Section 5.2.2 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) MinACFrequency() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	scpiFunc, err := d.acFilterFunction()
	if err != nil {
		return 0, err
	}

	return query.Float64f(ctx, d.inst, "%s:BAND?", scpiFunc)
}

// SetMinACFrequency configures the minimum AC input frequency component. The
// instrument selects the 3 Hz, 20 Hz, or 200 Hz filter based on the lowest
// frequency the caller expects to encounter.
//
// SetMinACFrequency is the setter for the read-write IviDmmACMeasurement
// Attribute AC Min Freq described in Section 5.2.2 of IVI-4.2: IviDmm Class
// Specification.
func (d *Driver) SetMinACFrequency(minFreq float64) error {
	if minFreq < minACInputFrequency {
		return fmt.Errorf(
			"SetMinACFrequency: %g Hz below minimum filter cutoff %g Hz: %w",
			minFreq, minACInputFrequency, ivi.ErrValueNotSupported,
		)
	}

	ctx, cancel := d.newContext()
	defer cancel()

	scpiFunc, err := d.acFilterFunction()
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, "%s:BAND %g", scpiFunc, minFreq)
}

// ConfigureACBandwidth configures both the minimum and maximum AC input
// frequency components. Because the instrument has a fixed 300 kHz upper
// bandwidth, the maxFreq argument is validated but does not issue any SCPI.
//
// ConfigureACBandwidth implements the IviDmmACMeasurement function described
// in Section 5.3.1 of IVI-4.2: IviDmm Class Specification.
func (d *Driver) ConfigureACBandwidth(minFreq, maxFreq float64) error {
	if err := d.SetMaxACFrequency(maxFreq); err != nil {
		return err
	}

	return d.SetMinACFrequency(minFreq)
}

// acFilterFunction returns the SCPI branch (VOLT:AC or CURR:AC) that the AC
// bandwidth filter applies to, based on the currently selected measurement
// function. Non-AC measurement functions default to VOLT:AC because that is
// where the attribute is most commonly read or programmed before switching the
// function.
func (d *Driver) acFilterFunction() (string, error) {
	fcn, err := d.MeasurementFunction()
	if err != nil {
		return "", err
	}

	switch fcn {
	case dmm.ACCurrent, dmm.ACPlusDCCurrent:
		return "CURR:AC", nil
	default:
		return "VOLT:AC", nil
	}
}
