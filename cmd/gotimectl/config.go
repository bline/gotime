
package ctl

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/bline/gotime/config"
)

type Config struct {
	cfg *config.Config
	interactive bool
}

(d *Config) func Init(c *cli.Context) error {
	fmt.PrintLn("Interactive: ", d.interactive);
}

(d *Config) func Set(c *cli.Context) error {
	var (
		args []string = c.Args()
		cfgVar string = args.First()
		cfgVal string = args[1]
	)
	fmt.PrintLn("Set: ", cfgVar, " => ", cfgVal);
}

(d *Config) func Get(c *cli.Context) error {
	var (
		cfgVar string = c.Args().First()
	)
	fmt.PrintLn("Get: ", cfgVar);
}
