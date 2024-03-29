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
Specifications, which are specified for C, COM, and .NET.

The main advantage of the [ivi][] package is not having to learn the [SCPI][]
commands for each individual piece of test equipment. For instance, by using the
[ivi][] package both the Agilent 33220A and the Stanford Research Systems DS345
function generators can be programmed using one standard API. The only
requirement for this is having an IVI driver for the desired test equipment.

If an [ivi][] driver doesn't exist for a peice of test equipment that you want
to use, please open an issue and/or submit a pull request. The [IVI
Specifications][] don't provide APIs for every type of test equipment (e.g.,
they don't specify an API for electronic loads) in which case a set of APIs will
be developed as needed for new types of test equipment.

Development focus is currently on solidifying the APIs and creating a few IVI
drivers for each instrument type.


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
- [prologix][] —  Used to communicate with instruments using a Prologix GPIB
    controller.

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

[ivi]: https://github.com/gotmc/ivi
[ivi-foundation]: http://www.ivifoundation.org/
[ivi-specs]: http://www.ivifoundation.org/specifications/
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
