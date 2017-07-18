// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kikusuipmx

import (
	"errors"
	"log"

	"github.com/gotmc/ivi/dcpwr"
)

// CurrentLimit determines the output current limit. The units are Amps.
// CurrentLimit implements the getter for the read-write IviDCPwrBase Attribute
// Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) CurrentLimit() (float64, error) {
	return ch.queryFloat64("CURR?\n")
}

// SetCurrentLimit specifies the output current limit. The units are Amps.
// SetCurrentLimit implements the setter for the read-write IviDCPwrBase
// Attribute Current Limit described in Section 4.2.1 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) SetCurrentLimit(limit float64) error {
	if ch.currentLimitBehavior == dcpwr.Regulate {
		return ch.Set("CURR %f;:CURR:PROT MAX\n", limit)
	} else if ch.currentLimitBehavior == dcpwr.Trip {
		return ch.Set("CURR %f;:CURR:PROT %f\n", limit, limit)
	}
	return errors.New("current limit behavior not set")
}

// CurrentLimitBehavior determines the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. The E3631A only supports the CurrentRegulate behavior.
// CurrentLimitBehavior implements the getter for the read-write IviDCPwrBase
// Attribute Current Limit Behavior described in Section 4.2.2 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) CurrentLimitBehavior() (dcpwr.CurrentLimitBehavior, error) {
	return ch.currentLimitBehavior, nil
}

// SetCurrentLimitBehavior specifies the behavior of the power supply when the
// output current is equal to or greater than the value of the Current Limit
// attribute. Since the PMX series always has the OCP enabled, if the current
// limit behavior is set to regulate, the OCP is set to MAX (110% of the max
// current value of the device).  Whereas, setting the current limit behavior
// to trip, the OCP is set equal to the current limit.  CurrentLimitBehavior
// implements the getter for the read-write IviDCPwrBase Attribute Current
// Limit Behavior described in Section 4.2.2 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetCurrentLimitBehavior(behavior dcpwr.CurrentLimitBehavior) error {
	if behavior == dcpwr.Regulate {
		ch.currentLimitBehavior = dcpwr.Regulate
		return ch.Set("CURR:PROT MAX\n")
	} else if behavior == dcpwr.Trip {
		ch.currentLimitBehavior = dcpwr.Trip
		limit, err := ch.queryFloat64("CURR?\n")
		if err != nil {
			return err
		}
		return ch.Set("CURR:PROT %f\n", limit)
	}
	return errors.New("unknown current limit behavior")
}

// OutputEnabled determines if all three output channels are enabled or
// disabled.  OutputEnabled is the getter for the read-write IviDCPwrBase
// Attribute Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr
// Class Specification.
func (ch *Channel) OutputEnabled() (bool, error) {
	return ch.queryBool("OUTP?\n")
}

// SetOutputEnabled sets all the output channel to enabled or disabled.
// SetOutputEnabled is the setter for the read-write IviDCPwrBase Attribute
// Output Enabled described in Section 4.2.3 of IVI-4.4: IviDCPwr Class
// Specification.
func (ch *Channel) SetOutputEnabled(v bool) error {
	if v {
		return ch.Set("OUTP 1\n")
	}
	return ch.Set("OUTP 0\n")
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

// OVPEnabled always returns true since the PMX series always has the OVP set.
// OVPEnabled is the getter for the read-write IviFgenBase Attribute OVP
// Enabled described in Section 4.2.4 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPEnabled() (bool, error) {
	return true, nil
}

// SetOVPEnabled always returns nil, since the PMX series always has the
// voltage protection set.  SetOVPEnabled is the setter for the read-write
// IviFgenBase Attribute OVP Enabled described in Section 4.2.4 of IVI-4.4:
// IviDCPwr Class Specification.
func (ch *Channel) SetOVPEnabled(v bool) error {
	return nil
}

// OVPLimit returns the current Over Voltage Protection (OVP) value. OPVLimit
// is the getter for the read-write IviDWPwrBase Attribute OVP Limit described
// in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) OVPLimit() (float64, error) {
	return ch.queryFloat64("VOLT:PROT?\n")
}

// SetOVPLimit returns an error since the E3631A doesn't support OVP.
// SetOVPLimit is the setter for the read-write IviDCPwrBase Attribute OVP
// Limit described in Section 4.2.5 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetOVPLimit(limit float64) error {
	return ch.Set("VOLT:PROT %f\n", limit)
}

// VoltageLevel reads the specified voltage level the DC power supply attempts
// to generate. The units are Volts.  VoltageLevel is the getter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) VoltageLevel() (float64, error) {
	return ch.queryFloat64("VOLT?\n")
}

// SetVoltageLevel specifies the voltage level the DC power supply attempts
// to generate. The units are Volts.  SetVoltageLevel is the setter for the
// read-write IviDCPwrBase Attribute Voltage Level described in Section 4.2.6
// of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) SetVoltageLevel(level float64) error {
	log.Printf("About to set voltage level to %f", level)
	return ch.Set("VOLT %f\n", level)
}

// OutputCount returns the number of available output channels. OutputCount is
// the getter for the read-only IviDCPwrBase Attribute Output Channel Count
// described in Section 4.2.7 of IVI-4.4: IviDCPwr Class Specification.
func (dcpwr *KikusuiPMX) OutputCount() int {
	return len(dcpwr.Channels)
}

// Name returns the name of the output channel. Name is the getter for the
// read-only IviDCPwrBase Attribute Output Channel Name described in Section
// 4.2.9 of IVI-4.4: IviDCPwr Class Specification.
func (ch *Channel) Name() string {
	return ch.name
}
