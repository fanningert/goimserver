package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdStandalone = cli.Command{
	Name:  "standalone",
	Usage: "Start GoIMServer in standalone mode",
	Description: `GoIMServer in standalone mode is good for testing`,
	Action: runStandalone,
	Flags: []cli.Flag{},
}

func runStandalone(ctx *cli.Context) error {
    return nil
}