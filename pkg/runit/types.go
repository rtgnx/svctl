package runit

import "time"

const (
	DefaultSVDIR = "/var/service"
	StateDown    = 0
	StateUp      = 1
	StateFinish  = 2
)

type Service struct {
	Name       string
	Pid        int
	Timestamp  time.Time
	Duration   time.Duration
	State      int
	NormallyUp bool
	Want       int
}

type SVDIR struct {
	Path     string
	Hostname string
	Services []Service
}
