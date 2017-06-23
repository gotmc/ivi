package dmm

type Base interface {
	DCVolts() (v float64, err error)
	ACVolts() (v float64, err error)
}
