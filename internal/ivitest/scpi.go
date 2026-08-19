// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivitest

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ErrMalformedSCPI indicates a command string is not a well-formed SCPI
// program message.
var ErrMalformedSCPI = errors.New("malformed SCPI")

// headerSegment matches one mnemonic in a command header. Trailing digits are
// part of the mnemonic, so channel-suffixed nodes such as OUTP1, SOUR2, and
// TRIG1 are accepted.
var headerSegment = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)

// commonCommand matches an IEEE 488.2 common command such as *RST or *IDN?.
var commonCommand = regexp.MustCompile(`^\*[A-Za-z]+$`)

// unitSuffix matches the suffix that may follow a numeric parameter, as in
// the "VPP" of "VOLT 2.5 VPP".
var unitSuffix = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9/]*$`)

// unformattedVerb matches a format verb left in a command string, which means
// a format string reached the instrument with an argument missing.
var unformattedVerb = regexp.MustCompile(`%[-+# 0-9.]*[a-zA-Z]`)

// CheckSCPI reports whether cmd is a well-formed SCPI program message.
//
// It validates structure only: that the header is a sequence of legal
// mnemonics, that exactly one space separates the header from its parameters,
// and that parameters are comma separated. It cannot know whether a given
// header exists on a given instrument or whether it does what the caller
// intended, so a nil return means "this could parse", not "this is correct".
//
// The check that matters most in practice is the parameter rule. A space
// inside what was meant to be a single parameter, as in "INST Output 1",
// produces a syntax error on every SCPI instrument, and no programming guide
// is needed to know that.
func CheckSCPI(cmd string) error {
	if strings.TrimSpace(cmd) == "" {
		return fmt.Errorf("%w: empty command", ErrMalformedSCPI)
	}

	if strings.Contains(cmd, "%!") {
		return fmt.Errorf(
			"%w: contains a formatting error, so a format string was given "+
				"the wrong arguments", ErrMalformedSCPI,
		)
	}

	if loc := unformattedVerb.FindString(cmd); loc != "" {
		return fmt.Errorf(
			"%w: contains the unformatted verb %q", ErrMalformedSCPI, loc,
		)
	}

	for unit := range strings.SplitSeq(cmd, ";") {
		if err := checkMessageUnit(strings.TrimSpace(unit)); err != nil {
			return err
		}
	}

	return nil
}

// checkMessageUnit validates one program message unit, meaning one header
// with its optional parameters.
func checkMessageUnit(unit string) error {
	if unit == "" {
		return fmt.Errorf("%w: empty program message unit", ErrMalformedSCPI)
	}

	header, params, hasParams := strings.Cut(unit, " ")

	if err := checkHeader(header); err != nil {
		return err
	}

	if !hasParams {
		return nil
	}

	if strings.HasPrefix(params, " ") {
		return fmt.Errorf(
			"%w: more than one space separates the header from its "+
				"parameters in %q", ErrMalformedSCPI, unit,
		)
	}

	return checkParams(params, unit)
}

// checkHeader validates a command header, with its optional leading colon and
// optional trailing question mark.
func checkHeader(header string) error {
	header = strings.TrimSuffix(header, "?")
	header = strings.TrimPrefix(header, ":")

	if commonCommand.MatchString(header) {
		return nil
	}

	if header == "" {
		return fmt.Errorf("%w: empty header", ErrMalformedSCPI)
	}

	for segment := range strings.SplitSeq(header, ":") {
		if !headerSegment.MatchString(segment) {
			return fmt.Errorf(
				"%w: %q is not a legal mnemonic in header %q",
				ErrMalformedSCPI, segment, header,
			)
		}
	}

	return nil
}

// checkParams validates the comma separated parameter list of one message
// unit.
func checkParams(params, unit string) error {
	for param := range strings.SplitSeq(params, ",") {
		param = strings.TrimSpace(param)
		if param == "" {
			return fmt.Errorf(
				"%w: empty parameter in %q", ErrMalformedSCPI, unit,
			)
		}

		if err := checkParam(param, unit); err != nil {
			return err
		}
	}

	return nil
}

// checkParam validates a single parameter. A parameter holds no spaces of its
// own, with one exception: a numeric value may be followed by a unit suffix,
// as in "2.5 VPP". A space after a non-numeric token means what was meant as
// one parameter is really two, which is the shape of a channel name that was
// never a legal SCPI token.
func checkParam(param, unit string) error {
	if strings.HasPrefix(param, `"`) || strings.HasPrefix(param, "'") {
		return nil
	}

	if strings.HasPrefix(param, "(@") {
		return nil
	}

	fields := strings.Fields(param)
	if len(fields) == 1 {
		return nil
	}

	if len(fields) != 2 {
		return fmt.Errorf(
			"%w: parameter %q in %q holds more than one space",
			ErrMalformedSCPI, param, unit,
		)
	}

	if _, err := strconv.ParseFloat(fields[0], 64); err != nil {
		return fmt.Errorf(
			"%w: parameter %q in %q holds a space, which is only legal "+
				"between a numeric value and its unit suffix, and %q is not "+
				"numeric", ErrMalformedSCPI, param, unit, fields[0],
		)
	}

	if !unitSuffix.MatchString(fields[1]) {
		return fmt.Errorf(
			"%w: %q is not a legal unit suffix in %q",
			ErrMalformedSCPI, fields[1], unit,
		)
	}

	return nil
}
