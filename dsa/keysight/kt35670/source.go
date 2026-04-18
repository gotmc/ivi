// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt35670

import (
	"fmt"
	"strings"

	"github.com/gotmc/ivi/dsa"
	"github.com/gotmc/query"
)

// SourceEnabled determines if the source output is enabled or disabled.
func (d *Driver) SourceEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "OUTP?")
}

// SetSourceEnabled sets the source output to enabled or disabled.
func (d *Driver) SetSourceEnabled(v bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if v {
		return d.inst.Command(ctx, "OUTP ON")
	}

	return d.inst.Command(ctx, "OUTP OFF")
}

// DisableSource is a convenience function for disabling the source output.
func (d *Driver) DisableSource() error {
	return d.SetSourceEnabled(false)
}

// EnableSource is a convenience function for enabling the source output.
func (d *Driver) EnableSource() error {
	return d.SetSourceEnabled(true)
}

// SourceShape queries the source output waveform shape.
func (d *Driver) SourceShape() (dsa.SourceShape, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "SOUR:FUNC:SHAP?")
	if err != nil {
		return "", err
	}

	shape, ok := dsa.SourceShapes[strings.TrimSpace(s)]
	if !ok {
		return "", fmt.Errorf("invalid source shape: %s", s)
	}

	return shape, nil
}

// SetSourceShape sets the source output waveform shape.
func (d *Driver) SetSourceShape(shape dsa.SourceShape) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SOUR:FUNC:SHAP %s", shape)
}

// SourceFrequency queries the source output frequency in Hz.
func (d *Driver) SourceFrequency() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SOUR:FREQ?")
}

// SetSourceFrequency sets the source output frequency in Hz. Allowable range
// is 0 to 115 kHz in 15.625 mHz increments.
func (d *Driver) SetSourceFrequency(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if freq < 0 || freq > 115000 {
		return fmt.Errorf("frequency out of allowable range: %f", freq)
	}

	return d.inst.Command(ctx, "SOUR:FREQ %f", freq)
}

// SourceOutputLevel queries the source output level.
func (d *Driver) SourceOutputLevel() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SOUR:VOLT:LEV:IMM:AMP?")
}

// SetSourceOutputLevel sets the source output level.
func (d *Driver) SetSourceOutputLevel(level float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SOUR:VOLT:LEV:IMM:AMP %f", level)
}
