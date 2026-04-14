// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package esa

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/specan"
	"github.com/gotmc/query"
)

// --- Amplitude Units ---

var amplitudeUnitsToSCPI = map[specan.AmplitudeUnits]string{
	specan.AmplitudeUnitsDBM:  "DBM",
	specan.AmplitudeUnitsDBMV: "DBMV",
	specan.AmplitudeUnitsDBUV: "DBUV",
	specan.AmplitudeUnitsVolt: "V",
	specan.AmplitudeUnitsWatt: "W",
}

var scpiToAmplitudeUnits = map[string]specan.AmplitudeUnits{
	"DBM":  specan.AmplitudeUnitsDBM,
	"DBMV": specan.AmplitudeUnitsDBMV,
	"DBUV": specan.AmplitudeUnitsDBUV,
	"DBUA": specan.AmplitudeUnitsDBUV, // treat dBuA as dBuV equivalent
	"V":    specan.AmplitudeUnitsVolt,
	"W":    specan.AmplitudeUnitsWatt,
	"A":    specan.AmplitudeUnitsWatt, // treat Amps as Watt equivalent
}

func (d *Driver) AmplitudeUnits(ctx context.Context) (specan.AmplitudeUnits, error) {
	s, err := query.String(ctx, d.inst, "UNIT:POW?")
	if err != nil {
		return 0, fmt.Errorf("AmplitudeUnits: %w", err)
	}

	units, err := ivi.ReverseLookup(scpiToAmplitudeUnits, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("AmplitudeUnits: %w", err)
	}

	return units, nil
}

func (d *Driver) SetAmplitudeUnits(ctx context.Context, units specan.AmplitudeUnits) error {
	cmd, err := ivi.LookupSCPI(amplitudeUnitsToSCPI, units)
	if err != nil {
		return fmt.Errorf("SetAmplitudeUnits: %w", err)
	}

	return d.inst.Command(ctx, "UNIT:POW %s", cmd)
}

// --- Reference Level ---

func (d *Driver) ReferenceLevel(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "DISP:WIND:TRAC:Y:RLEV?")
}

func (d *Driver) SetReferenceLevel(ctx context.Context, level float64) error {
	return d.inst.Command(ctx, "DISP:WIND:TRAC:Y:RLEV %f", level)
}

func (d *Driver) ReferenceLevelOffset(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "DISP:WIND:TRAC:Y:RLEV:OFFS?")
}

func (d *Driver) SetReferenceLevelOffset(ctx context.Context, offset float64) error {
	return d.inst.Command(ctx, "DISP:WIND:TRAC:Y:RLEV:OFFS %f", offset)
}

// --- Input Impedance ---

func (d *Driver) InputImpedance(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "CORR:IMP:INP:MAGN?")
}

func (d *Driver) SetInputImpedance(ctx context.Context, impedance float64) error {
	return d.inst.Command(ctx, "CORR:IMP:INP:MAGN %f", impedance)
}

// --- Vertical Scale ---

var verticalScaleToSCPI = map[specan.VerticalScale]string{
	specan.VerticalScaleLinear:      "LIN",
	specan.VerticalScaleLogarithmic: "LOG",
}

var scpiToVerticalScale = map[string]specan.VerticalScale{
	"LIN": specan.VerticalScaleLinear,
	"LOG": specan.VerticalScaleLogarithmic,
}

func (d *Driver) VerticalScale(ctx context.Context) (specan.VerticalScale, error) {
	s, err := query.String(ctx, d.inst, "DISP:WIND:TRAC:Y:SPAC?")
	if err != nil {
		return 0, fmt.Errorf("VerticalScale: %w", err)
	}

	scale, err := ivi.ReverseLookup(scpiToVerticalScale, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("VerticalScale: %w", err)
	}

	return scale, nil
}

func (d *Driver) SetVerticalScale(ctx context.Context, scale specan.VerticalScale) error {
	cmd, err := ivi.LookupSCPI(verticalScaleToSCPI, scale)
	if err != nil {
		return fmt.Errorf("SetVerticalScale: %w", err)
	}

	return d.inst.Command(ctx, "DISP:WIND:TRAC:Y:SPAC %s", cmd)
}

