// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
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

// mockInstrument simulates an instrument with configurable delays.
type mockInstrument struct {
	readDelay   time.Duration
	writeDelay  time.Duration
	queryDelay  time.Duration
	cmdDelay    time.Duration
	shouldError bool
}

func (m *mockInstrument) ReadBinary(_ context.Context, p []byte) (int, error) {
	if m.shouldError {
		return 0, errors.New("mock error")
	}
	time.Sleep(m.readDelay)
	copy(p, []byte("test response"))
	return len("test response"), nil
}

func (m *mockInstrument) WriteBinary(_ context.Context, p []byte) (int, error) {
	if m.shouldError {
		return 0, errors.New("mock error")
	}
	time.Sleep(m.writeDelay)
	return len(p), nil
}

func (m *mockInstrument) Close() error { return nil }

func (m *mockInstrument) Command(ctx context.Context, format string, a ...any) error {
	if m.shouldError {
		return errors.New("mock error")
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(m.cmdDelay):
		return nil
	}
}

func (m *mockInstrument) Query(ctx context.Context, s string) (value string, err error) {
	if m.shouldError {
		return "", errors.New("mock error")
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(m.queryDelay):
		return "query response", nil
	}
}

func TestQuery_Success(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 10 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	response, err := mock.Query(ctx, "*IDN?")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response != "query response" {
		t.Errorf("Expected 'query response', got '%s'", response)
	}
}

func TestQuery_Timeout(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 100 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := mock.Query(ctx, "*IDN?")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestCommand_Success(t *testing.T) {
	mock := &mockInstrument{
		cmdDelay: 10 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	err := mock.Command(ctx, "*RST")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCommand_Timeout(t *testing.T) {
	mock := &mockInstrument{
		cmdDelay: 100 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err := mock.Command(ctx, "*RST")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestConcurrentOperations(t *testing.T) {
	mock := &mockInstrument{
		queryDelay: 20 * time.Millisecond,
	}

	// Run multiple queries concurrently
	done := make(chan bool, 3)

	for range 3 {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			response, err := mock.Query(ctx, "*IDN?")
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
	for range 3 {
		select {
		case <-done:
			// Success
		case <-time.After(200 * time.Millisecond):
			t.Error("Concurrent operation timed out")
		}
	}
}
