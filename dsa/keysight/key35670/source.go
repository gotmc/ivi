// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key35670

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dsa"
	"github.com/gotmc/query"
)

// SetSourceOutputLevel sets the source output level.
func (dev *Key35670) SetSourceOutputLevel(freq float64) error {
	return ivi.Set(dev.inst, "sour:volt:lev:imm:amp %f\n", freq)
}

// SetSourceOutputLevelUnits sets the source output level and the units for the
// source output.
func (dev *Key35670) SetSourceOutputLevelUnits(freq float64, unit dsa.AmpUnits) error {
	return ivi.Set(dev.inst, "sour:volt:lev:imm:amp %f %s\n", freq, unit)
}

// SourceOutputLevel queries the source output level.
func (dev *Key35670) SourceOutputLevel() (float64, error) {
	return query.Float64(dev.inst, "sour:volt:lev:imm:amp?\n")
}

// SourceEnabled determines if the source output is enabled or disabled.
func (dev *Key35670) SourceEnabled() (bool, error) {
	return query.Bool(dev.inst, "OUTP?\n")
}

// SetSourceEnabled sets the source output to enabled or disabled.
func (dev *Key35670) SetSourceEnabled(v bool) error {
	if v {
		return ivi.Set(dev.inst, "OUTP ON\n")
	}
	return ivi.Set(dev.inst, "OUTP OFF\n")
}

// DisableSource is a convenience function for setting the Source Enabled
// attribute to false.
func (dev *Key35670) DisableSource() error {
	return dev.SetSourceEnabled(false)
}

// EnableSource is a convenience function for setting the Source Enabled
// attribute to true.
func (dev *Key35670) EnableSource() error {
	return dev.SetSourceEnabled(true)
}

// SetSourceFrequency sets the source output frequency of the sine source type
// in Hz. Allowable range is 0 to 115 kHz in 15.625 mHz increments.
func (dev *Key35670) SetSourceFrequency(f float64) error {
	if f < 0 || f > 115000 {
		return fmt.Errorf("frequency out of allowable range: %f", f)
	}
	return ivi.Set(dev.inst, "sour:freq %f\n", f)
}

// SourceFrequency queries the source output frequency in Hz.
func (dev *Key35670) SourceFrequency() (float64, error) {
	return query.Float64(dev.inst, "sour:freq?\n")
}

// SetSourceShape sets the source output shape.
func (dev *Key35670) SetSourceShape(shape dsa.SourceShape) error {
	return ivi.Set(dev.inst, "sour:func:shap %s\n", shape)
}

// SourceShape queries the source output shape.
func (dev *Key35670) SourceShape() (dsa.SourceShape, error) {
	s, err := query.String(dev.inst, "sour:func shap?\n")
	if err != nil {
		return "", err
	}
	shape, ok := dsa.SourceShapes[strings.TrimSpace(s)]
	if !ok {
		return "", fmt.Errorf("invalid source shape: %s", s)
	}
	return shape, nil
}
