package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Init() error {
	var (
		dsn string
		err error
	)

	dsn = "root:root@tcp(localhost:3306)/shorten?parseTime=true&charset=utf8"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("Database connect failed, " + err.Error())
		return err
	}
	return nil
}
