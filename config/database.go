package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"student_manage/utils/errors"
)

const (
	host     = "127.0.0.1"
	port     = "43306"
	user     = "root"
	password = "123456"
	dbname   = "student_manage"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	errors.ErrorPanic(err)
	return db
}
