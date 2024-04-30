// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"fmt"
	"time"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
	"github.com/gotmc/query"
)

// AcquisitionStartTime queries the length of time from the trigger event to
// the first point in the waveform record. If the value is positive, the first
// point in the waveform record occurs after the trigger event. If the value is
// negative, the first point in the waveform record occurs before the trigger
// event.
//
// AcquisitionStartTime is the getter for the read-write IviScopeBase
// Acquisition Start Time described in Section 4.2.1 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) AcquisitionStartTime() (time.Duration, error) {
	return 0, ivi.ErrNotImplemented
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
func (d *Driver) SetAcquisitionStartTime(startTime time.Duration) error {
	return ivi.ErrNotImplemented
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
func (d *Driver) AcquisitionStatus() (scope.AcquisitionStatus, error) {
	return 0, ivi.ErrNotImplemented
}

// AcquisitionType queries how the oscilloscope acquires data and fills the
// waveform record.
//
// AcquisitionType is the getter for the read-write IviScopeBase Acquisition
// Type described in Section 4.2.3 of the IVI-4.1: IviScope Class
// Specification.
func (d *Driver) AcquisitionType() (scope.AcquisitionType, error) {
	s, err := query.String(d.inst, ":ACQ:TYPE?")
	if err != nil {
		return 0, err
	}
	acType, ok := map[string]scope.AcquisitionType{
		"NORM": scope.NormalAcquisition,
		"AVER": scope.AverageAcquisition,
		"HRES": scope.HighResolutionAcquisition,
		"PEAK": scope.PeakDetectAcquisition,
	}[s]
	if !ok {
		return 0, ivi.ErrValueNotSupported
	}
	return acType, nil
}

// SetAcquisitionType specifices how the oscilloscope acquires data and fills
// the waveform record.
//
// SetAcquisitionType is the setter for the read-write IviScopeBase Acquisition
// Type described in Section 4.2.3 of the IVI-4.1: IviScope Class
// Specification.
func (d *Driver) SetAcquisitionType(acType scope.AcquisitionType) error {
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
	default:
		return ivi.ErrNotImplemented
	}

	return d.inst.Command(":ACQ:TYPE %s", cmd)
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
func (d *Driver) AcquisitionMinNumPoints() (int, error) {
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
func (d *Driver) SetAcquisitionMinNumPoints(numPoints int) error {
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
func (d *Driver) AcquisitionRecordLength() (int, error) {
	return query.Int(d.inst, ":ACQ:POIN?")
}

// AcquisitionSampleRate returns the effective sample rate of the acquired
// waveform using the current configuration. The units are samples per second.
//
// AcquisitionSampleRate is the getter for the read-only IviScopeBase
// Horizontal Sample Rate described in Section 4.2.10 of the IVI-4.1: IviScope
// Class Specification.
func (d *Driver) AcquisitionSampleRate() (float64, error) {
	return query.Float64(d.inst, ":ACQ:SRAT?")
}

// AcquisitionSampleRate queries the length of time that corresponds to the
// record length.
//
// AcquisitionTimePerRecord is the getter for the read-write IviScopeBase
// Horizontal Time Per Record described in Section 4.2.11 of the IVI-4.1:
// IviScope Class Specification.
func (d *Driver) AcquisitionTimePerRecord() (time.Duration, error) {
	return 0, ivi.ErrNotImplemented
}

// SetAcquisitionSampleRate specifies the length of time that corresponds to
// the record length.
//
// SetAcquisitionTimePerRecord is the setter for the read-write IviScopeBase
// Horizontal Time Per Record described in Section 4.2.11 of the IVI-4.1:
// IviScope Class Specification.
func (d *Driver) SetAcquisitionTimePerRecord(timePerRecord time.Duration) error {
	return ivi.ErrNotImplemented
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
func (d *Driver) TriggerHoldoff() (float64, error) {
	return 0.0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerHoldoff(holdoff float64) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerLevel() (float64, error) {
	return 0.0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerLevel(level float64) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerSlope() (scope.TriggerSlope, error) {
	return 0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerSlope(slope scope.TriggerSlope) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerSource() (scope.TriggerSource, error) {
	return 0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerSource(source scope.TriggerSource) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) TriggerType() (scope.TriggerType, error) {
	return 0, ivi.ErrNotImplemented
}

func (d *Driver) SetTriggerType(triggerType scope.TriggerType) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) AbortMeasurement() error {
	return ivi.ErrNotImplemented
}

func (d *Driver) ConfigureAcquisitionRecord(
	timePerRecord time.Duration,
	minNumPoints int,
	acquisitionStartTime time.Duration,
) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) CreateWaveform(numSamples int) error {
	return ivi.ErrNotImplemented
}

func (d *Driver) ConfigureEdgeTrigger(
	triggerType scope.TriggerType,
	level float64,
	slope scope.TriggerSlope,
) error {
	return ivi.ErrNotImplemented
}
func (d *Driver) ConfigureTrigger(triggerType scope.TriggerType, holdoff time.Duration) error {
	return ivi.ErrNotImplemented
}
func (d *Driver) InitiateMeasurement() error {
	return ivi.ErrNotImplemented
}

// ChannelEnabled queries whether or not the oscilloscope acquires a waveform for
// the channel.
//
// ChannelEnabled is the getter for the read-write IviScopeBase Attribute
// Channel Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) ChannelEnabled() (bool, error) {
	return query.Boolf(ch.inst, ":CHAN%d:DISPL?", ch.num)
}

// SetChannelEnabled sets the channel to either acquire (enabled) or not
// acquire (disabled) a waveform.
//
// Enabled is the setter for the read-write IviScopeBase Attribute Channel
// Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetChannelEnabled(b bool) error {
	cmd := "1"
	if !b {
		cmd = "0"
	}
	return ch.inst.Command(":CHAN%d:DISPL %s", ch.num, cmd)
}

// Name returns the name of the channel.
//
//	Name is the getter for the read-only IviScopeBase Attribute Channel Name
//	described in Section 4.2.7 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) Name() string {
	return fmt.Sprint("CH%d", ch.num)
}

// InputImpedance queries the input impedance for the channel in Ohms. Legal
// values are 50.0 and 1,000,000.0.

// InputImpedance is the getter for the read-write IviScopeBase Attribute Input
// Impedance described in Section 4.2.12 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) InputImpedance() (float64, error) {
	imped, err := query.Stringf(ch.inst, ":CHAN%d:IMP?", ch.num)
	if err != nil {
		return 0.0, err
	}
	switch imped {
	case "ONEM":
		return 1.0e6, nil
	case "FIFT":
		return 50.0, nil
	default:
		return 0.0, ivi.ErrValueNotSupported
	}
}

// SetInputImpedance sets the input impedance for the channel in Ohms. Legal
// values are 50.0 and 1,000,000.0.

// SetInputImpedance is the setter for the read-write IviScopeBase Attribute
// Input Impedance described in Section 4.2.12 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetInputImpedance(impedance float64) error {
	switch impedance {
	case 50.0:
		return ch.inst.Command(":CHAN%d:IMP FIFT", ch.num)
	case 1000000.0:
		return ch.inst.Command(":CHAN%d:IMP ONEM", ch.num)
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
func (ch *Channel) MaxInputFrequency() (float64, error) {
	return 0.0, ivi.ErrNotImplemented
}

func (ch *Channel) SetMaxInputFrequency(freq float64) error {
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
func (ch *Channel) ProbeAttenuation() (float64, error) {
	return query.Float64f(ch.inst, ":CHAN%d:PROBE?", ch.num)
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
func (ch *Channel) SetProbeAttenuation(atten float64) error {
	if atten < 0.001 || atten > 10000.0 {
		return ivi.ErrValueNotSupported
	}
	return ch.inst.Command(":CHAN%d:PROBE %E", ch.num, atten)
}

// ProbeAttenuationAuto always return false with no error since auto probe
// attenuation is not supported.

// ProbeAttenuationAuto is the getter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) ProbeAttenuationAuto() (bool, error) {
	return false, nil
}

// SetProbeAttenuationAuto if enabled will return an error since auto probe
// attenuation is not supported.

// SetProbeAttenuationAuto is the setter for the read-write IviScopeBase Probe
// Attenuation described in Section 4.2.16 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetProbeAttenuationAuto(b bool) error {
	if b {
		return ivi.ErrValueNotSupported
	}
	return nil
}

func (ch *Channel) TriggerCoupling() (scope.TriggerCoupling, error) {
	return 0, ivi.ErrNotImplemented
}

func (ch *Channel) SetTriggerCoupling(coupling scope.TriggerCoupling) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) VerticalCoupling() (scope.VerticalCoupling, error) {
	coupling, err := query.Stringf(ch.inst, ":CHAN%d:COUP?", ch.num)
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

func (ch *Channel) SetVerticalCoupling(coupling scope.VerticalCoupling) error {
	switch coupling {
	case scope.ACVerticalCoupling:
		return ch.inst.Command(":CHAN%d:COUP AC", ch.num)
	case scope.DCVerticalCoupling:
		return ch.inst.Command(":CHAN%d:COUP DC", ch.num)
	default:
		return ivi.ErrValueNotSupported
	}
}

// VerticalOffset queries the location of the center of the range that the
// Vertical Range attribute specifies. The value is with respect to ground and
// is in volts. For example, to acquire a sine wave that spans between on 0.0
// and 10.0 volts, set this attribute to 5.0 volts.

// VerticalOffset is the getter for the read-write IviScopeBase Vertical Offset
// described in Section 4.2.24 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) VerticalOffset() (float64, error) {
	return query.Float64f(ch.inst, ":CHAN%d:OFFS?", ch.num)
}

// SetVerticalOffset sets the location of the center of the range that the
// Vertical Range attribute specifies. The value is with respect to ground and
// is in volts. For example, to acquire a sine wave that spans between on 0.0
// and 10.0 volts, set this attribute to 5.0 volts.

// SetVerticalOffset is the setter for the read-write IviScopeBase Vertical
// Offset described in Section 4.2.24 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetVerticalOffset(offset float64) error {
	return ch.inst.Command(":CHAN%d:OFFS %E", ch.num, offset)
}

