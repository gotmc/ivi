// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package e4980

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/lcr"
	"github.com/gotmc/query"
)

// --- Measurement Function ---

var msrFuncToSCPI = map[lcr.MeasurementFunction]string{
	lcr.CpD:       "CPD",
	lcr.CpQ:       "CPQ",
	lcr.CpG:       "CPG",
	lcr.CpRp:      "CPRP",
	lcr.CsD:       "CSD",
	lcr.CsQ:       "CSQ",
	lcr.CsRs:      "CSRS",
	lcr.LpD:       "LPD",
	lcr.LpQ:       "LPQ",
	lcr.LpG:       "LPG",
	lcr.LpRp:      "LPRP",
	lcr.LsD:       "LSD",
	lcr.LsQ:       "LSQ",
	lcr.LsRs:      "LSRS",
	lcr.RX:        "RX",
	lcr.ZThetaDeg: "ZTD",
	lcr.ZThetaRad: "ZTR",
	lcr.GB:        "GB",
	lcr.YThetaDeg: "YTD",
	lcr.YThetaRad: "YTR",
}

var scpiToMsrFunc = map[string]lcr.MeasurementFunction{
	"CPD":  lcr.CpD,
	"CPQ":  lcr.CpQ,
	"CPG":  lcr.CpG,
	"CPRP": lcr.CpRp,
	"CSD":  lcr.CsD,
	"CSQ":  lcr.CsQ,
	"CSRS": lcr.CsRs,
	"LPD":  lcr.LpD,
	"LPQ":  lcr.LpQ,
	"LPG":  lcr.LpG,
	"LPRP": lcr.LpRp,
	"LSD":  lcr.LsD,
	"LSQ":  lcr.LsQ,
	"LSRS": lcr.LsRs,
	"RX":   lcr.RX,
	"ZTD":  lcr.ZThetaDeg,
	"ZTR":  lcr.ZThetaRad,
	"GB":   lcr.GB,
	"YTD":  lcr.YThetaDeg,
	"YTR":  lcr.YThetaRad,
}

func (d *Driver) MeasurementFunction() (lcr.MeasurementFunction, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "FUNC:IMP:TYPE?")
	if err != nil {
		return 0, fmt.Errorf("MeasurementFunction: %w", err)
	}

	fcn, err := ivi.ReverseLookup(scpiToMsrFunc, strings.TrimSpace(s))
	if err != nil {
		return 0, fmt.Errorf("MeasurementFunction: %w", err)
	}

	return fcn, nil
}

func (d *Driver) SetMeasurementFunction(fcn lcr.MeasurementFunction) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(msrFuncToSCPI, fcn)
	if err != nil {
		return fmt.Errorf("SetMeasurementFunction: %w", err)
	}

	return d.inst.Command(ctx, "FUNC:IMP:TYPE %s", cmd)
}

// --- Test Signal ---

func (d *Driver) Frequency() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "FREQ:CW?")
}

func (d *Driver) SetFrequency(freq float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "FREQ:CW %f", freq)
}

func (d *Driver) TestVoltageLevel() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "VOLT:LEV?")
}

func (d *Driver) SetTestVoltageLevel(volts float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "VOLT:LEV %f", volts)
}

func (d *Driver) TestCurrentLevel() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "CURR:LEV?")
}

func (d *Driver) SetTestCurrentLevel(amps float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "CURR:LEV %f", amps)
}

// --- Impedance Range ---

func (d *Driver) ImpedanceAutoRange() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "FUNC:IMP:RANG:AUTO?")
}

func (d *Driver) SetImpedanceAutoRange(auto bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if auto {
		return d.inst.Command(ctx, "FUNC:IMP:RANG:AUTO ON")
	}

	return d.inst.Command(ctx, "FUNC:IMP:RANG:AUTO OFF")
}

func (d *Driver) ImpedanceRange() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "FUNC:IMP:RANG?")
}

