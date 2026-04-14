// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key35670

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/dsa"
	"github.com/gotmc/query"
)

// --- Frequency ---

func (d *Key35670) FrequencyStart() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:STAR?")
}

func (d *Key35670) SetFrequencyStart(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:STAR %f", freq)
}

func (d *Key35670) FrequencyStop() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:STOP?")
}

func (d *Key35670) SetFrequencyStop(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:STOP %f", freq)
}

func (d *Key35670) FrequencySpan() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:SPAN?")
}

func (d *Key35670) SetFrequencySpan(span float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:SPAN %f", span)
}

func (d *Key35670) FrequencyCenter() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:CENT?")
}

func (d *Key35670) SetFrequencyCenter(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:CENT %f", freq)
}

func (d *Key35670) ConfigureFrequencyStartStop(
	startFreq, stopFreq float64,
) error {
	if err := d.SetFrequencyStart(startFreq); err != nil {
		return err
	}

	return d.SetFrequencyStop(stopFreq)
}

func (d *Key35670) ConfigureFrequencyCenterSpan(
	centerFreq, span float64,
) error {
	if err := d.SetFrequencyCenter(centerFreq); err != nil {
		return err
	}

	return d.SetFrequencySpan(span)
}

// --- Resolution (spectral lines) ---

func (d *Key35670) SpectralLines() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "SENS:FREQ:RES?")
}

func (d *Key35670) SetSpectralLines(lines int) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:RES %d", lines)
}

// --- Window function ---

var windowTypeToSCPI = map[dsa.WindowType]string{
	dsa.WindowHanning:     "HANN",
	dsa.WindowUniform:     "UNIF",
	dsa.WindowFlatTop:     "FLAT",
	dsa.WindowForce:       "FORC",
	dsa.WindowExponential: "EXP",
}

var scpiToWindowType = map[string]dsa.WindowType{
	"HANN": dsa.WindowHanning,
	"UNIF": dsa.WindowUniform,
	"FLAT": dsa.WindowFlatTop,
	"FORC": dsa.WindowForce,
	"EXP":  dsa.WindowExponential,
}

func (d *Key35670) WindowType() (dsa.WindowType, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "SENS:WIND:TYPE?")
	if err != nil {
		return 0, fmt.Errorf("WindowType: %w", err)
	}

	wt, err := ivi.ReverseLookup(scpiToWindowType, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("WindowType: %w", err)
	}

	return wt, nil
}

func (d *Key35670) SetWindowType(window dsa.WindowType) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(windowTypeToSCPI, window)
	if err != nil {
		return fmt.Errorf("SetWindowType: %w", err)
	}

	return d.inst.Command(ctx, "SENS:WIND:TYPE %s", cmd)
}

// --- Averaging ---

func (d *Key35670) AveragingEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "SENS:AVER?")
}

func (d *Key35670) SetAveragingEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "SENS:AVER ON")
	}

	return d.inst.Command(ctx, "SENS:AVER OFF")
}

func (d *Key35670) AveragingCount() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "SENS:AVER:COUN?")
}

func (d *Key35670) SetAveragingCount(count int) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:AVER:COUN %d", count)
}

var averagingTypeToSCPI = map[dsa.AveragingType]string{
	dsa.AveragingRMS:             "RMS",
	dsa.AveragingRMSExponential:  "RMSE",
	dsa.AveragingTime:            "TIME",
	dsa.AveragingTimeExponential: "TIMEE",
	dsa.AveragingPeakHold:        "PEAK",
}

var scpiToAveragingType = map[string]dsa.AveragingType{
	"RMS":   dsa.AveragingRMS,
	"RMSE":  dsa.AveragingRMSExponential,
	"TIME":  dsa.AveragingTime,
	"TIMEE": dsa.AveragingTimeExponential,
	"PEAK":  dsa.AveragingPeakHold,
}

func (d *Key35670) AveragingType() (dsa.AveragingType, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "SENS:AVER:TYPE?")
	if err != nil {
		return 0, fmt.Errorf("AveragingType: %w", err)
	}

	at, err := ivi.ReverseLookup(scpiToAveragingType, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("AveragingType: %w", err)
	}

	return at, nil
}

func (d *Key35670) SetAveragingType(avgType dsa.AveragingType) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(averagingTypeToSCPI, avgType)
	if err != nil {
		return fmt.Errorf("SetAveragingType: %w", err)
	}

	return d.inst.Command(ctx, "SENS:AVER:TYPE %s", cmd)
}

// --- Input range ---

func (d *Key35670) InputRange(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "INP%d:RANG?", channel+1)
}

func (d *Key35670) SetInputRange(channel int, rangeDBVrms float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INP%d:RANG %f", channel+1, rangeDBVrms)
}

func (d *Key35670) InputAutoRange(channel int) (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Boolf(ctx, d.inst, "INP%d:RANG:AUTO?", channel+1)
}

func (d *Key35670) SetInputAutoRange(channel int, auto bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if auto {
		return d.inst.Command(ctx, "INP%d:RANG:AUTO ON", channel+1)
	}

	return d.inst.Command(ctx, "INP%d:RANG:AUTO OFF", channel+1)
}

