# ivi
Go-based implementation of the Interchangeable Virtual Instrument (IVI)
standard.

[![GoDoc][godoc badge]][godoc link]
[![Go Report Card][report badge]][report card]
[![License Badge][license badge]][LICENSE.txt]

## Overview

The [IVI Specifications][ivi-specs] developed by the IVI Foundation
provide standardized APIs for programming test instruments. This package
is a partial, Go-based implementation of the IVI Specifications, which
are specified for C, COM, and .NET.

The main advantage of the ivi package is not having to learn the
[SCPI][] commands for each individual peice of test equipment. For
instance, both the Agilent 33220A and the Stanford Research Systems
DS345 function generators can be programmed using one standard API. The
only requirement for this is having an IVI driver for the desired test
equipment.

Currently, [ivi][] doesn't cache state. Every time an attribute is read
directly from the instrument. Development focus is currently on fleshing
out the APIs and creating a few IVI drivers for each instrument type.

## Installation

```bash
$ go get github.com/gotmc/ivi
```

## Usage

The [ivi][ivi] package requires receiving an Instrument interface. The
[visa][], [lxi][], and [usbtmc][] packages meet the Instrument
interface. You can either use [visa][], which will call [lxi][] and/or
[usbtmc][] as nescessary, or you can directly call [usbtmc][] or [lxi][]
as desired.

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
```

To update and view the test coverage report:

```bash
$ make cover
```

## License

[ivi][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[ivi]: https://github.com/gotmc/ivi
[ivi-specs]: http://www.ivifoundation.org/
[godoc badge]: https://godoc.org/github.com/gotmc/ivi?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/ivi
[LICENSE.txt]: https://github.com/gotmc/ivi/blob/master/LICENSE.txt
[license badge]: https://img.shields.io/badge/license-MIT-blue.svg
[lxi]: https://github.com/gotmc/lxi
[pull request]: https://help.github.com/articles/using-pull-requests
[report badge]: https://goreportcard.com/badge/github.com/gotmc/ivi
[report card]: https://goreportcard.com/report/github.com/gotmc/ivi
[scpi]: http://www.ivifoundation.org/scpi/
[usbtmc]: https://github.com/gotmc/usbtmc
[visa]: https://github.com/gotmc/visa
