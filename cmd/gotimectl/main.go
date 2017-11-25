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
		{
			Name:        "config",
			Aliases:     []string{"conf", "cfg"},
			Usage:       "set/get settings or initialize config"
			Subcommands: []cli.Command{
				{
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
				{
					Name: "get",
					Action: getConfig,
					Usage: "get configuration setting or all settings",
				},
				{
					Name: "set",
					Action: setConfig,
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
		{
			Name: "api",
			Usage: "make RPC calls to gotime's GRPC API",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "username, u",
					Usage: "specify the username (email) to login to the api",
					EnvVar: "GOTIME_APIUSER",
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
							Action: apiTSClockIn,
						},
						{
							Name: "ClockOut",
							Usage: "Clock the current user out of gotime.",
							Action: apiTSClockOut,
						},
						{
							Name: "Status",
							Usage: "Check clocked-in status of the current user.",
							Action: apiTSStatus,
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
							Action: apiAcctGetUser,
						},
						{
							Name: "GetUsers",
							Usage: "List users given a query and filters",
							Action: apiAcctGetUsers,
						},
						{
							Name: "DisableUser",
							Usage: "Disable a user account",
							Action: apiAcctDisableUser,
						},
						{
							Name: "DeleteUser",
							Usage: "Delete a user account",
							Action: apiAcctDeleteUser,
						},
						{
							Name: "LockUser",
							Usage: "Lock a user account",
							Action: apiAcctLockUser,
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

func apiTSClockIn (c *cli.Contaxt) error {
}

func apiTSClockOut (c *cli.Contaxt) error {
}

func apiTSStatus (c *cli.Contaxt) error {
}

func apiAcctGetUser (c *cli.Contaxt) error {
}

func apiAcctGetUsers (c *cli.Contaxt) error {
}

func apiAcctDisableUser (c *cli.Contaxt) error {
}

func apiAcctDeleteUser (c *cli.Contaxt) error {
}

func apiAcctLockUser (c *cli.Contaxt) error {
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

