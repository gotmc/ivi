// Copyright (c) 2017-2020 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"fmt"
)

// Set formats according to a format specifier and then writes the resulting
// string to the given StringWriter interface.
func Set(sw StringWriter, format string, a ...interface{}) error {
	cmd := format
	if a != nil {
		cmd = fmt.Sprintf(format, a...)
	}
	_, err := sw.WriteString(cmd)
	return err
}

// QueryID queries the identity of the instrument.
func QueryID(q Querier) (string, error) {
	return q.Query("*IDN?\n")
}