// VerticalRange queries the absolute value of the full-scale input range for a
// channel. The units are volts. For example, to acquire a sine wave that spans
// between -5.0 and 5.0 volts, set the Vertical Range attribute to 10.0 volts.

// VerticalRange is the getter for the read-write IviScopeBase Vertical Range
// described in Section 4.2.25 of the IVI-4.1: IviScope Class Specification.
func (ch *Channel) VerticalRange() (float64, error) {
	return query.Float64f(ch.inst, ":CHAN%d:RANG?", ch.num)
}

// SetVerticalRange sets the absolute value of the full-scale input range for a
// channel. The units are volts with valid ranges from 0.008 to 40.0. For
// example, to acquire a sine wave that spans between -5.0 and 5.0 volts, set
// the Vertical Range attribute to 10.0 volts.

// SetVerticalRange is the setter for the read-write IviScopeBase Vertical
// Range described in Section 4.2.25 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetVerticalRange(rng float64) error {
	if rng < 0.008 || rng > 40.0 {
		return ivi.ErrValueNotSupported
	}
	return ch.inst.Command(":CHAN%d:RANG %E", ch.num, rng)
}

func (ch *Channel) Configure(
	rng float64,
	offset float64,
	coupling scope.VerticalCoupling,
	autoProbeAttenuation bool,
	probeAttenuation float64,
	enabled bool,
) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) ConfigureCharacteristics(inputImepdance, inputFreqMax float64) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) FetchWaveform(waveform ivi.Waveform) error {
	return ivi.ErrNotImplemented
}

func (ch *Channel) ReadWaveform(maximumTime time.Duration, waveform ivi.Waveform) error {
	return ivi.ErrNotImplemented
}
