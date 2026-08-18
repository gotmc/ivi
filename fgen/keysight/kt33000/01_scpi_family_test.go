// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt33000

import (
	"context"
	"testing"
	"time"

	"github.com/gotmc/ivi/fgen"
	"github.com/gotmc/ivi/internal/ivitest"
)

// The 33200 series has no channel-suffixed nodes in its command tree, so
// every command it accepts is unprefixed and unsuffixed. The Trueform models
// require the suffixed form. These tests pin both spellings, since sending
// the wrong one produces an undefined header error on the instrument.

// setter names the channel operation under test and applies it.
type setter struct {
	name string
	call func(ch *Channel) error
}

var channelSetters = []setter{
	{"SetOutputEnabled", func(ch *Channel) error {
		return ch.SetOutputEnabled(true)
	}},
	{"SetOutputImpedance", func(ch *Channel) error {
		return ch.SetOutputImpedance(50.0)
	}},
	{"SetFrequency", func(ch *Channel) error {
		return ch.SetFrequency(1000.0)
	}},
	{"SetAmplitude", func(ch *Channel) error {
		return ch.SetAmplitude(2.5)
	}},
	{"SetDCOffset", func(ch *Channel) error {
		return ch.SetDCOffset(0.5)
	}},
	{"SetStandardWaveform", func(ch *Channel) error {
		return ch.SetStandardWaveform(fgen.Sine)
	}},
	{"SetStartTriggerDelay", func(ch *Channel) error {
		return ch.SetStartTriggerDelay(10 * time.Millisecond)
	}},
	{"SetStartTriggerSlope", func(ch *Channel) error {
		return ch.SetStartTriggerSlope(fgen.TriggerSlopePositive)
	}},
	{"SetStartTriggerSource", func(ch *Channel) error {
		return ch.SetStartTriggerSource(fgen.TriggerSourceExternal)
	}},
	{"SetBurstCount", func(ch *Channel) error {
		return ch.SetBurstCount(10)
	}},
	{"SetInternalTriggerRate", func(ch *Channel) error {
		return ch.SetInternalTriggerRate(1000.0)
	}},
	{"SetOperationMode", func(ch *Channel) error {
		return ch.SetOperationMode(fgen.ContinuousMode)
	}},
}

// TestSCPIFamily_Legacy33200Commands verifies that a 33220A receives the
// unsuffixed command forms it accepts. The 33220A rejects OUTP1, SOUR1:, and
// TRIG1:, so a regression here would break the instrument silently at the
// driver level and loudly at the instrument.
func TestSCPIFamily_Legacy33200Commands(t *testing.T) {
	wants := []string{
		"OUTP ON",
		"OUTP:LOAD 50.000000",
		"FREQ 1000.000000",
		"VOLT 2.500000 VPP",
		"VOLT:OFFS 0.500000",
		"FUNC SIN",
		"TRIG:DEL 0.010000",
		"TRIG:SLOP POS",
		"TRIG:SOUR EXT",
		"BURS:NCYC 10",
		"BURS:INT:PER 0.001",
		"BURS:STAT OFF",
	}

	assertChannelCommands(t, "33220A", 0, wants)
}

// TestSCPIFamily_TrueformCommands verifies the suffixed forms used by the
// 33500 and 33600 series.
func TestSCPIFamily_TrueformCommands(t *testing.T) {
	wants := []string{
		"OUTP1 ON",
		"OUTP1:LOAD 50.000000",
		"SOUR1:FREQ 1000.000000",
		"SOUR1:VOLT 2.500000 VPP",
		"SOUR1:VOLT:OFFS 0.500000",
		"SOUR1:FUNC SIN",
		"TRIG1:DEL 0.010000",
		"TRIG1:SLOP POS",
		"TRIG1:SOUR EXT",
		"SOUR1:BURS:NCYC 10",
		"SOUR1:BURS:INT:PER 0.001",
		"SOUR1:BURS:STAT OFF",
	}

	assertChannelCommands(t, "33522B", 0, wants)
}

// TestSCPIFamily_TrueformSecondChannel verifies that the second channel of a
// Trueform model addresses subsystem 2.
func TestSCPIFamily_TrueformSecondChannel(t *testing.T) {
	wants := []string{
		"OUTP2 ON",
		"OUTP2:LOAD 50.000000",
		"SOUR2:FREQ 1000.000000",
		"SOUR2:VOLT 2.500000 VPP",
		"SOUR2:VOLT:OFFS 0.500000",
		"SOUR2:FUNC SIN",
		"TRIG2:DEL 0.010000",
		"TRIG2:SLOP POS",
		"TRIG2:SOUR EXT",
		"SOUR2:BURS:NCYC 10",
		"SOUR2:BURS:INT:PER 0.001",
		"SOUR2:BURS:STAT OFF",
	}

	assertChannelCommands(t, "33522B", 1, wants)
}

