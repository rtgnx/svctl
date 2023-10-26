package agent

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rtgnx/svctl/pkg/proto"
	"github.com/rtgnx/svctl/pkg/runit"
)

type Monitor struct {
	roots  []string
	svdirs []runit.SVDIR
	sync.Mutex
}

func NewMonitor(roots ...string) *Monitor {
	return &Monitor{roots: roots}
}

func (m *Monitor) Watch() []runit.SVDIR {
	m.svdirs = make([]runit.SVDIR, len(m.roots))

	for i, root := range m.roots {
		svdir, err := runit.ReadSVDIR(root)
		if err != nil {
			log.Print(err)
			continue
		}
		m.Lock()
		m.svdirs[i] = svdir
		m.Unlock()
	}
	return m.svdirs
}

func (m *Monitor) Report(endpoint string) error {
	endpointURL, err := url.Parse(endpoint)

	if err != nil {
		return err
	}

	endpointURL = endpointURL.JoinPath(proto.APIV1Push)

	if err != nil {
		return err
	}
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	for {
		svdirs := m.Watch()
		m.Lock()

		msg := proto.Message{
			Hostname:  hostname,
			SVDIRs:    svdirs,
			Timestamp: time.Now(),
		}

		b, err := json.Marshal(msg)

		if err != nil {
			log.Print(err)
		}

		_, err = http.Post(endpointURL.String(), echo.MIMEApplicationJSON, bytes.NewBuffer(b))

		if err != nil {
			log.Print(err)
		}

		m.Unlock()
		<-time.After(time.Second * 2)
	}

}
