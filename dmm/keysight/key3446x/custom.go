// Copyright (c) 2017-2025 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package key3446x

import (
	"fmt"
	"strings"

	"github.com/gotmc/query"
)

// SelectTerminals queries if the front or rear terminals are selected on the
// 34461A front panel Front/Rear switch. This switch is not programmable; this
// query reports the position of the switch, but cannot change it.
func (d *Driver) SelectedTerminals() (Terminal, error) {
	term, err := query.String(d.inst, "rout:term?")
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

type Terminal int

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
