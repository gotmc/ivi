// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import "time"

// DriverOption configures the behavior of a driver's New constructor.
type DriverOption func(*DriverConfig)

// DriverConfig holds the configuration options applied by DriverOption
// functions. Driver constructors call [ApplyOptions] to obtain a populated
// config, then inspect the fields they care about. Fields that are
// driver-specific (e.g., Standalone) are ignored by drivers that don't use
// them.
type DriverConfig struct {
	Timeout     time.Duration
	SkipIDQuery bool
	Reset       bool
	Standalone  bool
}

// ApplyOptions returns a DriverConfig with all the given options applied.
// By default SkipIDQuery is false, meaning driver constructors will query
// *IDN? and validate the model against the driver's supported list.
func ApplyOptions(opts []DriverOption) DriverConfig {
	var cfg DriverConfig
	for _, opt := range opts {
		opt(&cfg)
	}

	return cfg
}

// WithoutIDQuery disables the default *IDN? query and model validation
// performed by driver constructors. Use this only when you know the attached
// instrument is supported but cannot respond to *IDN? (for example, during
// bring-up of a simulated instrument or a test fixture). Unsupported or
// unresponsive instruments will fail later when a method call issues SCPI.
func WithoutIDQuery() DriverOption {
	return func(cfg *DriverConfig) {
		cfg.SkipIDQuery = true
	}
}

// WithReset resets the instrument during construction.
func WithReset() DriverOption {
	return func(cfg *DriverConfig) {
		cfg.Reset = true
	}
}

// WithTimeout overrides the default timeout for instrument I/O operations
// during driver construction and subsequent method calls.
func WithTimeout(t time.Duration) DriverOption {
	return func(cfg *DriverConfig) {
		cfg.Timeout = t
	}
}

// WithStandalone marks the instrument as a standalone unit. This is
// driver-specific and only used by drivers that have different behavior
// depending on whether the instrument is standalone (e.g., the Keysight
// U2751A switch matrix uses different voltage ratings).
func WithStandalone() DriverOption {
	return func(cfg *DriverConfig) {
		cfg.Standalone = true
	}
}
