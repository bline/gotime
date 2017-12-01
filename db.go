package gotime

import (
	"log"
	"github.com/bline/gotime/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

func OpenDB() (*gorm.DB, error) {
	c := config.GetConfig()
	var err error
	if db != nil {
		return db, err
	}
	if db, err = gorm.Open(c.GetString("SQL.Driver"), c.GetString("SQL.Connection")); err != nil {
		log.Fatal(err)
	}
	return db, err
}

func GetDB() *gorm.DB {
	if db == nil {
		OpenDB()
	}
	return db
}

func CloseDB() {
	db.Close()
	db = nil
}

