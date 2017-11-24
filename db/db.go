
import (
	"log"
	"github.com/bline/gotime/gotime.go"
	"github.com/bline/gotime/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

func New() (*gorm.DB, error) {
	c := config.GetConfig()
	err error
	if db != nil {
		return db, err
	}
	if db, err = gorm.Open(c.GetString("SQL.Driver"), c.GetString("SQL.Connection")); err != nil {
		log.Fatal(err)
	}
	return db, err
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

