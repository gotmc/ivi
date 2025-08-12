// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"time"
)

// DefaultTimeout is the default timeout for instrument operations.
const DefaultTimeout = 10 * time.Second

// InstrumentWithTimeout extends the Instrument interface with context-aware
// methods that support timeouts and cancellation.
type InstrumentWithTimeout interface {
	Instrument
	ReadWithContext(ctx context.Context, p []byte) (n int, err error)
	WriteWithContext(ctx context.Context, p []byte) (n int, err error)
	WriteStringWithContext(ctx context.Context, s string) (n int, err error)
	CommandWithContext(ctx context.Context, format string, a ...any) error
	QueryWithContext(ctx context.Context, s string) (value string, err error)
	SetTimeout(timeout time.Duration)
	GetTimeout() time.Duration
}

// TimeoutConfig provides configuration for instrument timeouts.
type TimeoutConfig struct {
	// IOTimeout is the timeout for individual I/O operations (read/write).
	IOTimeout time.Duration

	// QueryTimeout is the timeout for query operations which involve both
	// write and read operations.
	QueryTimeout time.Duration

	// CommandTimeout is the timeout for command operations.
	CommandTimeout time.Duration

	// ResetTimeout is the timeout for reset operations which may take longer.
	ResetTimeout time.Duration

	// ClearTimeout is the timeout for clear operations.
	ClearTimeout time.Duration
}

// NewDefaultTimeoutConfig creates a TimeoutConfig with default values.
func NewDefaultTimeoutConfig() *TimeoutConfig {
	return &TimeoutConfig{
		IOTimeout:      10 * time.Second,
		QueryTimeout:   15 * time.Second,
		CommandTimeout: 10 * time.Second,
		ResetTimeout:   30 * time.Second,
		ClearTimeout:   5 * time.Second,
	}
}

// WithTimeout wraps an Instrument to add timeout support while maintaining
// backward compatibility with the existing Instrument interface.
type WithTimeout struct {
	inst   Instrument
	config *TimeoutConfig
}

// NewWithTimeout creates a new timeout-aware instrument wrapper.
func NewWithTimeout(inst Instrument, config *TimeoutConfig) *WithTimeout {
	if config == nil {
		config = NewDefaultTimeoutConfig()
	}
	return &WithTimeout{
		inst:   inst,
		config: config,
	}
}

// Read implements the Instrument interface, using the default timeout.
func (w *WithTimeout) Read(p []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.IOTimeout)
	defer cancel()
	return w.ReadWithContext(ctx, p)
}

// Write implements the Instrument interface, using the default timeout.
func (w *WithTimeout) Write(p []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.IOTimeout)
	defer cancel()
	return w.WriteWithContext(ctx, p)
}

// WriteString implements the Instrument interface, using the default timeout.
func (w *WithTimeout) WriteString(s string) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.IOTimeout)
	defer cancel()
	return w.WriteStringWithContext(ctx, s)
}

// Command implements the Instrument interface, using the default timeout.
func (w *WithTimeout) Command(format string, a ...any) error {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.CommandTimeout)
	defer cancel()
	return w.CommandWithContext(ctx, format, a...)
}

// Query implements the Instrument interface, using the default timeout.
func (w *WithTimeout) Query(s string) (value string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), w.config.QueryTimeout)
	defer cancel()
	return w.QueryWithContext(ctx, s)
}

// ReadWithContext performs a read operation with context support.
func (w *WithTimeout) ReadWithContext(ctx context.Context, p []byte) (n int, err error) {
	type result struct {
		n   int
		err error
	}

	ch := make(chan result, 1)
	go func() {
		n, err := w.inst.Read(p)
		ch <- result{n, err}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res := <-ch:
		return res.n, res.err
	}
}

// WriteWithContext performs a write operation with context support.
func (w *WithTimeout) WriteWithContext(ctx context.Context, p []byte) (n int, err error) {
	type result struct {
		n   int
		err error
	}

	ch := make(chan result, 1)
	go func() {
		n, err := w.inst.Write(p)
		ch <- result{n, err}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res := <-ch:
		return res.n, res.err
	}
}

// WriteStringWithContext performs a string write operation with context support.
func (w *WithTimeout) WriteStringWithContext(ctx context.Context, s string) (n int, err error) {
	type result struct {
		n   int
		err error
	}

	ch := make(chan result, 1)
	go func() {
		n, err := w.inst.WriteString(s)
		ch <- result{n, err}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res := <-ch:
		return res.n, res.err
	}
}

// CommandWithContext sends a command with context support.
func (w *WithTimeout) CommandWithContext(ctx context.Context, format string, a ...any) error {
	ch := make(chan error, 1)
	go func() {
		ch <- w.inst.Command(format, a...)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-ch:
		return err
	}
}

// QueryWithContext performs a query operation with context support.
func (w *WithTimeout) QueryWithContext(ctx context.Context, s string) (value string, err error) {
	type result struct {
		value string
		err   error
	}

	ch := make(chan result, 1)
	go func() {
		value, err := w.inst.Query(s)
		ch <- result{value, err}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res.value, res.err
	}
}

// SetTimeout updates the timeout configuration.
func (w *WithTimeout) SetTimeout(timeout time.Duration) {
	w.config.IOTimeout = timeout
	w.config.QueryTimeout = timeout + 5*time.Second
	w.config.CommandTimeout = timeout
}

// GetTimeout returns the current IO timeout.
func (w *WithTimeout) GetTimeout() time.Duration {
	return w.config.IOTimeout
}

// GetTimeoutConfig returns the full timeout configuration.
func (w *WithTimeout) GetTimeoutConfig() *TimeoutConfig {
	return w.config
}

// SetTimeoutConfig updates the full timeout configuration.
func (w *WithTimeout) SetTimeoutConfig(config *TimeoutConfig) {
	if config != nil {
		w.config = config
	}
}
