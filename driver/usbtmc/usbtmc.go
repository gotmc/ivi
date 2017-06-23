// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package usbtmc

import (
	"github.com/gotmc/ivi"
	"github.com/gotmc/usbtmc"
)

// Driver implements the ivi.Driver interface.
type Driver struct {
}

type Connection struct {
	ctx  *usbtmc.Context
	inst *usbtmc.Instrument
}

func (d Driver) Open(address string) (ivi.Instrument, error) {
	var c Connection
	c.ctx = usbtmc.NewContext()
	inst, err := c.ctx.NewInstrument(address)
	c.inst = inst
	return &c, err
}

func (c *Connection) Read(p []byte) (n int, err error) {
	return c.inst.Read(p)
}

func (c *Connection) Write(p []byte) (n int, err error) {
	return c.inst.Write(p)
}

func (c *Connection) Close() error {
	err := c.inst.Close()
	if err != nil {
		return err
	}
	return c.ctx.Close()
}

// init registers the driver with the program.
func init() {
	var driver Driver
	ivi.Register(ivi.Usbtmc, driver)
}
