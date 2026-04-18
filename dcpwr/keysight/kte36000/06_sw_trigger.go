// Copyright (c) 2017-2026 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package kte36000

import (
	"fmt"

	"github.com/gotmc/ivi/dcpwr"
)

// SendSoftwareTrigger sends the IEEE 488.2 common *TRG command to trigger a
// previously initiated transient once the trigger source has been set to
// [dcpwr.TriggerSourceSoftware]. SendSoftwareTrigger returns
// [dcpwr.ErrTriggerNotSoftware] if any channel's trigger source is not set
// to software.
//
// SendSoftwareTrigger implements the IviDCPwrSoftwareTrigger function
// described in Section 6.2.1 of IVI-4.4: IviDCPwr Class Specification.
func (d *Driver) SendSoftwareTrigger() error {
	ctx, cancel := d.newContext()
	defer cancel()

	for i := range d.channels {
		src, err := d.channels[i].TriggerSource()
		if err != nil {
			return fmt.Errorf("SendSoftwareTrigger: %w", err)
		}

		if src != dcpwr.TriggerSourceSoftware {
			return fmt.Errorf(
				"SendSoftwareTrigger: channel %q: %w",
				d.channels[i].name, dcpwr.ErrTriggerNotSoftware,
			)
		}
	}

	return d.inst.Command(ctx, "*TRG")
}