// assertChannelCommands runs each setter against the given model and channel
// index, checking that exactly the expected command reaches the instrument.
func assertChannelCommands(t *testing.T, model string, index int, wants []string) {
	t.Helper()

	if len(wants) != len(channelSetters) {
		t.Fatalf("got %d expectations, want %d", len(wants), len(channelSetters))
	}

	for i, s := range channelSetters {
		t.Run(s.name, func(t *testing.T) {
			mock := &ivitest.Mock{}
			d := newTestDriverForModel(mock, model)

			ch, err := d.Channel(index)
			if err != nil {
				t.Fatalf("Channel(%d) error: %v", index, err)
			}

			if err := s.call(ch); err != nil {
				t.Fatalf("%s error: %v", s.name, err)
			}

			if len(mock.CommandsSent) != 1 {
				t.Fatalf("sent %v, want exactly one command", mock.CommandsSent)
			}
			if mock.CommandsSent[0] != wants[i] {
				t.Errorf(
					"sent %q, want %q", mock.CommandsSent[0], wants[i],
				)
			}
		})
	}
}

// TestSCPIFamily_QueryCommands covers the query spellings, which use the same
// prefixes as the setters.
func TestSCPIFamily_QueryCommands(t *testing.T) {
	tests := []struct {
		name  string
		model string
		want  string
		call  func(ch *Channel) error
	}{
		{"33220A OutputEnabled", "33220A", "OUTP?", func(ch *Channel) error {
			_, err := ch.OutputEnabled()
			return err
		}},
		{"33522B OutputEnabled", "33522B", "OUTP1?", func(ch *Channel) error {
			_, err := ch.OutputEnabled()
			return err
		}},
		{"33220A OutputImpedance", "33220A", "OUTP:LOAD?", func(ch *Channel) error {
			_, err := ch.OutputImpedance()
			return err
		}},
		{"33522B OutputImpedance", "33522B", "OUTP1:LOAD?", func(ch *Channel) error {
			_, err := ch.OutputImpedance()
			return err
		}},
		{"33220A Frequency", "33220A", "FREQ?", func(ch *Channel) error {
			_, err := ch.Frequency()
			return err
		}},
		{"33522B Frequency", "33522B", "SOUR1:FREQ?", func(ch *Channel) error {
			_, err := ch.Frequency()
			return err
		}},
		{"33220A StartTriggerSlope", "33220A", "TRIG:SLOP?", func(ch *Channel) error {
			_, err := ch.StartTriggerSlope()
			return err
		}},
		{"33522B StartTriggerSlope", "33522B", "TRIG1:SLOP?", func(ch *Channel) error {
			_, err := ch.StartTriggerSlope()
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &recordingMock{Mock: ivitest.Mock{QueryResp: "POS"}}
			d := newTestDriverForModel(&mock.Mock, tt.model)

			ch, err := d.Channel(0)
			if err != nil {
				t.Fatalf("Channel(0) error: %v", err)
			}
			ch.inst = mock

			// The response is not under test here; only the query string is.
			_ = tt.call(ch)

			if len(mock.queries) != 1 {
				t.Fatalf("sent %v, want exactly one query", mock.queries)
			}
			if mock.queries[0] != tt.want {
				t.Errorf("queried %q, want %q", mock.queries[0], tt.want)
			}
		})
	}
}

// recordingMock extends ivitest.Mock to capture the query strings sent, which
// the shared mock discards.
type recordingMock struct {
	ivitest.Mock
	queries []string
}

func (m *recordingMock) Query(ctx context.Context, s string) (string, error) {
	m.queries = append(m.queries, s)

	return m.Mock.Query(ctx, s)
}

func TestSCPIFamily_LocalControlCommand(t *testing.T) {
	tests := []struct {
		family scpiFamily
		want   string
	}{
		{legacy33200, "SYST:LOC"},
		{trueform, "SYST:COMM:RLST LOC"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.family.localControlCommand(); got != tt.want {
				t.Errorf("localControlCommand() = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestNew_SetsLocalControlCommand verifies that New picks the local control
// command matching the connected model's family.
func TestNew_SetsLocalControlCommand(t *testing.T) {
	tests := []struct {
		model string
		want  string
	}{
		{"33220A", "SYST:LOC"},
		{"33522B", "SYST:COMM:RLST LOC"},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			mock := &ivitest.Mock{QueryResp: idn(tt.model)}
			d, err := New(mock)
			if err != nil {
				t.Fatalf("New() error: %v", err)
			}

			if got := d.LocalControlCommand; got != tt.want {
				t.Errorf("LocalControlCommand = %q, want %q", got, tt.want)
			}
		})
	}
}

// TestSupportedGenerators_Legacy33200Members guards the family assignment: the
// 33200 series models are the only ones using the legacy command set.
func TestSupportedGenerators_Legacy33200Members(t *testing.T) {
	legacy := map[string]bool{"33210A": true, "33220A": true}

	for _, gen := range supportedGenerators {
		wantLegacy := legacy[gen.model]
		if got := gen.family == legacy33200; got != wantLegacy {
			t.Errorf(
				"%s: family == legacy33200 is %t, want %t",
				gen.model, got, wantLegacy,
			)
		}
	}
}
