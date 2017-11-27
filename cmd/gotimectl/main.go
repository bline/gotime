package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	gtcli "github.com/bline/gotime/cli"
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

	var (
		ctlConfig *gtcli.Config = &gtcli.Config{}
		ctlDB *gtcli.DB = &CtlDB{}
		ctlApi *gtcli.Api = &CtlApi{}
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
		{
			Name:        "config",
			Aliases:     []string{"conf", "cfg"},
			Usage:       "set/get settings or initialize config"
			Subcommands: []cli.Command{
				{
					Name: "init",
					Action: ctlConfig.Init,
					Usage: "initialize configuration settings",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "interactive, i",
							Value: true,
							Usage: "interactive mode for configuration settings",
							Destination: &ctlConfig.interactive,
						},
					},
				},
				{
					Name: "get",
					Action: ctlConfig.Get,
					Usage: "get configuration setting or all settings",
				},
				{
					Name: "set",
					Action: ctlConfig.Set,
					Usage: "set configuration option",
				},
			},
		},
		{
			Name: "db",
			Aliases: []string{"database"},
			Usage: "work with the gotime database",
			Subcommands: []cli.Command{
				{
					Name: "init",
					Action: ctlDB.Init,
					Usage: "initialize the database, wipes all data",
					Flags: []cli.Flag{
						cli.BoolFlag{
							Name: "force, f",
							Value: false,
							Usage: "forces initializing when a database exists. wipes all data",
							Destination: *ctlDB.force,
						},
					},
				},
			},
		},
		{
			Name: "api",
			Usage: "make RPC calls to gotime's GRPC API",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "username, u",
					Usage: "specify the username (email) to login to the api",
					EnvVar: "GOTIME_APIUSER",
					Destination: &ctlApi.Username
				},
			}
			Subcommands: []cli.Command{
				{
					Name: "TimeSheet",
					Usage: "operations dealing with current user's timesheet",
					Subcommands: []cli.Command{
						{
							Name: "ClockIn",
							Usage: "Clock the current user into gotime.",
							Action: ctlApi.TSClockIn,
						},
						{
							Name: "ClockOut",
							Usage: "Clock the current user out of gotime.",
							Action: ctlApi.TSClockOut,
						},
						{
							Name: "Status",
							Usage: "Check clocked-in status of the current user.",
							Action: ctlApi.TSStatus,
						},
					},
				},
				{
					Name: "Accounts",
					Usage: "administration os user accounts",
					Subcommands: []cli.Command{
						{
							Name: "GetUser",
							Usage: "Get user by email or ID. requires admin privs",
							Action: ctlApi.AcctGetUser,
						},
						{
							Name: "GetUsers",
							Usage: "List users given a query and filters",
							Action: ctlApi.AcctGetUsers,
						},
						{
							Name: "DisableUser",
							Usage: "Disable a user account",
							Action: ctlApi.AcctDisableUser,
						},
						{
							Name: "DeleteUser",
							Usage: "Delete a user account",
							Action: ctlApi.AcctDeleteUser,
						},
						{
							Name: "LockUser",
							Usage: "Lock a user account",
							Action: ctlApi.AcctLockUser,
						},
					},
				},
			},
		},
	}
	app.Name = "gotimectl"
	app.Usage = "initialize and control the gotime server"
	c := config.New(env);
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

