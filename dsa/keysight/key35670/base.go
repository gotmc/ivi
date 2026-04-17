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

func (d *Driver) FrequencyStart() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:STAR?")
}

func (d *Driver) SetFrequencyStart(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:STAR %f", freq)
}

func (d *Driver) FrequencyStop() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:STOP?")
}

func (d *Driver) SetFrequencyStop(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:STOP %f", freq)
}

func (d *Driver) FrequencySpan() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:SPAN?")
}

func (d *Driver) SetFrequencySpan(span float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:SPAN %f", span)
}

func (d *Driver) FrequencyCenter() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "SENS:FREQ:CENT?")
}

func (d *Driver) SetFrequencyCenter(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "SENS:FREQ:CENT %f", freq)
}

func (d *Driver) ConfigureFrequencyStartStop(
	startFreq, stopFreq float64,
) error {
	if err := d.SetFrequencyStart(startFreq); err != nil {
		return err
	}

	return d.SetFrequencyStop(stopFreq)
}

func (d *Driver) ConfigureFrequencyCenterSpan(
	centerFreq, span float64,
) error {
	if err := d.SetFrequencyCenter(centerFreq); err != nil {
		return err
	}

	return d.SetFrequencySpan(span)
}

// --- Resolution (spectral lines) ---

func (d *Driver) SpectralLines() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "SENS:FREQ:RES?")
}

func (d *Driver) SetSpectralLines(lines int) error {
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

func (d *Driver) WindowType() (dsa.WindowType, error) {
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

func (d *Driver) SetWindowType(window dsa.WindowType) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(windowTypeToSCPI, window)
	if err != nil {
		return fmt.Errorf("SetWindowType: %w", err)
	}

	return d.inst.Command(ctx, "SENS:WIND:TYPE %s", cmd)
}

// --- Averaging ---

func (d *Driver) AveragingEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "SENS:AVER?")
}

func (d *Driver) SetAveragingEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "SENS:AVER ON")
	}

	return d.inst.Command(ctx, "SENS:AVER OFF")
}

func (d *Driver) AveragingCount() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "SENS:AVER:COUN?")
}

func (d *Driver) SetAveragingCount(count int) error {
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

func (d *Driver) AveragingType() (dsa.AveragingType, error) {
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

func (d *Driver) SetAveragingType(avgType dsa.AveragingType) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(averagingTypeToSCPI, avgType)
	if err != nil {
		return fmt.Errorf("SetAveragingType: %w", err)
	}

	return d.inst.Command(ctx, "SENS:AVER:TYPE %s", cmd)
}

// --- Input range ---

func (d *Driver) InputRange(channel int) (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64f(ctx, d.inst, "INP%d:RANG?", channel+1)
}

func (d *Driver) SetInputRange(channel int, rangeDBVrms float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INP%d:RANG %f", channel+1, rangeDBVrms)
}

func (d *Driver) InputAutoRange(channel int) (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Boolf(ctx, d.inst, "INP%d:RANG:AUTO?", channel+1)
}

func (d *Driver) SetInputAutoRange(channel int, auto bool) error {
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

func (d *Driver) InputCoupling(channel int) (dsa.InputCoupling, error) {
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

func (d *Driver) SetInputCoupling(
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

func (d *Driver) MeasurementMode() (dsa.MeasurementMode, error) {
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

func (d *Driver) SetMeasurementMode(mode dsa.MeasurementMode) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(measurementModeToSCPI, mode)
	if err != nil {
		return fmt.Errorf("SetMeasurementMode: %w", err)
	}

	return d.inst.Command(ctx, "INST:SEL %s", cmd)
}

func (d *Driver) ChannelCount() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Int(ctx, d.inst, "INST:NCHA?")
}

func (d *Driver) SetChannelCount(count int) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INST:NCHA %d", count)
}

// --- Acquisition control ---

func (d *Driver) Abort() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "ABOR")
}

func (d *Driver) Initiate() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INIT")
}

func (d *Driver) SweepModeContinuous() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "INIT:CONT?")
}

func (d *Driver) SetSweepModeContinuous(continuous bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if continuous {
		return d.inst.Command(ctx, "INIT:CONT ON")
	}

	return d.inst.Command(ctx, "INIT:CONT OFF")
}

// --- Trace data ---

// FetchYTrace returns the trace data as a slice of float64 values.
func (d *Driver) FetchYTrace(traceName string) ([]float64, error) {
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
func (d *Driver) ReadYTrace(
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
