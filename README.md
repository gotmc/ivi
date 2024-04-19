# ivi

Go-based implementation of the Interchangeable Virtual Instrument (IVI)
standard.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

## Overview

The [IVI Specifications][ivi-specs] developed by the [IVI
Foundation][ivi-foundation] provide standardized APIs for programming test
instruments. This package is a partial, Go-based implementation of the IVI
Specifications.

The main advantage of the [ivi][] package is not having to learn the [SCPI][]
commands for each individual piece of test equipment. For instance, by using the
[ivi][] package both the Agilent 33220A and the Stanford Research Systems DS345
function generators can be programmed using one standard API. The only
requirement for this is having an IVI driver for the desired test equipment.

If an [ivi][] driver doesn't exist for a peice of test equipment that you want
to use, please open an issue and/or submit a pull request. The [IVI
Specifications][ivi-specs] don't provide APIs for every type of test equipment
(e.g., they don't specify an API for electronic loads) in which case a set of
APIs will be developed as needed for new types of test equipment.

Development focus is currently on solidifying the APIs and creating a few IVI
drivers for each instrument type.

### IVI Driver Goal

Per Section 2.6 Capability Groups in IVI-3.1 Driver Architecture Specification:

> The fundamental goal of IVI drivers is to allow test developers to change the
> instrumentation hardware on their test systems without changing test program
> source code, recompiling, or re-linking. To achieve this goal, instrument
> drivers must have a standard programming interface...Because instruments do
> not have identical functionality or capability, it is impossible to create a
> single programming interface that covers all features of all instruments in a
> class. For this reason, the IVI Foundation recognizes different types of
> capabilities – Inherent Capabilities, Base Class Capabilities, Class Extension
> Capabilities and Instrument Specific Capabilities.

### IVI Specification Deviation

**TL;DR:** When developing Go-based IVI drivers follow the .NET methods and
prototypes as much as possible but deviate as required.

As stated in Section 1.5 Conformance Requirements of the _IVI-3.1: Driver
Architecture Specification, Revision 3.8_, "IVI drivers can be developed with a
COM, ANSI-C, or .NET API." In general, the Go method signatures try to be as
close to the .NET signatures as possible. However, given the desire to write
idiomatic Go, where necessary and where it makes sense, the Go-based IVI drivers
deviate from the detailed [IVI Specifications][ivi-specs] at times.

For instance, since Go does not provide [method overloading][go-overload],
the .NET method prototypes cannot be followed in all cases. For example,
in the IVI-4.2 IviDmm Class Specification, the .NET method prototypes show two
`Configure` methods with different function signatures based on whether
auto-range is specified or a manual range value is provided.

```csharp
void Configure(MeasurementFunction measurementFunction,
               Auto autoRange,
               Double resolution);

void Configure(MeasurementFunction measurementFunction,
               Double range, 
               Double resolution);
```

However, Go isn't C, so the Go-based IVI drivers don't have to rely on defined
values, such as specific negative numbers representing auto-range (e.g., -1.0 =
auto range on) and positive numbers representing the user specified manual
range. Using the same example, below is the C prototype for `Configure`, where
`Range` is a `ViReal64` and negative values provided defined values.

```c
ViStatus IviDmm_ConfigureMeasurement (ViSession Vi,
                                      ViInt32 Function,
                                      ViReal64 Range,
                                      ViReal64 Resolution);
```

Because of these differences, the Go function signatures may deviate from the
[IVI Specifications][ivi-specs], when required and where it makes sense to
enable writing idiomatic Go. Using the same example, below is the function
signature for `ConfigureMeasurement` in Go.

```go
ConfigureMeasurement(
    msrFunc MeasurementFunction,
    autoRange AutoRange,
    rangeValue float64,
    resolution float64,
) error
```

`MeasurementFunction` and `AutoRange` are user-defined types with [enumerated
constant][go-enums] values. Note: the function parameter `rangeValue` is used
instead of `range`, since `range` is a reserved keyword and cannot be used as an
identifier in Go.

```go
type MeaurementFunction int

const (
	DCVolts MeasurementFunction = iota
	ACVolts
    ...
	Period
	Temperature
)

type AutoRange int

const (
	AutoOn AutoRange = iota
	AutoOff
	AutoOnce
)
```

## Installation

```bash
$ go get github.com/gotmc/ivi
```

## Usage

The [ivi][] package requires receiving an Instrument interface. The following
gotmc packages meet the Instrument interface:

- [visa][] — Calls [lxi][] or [usbtmc][] as needed, so that you can identify
  instruments using a VISA resource address string.
- [lxi][] — Used to control LXI enabled instruments via Ethernet.
- [usbtmc][] — Used to control USBTMC compliant instruments via USB.
- [prologix][] — Used to communicate with instruments using a Prologix GPIB
  controller.
- [asrl][] — Used to control intstruments via serial.

## Examples

Examples can be found at <https://github.com/gotmc/ivi-examples>.

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/ivi>
- <http://localhost:6060/pkg/github.com/gotmc/ivi/> after running `$
godoc -http=:6060`

## Contributing

Contributions are welcome! To contribute please:

1. Fork the repository
2. Create a feature branch
3. Code
4. Submit a [pull request][]

### Testing

Prior to submitting a [pull request][], please run:

```bash
$ make check
$ make lint
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[ivi][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[asrl]: https://github.com/gotmc/asrl
[ivi]: https://github.com/gotmc/ivi
[ivi-foundation]: http://www.ivifoundation.org/
[ivi-specs]: http://www.ivifoundation.org/specifications/
[go-enums]: https://go.dev/doc/effective_go#constants
[go-overload]: https://go.dev/doc/faq#overloading
[godoc badge]: https://godoc.org/github.com/gotmc/ivi?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/ivi
[LICENSE.txt]: https://github.com/gotmc/ivi/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[lxi]: https://github.com/gotmc/lxi
[prologix]: https://github.com/gotmc/prologix
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/gotmc/ivi
[report card]: https://goreportcard.com/report/github.com/gotmc/ivi
[scpi]: http://www.ivifoundation.org/scpi/
[usbtmc]: https://github.com/gotmc/usbtmc
[visa]: https://github.com/gotmc/visa
