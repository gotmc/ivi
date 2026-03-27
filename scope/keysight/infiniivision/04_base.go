// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
	"github.com/gotmc/query"
)

const (
	oneMeg    = 1.0e6
	fiftyOhms = 50.0
)

// AcquisitionStartTime, also referred to as the Horizontal Time Per Record in
// the IVI specification, queries the length of time from the trigger event to
// the first point in the waveform record. If the value is positive, the first
// point in the waveform record occurs after the trigger event. If the value is
// negative, the first point in the waveform record occurs before the trigger
// event.
//
// AcquisitionStartTime is the getter for the read-write IviScopeBase
// Acquisition Start Time described in Section 4.2.1 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) AcquisitionStartTime(ctx context.Context) (time.Duration, error) {
	// The InfiniiVision 3000 X-series scopes have 10 divisions, and the
	// reference can either be the center, one division from the left, or one
	// division from the right. Therefore, find the current range and reference.
	timebaseInfo, err := query.String(ctx, d.inst, ":TIM?")
	if err != nil {
		return 0, err
	}

	timebase, err := decodeTimebase(timebaseInfo)
	if err != nil {
		return 0, err
	}

	if timebase.mode != "MAIN" && timebase.reference != "CENT" && timebase.position != 0.0 {
		// FIXME: I need to handle the abnormal situations.
		return 0, ivi.ErrValueNotSupported
	}

	// The reference is in the center, so per IVI scope, the acquisition start
	// time is the negative value of seconds from the left (waveform start) to
	// the center.
	return durationFromSeconds(-timebase.rng / 2), nil
}

// SetAcquisitionStartTime sets the length of time from the trigger event to
// the first point in the waveform record. If the value is positive, the first
// point in the waveform record occurs after the trigger event. If the value is
// negative, the first point in the waveform record occurs before the trigger
// event.
//
// SetAcquisitionStartTime is the setter for the read-write IviScopeBase
// Acquisition Start Time described in Section 4.2.1 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) SetAcquisitionStartTime(ctx context.Context, delay time.Duration) error {
	return ivi.ErrFunctionNotSupported
}

// AcquisitionStatus indicates whether an acquisition is in progress, complete,
// or if the status is unknown. Acquisition status is not the same as
// instrument status, and does not necessarily check for instrument errors. To
// make sure that the instrument is checked for errors after getting the
// acquisition status, call the Error Query method. (Note that the end user may
// want to call Error Query at the end of a sequence of other calls which
// include getting the acquisition status, - it does not necessarily need to be
// called immediately.) If the driver cannot determine whether the acquisition
// is complete or not, it returns the Acquisition Status Unknown value.
//
// AcquisitionStatus is the getter for the read-only IviScopeBase Acquisition
// Status described in Section 4.2.2 of the IVI-4.1: IviScope Class
// Specification.
func (d *Driver) AcquisitionStatus(ctx context.Context) (scope.AcquisitionStatus, error) {
	return 0, ivi.ErrNotImplemented
}

// AcquisitionType queries how the oscilloscope acquires data and fills the
// waveform record.
//
// AcquisitionType is the getter for the read-write IviScopeBase Acquisition
// Type described in Section 4.2.3 of the IVI-4.1: IviScope Class
// Specification.
func (d *Driver) AcquisitionType(ctx context.Context) (scope.AcquisitionType, error) {
	s, err := query.String(ctx, d.inst, ":ACQ:TYPE?")
	if err != nil {
		return 0, err
	}

	acType, ok := map[string]scope.AcquisitionType{
		"NORM": scope.NormalAcquisition,
		"AVER": scope.AverageAcquisition,
		"HRES": scope.HighResolutionAcquisition,
		"PEAK": scope.PeakDetectAcquisition,
	}[strings.TrimSpace(s)]
	if !ok {
		return 0, fmt.Errorf("%w: %s", ivi.ErrValueNotSupported, s)
	}

	return acType, nil
}

