# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Development Commands

### Building and Testing

- **Run tests**: `just unit` - formats, vets, and runs unit tests with coverage
- **Integration tests**: `just int` - runs integration tests
- **Linting**: `just lint` - runs golangci-lint with the project's configuration
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
instrument manufacturers.

### Key Components

#### Core Interface Layer (`ivi.go`, `inherent.go`)

- `Instrument` interface: Core abstraction for all test instruments requiring
  Read/Write/Command/Query methods
- `Inherent` struct: Base capabilities common to all IVI instruments (reset,
  clear, identification)
- Inherent capabilities follow IVI-3.2 specification for standardized instrument
  behavior

#### Instrument Class Packages

Each directory represents an IVI instrument class with numbered files organized
by capability groups:

**DMM (Digital Multimeter)** - `dmm/`

- Implements IVI-4.2 IviDmm Class Specification
- Base capabilities in `04_base.go`, measurements in `05_ac_msrmnt.go`, etc.
- Keysight 3446x family driver in `dmm/keysight/key3446x/`

**Function Generators** - `fgen/`

- Implements IVI-4.3 IviFgen Class Specification
- Standard functions (`05_std_func.go`), arbitrary waveforms (`06_arb_wfm.go`),
  triggers (`09_trigger.go`)
- Keysight 33220A driver in `fgen/keysight/key33220/`

**DC Power Supplies** - `dcpwr/`

- Implements IVI-4.4 IviDCPwr Class Specification
- Base functions, triggers, measurements organized by capability groups
- Keysight E36xx series driver in `dcpwr/keysight/e36xx/`

**Oscilloscopes** - `scope/`

- Implements IVI-4.1 IviScope Class Specification
- Waveform measurements, triggers, acquisition modes
- Keysight InfiniiVision driver in `scope/keysight/infiniivision/`

#### Driver Implementation Pattern

All instrument drivers follow consistent patterns:

- `01_<instrument>.go`: Package documentation and main driver struct
- `04_base.go`: Base capability group implementation
- Numbered files correspond to IVI capability groups from specifications
- Each driver implements relevant interfaces and embeds `ivi.Inherent`
- Channel-based instruments have separate Channel structs for repeated capabilities

### Design Philosophy

- **Idiomatic Go over strict IVI compliance**: Uses Go enums and type safety
  instead of magic numbers
- **No state caching**: All attributes read directly from instruments for
  reliability
- **Interface-based**: Drivers implement capability interfaces, enabling
  instrument interchangeability
- **Capability groups**: Code organized by IVI specification capability groups
  for maintainability

### Dependencies

- `github.com/gotmc/query`: SCPI query utilities
- `github.com/gotmc/convert`: Data conversion helpers
- Go 1.21+ required

## Code Style and Conventions

### File Organization

- Files prefixed with numbers (01*, 04*, 05\_) correspond to IVI specification
  sections
- Capability groups determine file boundaries (Base, AC Measurements, Triggers,
  etc.)
- Driver packages use manufacturer/model naming: `keysight/key3446x/`,
  `srs/ds345/`

### Interface Compliance Verification

All drivers include interface compliance checks using blank identifiers:

```go
var _ dmm.Base = (*Driver)(nil)
var _ dmm.BaseChannel = (*Channel)(nil)
```

### Error Handling

- Package-specific error variables defined in main package files
- Standard Go error handling patterns throughout
- IVI-specific errors like `ErrNotImplemented` for unsupported features

### Code Quality

- golangci-lint configuration in `.golangci.yaml` includes staticcheck, gosec,
  misspell
- Comprehensive linting rules ensure code quality and security
- Test coverage reporting available via `just cover`

