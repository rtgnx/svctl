package main

import (
	"log"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/rtgnx/svctl/internal/agent"
	"github.com/rtgnx/svctl/internal/util"
	"github.com/rtgnx/svctl/pkg/runit"
)

func main() {
	app := cli.App("svagent", "svdir monitoring agent")
	var (
		endpoint = app.StringOpt("addr", "http://127.0.0.1:8080", "dest addr to push data to")
		svroots  = app.StringsOpt("roots", []string{util.Env("SVDIR", runit.DefaultSVDIR)}, "svdir paths")
	)

	app.Action = func() {
		monitor := agent.NewMonitor(*svroots...)
		go monitor.Watch()

		if err := monitor.Report(*endpoint); err != nil {
			log.Fatal(err)
		}
	}

	app.Run(os.Args)
}