// SetAcquisitionType specifies how the oscilloscope acquires data and fills
// the waveform record.
//
// SetAcquisitionType is the setter for the read-write IviScopeBase Acquisition
// Type described in Section 4.2.3 of the IVI-4.1: IviScope Class
// Specification.
func (d *Driver) SetAcquisitionType(ctx context.Context, acType scope.AcquisitionType) error {
	var cmd string

	switch acType {
	case scope.NormalAcquisition:
		cmd = "NORM"
	case scope.AverageAcquisition:
		cmd = "AVER"
	case scope.HighResolutionAcquisition:
		cmd = "HRES"
	case scope.PeakDetectAcquisition:
		cmd = "PEAK"
	case scope.EnvelopeAcquisition:
		return ivi.ErrNotImplemented
	default:
		return ivi.ErrNotImplemented
	}

	return d.inst.Command(ctx, ":ACQ:TYPE %s", cmd)
}

// ChnnalCount returns the number of currently available channels. The count
// returned includes any of the supported reserved repeated capability names
// defined in Section 2.3, Repeated Capability Names as well as any custom
// repeated capability identifiers.
//
// ChannelCount is the getter for the read-only IviScopeBase Channel Count
// described in Section 4.2.4 of the IVI-4.1: IviScope Class Specification.
func (d *Driver) ChannelCount() int {
	return len(d.Channels)
}

// AcquisitionMinNumPoints returns the minimum number of points the end-user
// requires in the waveform record for each channel. The instrument driver uses
// the value the end-user specifies to configure the record length that the
// oscilloscope uses for waveform acquisition. If the instrument cannot support
// the requested record length, the driver shall configure the instrument to
// the closest bigger record length. The Horizontal Record Length attribute
// returns the actual record length.
//
// AcquisitionMinNumPoints is the getter for the read-write IviScopeBase
// Horizontal Minimum Number of Points described in Section 4.2.8 of the
// IVI-4.1: IviScope Class Specification.
func (d *Driver) AcquisitionMinNumPoints(ctx context.Context) (int, error) {
	return 0, ivi.ErrNotImplemented
}

// SetAcquisitionMinNumPoints sets the minimum number of points the end-user
// requires in the waveform record for each channel. The instrument driver uses
// the value the end-user specifies to configure the record length that the
// oscilloscope uses for waveform acquisition. If the instrument cannot support
// the requested record length, the driver shall configure the instrument to
// the closest bigger record length. The Horizontal Record Length attribute
// returns the actual record length.
//
// SetAcquisitionMinNumPoints is the setter for the read-write IviScopeBase
// Horizontal Minimum Number of Points described in Section 4.2.8 of the
// IVI-4.1: IviScope Class Specification.
func (d *Driver) SetAcquisitionMinNumPoints(ctx context.Context, numPoints int) error {
	return ivi.ErrNotImplemented
}

// AcquisitionRecordLength queries the actual number of points the oscilloscope
// acquires for each channel. The value is equal to or greater than the minimum
// number of points the end-user specifies with the Horizontal Minimum Number
// of Points attribute.
//
// AcquisitionRecordLength is the getter for the read-only IviScopeBase
// Horizontal Record Length described in Section 4.2.9 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) AcquisitionRecordLength(ctx context.Context) (int, error) {
	return query.Int(ctx, d.inst, ":ACQ:POIN?")
}

// AcquisitionSampleRate returns the effective sample rate of the acquired
// waveform using the current configuration. The units are samples per second.
//
// AcquisitionSampleRate is the getter for the read-only IviScopeBase
// Horizontal Sample Rate described in Section 4.2.10 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) AcquisitionSampleRate(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, ":ACQ:SRAT?")
}

// AcquisitionTimePerRecord queries the length of time that corresponds to the
// record length.
//
// AcquisitionTimePerRecord is the getter for the read-write IviScopeBase
// Horizontal Time Per Record described in Section 4.2.11 of the IVI-4.1:
// IviScope Class Specification.
func (d *Driver) AcquisitionTimePerRecord(ctx context.Context) (time.Duration, error) {
	seconds, err := query.Float64(ctx, d.inst, ":TIM:RANG?")
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ivi.ErrValueNotSupported, err)
	}

	return durationFromSeconds(seconds), nil
}

