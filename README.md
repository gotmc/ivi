# ivi
Go-based implementation of the Interchangeable Virtual Instrument (IVI) standard

## Overview

The IVI standard wasn't written for Go, so the [ivi][] package doesn't
fully implement the IVI standard. However, the value of the [ivi][]
package is in providing a standardized API for test equipment. For
instance, instead of sending SCPI commands to an Agilent 33220A function
generator, which would be different than the SCPI commands for an SRS
DS345 function generator, the [ivi][] package provides a common API for
function generators.