// --- Attenuation ---

func (d *Driver) Attenuation(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "POW:ATT?")
}

func (d *Driver) SetAttenuation(ctx context.Context, attenuation float64) error {
	return d.inst.Command(ctx, "POW:ATT %f", attenuation)
}

func (d *Driver) AttenuationAuto(ctx context.Context) (bool, error) {
	return query.Bool(ctx, d.inst, "POW:ATT:AUTO?")
}

func (d *Driver) SetAttenuationAuto(ctx context.Context, auto bool) error {
	if auto {
		return d.inst.Command(ctx, "POW:ATT:AUTO ON")
	}

	return d.inst.Command(ctx, "POW:ATT:AUTO OFF")
}

// --- Frequency ---

func (d *Driver) FrequencyStart(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "FREQ:STAR?")
}

func (d *Driver) SetFrequencyStart(ctx context.Context, freq float64) error {
	return d.inst.Command(ctx, "FREQ:STAR %f", freq)
}

func (d *Driver) FrequencyStop(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "FREQ:STOP?")
}

func (d *Driver) SetFrequencyStop(ctx context.Context, freq float64) error {
	return d.inst.Command(ctx, "FREQ:STOP %f", freq)
}

func (d *Driver) FrequencyOffset(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "DISP:WIND:TRAC:X:OFFS?")
}

func (d *Driver) SetFrequencyOffset(ctx context.Context, offset float64) error {
	return d.inst.Command(ctx, "DISP:WIND:TRAC:X:OFFS %f", offset)
}

// --- Bandwidth ---

func (d *Driver) ResolutionBandwidth(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "BAND?")
}

func (d *Driver) SetResolutionBandwidth(ctx context.Context, bw float64) error {
	return d.inst.Command(ctx, "BAND %f", bw)
}

func (d *Driver) ResolutionBandwidthAuto(ctx context.Context) (bool, error) {
	return query.Bool(ctx, d.inst, "BAND:AUTO?")
}

func (d *Driver) SetResolutionBandwidthAuto(ctx context.Context, auto bool) error {
	if auto {
		return d.inst.Command(ctx, "BAND:AUTO ON")
	}

	return d.inst.Command(ctx, "BAND:AUTO OFF")
}

func (d *Driver) VideoBandwidth(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "BAND:VID?")
}

func (d *Driver) SetVideoBandwidth(ctx context.Context, bw float64) error {
	return d.inst.Command(ctx, "BAND:VID %f", bw)
}

func (d *Driver) VideoBandwidthAuto(ctx context.Context) (bool, error) {
	return query.Bool(ctx, d.inst, "BAND:VID:AUTO?")
}

func (d *Driver) SetVideoBandwidthAuto(ctx context.Context, auto bool) error {
	if auto {
		return d.inst.Command(ctx, "BAND:VID:AUTO ON")
	}

	return d.inst.Command(ctx, "BAND:VID:AUTO OFF")
}

// --- Sweep ---

func (d *Driver) SweepModeContinuous(ctx context.Context) (bool, error) {
	return query.Bool(ctx, d.inst, "INIT:CONT?")
}

func (d *Driver) SetSweepModeContinuous(ctx context.Context, continuous bool) error {
	if continuous {
		return d.inst.Command(ctx, "INIT:CONT ON")
	}

	return d.inst.Command(ctx, "INIT:CONT OFF")
}

func (d *Driver) SweepTime(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, "SWE:TIME?")
}

func (d *Driver) SetSweepTime(ctx context.Context, sweepTime float64) error {
	return d.inst.Command(ctx, "SWE:TIME %f", sweepTime)
}

func (d *Driver) SweepTimeAuto(ctx context.Context) (bool, error) {
	return query.Bool(ctx, d.inst, "SWE:TIME:AUTO?")
}

