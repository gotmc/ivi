// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33000

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gotmc/ivi/fgen"

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

// exerciseChannel calls every Channel method that reaches the transport.
// Errors are ignored: only the strings that reach the wire are under test.
func exerciseChannel(ch *Channel) {
	_, _ = ch.OperationMode()
	_ = ch.SetOperationMode(fgen.BurstMode)
	_ = ch.SetOperationMode(fgen.ContinuousMode)
	_, _ = ch.OutputEnabled()
	_ = ch.SetOutputEnabled(true)
	_ = ch.DisableOutput()
	_ = ch.EnableOutput()
	_, _ = ch.OutputImpedance()
	_ = ch.SetOutputImpedance(50.0)
	_ = ch.AbortGeneration()
	_, _ = ch.Amplitude()
	_ = ch.SetAmplitude(2.5)
	_, _ = ch.DCOffset()
	_ = ch.SetDCOffset(0.5)
	_, _ = ch.DutyCycleHigh()
	_ = ch.SetDutyCycleHigh(50.0)
	_, _ = ch.Frequency()
	_ = ch.SetFrequency(1000.0)
	_, _ = ch.StartPhase()
	_ = ch.SetStartPhase(0.0)
	_, _ = ch.StandardWaveform()
	_ = ch.ConfigureStandardWaveform(fgen.Sine, 0.5, 0.0, 100.0, 0.0)
	_, _ = ch.StartTriggerDelay()
	_ = ch.SetStartTriggerDelay(10 * time.Millisecond)
	_, _ = ch.StartTriggerSlope()
	_ = ch.SetStartTriggerSlope(fgen.TriggerSlopePositive)
	_, _ = ch.StartTriggerSource()
	_ = ch.SetStartTriggerSource(fgen.TriggerSourceExternal)
	_, _ = ch.TriggerSource()
	_, _ = ch.InternalTriggerRate()
	_ = ch.SetInternalTriggerRate(1000.0)
	_, _ = ch.BurstCount()
	_ = ch.SetBurstCount(10)
	_, _ = ch.ArbitraryGain()
	_ = ch.SetArbitraryGain(1.0)
	_, _ = ch.ArbitraryOffset()
	_ = ch.SetArbitraryOffset(0.0)

	for _, wave := range []fgen.StandardWaveform{
		fgen.Sine, fgen.Square, fgen.Triangle, fgen.RampUp, fgen.RampDown,
		fgen.DC,
	} {
		_ = ch.SetStandardWaveform(wave)
	}
}

// TestAllModels_EmitValidSCPI drives every supported model, on every one of
// its channels, through every channel-scoped method and checks that each
// string reaching the transport is well-formed SCPI.
//
// The 33200 series and the Trueform models spell channel-scoped commands
// differently, so this sweep is what keeps a new table entry from inheriting
// the wrong command set silently.
func TestAllModels_EmitValidSCPI(t *testing.T) {
	for _, gen := range supportedGenerators {
		t.Run(gen.model, func(t *testing.T) {
			for i, name := range gen.channels {
				t.Run(name, func(t *testing.T) {
					strict := &ivitest.Strict{
						Mock: ivitest.Mock{QueryResp: "1"},
					}

					channels := make([]Channel, len(gen.channels))
					for j, chName := range gen.channels {
						channels[j] = Channel{
							name: chName, inst: strict, num: j,
							family: gen.family, timeout: ivi.DefaultTimeout,
						}
					}

					exerciseChannel(&channels[i])

					strict.Check(t)

					if len(strict.CommandsSent)+len(strict.QueriesSent) == 0 {
						t.Error("no SCPI reached the transport")
					}
				})
			}
		})
	}
}
