package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/rtgnx/svctl/internal/agent/cmd"
)

func main() {
	app := cli.App("svctl", "runit control")

	app.Command("list", "list running services", cmd.CmdList)

	app.Run(os.Args)
}
