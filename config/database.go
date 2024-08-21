package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"neptune/utils/errors"
	"neptune/utils/logger"
)

const (
	host     = "127.0.0.1"
	port     = "43306"
	user     = "root"
	password = "123456"
	dbname   = "neptune"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.NewGormLogger(), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	errors.ErrorPanic(err)
	return db
}
