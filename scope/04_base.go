// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

import (
	"context"
	"time"

	"github.com/gotmc/ivi"
)

/*

# Section 4 IviScopeBase Capability Group

## Section 4.1 IviScopeBase Overview

The IviScope base capabilities support oscilloscopes that can acquire waveforms
from multiple channels with an edge trigger. The IviScope base capabilities
define attributes and their values to configure the oscilloscope’s channel,
acquisition, and trigger sub-systems. The IviScope base capabilities also
include functions for configuring the oscilloscope as well as initiating
waveform acquisition and retrieving waveforms.

The IviScope base capabilities organize the configurable settings into three
main categories: the channel sub-system, the acquisition sub-system, and the
trigger sub-system.

## Section 4.2 IviScopeBase Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|   4.2.1 | Acquisition Start Time | TimeSpan | R/W    | N/A         |
|   4.2.2 | Acquisition Status     | AqStatus | R/O    | Measurement |
|   4.2.3 | Acquisition Type       | Int32    | R/W    | N/A         |
|   4.2.4 | Channel Count          | Int32    | RO     | Channel     |
|   4.2.5 | Channel Enabled        | Bool     | R/W    | Channel     |
|   4.2.6 | Channel Item           | Channel  | RO     | Channel     |
|   4.2.7 | Channel Name           | String   | RO     | Channel     |
|   4.2.8 | Horiz Min Num Points   | Int32    | R/W    | N/A         |
|   4.2.9 | Horiz Record Length    | Int32    | RO     | N/A         |
|  4.2.10 | Horiz Sample Rate      | Real64   | RO     | N/A         |
|  4.2.11 | Horiz Time per Record  | TimeSpan | R/W    | N/A         |
|  4.2.12 | Input Impedance        | Real64   | R/W    | Channel     |
|  4.2.13 | Max Input Frequency    | Real64   | R/W    | Channel     |
|  4.2.16 | Probe Attenuation      | Real64   | R/W    | Channel     |
|  4.2.17 | Trigger Coupling       | Int32    | R/W    | Channel     |
|  4.2.18 | Trigger Holdoff        | TimeSpan | R/W    | N/A         |
|  4.2.19 | Trigger Level          | Real64   | R/W    | N/A         |
|  4.2.20 | Trigger Slope          | Int32    | R/W    | N/A         |
|  4.2.21 | Trigger Source         | String   | R/W    | N/A         |
|  4.2.22 | Trigger Type           | Int32    | R/W    | N/A         |
|  4.2.23 | Vertical Coupling      | Int32    | R/W    | Channel     |
|  4.2.24 | Vertical Offset        | Real64   | R/W    | Channel     |
|  4.2.25 | Vertical Range         | Real64   | R/W    | Channel     |

## Section 4.3 IviScopeBase Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

4.3.1 void Measurement.Abort();
4.3.4 void Acquisition.ConfigureRecord(
                          Ivi.Driver.PrecisionTimeSpan timePerRecord,
                          Int32 minimumNumberPoints,
                          Ivi.Driver.PrecisionTimeSpan acquisitionStartTime);
4.3.6 void Channels[].Configure (Double range,
                                 Double offset,
                                 VerticalCoupling coupling,
                                 Double probeAttenuation,
                                 Boolean enabled);
      void Channels[].Configure (Double range,
                                 Double offset,
                                 VerticalCoupling coupling,
                                 Boolean probeAttenuationAuto,
                                 Boolean enabled);
4.3.7 IWaveform<Double> Measurement.CreateWaveformDouble(Int32 numberSamples);
      IWaveform<Int32>  Measurement.CreateWaveformInt32(Int32 numberSamples);
      IWaveform<Int16>  Measurement.CreateWaveformInt16(Int32 numberSamples);
      IWaveform<Byte>   Measurement.CreateWaveformByte(Int32 numberSamples);
4.3.8 void Channels[].ConfigureCharacteristics (Double inputImpedance,
                                                Double inputFrequencyMaximum);
4.3.9 void Trigger.Edge.Configure (String source,
                                   Double level,
                                   Slope slope);
4.3.10 void Trigger.Configure (TriggerType type,
                               Ivi.Driver.PrecisionTimeSpan holdoff);
4.3.13 IWaveform<Double> Channels[].Measurement.FetchWaveform(
                                                   IWaveform<Double> waveform);
       IWaveform<Int32> Channels[].Measurement.FetchWaveform(
                                                   IWaveform<Int32> waveform);
       IWaveform<Int16> Channels[].Measurement.FetchWaveform(
                                                   IWaveform<Int16> waveform);
       IWaveform<Byte> Channels[].Measurement.FetchWaveform(
                                                   IWaveform<Byte> waveform);
4.3.14 void Measurement.Initiate();
4.3.16 IWaveform<Double> Channels[].Measurement.ReadWaveform(
                                                 PrecisionTimeSpan maximumTime
                                                 IWaveform<Double> waveform);
       IWaveform<Int32>  Channels[].Measurement.ReadWaveform(
                                                 PrecisionTimeSpan maximumTime
                                                 IWaveform<Int32> waveform);
       IWaveform<Int16>  Channels[].Measurement.ReadWaveform(
                                                 Precision TimeSpan maximumTime
                                                 IWaveform<Int16> waveform);
       IWaveform<Byte>   Channels[].Measurement.ReadWaveform(
                                                 Precision TimeSpan maximumTime
                                                 IWaveform<Byte> waveform);


*/

