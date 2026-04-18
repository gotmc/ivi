// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kt34400

import (
	"fmt"
	"strings"

	"github.com/gotmc/query"
)

// SelectedTerminals queries if the front or rear terminals are selected on the
// 34461A front panel Front/Rear switch. This switch is not programmable; this
// query reports the position of the switch, but cannot change it.
func (d *Driver) SelectedTerminals() (Terminal, error) {
	ctx, cancel := d.newContext()
	defer cancel()

	term, err := query.String(ctx, d.inst, "rout:term?")
	if err != nil {
		return 0, err
	}

	switch strings.TrimSpace(term) {
	case "FRON":
		return FrontTerminals, nil
	case "REAR":
		return RearTerminals, nil
	}

	return 0, fmt.Errorf("illegal terminal value: %s", term)
}

// Terminal identifies which pair of input terminals (front or rear) the
// instrument's front-panel switch has selected. This is a Truevolt-specific
// capability not covered by the IVI-4.2 DMM class specification.
type Terminal int

// Available Terminal values.
const (
	FrontTerminals Terminal = iota
	RearTerminals
)

var terminalStringer = map[Terminal]string{
	FrontTerminals: "Front terminals",
	RearTerminals:  "Rear terminals",
}

func (t Terminal) String() string {
	return terminalStringer[t]
}
