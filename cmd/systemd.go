package cmd

// https://github.com/coreos/go-systemd

import (
	"github.com/codegangsta/cli"
    //"github.com/coreos/go-systemd"
)

var CmdSystemd = cli.Command{
	Name:  "systemd",
	Usage: "Start GoIMServer in systemd mode",
	Description: `GoIMServer in systemd mode has special systemd integration`,
	Action: runSystemd,
	Flags: []cli.Flag{
		boolFlag("start", "Start service"),
		boolFlag("stop", "Stop service"),
		boolFlag("restart", "Restart service"),
		stringFlag("profile", "config", "Configuration profile"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func runSystemd(ctx *cli.Context) error {
    return nil
}