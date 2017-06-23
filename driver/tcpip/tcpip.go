// Project site: https://github.com/gotmc/ivi
// Copyright (c) 2017 The ivi developers. All rights reserved.
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tcpip

import (
	"fmt"
	"net"

	"github.com/gotmc/ivi"
)

// tcpipDriver implements the ivi.Driver interface.
type Driver struct{}

type Connection struct {
	conn net.Conn
}

func (d Driver) Open(address string) (ivi.Instrument, error) {
	var c Connection
	tcpAddress, err := getTcpAddress(address)
	if err != nil {
		return nil, fmt.Errorf("%s is not a TCPIP VISA resource address: %s", address, err)
	}
	c.conn, err = net.Dial("tcp", tcpAddress)
	if err != nil {
		return nil, fmt.Errorf("Problem connecting to TCP instrument at %s: %s", tcpAddress, err)
	}
	return &c, nil
}

func (c *Connection) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (c *Connection) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}

func getTcpAddress(address string) (string, error) {
	return "127.0.0.1:5025", nil
}

// init registers the driver with the program.
func init() {
	var driver Driver
	ivi.Register(ivi.Tcpip, driver)
}