// SetAcquisitionSampleRate specifies the length of time that corresponds to
// the record length.
//
// SetAcquisitionTimePerRecord is the setter for the read-write IviScopeBase
// Horizontal Time Per Record described in Section 4.2.11 of the IVI-4.1:
// IviScope Class Specification.
func (d *Driver) SetAcquisitionTimePerRecord(
	ctx context.Context,
	timePerRecord time.Duration,
) error {
	return d.inst.Command(ctx, ":TIM:RANG %.4e", timePerRecord.Seconds())
}

// TriggerHoldoff queries the length of time the oscilloscope waits after it
// detects a trigger until the oscilloscope enables the trigger subsystem to
// detect another trigger. For C and COM, the units are seconds. For IVI.NET,
// the units are implicit in the definition of PrecisionTimeSpan. The Trigger
// Holdoff attribute affects instrument operation only when the oscilloscope
// requires multiple acquisitions to build a complete waveform. The
// oscilloscope requires multiple waveform acquisitions when it uses
// equivalent-time sampling or when the Acquisition Type attribute is set to
// Envelope or Average.
//
// TriggerHoldoff is the getter for the read-write IviScopeBase Trigger Holdoff
// described in Section 4.2.18 of the IVI-4.1: IviScope Class Specification.
func (d *Driver) TriggerHoldoff(ctx context.Context) (time.Duration, error) {
	seconds, err := query.Float64(ctx, d.inst, ":TRIG:HOLD?")
	if err != nil {
		return 0, fmt.Errorf("%w: %v", ivi.ErrValueNotSupported, err)
	}

	return durationFromSeconds(seconds), nil
}

func (d *Driver) SetTriggerHoldoff(ctx context.Context, holdoff time.Duration) error {
	const (
		minHoldoff = 40 * time.Nanosecond
		maxHoldoff = 10 * time.Second
	)

	if holdoff < minHoldoff || holdoff > maxHoldoff {
		return fmt.Errorf(
			"%w: holdoff must be between %s and %s, received %s",
			ivi.ErrValueNotSupported,
			minHoldoff,
			maxHoldoff,
			holdoff,
		)
	}

	return d.inst.Command(ctx, ":TRIG:HOLD %e", holdoff.Seconds())
}

func (d *Driver) TriggerLevel(ctx context.Context) (float64, error) {
	return query.Float64(ctx, d.inst, ":TRIG:EDGE:LEV?")
}

func (d *Driver) SetTriggerLevel(ctx context.Context, level float64) error {
	return d.inst.Command(ctx, ":TRIG:EDGE:LEV %e", level)
}

func (d *Driver) TriggerSlope(ctx context.Context) (scope.TriggerSlope, error) {
	return 0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerSlope(ctx context.Context, slope scope.TriggerSlope) error {
	// Need to determine if in TV Trigger mode, because that has a different
	// command.
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerSource(ctx context.Context) (scope.TriggerSource, error) {
	return 0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerSource(ctx context.Context, source scope.TriggerSource) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerType(ctx context.Context) (scope.TriggerType, error) {
	mode, err := query.String(ctx, d.inst, ":TRIG:MODE?")
	if err != nil {
		return 0, ivi.ErrValueNotSupported
	}

	switch strings.TrimSpace(mode) {
	case "EDGE":
		return scope.EdgeTrigger, nil
	case "GLIT":
		return scope.GlitchTrigger, nil
	case "PATT":
		return scope.WidthTrigger, nil
	case "TV":
		return scope.TVTrigger, nil
	case "RUNT":
		return scope.RuntTrigger, nil
	default:
		return 0, ivi.ErrValueNotSupported
	}
}

func (d *Driver) SetTriggerType(ctx context.Context, triggerType scope.TriggerType) error {
	switch triggerType {
	case scope.EdgeTrigger:
		return d.inst.Command(ctx, ":TRIG:MODE EDGE")
	case scope.WidthTrigger:
		return d.inst.Command(ctx, ":TRIG:MODE PATT")
	case scope.RuntTrigger:
		return d.inst.Command(ctx, ":TRIG:MODE RUNT")
	case scope.GlitchTrigger:
		return d.inst.Command(ctx, ":TRIG:MODE GLIT")
	case scope.TVTrigger:
		return d.inst.Command(ctx, ":TRIG:MODE TV")
	case scope.ImmediateTrigger:
		return d.inst.Command(ctx, ":TRIG:FORC")
	case scope.ACLineTrigger:
		return fmt.Errorf("%s not supported: %w", scope.TVTrigger, ivi.ErrValueNotSupported)
	default:
		return fmt.Errorf("%s not supported: %w", triggerType, ivi.ErrValueNotSupported)
	}
}

func (d *Driver) AbortMeasurement(ctx context.Context) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) ConfigureAcquisitionRecord(
	ctx context.Context,
	timePerRecord time.Duration,
	minNumPoints int,
	acquisitionStartTime time.Duration,
) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) CreateWaveform(ctx context.Context, numSamples int) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) ConfigureEdgeTrigger(
	ctx context.Context,
	triggerType scope.TriggerType,
	level float64,
	slope scope.TriggerSlope,
) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) ConfigureTrigger(
	ctx context.Context,
	triggerType scope.TriggerType,
	holdoff time.Duration,
) error {
	if err := d.SetTriggerType(ctx, triggerType); err != nil {
		return err
	}

	return d.SetTriggerHoldoff(ctx, holdoff)
}

