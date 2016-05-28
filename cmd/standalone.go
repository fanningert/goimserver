package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdStandalone = cli.Command{
	Name:  "standalone",
	Usage: "Start GoIMServer in standalone mode",
	Description: `GoIMServer in standalone mode is good for testing`,
	Action: runStandalone,
	Flags: []cli.Flag{
		stringFlag("config, c", "/etc/goimserver/config.ini", "Custom configuration file path"),
	},
}

func runStandalone(ctx *cli.Context) error {
    return nil
}