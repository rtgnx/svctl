package cmd

import (
	"fmt"
	"log"

	cli "github.com/jawher/mow.cli"
	"github.com/rtgnx/svctl/internal/util"
	"github.com/rtgnx/svctl/pkg/runit"
)

func CmdList(cmd *cli.Cmd) {

	var (
		svroot = cmd.StringOpt("svdir", util.Env("SVDIR", runit.DefaultSVDIR), "service directory")
	)

	cmd.Action = func() {
		svdir, err := runit.ReadSVDIR(*svroot)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Hostname: %s\nSVDIR: %s\n\n", svdir.Hostname, svdir.Path)

		for _, svc := range svdir.Services {
			fmt.Printf("%s\n", svc.Name)
		}
	}
}