func (d *Driver) InitiateMeasurement(ctx context.Context) error {
	return ivi.ErrNotImplemented
}

// ChannelEnabled queries whether or not the oscilloscope acquires a waveform for
// the channel.
//
// ChannelEnabled is the getter for the read-write IviScopeBase Attribute
// Channel Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) ChannelEnabled(ctx context.Context) (bool, error) {
	return query.Boolf(ctx, ch.inst, ":CHAN%d:DISP?", ch.num)
}

// SetChannelEnabled sets the channel to either acquire (enabled) or not
// acquire (disabled) a waveform.
//
// Enabled is the setter for the read-write IviScopeBase Attribute Channel
// Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetChannelEnabled(ctx context.Context, b bool) error {
	cmd := "1"
	if !b {
		cmd = "0"
	}

	return ch.inst.Command(ctx, ":CHAN%d:DISP %s", ch.num, cmd)
}

// Name returns the name of the channel.
//
//	Name is the getter for the read-only IviScopeBase Attribute Channel Name
//	described in Section 4.2.7 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) Name() string {
	return fmt.Sprintf("CH%d", ch.num)
}

// InputImpedance queries the input impedance for the channel in Ohms. Legal
// values are 50.0 and 1,000,000.0.

// InputImpedance is the getter for the read-write IviScopeBase Attribute Input
// Impedance described in Section 4.2.12 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) InputImpedance(ctx context.Context) (float64, error) {
	imped, err := query.Stringf(ctx, ch.inst, ":CHAN%d:IMP?", ch.num)
	if err != nil {
		return 0.0, err
	}

	switch imped {
	case "ONEM":
		return oneMeg, nil
	case "FIFT":
		return fiftyOhms, nil
	default:
		return 0.0, ivi.ErrValueNotSupported
	}
}

// SetInputImpedance sets the input impedance for the channel in Ohms. Legal
// values are 50.0 and 1,000,000.0.

// SetInputImpedance is the setter for the read-write IviScopeBase Attribute
// Input Impedance described in Section 4.2.12 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetInputImpedance(ctx context.Context, impedance float64) error {
	switch impedance {
	case fiftyOhms:
		return ch.inst.Command(ctx, ":CHAN%d:IMP FIFT", ch.num)
	case oneMeg:
		return ch.inst.Command(ctx, ":CHAN%d:IMP ONEM", ch.num)
	default:
		return ivi.ErrValueNotSupported
	}
}

