// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilente36xx

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotmc/ivi"
)

func queryBool(q ivi.Querier, query string) (bool, error) {
	s, err := q.Query(query)
	if err != nil {
		return false, err
	}
	switch s {
	case "OFF":
		return false, nil
	case "ON":
		return true, nil
	default:
		return false, fmt.Errorf("could not determine boolean status from %s", s)
	}
}

func queryFloat64(q ivi.Querier, query string) (float64, error) {
	s, err := q.Query(query)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(strings.TrimSpace(s), 64)
}

func queryString(q ivi.Querier, query string) (string, error) {
	return q.Query(query)
}

func setFloat64(sw ivi.StringWriter, cmd string, value float64) error {
	send := fmt.Sprintf(cmd, value)
	_, err := sw.WriteString(send)
	return err
}