// Base provides the interface required for the IviScopeBase capability group.
type Base interface {
	AcquisitionStartTime(ctx context.Context) (time.Duration, error)
	SetAcquisitionStartTime(ctx context.Context, startTime time.Duration) error
	AcquisitionStatus(ctx context.Context) (AcquisitionStatus, error)
	AcquisitionType(ctx context.Context) (AcquisitionType, error)
	SetAcquisitionType(ctx context.Context, acquisitionType AcquisitionType) error
	ChannelCount() int
	AcquisitionMinNumPoints(ctx context.Context) (int, error)
	SetAcquisitionMinNumPoints(ctx context.Context, numPoints int) error
	AcquisitionRecordLength(ctx context.Context) (int, error)
	AcquisitionSampleRate(ctx context.Context) (float64, error)
	AcquisitionTimePerRecord(ctx context.Context) (time.Duration, error)
	SetAcquisitionTimePerRecord(ctx context.Context, timePerRecord time.Duration) error
	TriggerHoldoff(ctx context.Context) (time.Duration, error)
	SetTriggerHoldoff(ctx context.Context, holdoff time.Duration) error
	TriggerLevel(ctx context.Context) (float64, error)
	SetTriggerLevel(ctx context.Context, level float64) error
	TriggerSlope(ctx context.Context) (TriggerSlope, error)
	SetTriggerSlope(ctx context.Context, slope TriggerSlope) error
	TriggerSource(ctx context.Context) (TriggerSource, error)
	SetTriggerSource(ctx context.Context, source TriggerSource) error
	TriggerType(ctx context.Context) (TriggerType, error)
	SetTriggerType(ctx context.Context, triggerType TriggerType) error
	AbortMeasurement(ctx context.Context) error
	ConfigureAcquisitionRecord(
		ctx context.Context,
		timePerRecord time.Duration,
		minNumPoints int,
		acquisitionStartTime time.Duration,
	) error
	CreateWaveform(ctx context.Context, numSamples int) error
	ConfigureEdgeTrigger(
		ctx context.Context,
		triggerType TriggerType,
		level float64,
		slope TriggerSlope,
	) error
	ConfigureTrigger(ctx context.Context, triggerType TriggerType, holdoff time.Duration) error
	InitiateMeasurement(ctx context.Context) error
}

// BaseChannel provides the interface required for the channel repeated
// capability for the IviScopeBase capability group.
type BaseChannel interface {
	ChannelEnabled(ctx context.Context) (bool, error)
	SetChannelEnabled(ctx context.Context, b bool) error
	Name() string
	InputImpedance(ctx context.Context) (float64, error)
	SetInputImpedance(ctx context.Context, impedance float64) error
	MaxInputFrequency(ctx context.Context) (float64, error)
	SetMaxInputFrequency(ctx context.Context, freq float64) error
	ProbeAttenuation(ctx context.Context) (float64, error)
	SetProbeAttenuation(ctx context.Context, atten float64) error
	ProbeAttenuationAuto(ctx context.Context) (bool, error)
	SetProbeAttenuationAuto(ctx context.Context, b bool) error
	TriggerCoupling(ctx context.Context) (TriggerCoupling, error)
	SetTriggerCoupling(ctx context.Context, coupling TriggerCoupling) error
	VerticalCoupling(ctx context.Context) (VerticalCoupling, error)
	SetVerticalCoupling(ctx context.Context, coupling VerticalCoupling) error
	VerticalOffset(ctx context.Context) (float64, error)
	SetVerticalOffset(ctx context.Context, offset float64) error
	VerticalRange(ctx context.Context) (float64, error)
	SetVerticalRange(ctx context.Context, rng float64) error
	Configure(
		ctx context.Context,
		rng float64,
		offset float64,
		coupling VerticalCoupling,
		autoProbeAttenuation bool,
		probeAttenuation float64,
		enabled bool,
	) error
	ConfigureCharacteristics(ctx context.Context, inputImpedance, inputFreqMax float64) error
	FetchWaveform(ctx context.Context, waveform *ivi.Waveform) error
	ReadWaveform(ctx context.Context, maximumTime time.Duration, waveform *ivi.Waveform) error
}
