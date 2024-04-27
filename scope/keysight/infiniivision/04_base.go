// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package infiniivision

import (
	"fmt"

	"github.com/gotmc/ivi"
	"github.com/gotmc/ivi/scope"
	"github.com/gotmc/query"
)

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

// Enabled queries whether or not the oscilloscope acquires a waveform for
// the channel.
//
// Enabled is the getter for the read-write IviScopeBase Attribute Channel
// Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) Enabled() (bool, error) {
	return query.Boolf(ch.inst, ":CHAN%d:DISPL?", ch.num)
}

// Enabled sets the channel to either acquire (enabled) or not acquire
// (disabled) a waveform.
//
// Enabled is the setter for the read-write IviScopeBase Attribute Channel
// Enabled described in Section 4.2.5 of the IVI-4.1: IviScope Class
// Specification.
func (ch *Channel) SetEnabled(b bool) error {
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

func (ch *Channel) VerticalRange() (float64, error) {
	return 0.0, ivi.ErrNotImplemented
}

func (ch *Channel) SetVerticalRange(rng float64) error {
	return ivi.ErrNotImplemented
}
