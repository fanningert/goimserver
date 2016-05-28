package main

//https://github.com/tam7t/xmpp/tree/master

import (
	"os"
	"runtime"
	"github.com/codegangsta/cli"
	
	"./cmd"
	//"./modules"
	"./modules/setting"
)

const APP_VER = "0.0.1.0001"

func init(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = APP_VER
}

func main() {
	app := cli.NewApp()
	app.Name = "GoIMServer"
	app.Usage = "GoIMServer Service"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdStandalone,
		cmd.CmdSystemd,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}