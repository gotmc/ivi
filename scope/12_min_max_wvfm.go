// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package scope

/*

# Section 12 IviScopeMinMaxWaveform Extension Group

## Section 12.1 IviScopeMinMaxWaveform Overview

The IviScopeMinMaxWaveform extension group provides support for oscilloscopes
that can acquire minimum and maximum waveforms that correspond to the same
range of time. The two most common acquisition types in which oscilloscopes
return minimum and maximum waveforms are envelope and peak detect.

#  Section 12.2 IviScopeMinMaxWaveform Attributes

Below are the .NET attributes, since they are the basis for the Go interfaces.

| Section | Attribute              | Type     | Access | AppliesTo   |
| ------- | ---------------------- | -------- | ------ | ----------- |
|  12.2.1 | Number of Envelopes    | Int32    | R/W    | N/A         |

## Section 12.3 IviScopeMinMaxWaveform Functions

Below are the .NET functions, since they are the basis for the Go interfaces.

12.3.2 MinMaxWaveform<Double> Channels[].Measurement.FetchWaveformMinMax (
                                               MinMaxWaveform<Double> minMaxWaveform);
       MinMaxWaveform<Int32> Channels[].Measurement.FetchWaveformMinMax (
                                               MinMaxWaveform<Int32> minMaxWaveform);
       MinMaxWaveform<Int16> Channels[].Measurement.FetchWaveformMinMax (
                                               MinMaxWaveform<Int16> minMaxWaveform);
       MinMaxWaveform<Byte> Channels[].Measurement.FetchWaveformMinMax (
                                               MinMaxWaveform<Byte> minMaxWaveform);

      Struct MinMaxWaveform<T>
      {
         Public MinMaxWaveform(IWaveform<T> minWaveform, IWaveform<T> maxWaveform);
         public IWaveform<T> MinWaveform { get; }
         public IWaveform<T> MaxWaveform { get; }
      }

12.3.3 MinMaxWaveform<Double> Channels[].Measurement.ReadWaveformMinMax (
                                          PrecisionTimeSpan maximumTime,
                                          MinMaxWaveform<Double> minMaxWaveform);
       MinMaxWaveform<Int32> Channels[].Measurement.ReadWaveformMinMax (
                                           PrecisionTimeSpan maximumTime,
                                           MinMaxWaveform<Int32> minMaxWaveform);
       MinMaxWaveform<Int16> Channels[].Measurement.ReadWaveformMinMax (
                                           PrecisionTimeSpan maximumTime,
                                           MinMaxWaveform<Int16> minMaxWaveform);
       MinMaxWaveform<Byte> Channels[].Measurement.ReadWaveformMinMax (
                                           PrecisionTimeSpan maximumTime,
                                           MinMaxWaveform<Byte> minMaxWaveform);

      Struct MinMaxWaveform<T>
      {
         Public MinMaxWaveform(IWaveform<T> minWaveform, IWaveform<T> maxWaveform);
         public IWaveform<T> MinWaveform { get; }
         public IWaveform<T> MaxWaveform { get; }
      }

*/