func (d *Driver) SetSweepTimeAuto(ctx context.Context, auto bool) error {
	if auto {
		return d.inst.Command(ctx, "SWE:TIME:AUTO ON")
	}

	return d.inst.Command(ctx, "SWE:TIME:AUTO OFF")
}

func (d *Driver) NumberOfSweeps(ctx context.Context) (int, error) {
	return query.Int(ctx, d.inst, "AVER:COUN?")
}

func (d *Driver) SetNumberOfSweeps(ctx context.Context, num int) error {
	return d.inst.Command(ctx, "AVER:COUN %d", num)
}

// --- Trace ---

// TraceCount returns the number of traces available. The ESA series supports 3
// traces (TRACE1, TRACE2, TRACE3).
func (d *Driver) TraceCount() int {
	return 3
}

var traceTypeToSCPI = map[specan.TraceType]string{
	specan.TraceTypeClearWrite:   "WRIT",
	specan.TraceTypeMaxHold:      "MAXH",
	specan.TraceTypeMinHold:      "MINH",
	specan.TraceTypeVideoAverage: "AVER",
	specan.TraceTypeView:         "VIEW",
	specan.TraceTypeStore:        "BLAN",
}

var scpiToTraceType = map[string]specan.TraceType{
	"WRIT": specan.TraceTypeClearWrite,
	"MAXH": specan.TraceTypeMaxHold,
	"MINH": specan.TraceTypeMinHold,
	"VIEW": specan.TraceTypeView,
	"BLAN": specan.TraceTypeStore,
}

func (d *Driver) TraceType(ctx context.Context, traceName string) (specan.TraceType, error) {
	s, err := query.Stringf(ctx, d.inst, "TRAC%s:MODE?", traceName)
	if err != nil {
		return 0, fmt.Errorf("TraceType: %w", err)
	}

	tt, err := ivi.ReverseLookup(scpiToTraceType, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("TraceType: %w", err)
	}

	return tt, nil
}

func (d *Driver) SetTraceType(
	ctx context.Context, traceName string, traceType specan.TraceType,
) error {
	cmd, err := ivi.LookupSCPI(traceTypeToSCPI, traceType)
	if err != nil {
		return fmt.Errorf("SetTraceType: %w", err)
	}

	return d.inst.Command(ctx, "TRAC%s:MODE %s", traceName, cmd)
}

var detectorTypeToSCPI = map[specan.DetectorType]string{
	specan.DetectorTypeAutoPeak: "NORM",
	specan.DetectorTypeAverage:  "AVER",
	specan.DetectorTypeMaxPeak:  "POS",
	specan.DetectorTypeMinPeak:  "NEG",
	specan.DetectorTypeSample:   "SAMP",
	specan.DetectorTypeRMS:      "RMS",
}

var scpiToDetectorType = map[string]specan.DetectorType{
	"NORM": specan.DetectorTypeAutoPeak,
	"AVER": specan.DetectorTypeAverage,
	"POS":  specan.DetectorTypeMaxPeak,
	"NEG":  specan.DetectorTypeMinPeak,
	"SAMP": specan.DetectorTypeSample,
	"RMS":  specan.DetectorTypeRMS,
}

func (d *Driver) DetectorType(ctx context.Context, _ string) (specan.DetectorType, error) {
	s, err := query.String(ctx, d.inst, "DET?")
	if err != nil {
		return 0, fmt.Errorf("DetectorType: %w", err)
	}

	dt, err := ivi.ReverseLookup(scpiToDetectorType, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("DetectorType: %w", err)
	}

	return dt, nil
}

func (d *Driver) SetDetectorType(
	ctx context.Context, _ string, detector specan.DetectorType,
) error {
	cmd, err := ivi.LookupSCPI(detectorTypeToSCPI, detector)
	if err != nil {
		return fmt.Errorf("SetDetectorType: %w", err)
	}

	return d.inst.Command(ctx, "DET %s", cmd)
}

func (d *Driver) DetectorTypeAuto(ctx context.Context, _ string) (bool, error) {
	return query.Bool(ctx, d.inst, "DET:AUTO?")
}