// MaxInputFrequency queries the Specifies the maximum frequency for the input
// signal you want the instrument to accommodate without attenuating it by more
// than 3dB. If the bandwidth limit frequency of the instrument is greater than
// this maximum frequency, the driver enables the bandwidth limit. This
// attenuates the input signal by at least 3dB at frequencies greater than the
// bandwidth limit.
//
// MaxInputImpedance is the getter for the read-write IviScopeBase Attribute
// Maximum Input Impedance described in Section 4.2.13 of the IVI-4.1: IviScope
// Class Specification.
func (ch *Channel) MaxInputFrequency(ctx context.Context) (float64, error) {
	return 0.0, ivi.ErrNotImplemented
}

func (ch *Channel) SetMaxInputFrequency(ctx context.Context, _ float64) error {
	return ivi.ErrNotImplemented
}

// ProbeAttenuation queries the scaling factor by which the probe the end-user
// attaches to the channel attenuates the input. The probe attenuation factor
// may be 0.001 to 10,000.0.

// For example, for a 10:1 probe, ProbeAttenuation would return 10.0. Note that
// if the probe is changed to one with a different attenuation, and this
// attribute is not set, the amplitude calculations will be incorrect.
//
// ProbeAttenuation is the getter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) ProbeAttenuation(ctx context.Context) (float64, error) {
	return query.Float64f(ctx, ch.inst, ":CHAN%d:PROBE?", ch.num)
}

// ProbeAttenuation sets the scaling factor by which the probe the end-user
// attaches to the channel attenuates the input and disables auto probe
// attenuation. The probe attenuation factor may be 0.001 to 10,000.0.
//
// For example, for a 10:1 probe, the end-user sets this attribute to 10.0.
// Note that if the probe is changed to one with a different attenuation, and
// this attribute is not set, the amplitude calculations will be incorrect.
//
// ProbeAttenuation is the getter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetProbeAttenuation(ctx context.Context, atten float64) error {
	if atten < 0.001 || atten > 10000.0 {
		return ivi.ErrValueNotSupported
	}
	return ch.inst.Command(ctx, ":CHAN%d:PROBE %E", ch.num, atten)
}

// ProbeAttenuationAuto always return false with no error since auto probe
// attenuation is not supported.

// ProbeAttenuationAuto is the getter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) ProbeAttenuationAuto(ctx context.Context) (bool, error) {
	return false, nil
}

// SetProbeAttenuationAuto if enabled will return an error since auto probe
// attenuation is not supported.

// SetProbeAttenuationAuto is the setter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetProbeAttenuationAuto(ctx context.Context, b bool) error {
	if b {
		return ivi.ErrValueNotSupported
	}

	return nil
}

func (ch *Channel) TriggerCoupling(ctx context.Context) (scope.TriggerCoupling, error) {
	return 0, ivi.ErrNotImplemented
}

func (ch *Channel) SetTriggerCoupling(ctx context.Context, coupling scope.TriggerCoupling) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) VerticalCoupling(ctx context.Context) (scope.VerticalCoupling, error) {
	coupling, err := query.Stringf(ctx, ch.inst, ":CHAN%d:COUP?", ch.num)
	if err != nil {
		return 0, err
	}

	switch coupling {
	case "AC":
		return scope.ACVerticalCoupling, nil
	case "DC":
		return scope.DCVerticalCoupling, nil
	default:
		return 0, ivi.ErrValueNotSupported
	}
}

func (ch *Channel) SetVerticalCoupling(ctx context.Context, coupling scope.VerticalCoupling) error {
	switch coupling {
	case scope.ACVerticalCoupling:
		return ch.inst.Command(ctx, ":CHAN%d:COUP AC", ch.num)
	case scope.DCVerticalCoupling:
		return ch.inst.Command(ctx, ":CHAN%d:COUP DC", ch.num)
	case scope.GndVerticalCoupling:
		return ivi.ErrValueNotSupported
	default:
		return ivi.ErrValueNotSupported
	}
}

// VerticalOffset queries the location of the center of the range that the
// Vertical Range attribute specifies. The value is with respect to ground and
// is in volts. For example, to acquire a sine wave that spans between on 0.0
// and 10.0 volts, set this attribute to 5.0 volts.
//
// VerticalOffset is the getter for the read-write IviScopeBase Vertical Offset
// described in Section 4.2.24 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) VerticalOffset(ctx context.Context) (float64, error) {
	return query.Float64f(ctx, ch.inst, ":CHAN%d:OFFS?", ch.num)
}