func (d *Driver) SetImpedanceRange(ohms float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "FUNC:IMP:RANG %f", ohms)
}

// --- Measurement Speed and Averaging ---

var speedToSCPI = map[lcr.MeasurementSpeed]string{
	lcr.MeasurementSpeedShort:  "SHOR",
	lcr.MeasurementSpeedMedium: "MED",
	lcr.MeasurementSpeedLong:   "LONG",
}

var scpiToSpeed = map[string]lcr.MeasurementSpeed{
	"SHOR": lcr.MeasurementSpeedShort,
	"MED":  lcr.MeasurementSpeedMedium,
	"LONG": lcr.MeasurementSpeedLong,
}

func (d *Driver) MeasurementSpeed() (lcr.MeasurementSpeed, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "APER?")
	if err != nil {
		return 0, fmt.Errorf("MeasurementSpeed: %w", err)
	}

	// Response is "SHOR,1" or "MED,10" — speed and averaging count
	// comma-separated.
	parts := strings.SplitN(strings.TrimSpace(s), ",", 2)

	speed, err := ivi.ReverseLookup(scpiToSpeed, parts[0])
	if err != nil {
		return 0, fmt.Errorf("MeasurementSpeed: %w", err)
	}

	return speed, nil
}

func (d *Driver) SetMeasurementSpeed(speed lcr.MeasurementSpeed) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(speedToSCPI, speed)
	if err != nil {
		return fmt.Errorf("SetMeasurementSpeed: %w", err)
	}

	// Preserve the current averaging count.
	count, err := d.AveragingCount()
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, "APER %s,%d", cmd, count)
}

func (d *Driver) AveragingCount() (int, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "APER?")
	if err != nil {
		return 0, fmt.Errorf("AveragingCount: %w", err)
	}

	parts := strings.SplitN(strings.TrimSpace(s), ",", 2)
	if len(parts) < 2 {
		return 1, nil
	}

	count, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, fmt.Errorf("AveragingCount: %w", err)
	}

	return count, nil
}

func (d *Driver) SetAveragingCount(count int) error {
	ctx, cancel := d.newContext()
	defer cancel()

	// Preserve the current speed setting.
	speed, err := d.MeasurementSpeed()
	if err != nil {
		return err
	}

	cmd, err := ivi.LookupSCPI(speedToSCPI, speed)
	if err != nil {
		return err
	}

	return d.inst.Command(ctx, "APER %s,%d", cmd, count)
}

// --- Trigger ---

var triggerSourceToSCPI = map[lcr.TriggerSource]string{
	lcr.TriggerSourceInternal: "INT",
	lcr.TriggerSourceExternal: "EXT",
	lcr.TriggerSourceBus:      "BUS",
	lcr.TriggerSourceHold:     "HOLD",
}

var scpiToTriggerSource = map[string]lcr.TriggerSource{
	"INT":  lcr.TriggerSourceInternal,
	"EXT":  lcr.TriggerSourceExternal,
	"BUS":  lcr.TriggerSourceBus,
	"HOLD": lcr.TriggerSourceHold,
}

func (d *Driver) TriggerSource() (lcr.TriggerSource, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "TRIG:SOUR?")
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	src, err := ivi.ReverseLookup(
		scpiToTriggerSource, strings.TrimSpace(s),
	)
	if err != nil {
		return 0, fmt.Errorf("TriggerSource: %w", err)
	}

	return src, nil
}

func (d *Driver) SetTriggerSource(src lcr.TriggerSource) error {
	ctx, cancel := d.newContext()
	defer cancel()

	cmd, err := ivi.LookupSCPI(triggerSourceToSCPI, src)
	if err != nil {
		return fmt.Errorf("SetTriggerSource: %w", err)
	}

	return d.inst.Command(ctx, "TRIG:SOUR %s", cmd)
}

