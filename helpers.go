// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// QueryBool is used to query a Querier interface and return a bool.
func QueryBool(q Querier, query string) (bool, error) {
	s, err := q.Query(query)
	if err != nil {
		return false, err
	}
	switch s {
	case "OFF", "0":
		return false, nil
	case "ON", "1":
		return true, nil
	default:
		return false, fmt.Errorf("could not determine boolean status from %s", s)
	}
}

// QueryFloat64 is used to query a Querier interface and return a float64.
func QueryFloat64(q Querier, query string) (float64, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

// QueryInt is used to query a Querier interface and return an int.
func QueryInt(q Querier, query string) (int, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 32)
	return int(i), err
}

// QueryString is used to query a Querier interface and return a string.
func QueryString(q Querier, query string) (string, error) {
	return q.Query(query)
}

// Set formats according to a format specifier and then writes the resulting
// string to the given StringWriter interface.
func Set(sw StringWriter, format string, a ...interface{}) error {
	cmd := format
	if a != nil {
		cmd = fmt.Sprintf(format, a...)
	}
	log.Printf("writing cmd: %s", cmd)
	_, err := sw.WriteString(cmd)
	return err
}

// QueryID queries the identity of the instrument.
func QueryID(q Querier) (string, error) {
	return q.Query("*IDN?\n")
}
