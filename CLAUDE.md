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
- **Linting**: `just lint` - runs golangci-lint v2 with the project's configuration
  (includes `gosec`, `bodyclose`, `contextcheck` among enabled linters)
- **Coverage report**: `just cover` - generates HTML coverage report
- **Format and vet**: `just check` - formats and vets code (runs before tests
  automatically)

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

- `Instrument` interface: Core abstraction requiring Read, Write, WriteString,
  Command, and Query methods. Command and Query take `context.Context` as their
  first parameter. Sub-interfaces `Commander`, `Querier`, and `StringWriter`
  allow accepting narrower types.
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

Each class package defines Go interfaces for capability groups (Base,
StdFunc, Trigger, etc.) that drivers must implement. All interface methods
that communicate with instruments take `context.Context` as first parameter.

#### Driver Implementation Pattern

Drivers live under `<class>/<manufacturer>/<model>/` (e.g.,
`fgen/keysight/key33220/`). All drivers follow this structure:

- `01_<model>.go`: Package doc, Driver struct, Channel struct, `New()` constructor
- `04_base.go` and subsequent numbered files: One file per IVI capability group

The Driver struct always follows this pattern:

```go
type Driver struct {
    inst     ivi.Instrument
    Channels []Channel
    ivi.Inherent  // Embedded for inherent capabilities
}
```

The `New(inst ivi.Instrument, reset bool) (*Driver, error)` constructor creates
channels, populates `InherentBase` metadata, and calls `ivi.NewInherent()`.
When `reset` is true, the constructor uses `context.Background()` internally.

Channel structs hold an `ivi.Instrument` reference and a channel `name` string,
implementing per-channel capability interfaces.

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
