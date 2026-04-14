# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Development Commands

### Building and Testing

- **Run tests**: `just unit` - formats, vets, and runs unit tests with coverage
- **Run single test**: `just unit -run TestName` - run a specific test
- **Run single package tests**: `just unit ./dmm/...` - run tests in a package
- **Integration tests**: `just int` - runs integration tests
- **E2E tests**: `just e2e` - runs end-to-end tests
- **Linting**: `just lint` - runs golangci-lint v2 with `.golangci.yaml` config
  (staticcheck, govet, errcheck, ineffassign, unused, gosec, misspell,
  bodyclose, contextcheck)
- **Coverage report**: `just cover` - generates HTML coverage report (also
  accepts `just cover int`, `just cover e2e`, or `just cover all`)
- **Format and vet**: `just check` - formats and vets code (runs before tests
  automatically)
- **Lines of code**: `just loc` - count lines of code using scc
- **Local docs**: `just docs` - browse documentation locally using pkgsite

### Dependencies

- **Update single module**: `just update <module>`
- **Update all**: `just updateall`
- **Check outdated**: `just outdated` (requires go-mod-outdated)
- **Tidy and verify**: `just tidy`

## Code Architecture

### High-Level Structure

This is a Go implementation of the IVI (Interchangeable Virtual Instrument)
Foundation specifications for programming test instruments. The architecture
follows IVI class specifications with standardized APIs across different
instrument manufacturers. The main advantage is abstracting SCPI commands behind
a standard API so different instruments (e.g., Agilent 33220A and SRS DS345
function generators) can be programmed identically.

### Key Components

#### Core Interface Layer (`ivi.go`, `inherent.go`)

- `Transport` interface: Core abstraction requiring Command, Query, ReadBinary,
  WriteBinary, and Close methods. The I/O methods all take `context.Context` as
  their first parameter. Sub-interfaces `Commander`, `Querier`, `BinaryReader`,
  and `BinaryWriter` allow accepting narrower types.
- `Inherent` struct: Base capabilities common to all IVI instruments (reset,
  clear, identification, timeout, local control). All methods that communicate
  with instruments take `context.Context`.
- `InherentBase` struct: Metadata about the driver (class spec version, supported
  models, bus interfaces, reset/clear delays)
- Inherent capabilities follow IVI-3.2 specification

#### Instrument Class Packages

Each directory defines interfaces for an IVI instrument class:

- `dmm/` — IVI-4.2 IviDmm Class (Digital Multimeter)
- `fgen/` — IVI-4.3 IviFgen Class (Function Generator)
- `dcpwr/` — IVI-4.4 IviDCPwr Class (DC Power Supply)
- `scope/` — IVI-4.1 IviScope Class (Oscilloscope)

Additional class packages exist with varying levels of implementation:
`acpwr`, `counter`, `digitizer`, `downcnvrtr`, `dsa`, `lcr`, `load`,
`pwrmeter`, `rfsiggen`, `specan`, `swtch`, `upcnvrtr`.

The `com/` package provides communication port types (serial modes, data frames,
parity, equipment class) used by transport layers.

Each class package defines Go interfaces for capability groups (Base,
StdFunc, Trigger, etc.) that drivers must implement. All interface methods
that communicate with instruments take `context.Context` as first parameter.
All enum types implement the `Stringer` interface with `String()` methods.

#### Driver Implementation Pattern

Drivers live under `<class>/<manufacturer>/<model>/` (e.g.,
`fgen/keysight/key33220/`). All drivers follow this structure:

- `01_<model>.go`: Package doc, Driver struct, Channel struct, `New()` constructor,
  `Close()` method
- `04_base.go`: Base capability group
- `05_*.go` onwards: Additional capability groups (StdFunc, AC Measurements,
  Triggers, etc.) — one file per IVI specification section

The Driver struct always follows this pattern:

```go
type Driver struct {
    inst     ivi.Instrument
    channels []Channel           // unexported; access via Channel(index)
    ivi.Inherent                 // Embedded for inherent capabilities
}
```

The `New(inst ivi.Instrument, idQuery, reset bool) (*Driver, error)` constructor
creates channels, populates `InherentBase` metadata, and calls
`ivi.NewInherent()`. When `idQuery` is true, the constructor queries `*IDN?`
and validates the instrument model against the driver's supported models list
via `Inherent.CheckID()`. When `reset` is true, the instrument is reset.

