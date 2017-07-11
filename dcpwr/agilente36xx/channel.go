// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package agilente36xx

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gotmc/ivi"
)

// CurrentLimitBehavior provides the defined values for the Current Limit
// Behavior defined in Section 4.2.2 of IVI-4.4: IviDCPwr Class Specification.
type CurrentLimitBehavior int

const (
	CurrentTrip CurrentLimitBehavior = iota
	CurrentRegulate
)

type Channel struct {
	id   int
	name string
	inst ivi.Instrument
}

// String implements the stringer interface for channel.
func (ch *Channel) String() string {
	return ch.name
}

// CurrentLimit specifies the output current limit. The units are Amp.s
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	return ch.queryFloat64(fmt.Sprintf("INST %s;:SOUR:CURR?\n", ch.name))
}

// CurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute.  CurrentLimitBehavior implements the getter for the read-write
// IviDCPwrBase Attribute Current Limit Behavior described in Section 4.2.2 of
// IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (CurrentLimitBehavior, error) {
	// FIXME(mdr): Need to implement!
	return CurrentRegulate, nil
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.  OutputEnabled is the getter for the read-write IviDCPwrBase
// Attribute Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return ch.queryBool("OUTP?\n")
}

// SetOutputEnabled sets all three output channels to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviFgenBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	var send string
	if v {
		send = "OUTP ON\n"
	} else {
		send = "OUTP OFF\n"
	}
	_, err := ch.inst.WriteString(send)
	return err
}

// DisableOutput is a convenience function for setting the Output Enabled
// attribute to false.
func (ch *Channel) DisableOutput() error {
	return ch.SetOutputEnabled(false)
}

// EnableOutput is a convenience function for setting the Output Enabled
// attribute to true.
func (ch *Channel) EnableOutput() error {
	return ch.SetOutputEnabled(true)
}

// OVPEnabled always returns false for the E3631A since it doesn't have OVP.
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	return false, nil
}

// SetOVPEnabled always returns an error for the E3631A since it doesn't have
// OVP.  SetOVPEnabled is the setter for the read-write IviFgenBase Attribute
// OVP Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	return errors.New("OVP not supported")
}

// OVPLimit returns an error, since the E3631A doesn't support OVP. OVPLimit is
// the getter for the read-write IviDWPwrBase Attribute OVP Limit described in
// Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	return 0, errors.New("OVP not supported")
}

// SetOVPLimit returns an error since the E3631A doesn't support OVP.
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	return errors.New("OVP not supported")
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate. The units are Volts.  VoltageLevel is the getter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	cmd := fmt.Sprintf("APPL? %s", ch.name)
	s, err := ch.inst.Query(cmd)
	if err != nil {
		return 0.0, err
	}
	ret := strings.Split(s, ",")
	return strconv.ParseFloat(ret[0], 64)
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate. The units are Volts.  SetVoltageLevel is the setter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetVoltageLevel(amp float64) error {
	cmd := fmt.Sprintf("APPL %s, %f\n", ch.name, amp)
	_, err := ch.inst.WriteString(cmd)
	return err
}

func (ch *Channel) setFloat64(cmd string, value float64) error {
	return setFloat64(ch.inst, cmd, value)
}

func (ch *Channel) queryBool(query string) (bool, error) {
	return queryBool(ch.inst, query)
}

func (ch *Channel) queryFloat64(query string) (float64, error) {
	return queryFloat64(ch.inst, query)
}

func (ch *Channel) queryString(query string) (string, error) {
	return queryString(ch.inst, query)
}