func (d *Driver) TriggerDelay() (time.Duration, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	sec, err := query.Float64(ctx, d.inst, "TRIG:DEL?")
	if err != nil {
		return 0, fmt.Errorf("TriggerDelay: %w", err)
	}

	return time.Duration(sec * float64(time.Second)), nil
}

func (d *Driver) SetTriggerDelay(delay time.Duration) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TRIG:DEL %f", delay.Seconds())
}

// --- Measurement Control ---

func (d *Driver) Initiate() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "INIT:IMM")
}

func (d *Driver) Trigger() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "TRIG:IMM")
}

func (d *Driver) Abort() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "ABOR")
}

// --- Results ---

// FetchMeasurement returns the primary value, secondary value, and measurement
// status from the most recent measurement. The trigger source should be set to
// BUS and a measurement initiated before calling this method.
func (d *Driver) FetchMeasurement() (float64, float64, lcr.MeasurementStatus, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	s, err := query.String(ctx, d.inst, "FETC:IMP:FORM?")
	if err != nil {
		return 0, 0, lcr.MeasurementStatusNoData,
			fmt.Errorf("FetchMeasurement: %w", err)
	}

	return parseMeasurementResult(strings.TrimSpace(s))
}

func parseMeasurementResult(
	s string,
) (float64, float64, lcr.MeasurementStatus, error) {
	parts := strings.Split(s, ",")
	if len(parts) < 3 {
		return 0, 0, lcr.MeasurementStatusNoData,
			fmt.Errorf("unexpected response format: %q", s)
	}

	primary, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, 0, lcr.MeasurementStatusNoData,
			fmt.Errorf("parsing primary value: %w", err)
	}

	secondary, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, 0, lcr.MeasurementStatusNoData,
			fmt.Errorf("parsing secondary value: %w", err)
	}

	statusVal, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return 0, 0, lcr.MeasurementStatusNoData,
			fmt.Errorf("parsing status: %w", err)
	}

	return primary, secondary, lcr.MeasurementStatus(statusVal), nil
}

// --- DC Bias ---

func (d *Driver) DCBiasEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "BIAS:STAT?")
}

func (d *Driver) SetDCBiasEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "BIAS:STAT ON")
	}

	return d.inst.Command(ctx, "BIAS:STAT OFF")
}

func (d *Driver) DCBiasVoltageLevel() (float64, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Float64(ctx, d.inst, "BIAS:VOLT:LEV?")
}

func (d *Driver) SetDCBiasVoltageLevel(volts float64) error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "BIAS:VOLT:LEV %f", volts)
}

// --- Compensation ---

func (d *Driver) OpenCorrectionEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "CORR:OPEN:STAT?")
}

func (d *Driver) SetOpenCorrectionEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "CORR:OPEN:STAT ON")
	}

	return d.inst.Command(ctx, "CORR:OPEN:STAT OFF")
}

func (d *Driver) ExecuteOpenCorrection() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "CORR:OPEN")
}

func (d *Driver) ShortCorrectionEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "CORR:SHOR:STAT?")
}

func (d *Driver) SetShortCorrectionEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "CORR:SHOR:STAT ON")
	}

	return d.inst.Command(ctx, "CORR:SHOR:STAT OFF")
}

func (d *Driver) ExecuteShortCorrection() error {
	ctx, cancel := d.newContext()
	defer cancel()

	return d.inst.Command(ctx, "CORR:SHOR")
}

func (d *Driver) LoadCorrectionEnabled() (bool, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	return query.Bool(ctx, d.inst, "CORR:LOAD:STAT?")
}

func (d *Driver) SetLoadCorrectionEnabled(enabled bool) error {
	ctx, cancel := d.newContext()
	defer cancel()

	if enabled {
		return d.inst.Command(ctx, "CORR:LOAD:STAT ON")
	}

	return d.inst.Command(ctx, "CORR:LOAD:STAT OFF")
}
