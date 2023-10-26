package runit

import (
	"log"
	"os"
)

func ReadSVDIR(root string) (svdir SVDIR, err error) {
	svdir = SVDIR{Path: root}

	svcNames, err := ListServiceNames(root)
	if err != nil {
		return svdir, err
	}

	svdir.Services = make([]Service, len(svcNames))

	for i, name := range svcNames {
		svc, err := ReadService(root, name)
		if err != nil {
			log.Print(err)
			continue
		}

		svdir.Services[i] = svc
	}

	svdir.Hostname, err = os.Hostname()
	return svdir, err
}
