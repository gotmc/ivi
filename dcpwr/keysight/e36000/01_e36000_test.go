// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e36000

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/gotmc/ivi"
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

// TestChannelPrefix covers the difference that motivates [scpiFamily]. The
// E3631A selects the output it programs with INSTrument[:SELect], while the
// E36100B series has no INSTrument subsystem at all, so the same prefix would
// produce error -113 (Undefined header) on an E36102B.
func TestChannelPrefix(t *testing.T) {
	tests := []struct {
		name      string
		model     string
		index     int
		wantPre   string
		wantParam string
	}{
		{"E3631A 6V", "E3631A", 0, "INST P6V; ", " P6V"},
		{"E3631A -25V", "E3631A", 2, "INST N25V; ", " N25V"},
		{"E36102B", "E36102B", 0, "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := channelForModel(t, &ivitest.Mock{}, tt.model, tt.index)
			if got := ch.prefix(); got != tt.wantPre {
				t.Errorf("prefix() = %q, want %q", got, tt.wantPre)
			}
			if got := ch.measureParameter(); got != tt.wantParam {
				t.Errorf("measureParameter() = %q, want %q", got, tt.wantParam)
			}
		})
	}
}
