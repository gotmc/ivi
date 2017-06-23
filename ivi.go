// Copyright (c) 2017 The ivi developers. All rights reserved.
// Project site: https://github.com/gotmc/ivi
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package ivi

import (
	"errors"
	"fmt"
	"io"
	"log"
)

type Instrument interface {
	io.ReadWriteCloser
	WriterString(s string) (n int, err error)
	Query(s string) (value string, err error)
}

type InterfaceType int

const (
	Usbtmc InterfaceType = iota
	Tcpip
	Asrl
)

var interfaceDescription = map[InterfaceType]string{
	Usbtmc: "USBTMC",
	Tcpip:  "TCP-IP",
	Asrl:   "Serial",
}

func (i InterfaceType) String() string {
	return interfaceDescription[i]
}

// A map of registered matchers for searching.
var drivers = make(map[InterfaceType]Driver)

// Driver defines the behavior required by types that want
// to implement a new search type.
type Driver interface {
	Open(address string) (Instrument, error)
}

// Register is called to register a driver for use by the program.
func Register(interfaceType InterfaceType, driver Driver) {
	if _, exists := drivers[interfaceType]; exists {
		log.Fatalln(interfaceType, "Driver already registered")
	}

	log.Println("Register", interfaceType, "driver")
	drivers[interfaceType] = driver
}

func NewInstrument(address string) (Instrument, error) {
	interfaceType, err := determineInterfaceType(address)
	if err != nil {
		return nil, errors.New("Problem determining interface type in address.")
	}
	driver, exists := drivers[interfaceType]
	if !exists {
		return nil, fmt.Errorf("The %s interface hasn't been registered.", interfaceType)
	}
	return driver.Open(address)
}

func determineInterfaceType(address string) (InterfaceType, error) {
	return Usbtmc, nil
}
