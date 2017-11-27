
package ctl

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/bline/gotime/api"
	"github.com/bline/gotime/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Api struct {
	cfg *config.Config
	Username string
}

func New(cfg *config.Config, user string) (Api, error) {
}

(a *Api) func TSClockIn(c *cli.Context) error {
}

(a *Api) func TSClockOut(c *cli.Context) error {
}

(a *Api) func TSStatus(c *cli.Context) error {
}

(a *Api) func AcctGetUser(c *cli.Context) error {
}

(a *Api) func AcctGetUsers(c *cli.Context) error {
}

(a *Api) func AcctDisableUser(c *cli.Context) error {
}

(a *Api) func AcctDeleteUser(c *cli.Context) error {
}

(a *Api) func AcctLockUser(c *cli.Context) error {
}


