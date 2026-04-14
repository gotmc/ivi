// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"context"
	"fmt"
	"slices"
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
	inst    Transport
	timeout time.Duration
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

// NewInherent creates a new Inherent struct using the given Transport
// interface, InherentBase struct, and timeout. If timeout is zero,
// DefaultTimeout is used. Note: callers should explicitly set ReturnToLocal in
// InherentBase (default Go zero value is false).
func NewInherent(inst Transport, base InherentBase, timeout time.Duration) Inherent {
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return Inherent{
		inst:         inst,
		timeout:      timeout,
		InherentBase: base,
	}
}

// Timeout returns the timeout used for instrument I/O operations.
func (inherent *Inherent) Timeout() time.Duration {
	return inherent.timeout
}

// newContext creates a context with the driver's configured timeout.
func (inherent *Inherent) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), inherent.timeout)
}

// CheckID queries the instrument for its identification string (*IDN?) and
// verifies that the instrument model is in the list of SupportedInstrumentModels.
// On success, the IDNString field is populated with the full *IDN? response and
// the trimmed model string is returned. CheckID implements the IdQuery behavior
// described in Section 6.16 of IVI-3.2: Inherent Capabilities Specification.
func (inherent *Inherent) CheckID() (string, error) {
	ctx, cancel := inherent.newContext()
	defer cancel()

	idn, err := query.String(ctx, inherent.inst, "*IDN?")
	if err != nil {
		return "", fmt.Errorf("error querying instrument identity: %w", err)
	}

	inherent.IDNString = strings.TrimSpace(idn)

	model, err := parseIdentification(inherent.IDNString, modelID)
	if err != nil {
		return "", err
	}

	model = strings.TrimSpace(model)

	if !slices.Contains(inherent.SupportedInstrumentModels, model) {
		return model, fmt.Errorf(
			"%w: %q is not supported by this driver",
			ErrUnsupportedModel,
			model,
		)
	}

	return model, nil
}

// FirmwareRevision queries the instrument and returns the firmware revision of
// the instrument. FirmwareRevision is the getter for the read-only inherent
// attribute Instrument Firmware Revision described in Section 5.18 of IVI-3.2:
// Inherent Capabilities Specification.
func (inherent *Inherent) FirmwareRevision() (string, error) {
	return inherent.queryIdentification(fwrID)
}

// InstrumentManufacturer queries the instrument and returns the manufacturer
// of the instrument. InstrumentManufacturer is the getter for the read-only
// inherent attribute Instrument Manufacturer described in Section 5.19 of
// IVI-3.2: Inherent Capabilities Specification.
func (inherent *Inherent) InstrumentManufacturer() (string, error) {
	return inherent.queryIdentification(mfrID)
}

// InstrumentModel queries the instrument and returns the model of the
// instrument.  InstrumentModel is the getter for the read-only inherent
// attribute Instrument Model described in Section 5.20 of IVI-3.2: Inherent
// Capabilities Specification.
func (inherent *Inherent) InstrumentModel() (string, error) {
	return inherent.queryIdentification(modelID)
}

// InstrumentSerialNumber queries the instrument and returns the S/N of the
// instrument.
func (inherent *Inherent) InstrumentSerialNumber() (string, error) {
	return inherent.queryIdentification(snID)
}

// Reset resets the instrument.
func (inherent *Inherent) Reset() error {
	ctx, cancel := inherent.newContext()
	defer cancel()

	if err := inherent.inst.Command(ctx, "*rst"); err != nil {
		return err
	}
	// Wait for the device to finish resetting.
	time.Sleep(inherent.ResetDelay)

	return nil
}

// Clear clears the instrument.
func (inherent *Inherent) Clear() error {
	ctx, cancel := inherent.newContext()
	defer cancel()

	if err := inherent.inst.Command(ctx, "*cls"); err != nil {
		return err
	}
	// Wait for the device to finish clearing.
	time.Sleep(inherent.ClearDelay)

	return nil
}

// Disable places the instrument in a quiescent state as quickly as possible.
// If ReturnToLocal is true, this method also returns the instrument to local
// control by sending the SYST:LOC command, allowing the front panel to regain
// control. Disable provides the method described in Section 6.4 of IVI-3.2:
// Inherent Capabilities Specification.
//
// Errors from the local control command are intentionally ignored because not
// all instruments support SYST:LOC, and a failure to return to local should
// not prevent the instrument from being closed.
func (inherent *Inherent) Disable() error {
	if !inherent.ReturnToLocal {
		return nil
	}

	ctx, cancel := inherent.newContext()
	defer cancel()

	// Best-effort return to local control so the front panel regains control.
	_ = inherent.inst.Command(ctx, "SYST:LOC")

	return nil
}

func (inherent *Inherent) queryIdentification(part idPart) (string, error) {
	ctx, cancel := inherent.newContext()
	defer cancel()

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
	disableErr := inherent.Disable()

	// Close the underlying transport connection
	closeErr := inherent.inst.Close()
	// Prioritize close errors over disable errors
	if closeErr != nil {
		return closeErr
	}

	return disableErr
}
