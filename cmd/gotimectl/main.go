package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/bline/gotime/api"
	"github.com/bline/gotime/config"
	"github.com/bline/gotime/db"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	var (
		env  string = "production"
		host string = "127.0.0.1"
		port int    = 9443
	)
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:        "environment, e",
			Value:       env,
			Usage:       "Either development or production",
			EnvVar:      "GOENV",
			Destination: &env,
		},
		cli.StringFlag{
			Name:        "host, h",
			Value:       host,
			Usage:       "gotime server to connect to"
			Destination: &host,
		},
		cli.IntFlag{
			Name:        "port, p",
			Value:       port,
			Usage:       "gotime server to connect to"
			Destination: &port,
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "config",
			Aliases:     []string{"conf", "cfg"},
			Usage:       "set/get settings or initialize config"
			Subcommands: cli.Commands{
				cli.Command{
					Name: "init",
					Action: initConfig,
					Usage: "initialize configuration settings",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "interactive, i",
							Value: true,
							Usage: "interactive mode for configuration settings",
						},
					},
				},
				cli.Command{
					Name: "get",
					Action: getConfig,
					Usage: "get configuration setting or all settings",
				},
				cli.Command{
					Name: "set",
					Action: setConfig,
					Usage: "set configuration option",
				},
			},
		},
		cli.Command{
			Name: "db",
			Aliases: []string{"database"},
			Usage: "work with the gotime database",
			Subcommands: cli.Commands{
				cli.Command{
					Name: "init",
					Action: initDatabase,
					Usage: "initialize the database, wipes all data",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "force, f",
							Value: false,
							Usage: "forces initializing when a database exists. wipes all data",
						},
					},
				},
			},
		},
		cli.Command{
			Name: "api",
			Usage: "make RPC calls to gotime's GRPC API",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "username, u",
					Usage: "specify the username (email) to login to the api",
					EnvVar: "GOTIME_APIUSER",
				},
			}
			Subcommands: cli.Commands{
				cli.Command{
					Name: "TimeSheet",
					Usage: "operations dealing with current user's timesheet",
					Subcommands: cli.Commands{
						cli.Command{
							Name: "ClockIn",
							Usage: "Clock the current user into gotime.",
							Action: apiTSClockIn,
						},
						cli.Command{
							Name: "ClockOut",
							Usage: "Clock the current user out of gotime.",
							Action: apiTSClockOut,
						},
						cli.Command{
							Name: "Status",
							Usage: "Check clocked-in status of the current user.",
							Action: apiTSStatus,
						},
					},
				},
			},
		},
	}
	app.Name = "gotimectl"
	app.Usage = "initialize and control the gotime server"
	c := config.New(
}

func apiClockIn (c *cli.Contaxt) error {
}

func initDatabase (c *cli.Context) error {
	var force bool = c.flagSet.Lookup("force").Value.(flag.Getter).Get()
	return nil
}

func initConfig (c *cli.Context) error {
	var interactive bool = c.flagSet.Lookup("interactive").Value.(flag.Getter).Get()
	return nil
}

func setConfig (c *cli.Context) error {
	var (
		args []string = c.Args()
		cfgVar string = args.First()
		cfgVal string = args[1]
	)

	return nil
}

func getConfig (c *cli.Context) error {
	var (
		cfgVar string = c.Args().First()
	)
	return nil
}

