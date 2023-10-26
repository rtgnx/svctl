package proto

import (
	"time"

	"github.com/rtgnx/svctl/pkg/runit"
)

const (
	APIV1Push    = "/api/v1/push"
	APIV1Control = "/api/v1/control"
	APIV1State   = "/api/v1/state"
	APIV1Metrics = "/api/v1/metrics"
)

type Message struct {
	Hostname  string        `json:"hostname,omitempty"`
	SVDIRs    []runit.SVDIR `json:"data,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}