Channel structs hold an `ivi.Instrument` reference and a channel `name` string,
implementing per-channel capability interfaces. Channels are accessed via a
bounds-checked accessor method:

```go
ch, err := driver.Channel(0)  // returns (*Channel, error)
```

Drivers implement a `Close()` method that delegates to the embedded Inherent.
`Inherent.Close()` calls `Disable()` then closes the underlying connection if
it implements `io.Closer`:

```go
func (d *Driver) Close() error {
    return d.Inherent.Close()
}
```

#### SCPI Command Mapping Pattern

Drivers use dual maps for bidirectional enum↔SCPI conversion, paired with
generic helpers from the root `ivi` package:

```go
// Forward map: Go enum → SCPI command string (for setting values)
var outputModeToSCPI = map[fgen.OutputMode]string{
    fgen.OutputModeFunction: "FUNC SIN",
    fgen.OutputModeNoise:    "FUNC NOIS",
}

// Reverse map: SCPI response string → Go enum (for querying values)
var scpiToOutputMode = map[string]fgen.OutputMode{
    "SIN":  fgen.OutputModeFunction,
    "NOIS": fgen.OutputModeNoise,
}

// Usage via generic helpers:
cmd, err := ivi.LookupSCPI(outputModeToSCPI, outputMode)   // returns ErrValueNotSupported on miss
mode, err := ivi.ReverseLookup(scpiToOutputMode, scpiStr)   // returns ErrUnexpectedResponse on miss
```

#### Timeout Support (`timeout.go`)

`ivi.DefaultTimeout` (10 seconds) is a convenience constant for creating
timeout contexts. All instrument I/O uses `context.Context` for per-operation
timeout and cancellation control.

#### Core Helpers and Errors (`helpers.go`, `errors.go`)

- `ivi.Set(ctx, cmdr, format, args...)` — convenience for sending formatted SCPI
  command strings via the Commander interface
- `ivi.QueryID(ctx, q)` — standard `*IDN?` query
- Error sentinels: `ErrNotImplemented`, `ErrFunctionNotSupported`,
  `ErrValueNotSupported`, `ErrUnexpectedResponse`, `ErrChannelNotFound`,
  `ErrUnsupportedModel`
- Class packages may define additional error sentinels (e.g., `dcpwr` defines
  `ErrOVPUnsupported` and `ErrTriggerNotSoftware`)

### Design Philosophy

- **Idiomatic Go over strict IVI compliance**: Uses Go enums and type safety
  instead of magic numbers. Go method signatures follow .NET prototypes where
  possible, but deviate when needed for idiomatic Go (e.g., separate AutoRange
  type instead of negative sentinel values).
- **No state caching**: All attributes read directly from instruments for
  reliability
- **Interface-based**: Drivers implement capability interfaces, enabling
  instrument interchangeability
- **Capability groups**: Code organized by IVI specification capability groups

### Transport Layer

The `Instrument` interface is satisfied by external transport packages:
`visa`, `lxi`, `usbtmc`, `prologix`, `asrl` (all under `github.com/gotmc/`).
Drivers are transport-agnostic — they only depend on the `ivi.Instrument`
interface.

### Dependencies

- `github.com/gotmc/query`: SCPI query utilities
- `github.com/gotmc/convert`: Data conversion helpers
- Go 1.21+ required

## Code Style and Conventions

### File Organization

- Files prefixed with numbers (01_, 04_, 05_) correspond to IVI specification
  sections
- Capability groups determine file boundaries (Base, AC Measurements, Triggers,
  etc.)
- Driver packages use manufacturer/model naming: `keysight/key3446x/`,
  `srs/ds345/`

### Interface Compliance Verification

All drivers include compile-time interface compliance checks:

```go
var _ dmm.Base = (*Driver)(nil)
var _ dmm.BaseChannel = (*Channel)(nil)
```

### Testing Patterns

- Table-driven tests with `t.Run()` subtests
- Mock instruments implementing the `Instrument` interface for unit tests
- Integration tests use `Integration` prefix, E2E tests use `E2E` prefix in
  test names (filtered by `-run` flag)

### Error Handling

- Package-specific error variables defined in main package files
- IVI-specific errors like `ErrNotImplemented` for unsupported features
- Wrap errors with `fmt.Errorf("context: %w", err)` to preserve the error chain

### Formatting

- `golines` is enabled via golangci-lint, so long lines will be wrapped
  automatically on `just check` or `just lint`
- `goimports` is enabled, so imports are organized automatically
