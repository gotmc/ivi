package dmm

// Base may not be needed, but the idea is that there are some functions that
// could be abstracted out to the general DMM.
type Base interface {
	DCVolts() (v float64, err error)
	ACVolts() (v float64, err error)
}
