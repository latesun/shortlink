package db

import (
	"bindolabs/bindocommon/env"
	"bindolabs/golib/logger"

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

	switch env.Env.RunMode {
	case env.RunModeStaging:
		dsn = "gateway:pass4u@tcp(db-master.stg.tw.bindo.in:3306)/gateway?parseTime=true&charset=utf8"
	case env.RunModeProduction:
		dsn = "sales-team:ev3tMAtY@tcp(bi-readonly-database.prd.tw.bindo.in:3306)/gateway?parseTime=true&charset=utf8"
	default:
		dsn = "root:root@tcp(localhost:34561)/gateway?parseTime=true&charset=utf8"
	}

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		logger.Error("Database connect failed, " + err.Error())
		return err
	}
	return nil
}