var inputCouplingToSCPI = map[dsa.InputCoupling]string{
	dsa.InputCouplingAC: "AC",
	dsa.InputCouplingDC: "DC",
}

var scpiToInputCoupling = map[string]dsa.InputCoupling{
	"AC": dsa.InputCouplingAC,
	"DC": dsa.InputCouplingDC,
}

func (d *Key35670) InputCoupling(channel int) (dsa.InputCoupling, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "INP%d:COUP?", channel+1)
	if err != nil {
		return 0, fmt.Errorf("InputCoupling: %w", err)
	}

	coupling, err := ivi.ReverseLookup(
		scpiToInputCoupling, strings.TrimSpace(s),
	)
	if err != nil {
		return 0, fmt.Errorf("InputCoupling: %w", err)
	}

	return coupling, nil
}

func (d *Key35670) SetInputCoupling(
	channel int, coupling dsa.InputCoupling,
) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(inputCouplingToSCPI, coupling)
	if err != nil {
		return fmt.Errorf("SetInputCoupling: %w", err)
	}

	return d.inst.Command(ctx, "INP%d:COUP %s", channel+1, cmd)
}

// --- Measurement mode ---

var measurementModeToSCPI = map[dsa.MeasurementMode]string{
	dsa.MeasurementModeFFT:         "FFT",
	dsa.MeasurementModeCorrelation: "CORR",
	dsa.MeasurementModeHistogram:   "HIST",
	dsa.MeasurementModeOctave:      "OCT",
	dsa.MeasurementModeOrder:       "ORD",
	dsa.MeasurementModeSweptSine:   "SWEP",
}

var scpiToMeasurementMode = map[string]dsa.MeasurementMode{
	"FFT":  dsa.MeasurementModeFFT,
	"CORR": dsa.MeasurementModeCorrelation,
	"HIST": dsa.MeasurementModeHistogram,
	"OCT":  dsa.MeasurementModeOctave,
	"ORD":  dsa.MeasurementModeOrder,
	"SWEP": dsa.MeasurementModeSweptSine,
}

func (d *Key35670) MeasurementMode() (dsa.MeasurementMode, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "INST:SEL?")
	if err != nil {
		return 0, fmt.Errorf("MeasurementMode: %w", err)
	}

	mode, err := ivi.ReverseLookup(
		scpiToMeasurementMode, strings.TrimSpace(s),
	)
	if err != nil {
		return 0, fmt.Errorf("MeasurementMode: %w", err)
	}

	return mode, nil
}

func (d *Key35670) SetMeasurementMode(mode dsa.MeasurementMode) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(measurementModeToSCPI, mode)
	if err != nil {
		return fmt.Errorf("SetMeasurementMode: %w", err)
	}

	return d.inst.Command(ctx, "INST:SEL %s", cmd)
}

func (d *Key35670) ChannelCount() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "INST:NCHA?")
}

func (d *Key35670) SetChannelCount(count int) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INST:NCHA %d", count)
}

// --- Acquisition control ---

func (d *Key35670) Abort() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "ABOR")
}

func (d *Key35670) Initiate() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INIT")
}

func (d *Key35670) SweepModeContinuous() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "INIT:CONT?")
}

func (d *Key35670) SetSweepModeContinuous(continuous bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if continuous {
		return d.inst.Command(ctx, "INIT:CONT ON")
	}

	return d.inst.Command(ctx, "INIT:CONT OFF")
}

// --- Trace data ---

// FetchYTrace returns the trace data as a slice of float64 values.
func (d *Key35670) FetchYTrace(traceName string) ([]float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.Stringf(ctx, d.inst, "CALC%s:DATA?", traceName)
	if err != nil {
		return nil, fmt.Errorf("FetchYTrace: %w", err)
	}

	return parseCSVFloat64(strings.TrimSpace(s))
}

// ReadYTrace initiates a measurement, waits for completion, and returns the
// trace data.
func (d *Key35670) ReadYTrace(
	traceName string, maxTime time.Duration,
) ([]float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	// Use a longer timeout for the measurement if needed.
	if maxTime > d.timeout {
		var longCancel context.CancelFunc
		ctx, longCancel = context.WithTimeout(context.Background(), maxTime)
		defer longCancel()
	}

	if err := d.inst.Command(ctx, "INIT:CONT OFF"); err != nil {
		return nil, err
	}

	if err := d.inst.Command(ctx, "INIT"); err != nil {
		return nil, err
	}

	// Wait for operation complete.
	if _, err := query.String(ctx, d.inst, "*OPC?"); err != nil {
		return nil, fmt.Errorf(
			"ReadYTrace: waiting for measurement complete: %w", err,
		)
	}

	return d.FetchYTrace(traceName)
}

func parseCSVFloat64(s string) ([]float64, error) {
	parts := strings.Split(s, ",")
	result := make([]float64, 0, len(parts))

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		v, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing trace value %q: %w", p, err)
		}

		result = append(result, v)
	}

	return result, nil
}
