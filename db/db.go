package db

import (
	"log"
	"github.com/bline/gotime/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

func Open() (*gorm.DB, error) {
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
	return db
}

func Close() {
	db.Close()
	db = nil
}

