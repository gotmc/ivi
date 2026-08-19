// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36000

import (
	"context"
	"errors"
	"regexp"
	"slices"
	"testing"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dcpwr"
	"github.com/gotmc/ivi/internal/ivitest"
)

// queryRecorder wraps [ivitest.Mock] to capture the query strings sent, which
// the mock itself discards. How a query is spelled is exactly what the
// command set tests need to check.
type queryRecorder struct {
	ivitest.Mock
	QueriesSent []string
}

func (q *queryRecorder) Query(ctx context.Context, cmd string) (string, error) {
	q.QueriesSent = append(q.QueriesSent, cmd)

	return q.Mock.Query(ctx, cmd)
}

// channelForModel builds the Channel at the given index for a model exactly
// as New would, so the tests exercise the same family and protection wiring
// the driver assigns from the model reported by *IDN?.
func channelForModel(
	t *testing.T,
	inst ivi.Transport,
	model string,
	index int,
) *Channel {
	t.Helper()

	supply, err := supplyForModel(model)
	if err != nil {
		t.Fatalf("supplyForModel(%q) error: %v", model, err)
	}

	if index >= len(supply.channels) {
		t.Fatalf(
			"model %q has %d channels, want index %d",
			model, len(supply.channels), index,
		)
	}

	return &Channel{
		inst:       inst,
		name:       supply.channels[index],
		num:        index,
		family:     supply.family,
		protection: supply.protection,
	}
}

func TestSupplyForModel(t *testing.T) {
	tests := []struct {
		model          string
		wantChannels   int
		wantFamily     scpiFamily
		wantProtection protectionSupport
	}{
		{"E3631A", 3, instSelect, protectionSupport{}},
		{"E36102B", 1, singleOutput, e36100Protection},
		{"E36106B", 1, singleOutput, e36100Protection},
		{"E36102A", 1, singleOutput, protectionSupport{}},
		{"E3646A", 2, instSelect, protectionSupport{}},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			supply, err := supplyForModel(tt.model)
			if err != nil {
				t.Fatalf("supplyForModel(%q) error: %v", tt.model, err)
			}
			if got := len(supply.channels); got != tt.wantChannels {
				t.Errorf("channels = %d, want %d", got, tt.wantChannels)
			}
			if supply.family != tt.wantFamily {
				t.Errorf("family = %v, want %v", supply.family, tt.wantFamily)
			}
			if supply.protection != tt.wantProtection {
				t.Errorf(
					"protection = %+v, want %+v",
					supply.protection, tt.wantProtection,
				)
			}
		})
	}
}

func TestSupplyForModel_Unsupported(t *testing.T) {
	_, err := supplyForModel("E3631B")
	if !errors.Is(err, ivi.ErrUnsupportedModel) {
		t.Errorf("supplyForModel() = %v, want ErrUnsupportedModel", err)
	}
}

func TestSupportedModels(t *testing.T) {
	models := supportedModels()
	if len(models) != len(supportedSupplies) {
		t.Fatalf(
			"supportedModels() returned %d models, want %d",
			len(models), len(supportedSupplies),
		)
	}

	for _, want := range []string{"E3631A", "E36102B", "EDU36311A"} {
		if !slices.Contains(models, want) {
			t.Errorf("supportedModels() missing %q", want)
		}
	}
}

// TestChannelCommandSpelling covers the differences that motivate
// [scpiFamily]. The E3631A selects the output it programs with
// INSTrument[:SELect]; the E36100B series has no INSTrument subsystem at all,
// so the same prefix would produce error -113 (Undefined header) on an
// E36102B; and the E36400 series names the output with a trailing channel
// list instead of selecting one.
//
// The set and query forms differ for the channel list family, since the list
// follows a comma when the command carries a value and a space when it does
// not.
func TestChannelCommandSpelling(t *testing.T) {
	tests := []struct {
		name      string
		model     string
		index     int
		wantSet   string
		wantGet   string
		wantMeas  string
		wantOutSt string
	}{
		{
			name: "E3631A 6V", model: "E3631A", index: 0,
			wantSet:  "INST P6V; VOLT %.4f",
			wantGet:  "INST P6V; VOLT?",
			wantMeas: "MEAS:VOLT? P6V",
			// OUTPut:STATe is global on the E3631A.
			wantOutSt: "OUTP ON",
		},
		{
			name: "E3631A -25V", model: "E3631A", index: 2,
			wantSet:   "INST N25V; VOLT %.4f",
			wantGet:   "INST N25V; VOLT?",
			wantMeas:  "MEAS:VOLT? N25V",
			wantOutSt: "OUTP ON",
		},
		{
			name: "E36102B", model: "E36102B", index: 0,
			wantSet:   "VOLT %.4f",
			wantGet:   "VOLT?",
			wantMeas:  "MEAS:VOLT?",
			wantOutSt: "OUTP ON",
		},
		{
			name: "E36441A channel 1", model: "E36441A", index: 0,
			wantSet:   "VOLT %.4f,(@1)",
			wantGet:   "VOLT? (@1)",
			wantMeas:  "MEAS:VOLT? (@1)",
			wantOutSt: "OUTP ON,(@1)",
		},
		{
			name: "E36441A channel 3", model: "E36441A", index: 2,
			wantSet:   "VOLT %.4f,(@3)",
			wantGet:   "VOLT? (@3)",
			wantMeas:  "MEAS:VOLT? (@3)",
			wantOutSt: "OUTP ON,(@3)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := channelForModel(t, &ivitest.Mock{}, tt.model, tt.index)

			if got := ch.setCmd("VOLT %.4f"); got != tt.wantSet {
				t.Errorf("setCmd() = %q, want %q", got, tt.wantSet)
			}
			if got := ch.getCmd("VOLT?"); got != tt.wantGet {
				t.Errorf("getCmd() = %q, want %q", got, tt.wantGet)
			}
			if got := ch.measCmd("MEAS:VOLT?"); got != tt.wantMeas {
				t.Errorf("measCmd() = %q, want %q", got, tt.wantMeas)
			}
			if got := ch.outputSetCmd("OUTP ON"); got != tt.wantOutSt {
				t.Errorf("outputSetCmd() = %q, want %q", got, tt.wantOutSt)
			}
		})
	}
}

