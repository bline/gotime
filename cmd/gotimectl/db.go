
package ctl

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/bline/gotime/api"
	"github.com/bline/gotime/config"
	"github.com/bline/gotime/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	DB *gorm.DB
	force bool
}

(d *DB) func Init(c *cli.Context) error {

}
