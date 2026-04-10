// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotmc/query"
)

type idPart int

const (
	mfrID idPart = iota
	modelID
	snID
	fwrID
)

// Inherent provides the inherent capabilities for all IVI instruments.
type Inherent struct {
	inst Instrument
	InherentBase
}

// InherentBase provides the exported properties for the inherent capabilities
// of all IVI instruments.
type InherentBase struct {
	ClassSpecRevision         string
	IDNString                 string
	GroupCapabilities         []string
	SupportedInstrumentModels []string
	SupportedBusInterfaces    []string
	ClassSpecMajorVersion     int
	ClassSpecMinorVersion     int
	ResetDelay                time.Duration
	ClearDelay                time.Duration
	ReturnToLocal             bool // Whether to return to local control on Close/Disable
}

// NewInherent creates a new Inherent struct using the given Instrument
// interface and the InherentBase struct. Note: callers should explicitly set
// ReturnToLocal in InherentBase (default Go zero value is false).
func NewInherent(inst Instrument, base InherentBase) Inherent {
	return Inherent{
		inst:         inst,
		InherentBase: base,
	}
}

// FirmwareRevision queries the instrument and returns the firmware revision of
// the instrument. FirmwareRevision is the getter for the read-only inherent
// attribute Instrument Firmware Revision described in Section 5.18 of IVI-3.2:
// Inherent Capabilities Specification.
func (inherent *Inherent) FirmwareRevision(ctx context.Context) (string, error) {
	return inherent.queryIdentification(ctx, fwrID)
}

// InstrumentManufacturer queries the instrument and returns the manufacturer
// of the instrument. InstrumentManufacturer is the getter for the read-only
// inherent attribute Instrument Manufacturer described in Section 5.19 of
// IVI-3.2: Inherent Capabilities Specification.
func (inherent *Inherent) InstrumentManufacturer(ctx context.Context) (string, error) {
	return inherent.queryIdentification(ctx, mfrID)
}

// InstrumentModel queries the instrument and returns the model of the
// instrument.  InstrumentModel is the getter for the read-only inherent
// attribute Instrument Model described in Section 5.20 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) InstrumentModel(ctx context.Context) (string, error) {
	return inherent.queryIdentification(ctx, modelID)
}

// InstrumentSerialNumber queries the instrument and returns the S/N of the
// instrument.
func (inherent *Inherent) InstrumentSerialNumber(ctx context.Context) (string, error) {
	return inherent.queryIdentification(ctx, snID)
}

// Reset resets the instrument.
func (inherent *Inherent) Reset(ctx context.Context) error {
	if err := inherent.inst.Command(ctx, "*rst"); err != nil {
		return err
	}
	// Wait for the device to finish resetting.
	select {
	case <-time.After(inherent.ResetDelay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Clear clears the instrument.
func (inherent *Inherent) Clear(ctx context.Context) error {
	if err := inherent.inst.Command(ctx, "*cls"); err != nil {
		return err
	}
	// Wait for the device to finish clearing.
	select {
	case <-time.After(inherent.ClearDelay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Disable places the instrument in a quiescent state as quickly as possible.
// If ReturnToLocal is true, this method also returns the instrument to local
// control by sending the SYST:LOC command, allowing the front panel to regain
// control. Disable provides the method described in Section 6.4 of IVI-3.2:
// Inherent Capabilities Specification.
func (inherent *Inherent) Disable(ctx context.Context) error {
	if !inherent.ReturnToLocal {
		// Skip sending local control command if not requested
		return nil
	}

	// Send the system local command to return control to the front panel
	// This addresses the issue where instruments remain in remote mode
	// after the program terminates.
	err := inherent.inst.Command(ctx, "SYST:LOC")
	if err != nil {
		// If SYST:LOC fails, try the alternative SCPI command
		fallbackErr := inherent.inst.Command(ctx, "SYSTem:LOCal")
		if fallbackErr != nil {
			// Return the original error if both fail
			return err
		}
		// Fallback succeeded
		return nil
	}

	return err
}

func (inherent *Inherent) queryIdentification(
	ctx context.Context, part idPart,
) (string, error) {
	s, err := query.String(ctx, inherent.inst, "*IDN?")
	if err != nil {
		return "", err
	}

	return parseIdentification(strings.TrimSpace(s), part)
}

func parseIdentification(idn string, part idPart) (string, error) {
	const numIdentificationParts = 4

	parts := strings.Split(idn, ",")

	if len(parts) != numIdentificationParts {
		return "", fmt.Errorf("idn string (`%s`) could not be split in four", idn)
	}

	return parts[part], nil
}

// SetReturnToLocal controls whether the instrument returns to local control
// when Disable() or Close() is called. Default is true.
func (inherent *Inherent) SetReturnToLocal(enabled bool) {
	inherent.ReturnToLocal = enabled
}

// IsReturnToLocal returns whether the instrument will return to local control
// when Disable() or Close() is called.
func (inherent *Inherent) IsReturnToLocal() bool {
	return inherent.ReturnToLocal
}

// Close properly shuts down the instrument connection by returning it to local
// control and then closing the underlying transport connection. This method
// should be called when finished with an instrument to ensure it returns to
// local control for front panel operation.
func (inherent *Inherent) Close() error {
	// First, return the instrument to local control
	disableErr := inherent.Disable(context.Background())

	// Close the underlying transport connection
	closeErr := inherent.inst.Close()
	// Prioritize close errors over disable errors
	if closeErr != nil {
		return closeErr
	}

	return disableErr
}