// SetVerticalOffset sets the location of the center of the range that the
// Vertical Range attribute specifies. The value is with respect to ground and
// is in volts. For example, to acquire a sine wave that spans between on 0.0
// and 10.0 volts, set this attribute to 5.0 volts.
//
// SetVerticalOffset is the setter for the read-write IviScopeBase Vertical
// Offset described in Section 4.2.24 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetVerticalOffset(ctx context.Context, offset float64) error {
	return ch.inst.Command(ctx, ":CHAN%d:OFFS %E", ch.num, offset)
}

// VerticalRange queries the absolute value of the full-scale input range for a
// channel. The units are volts. For example, to acquire a sine wave that spans
// between -5.0 and 5.0 volts, set the Vertical Range attribute to 10.0 volts.
//
// VerticalRange is the getter for the read-write IviScopeBase Vertical Range
// described in Section 4.2.25 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) VerticalRange(ctx context.Context) (float64, error) {
	return query.Float64f(ctx, ch.inst, ":CHAN%d:RANG?", ch.num)
}

// SetVerticalRange sets the absolute value of the full-scale input range for a
// channel. The units are volts with valid ranges from 0.008 to 40.0. For
// example, to acquire a sine wave that spans between -5.0 and 5.0 volts, set
// the Vertical Range attribute to 10.0 volts.
//
// SetVerticalRange is the setter for the read-write IviScopeBase Vertical
// Range described in Section 4.2.25 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetVerticalRange(ctx context.Context, rng float64) error {
	if rng < 0.008 || rng > 40.0 {
		return ivi.ErrValueNotSupported
	}

	return ch.inst.Command(ctx, ":CHAN%d:RANG %E", ch.num, rng)
}

func (ch *Channel) Configure(
	ctx context.Context,
	rng float64,
	offset float64,
	coupling scope.VerticalCoupling,
	autoProbeAttenuation bool,
	probeAttenuation float64,
	enabled bool,
) error {
	if err := ch.SetVerticalRange(ctx, rng); err != nil {
		return err
	}

	if err := ch.SetVerticalOffset(ctx, offset); err != nil {
		return err
	}

	if err := ch.SetVerticalCoupling(ctx, coupling); err != nil {
		return err
	}

	if !autoProbeAttenuation {
		if err := ch.SetProbeAttenuation(ctx, probeAttenuation); err != nil {
			return err
		}
	}

	return ch.SetChannelEnabled(ctx, enabled)
}

func (ch *Channel) ConfigureCharacteristics(
	ctx context.Context,
	inputImepdance, inputFreqMax float64,
) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) FetchWaveform(ctx context.Context, waveform *ivi.Waveform) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) ReadWaveform(
	ctx context.Context,
	maximumTime time.Duration,
	waveform *ivi.Waveform,
) error {
	return ivi.ErrNotImplemented
}

func durationFromSeconds(seconds float64) time.Duration {
	return time.Duration(seconds * float64(time.Second))
}

type timebase struct {
	mode      string
	reference string
	rng       float64
	position  float64
}

func decodeTimebase(s string) (timebase, error) {
	parts := strings.Split(s, ";")
	if len(parts) != 4 {
		return timebase{}, fmt.Errorf("should have received four responses but got %d", len(parts))
	}

	mode := strings.TrimPrefix(parts[0], ":TIM:MODE ")

	ref := strings.TrimPrefix(parts[1], "REF ")

	rngString := strings.TrimPrefix(parts[2], "MAIN:RANG ")
	rng, err := strconv.ParseFloat(strings.TrimSpace(rngString), 64)
	if err != nil {
		return timebase{}, err
	}

	posString := strings.TrimPrefix(parts[3], "POS ")
	pos, err := strconv.ParseFloat(strings.TrimSpace(posString), 64)
	if err != nil {
		return timebase{}, err
	}

	return timebase{
		mode:      mode,
		reference: ref,
		rng:       rng,
		position:  pos,
	}, nil
}
