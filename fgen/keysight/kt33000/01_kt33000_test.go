// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33000

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/internal/ivitest"
)

// idn returns an *IDN? response for the given model.
func idn(model string) string {
	return fmt.Sprintf("Keysight Technologies,%s,MY12345678,5.03-1.19", model)
}

func TestGeneratorForModel(t *testing.T) {
	tests := []struct {
		model        string
		channels     int
		bandwidthHz  int
		arbWaveforms bool
	}{
		{"33210A", 1, 10_000_000, false},
		{"33220A", 1, 20_000_000, true},
		{"33509B", 1, 20_000_000, false},
		{"33512B", 2, 20_000_000, true},
		{"33520B", 2, 30_000_000, false},
		{"33522B", 2, 30_000_000, true},
		{"33611A", 1, 80_000_000, true},
		{"33622A", 2, 120_000_000, true},
		{"EDU33212A", 2, 20_000_000, true},
		{"33502A", 2, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			gen, err := generatorForModel(tt.model)
			if err != nil {
				t.Fatalf("generatorForModel(%q) error: %v", tt.model, err)
			}
			if got := len(gen.channels); got != tt.channels {
				t.Errorf("channels = %d, want %d", got, tt.channels)
			}
			if gen.bandwidthHz != tt.bandwidthHz {
				t.Errorf(
					"bandwidthHz = %d, want %d", gen.bandwidthHz, tt.bandwidthHz,
				)
			}
			if gen.arbWaveforms != tt.arbWaveforms {
				t.Errorf(
					"arbWaveforms = %t, want %t",
					gen.arbWaveforms, tt.arbWaveforms,
				)
			}
		})
	}
}

func TestGeneratorForModel_Unsupported(t *testing.T) {
	_, err := generatorForModel("33220B")
	if !errors.Is(err, ivi.ErrUnsupportedModel) {
		t.Errorf("generatorForModel() error = %v, want ErrUnsupportedModel", err)
	}
}

// TestSupportedGenerators_Table guards the table against duplicate or empty
// entries, since generatorForModel returns the first match.
func TestSupportedGenerators_Table(t *testing.T) {
	seen := make(map[string]bool, len(supportedGenerators))
	for _, gen := range supportedGenerators {
		if gen.model == "" {
			t.Error("table contains an entry with an empty model")
		}
		if seen[gen.model] {
			t.Errorf("table contains duplicate entries for %q", gen.model)
		}
		seen[gen.model] = true

		if len(gen.channels) == 0 {
			t.Errorf("%s: no channels listed", gen.model)
		}
	}

	if got := len(supportedModels()); got != len(supportedGenerators) {
		t.Errorf(
			"supportedModels() returned %d models, want %d",
			got, len(supportedGenerators),
		)
	}
}

func TestNew_ChannelsFromModel(t *testing.T) {
	tests := []struct {
		model string
		names []string
	}{
		{"33220A", []string{"Output"}},
		{"33622A", []string{"Output 1", "Output 2"}},
		{"33502A", []string{"Channel 1", "Channel 2"}},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: idn(tt.model)}
			d, err := New(mock)
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			if got := d.OutputCount(); got != len(tt.names) {
				t.Fatalf("OutputCount() = %d, want %d", got, len(tt.names))
			}

			for i, want := range tt.names {
				ch, err := d.Channel(i)
				if err != nil {
					t.Fatalf("Channel(%d) error: %v", i, err)
				}
				if got := ch.Name(); got != want {
					t.Errorf("Channel(%d).Name() = %q, want %q", i, got, want)
				}
			}
		})
	}
}

func TestNew_UnsupportedModel(t *testing.T) {
	mock := &ivitest.Mock{QueryResp: idn("33220B")}
	if _, err := New(mock); !errors.Is(err, ivi.ErrUnsupportedModel) {
		t.Errorf("New() error = %v, want ErrUnsupportedModel", err)
	}
}

// TestNew_UnsupportedModelWithoutIDQuery verifies that skipping the *IDN?
// validation still rejects a model the driver has no channel configuration
// for, since New cannot build a usable driver without one.
func TestNew_UnsupportedModelWithoutIDQuery(t *testing.T) {
	mock := &ivitest.Mock{QueryResp: idn("33220B")}
	_, err := New(mock, ivi.WithoutIDQuery())
	if !errors.Is(err, ivi.ErrUnsupportedModel) {
		t.Errorf("New() error = %v, want ErrUnsupportedModel", err)
	}
}

func TestDriver_MaxFrequency(t *testing.T) {
	mock := &ivitest.Mock{QueryResp: idn("33611A")}
	d, err := New(mock)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	if got := d.MaxFrequency(); got != 80_000_000 {
		t.Errorf("MaxFrequency() = %v, want %v", got, 80_000_000)
	}
	if !d.SupportsArbWaveform() {
		t.Error("SupportsArbWaveform() = false, want true")
	}
}
