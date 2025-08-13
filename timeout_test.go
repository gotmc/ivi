// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"errors"
	"testing"
	"time"
)

// mockInstrument simulates an instrument with configurable delays
type mockInstrument struct {
	readDelay   time.Duration
	writeDelay  time.Duration
	queryDelay  time.Duration
	cmdDelay    time.Duration
	shouldError bool
}

func (m *mockInstrument) Read(p []byte) (n int, err error) {
	if m.shouldError {
		return 0, errors.New("mock error")
	}
	time.Sleep(m.readDelay)
	copy(p, []byte("test response"))
	return len("test response"), nil
}

func (m *mockInstrument) Write(p []byte) (n int, err error) {
	if m.shouldError {
		return 0, errors.New("mock error")
	}
	time.Sleep(m.writeDelay)
	return len(p), nil
}

func (m *mockInstrument) WriteString(s string) (n int, err error) {
	if m.shouldError {
		return 0, errors.New("mock error")
	}
	time.Sleep(m.writeDelay)
	return len(s), nil
}

func (m *mockInstrument) Command(format string, a ...any) error {
	if m.shouldError {
		return errors.New("mock error")
	}
	time.Sleep(m.cmdDelay)
	return nil
}

func (m *mockInstrument) Query(s string) (value string, err error) {
	if m.shouldError {
		return "", errors.New("mock error")
	}
	time.Sleep(m.queryDelay)
	return "query response", nil
}

func TestWithTimeout_Read_Success(t *testing.T) {
	mock := &mockInstrument{
		readDelay: 10 * time.Millisecond,
	}

	config := &TimeoutConfig{
		IOTimeout: 100 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	buf := make([]byte, 100)
	n, err := inst.Read(buf)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if n != len("test response") {
		t.Errorf("Expected %d bytes, got %d", len("test response"), n)
	}
}

func TestWithTimeout_Read_Timeout(t *testing.T) {
	mock := &mockInstrument{
		readDelay: 100 * time.Millisecond,
	}

	config := &TimeoutConfig{
		IOTimeout: 10 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	buf := make([]byte, 100)
	_, err := inst.Read(buf)

	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout_Query_Success(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 10 * time.Millisecond,
	}

	config := &TimeoutConfig{
		QueryTimeout: 100 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	response, err := inst.Query("*IDN?")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response != "query response" {
		t.Errorf("Expected 'query response', got '%s'", response)
	}
}

func TestWithTimeout_Query_Timeout(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 100 * time.Millisecond,
	}

	config := &TimeoutConfig{
		QueryTimeout: 10 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	_, err := inst.Query("*IDN?")

	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout_Command_Success(t *testing.T) {
	mock := &mockInstrument{
		cmdDelay: 10 * time.Millisecond,
	}

	config := &TimeoutConfig{
		CommandTimeout: 100 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	err := inst.Command("*RST")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestWithTimeout_Command_Timeout(t *testing.T) {
	mock := &mockInstrument{
		cmdDelay: 100 * time.Millisecond,
	}

	config := &TimeoutConfig{
		CommandTimeout: 10 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	err := inst.Command("*RST")

	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestWithTimeout_SetTimeout(t *testing.T) {
	mock := &mockInstrument{}
	inst := NewWithTimeout(mock, nil)

	// Test default timeout
	if inst.GetTimeout() != 10*time.Second {
		t.Errorf("Expected default timeout of 10s, got %v", inst.GetTimeout())
	}

	// Set new timeout
	newTimeout := 5 * time.Second
	inst.SetTimeout(newTimeout)

	if inst.GetTimeout() != newTimeout {
		t.Errorf("Expected timeout of %v, got %v", newTimeout, inst.GetTimeout())
	}

	// Check that query timeout is adjusted
	if inst.config.QueryTimeout != newTimeout+5*time.Second {
		t.Errorf(
			"Expected query timeout of %v, got %v",
			newTimeout+5*time.Second,
			inst.config.QueryTimeout,
		)
	}
}

func TestWithTimeout_WithContext(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 50 * time.Millisecond,
	}

	config := &TimeoutConfig{
		QueryTimeout: 1 * time.Second,
	}

	inst := NewWithTimeout(mock, config)

	// Test with custom context timeout that's shorter than the operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := inst.QueryWithContext(ctx, "*IDN?")

	if err != context.DeadlineExceeded {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}

	// Test with context that has enough time
	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()

	response, err := inst.QueryWithContext(ctx2, "*IDN?")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response != "query response" {
		t.Errorf("Expected 'query response', got '%s'", response)
	}
}

func TestNewDefaultTimeoutConfig(t *testing.T) {
	config := NewDefaultTimeoutConfig()

	if config.IOTimeout != 10*time.Second {
		t.Errorf("Expected IOTimeout of 10s, got %v", config.IOTimeout)
	}

	if config.QueryTimeout != 15*time.Second {
		t.Errorf("Expected QueryTimeout of 15s, got %v", config.QueryTimeout)
	}

	if config.CommandTimeout != 10*time.Second {
		t.Errorf("Expected CommandTimeout of 10s, got %v", config.CommandTimeout)
	}

	if config.ResetTimeout != 30*time.Second {
		t.Errorf("Expected ResetTimeout of 30s, got %v", config.ResetTimeout)
	}

	if config.ClearTimeout != 5*time.Second {
		t.Errorf("Expected ClearTimeout of 5s, got %v", config.ClearTimeout)
	}
}

func TestWithTimeout_ConcurrentOperations(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 20 * time.Millisecond,
	}

	config := &TimeoutConfig{
		QueryTimeout: 100 * time.Millisecond,
	}

	inst := NewWithTimeout(mock, config)

	// Run multiple queries concurrently
	done := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func() {
			response, err := inst.Query("*IDN?")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if response != "query response" {
				t.Errorf("Expected 'query response', got '%s'", response)
			}
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 3; i++ {
		select {
		case <-done:
			// Success
		case <-time.After(200 * time.Millisecond):
			t.Error("Concurrent operation timed out")
		}
	}
}