func (d *Driver) SetDetectorTypeAuto(ctx context.Context, _ string, auto bool) error {
	if auto {
		return d.inst.Command(ctx, "DET:AUTO ON")
	}

	return d.inst.Command(ctx, "DET:AUTO OFF")
}

// --- Acquisition Control ---

func (d *Driver) Abort(ctx context.Context) error {
	return d.inst.Command(ctx, "ABOR")
}

func (d *Driver) AcquisitionStatus(ctx context.Context) (specan.AcquisitionStatus, error) {
	s, err := query.String(ctx, d.inst, "STAT:OPER:COND?")
	if err != nil {
		return specan.AcquisitionStatusUnknown, fmt.Errorf("AcquisitionStatus: %w", err)
	}

	val, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return specan.AcquisitionStatusUnknown, fmt.Errorf("AcquisitionStatus: %w", err)
	}

	// Bit 3 (value 8) of the operation condition register indicates sweeping.
	if val&8 != 0 {
		return specan.AcquisitionStatusInProgress, nil
	}

	return specan.AcquisitionStatusComplete, nil
}

func (d *Driver) Initiate(ctx context.Context) error {
	return d.inst.Command(ctx, "INIT")
}

// --- Configuration Helpers ---

func (d *Driver) ConfigureFrequencyCenterSpan(ctx context.Context, centerFreq, span float64) error {
	if err := d.inst.Command(ctx, "FREQ:CENT %f", centerFreq); err != nil {
		return err
	}

	return d.inst.Command(ctx, "FREQ:SPAN %f", span)
}

func (d *Driver) ConfigureFrequencyStartStop(
	ctx context.Context, startFreq, stopFreq float64,
) error {
	if err := d.SetFrequencyStart(ctx, startFreq); err != nil {
		return err
	}

	return d.SetFrequencyStop(ctx, stopFreq)
}

func (d *Driver) ConfigureLevel(
	ctx context.Context, units specan.AmplitudeUnits, refLevel float64,
) error {
	if err := d.SetAmplitudeUnits(ctx, units); err != nil {
		return err
	}

	return d.SetReferenceLevel(ctx, refLevel)
}

func (d *Driver) ConfigureSweepCoupling(
	ctx context.Context, resBW, videoBW, sweepTime float64,
) error {
	if err := d.SetResolutionBandwidth(ctx, resBW); err != nil {
		return err
	}

	if err := d.SetVideoBandwidth(ctx, videoBW); err != nil {
		return err
	}

	return d.SetSweepTime(ctx, sweepTime)
}

// --- Trace Data ---

// FetchYTrace returns the trace amplitude data as a slice of float64 values.
// The traceName should be "1", "2", or "3" for the ESA series.
func (d *Driver) FetchYTrace(ctx context.Context, traceName string) ([]float64, error) {
	s, err := query.Stringf(ctx, d.inst, "TRAC:DATA? TRACE%s", traceName)
	if err != nil {
		return nil, fmt.Errorf("FetchYTrace: %w", err)
	}

	return parseCSVFloat64(strings.TrimSpace(s))
}

// ReadYTrace initiates a sweep, waits for it to complete (up to maxTime), and
// returns the trace amplitude data. The traceName should be "1", "2", or "3".
func (d *Driver) ReadYTrace(
	ctx context.Context,
	traceName string,
	maxTime time.Duration,
) ([]float64, error) {
	ctx, cancel := context.WithTimeout(ctx, maxTime)
	defer cancel()

	if err := d.SetSweepModeContinuous(ctx, false); err != nil {
		return nil, err
	}

	if err := d.Initiate(ctx); err != nil {
		return nil, err
	}

	// Wait for operation complete.
	if _, err := query.String(ctx, d.inst, "*OPC?"); err != nil {
		return nil, fmt.Errorf("ReadYTrace: waiting for sweep complete: %w", err)
	}

	return d.FetchYTrace(ctx, traceName)
}

// parseCSVFloat64 parses a comma-separated string of floating point values.
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