// exerciseChannel calls every Channel method that reaches the transport.
// Errors are ignored on purpose: a model that lacks a capability returns a
// sentinel error without sending anything, which is exactly the behavior the
// sweep should tolerate. Only the strings that do reach the wire matter here.
func exerciseChannel(ch *Channel) {
	_, _ = ch.CurrentLimit()
	_ = ch.SetCurrentLimit(1.2)
	_, _ = ch.CurrentLimitBehavior()
	_ = ch.SetCurrentLimitBehavior(dcpwr.CurrentRegulate)
	_ = ch.SetCurrentLimitBehavior(dcpwr.CurrentTrip)
	_, _ = ch.OutputEnabled()
	_ = ch.SetOutputEnabled(true)
	_ = ch.DisableOutput()
	_ = ch.EnableOutput()
	_, _ = ch.OVPEnabled()
	_ = ch.SetOVPEnabled(true)
	_ = ch.DisableOVP()
	_ = ch.EnableOVP()
	_, _ = ch.OVPLimit()
	_ = ch.SetOVPLimit(6.0)
	_, _ = ch.VoltageLevel()
	_ = ch.SetVoltageLevel(4.1)
	_ = ch.ConfigureCurrentLimit(dcpwr.CurrentRegulate, 1.2)
	_ = ch.ConfigureOVP(true, 6.0)
	_, _ = ch.QueryOutputState(dcpwr.ConstantVoltage)
	_ = ch.ResetOutputProtection()
	_, _ = ch.Measure(dcpwr.VoltageMeasurement)
	_, _ = ch.Measure(dcpwr.CurrentMeasurement)
	_, _ = ch.MeasureVoltage()
	_, _ = ch.MeasureCurrent()
	_, _ = ch.TriggerSource()
	_ = ch.SetTriggerSource(dcpwr.TriggerSourceImmediate)
	_, _ = ch.TriggeredCurrentLimit()
	_ = ch.SetTriggeredCurrentLimit(1.2)
	_, _ = ch.TriggeredVoltageLevel()
	_ = ch.SetTriggeredVoltageLevel(4.1)
}

// TestAllModels_EmitValidSCPI drives every supported model, on every one of
// its channels, through every channel-scoped method, and checks that each
// string reaching the transport is well-formed SCPI.
//
// This is the sweep that catches what per-model tests miss. A model nobody
// wrote a test for still gets every command it can emit validated, so a
// channel name that is not a legal SCPI token, a forgotten command set, or a
// format string given the wrong arguments all fail here without hardware.
//
// It cannot confirm that a well-formed command is the right command. That an
// E3646A wants OUT1 rather than OUTPUT1 is a fact about the instrument, not
// about the string, and belongs to the integration tests.
func TestAllModels_EmitValidSCPI(t *testing.T) {
	for _, supply := range supportedSupplies {
		t.Run(supply.model, func(t *testing.T) {
			for i, name := range supply.channels {
				t.Run(name, func(t *testing.T) {
					strict := &ivitest.Strict{
						Mock: ivitest.Mock{QueryResp: "1"},
					}
					ch := channelForModel(t, strict, supply.model, i)

					exerciseChannel(ch)

					strict.Check(t)

					if len(strict.CommandsSent)+len(strict.QueriesSent) == 0 {
						t.Error("no SCPI reached the transport")
					}
				})
			}
		})
	}
}

// TestSupportedSupplies_ChannelNamesAreSCPITokens checks the table invariant
// behind the sweep above: for the instSelect family a channel name is spliced
// into "INST <name>; ", so it has to be a bare SCPI token. A descriptive name
// such as "Output 1" cannot work no matter which instrument receives it.
//
// The singleOutput family is exempt because its names never reach the wire.
func TestSupportedSupplies_ChannelNamesAreSCPITokens(t *testing.T) {
	token := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)

	for _, supply := range supportedSupplies {
		// Only the instSelect family sends its channel names. singleOutput
		// has no way to name an output, and channelList names one by number,
		// so both may use descriptive names.
		if supply.family != instSelect {
			continue
		}

		for _, name := range supply.channels {
			if !token.MatchString(name) {
				t.Errorf(
					"%s: channel name %q is not a legal SCPI token, but the "+
						"instSelect family sends it as the INSTrument "+
						"parameter", supply.model, name,
				)
			}
		}
	}
}
