// Copyright (c) 2017-2024 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package dmm

import "time"

/*

# Section 11 IviDmmMultiPoint Capability Group

## Section 11.1 IviDmmMultiPoint Overview

The IviDmmMultiPoint extension group defines extensions for DMMs capable of
acquiring measurements based on multiple triggers, and acquiring multiple
measurements for each trigger.

The IviDmmMultiPoint extension group defines additional attributes such sample
count, sample trigger, trigger count, and trigger delay to control
“multi-point” DMMs. The IviDmmMultiPoint extension group also adds functions
for configuring the DMM as well as starting acquisitions and retrieving
multiple measured values.

## Section 11.2 IviDmmMultiPoint Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute                    | Type     | Access | AppliesTo |
| ------- | ---------------------------- | -------- | ------ | --------- |
|  11.2.1 | Measure Complete Destination | String   | R/W    | N/A       |
|  11.2.2 | Sample Count                 | Int32    | R/W    | N/A       |
|  11.2.3 | Sample Interval              | TimeSpan | R/W    | N/A       |
|  11.2.4 | Sample Trigger               | String   | R/W    | N/A       |
|  11.2.5 | Trigger Count                | Int32    | R/W    | N/A       |

## Section 11.3 IviDmmMultiPoint Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

11.3.2 void Trigger.MultiPoint.Configure(
                                Double triggerCount,
                                Double sampleCount,
                                String sampleTrigger,
                                Ivi.Driver.PrecisionTimeSpan sampleInterval);

11.3.3 Double[] Measurement.FetchMultiPoint(PrecisionTimeSpan maximumTime);
       Double[] Measurement.FetchMultiPoint(PrecisionTimeSpan maximumTime,
                                            Int32 numberOfMeasurements);

11.3.4 Double[] Measurement.ReadMultiPoint(PrecisionTimeSpan maximumTime);
       Double[] Measurement.ReadMultiPoint(PrecisionTimeSpan maximumTime,
                                           Int32 numberOfMeasurements);

*/

// MultiPointExtension provides the interface required for the IviDmmMultiPoint
// extension group described in Section 11 of IVI-4.2 IviDmm Class
// Specification.
type MultiPointExtension interface {
	MeasureCompleteDestination() (MeasurementDestination, error)
	SetMeasureCompleteDestination(dest MeasurementDestination) error
	SampleCount() (int, error)
	SetSampleCount(count int) error
	SampleInterval() (time.Duration, error)
	SetSampleInterval(interval time.Duration) error
	SampleTrigger() (TriggerSource, error)
	SetSampleTrigger(triggerSource TriggerSource) error
	TriggerCount() (int, error)
	SetTriggerCount(count int) error
	ConfigureMultiPoint(
		triggerCount, sampleCount int,
		triggerSource TriggerSource,
		interval time.Duration,
	) error
	// FetchMultiPoint() ([]float64, error)
	// ReadMultiPoint() ([]float64, error)
}
