// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

import (
	"context"
	"time"
)

/*

# Section 4 IviDmmBase Capability Group

## Section 4.1 IviDmmBase Overview

The IviDmmBase Capability Group supports DMMs that take a single measurement at
a time. The IviDmmBase Capability Group defines attributes and their values to
configure the type of measurement and how the measurement is to be performed.
These attributes include the measurement function, range, resolution, and
trigger source. The IviDmmBase capability group also includes functions for
configuring the DMM as well as initiating and retrieving measurements.

## Section 4.2 IviDmmBase Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute           | Type     | Access | AppliesTo |
| ------- | ----------------    | ------   | ------ | --------- |
|   4.2.1 | Function            | Int32    | R/W    | N/A       |
|   4.2.2 | Range               | Real64   | R/W    | N/A       |
|   4.2.3 | AutoRange           | Int32    | R/W    | N/A       |
|   4.2.4 | Resolution Absolute | Real64   | R/W    | N/A       |
|   4.2.5 | Trigger Delay       | TimeSpan | R/W    | N/A       |
|   4.2.6 | Trigger Delay Auto  | Bool     | R/W    | N/A       |
|   4.2.7 | Trigger Source      | String   | R/W    | N/A       |

## Section 4.3 IviDmmBase Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

4.3.1 void Measurement.Abort();

4.3.2 void Configure(MeasurementFunction measurementFunction,
                     Auto autoRange,
                     Double resolution);
      void Configure(MeasurementFunction measurementFunction,
                     Double range,
                     Double resolution);

4.3.3 void Trigger.Configure(String triggerSource,
                              Boolean autoTriggerDelay)
       void Trigger.Configure(String triggerSource,
                              Ivi.Driver.PrecisionTimeSpan triggerDelay);

4.3.4 Double Measurement.Fetch(PrecisionTimeSpan maximumTime);

4.3.5 void Meaurement.Initiate();

4.3.6 Boolean Measurement.IsOutOfRange(Double MeasurementValue);

4.3.7 Boolean Measurement.IsOverRange(Double MeasurementValue);

4.3.8 Boolean Measurement.IsUnderRange(Double MeasurementValue);

4.3.9 Double Measurement.Read(PrecisionTimeSpan maximumTime);


*/

// Base provides the interface required for the IviDCPwrBase capability group.
type Base interface {
	MeasurementFunction(ctx context.Context) (MeasurementFunction, error)
	SetMeasurementFunction(ctx context.Context, msrFunc MeasurementFunction) error
	Range(ctx context.Context) (autoRange AutoRange, rangeValue float64, err error)
	SetRange(ctx context.Context, autoRange AutoRange, rangeValue float64) error
	ResolutionAbsolute(ctx context.Context) (float64, error)
	SetResolutionAbsolute(ctx context.Context, resolution float64) error
	TriggerDelay(ctx context.Context) (autoDelay bool, delay float64, err error)
	SetTriggerDelay(ctx context.Context, autoDelay bool, delay float64) error
	TriggerSource(ctx context.Context) (TriggerSource, error)
	SetTriggerSource(ctx context.Context, src TriggerSource) error
	Abort(ctx context.Context) error
	ConfigureMeasurement(
		ctx context.Context,
		msrFunc MeasurementFunction,
		autoRange AutoRange,
		rangeValue float64,
		resolution float64,
	) error
	ConfigureTrigger(ctx context.Context, src TriggerSource, delay time.Duration) error
	FetchMeasurement(ctx context.Context, maxTime time.Duration) (float64, error)
	InitiateMeasurement(ctx context.Context) error
	IsOutOfRange(ctx context.Context, value float64) (bool, error)
	IsOverRange(ctx context.Context, value float64) (bool, error)
	IsUnderRange(ctx context.Context, value float64) (bool, error)
	ReadMeasurement(ctx context.Context, maxTime time.Duration) (float64, error)
}
