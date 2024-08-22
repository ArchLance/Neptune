package main

import (
	log "github.com/sirupsen/logrus"
	"neptune/config"
	"neptune/global"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
	"neptune/utils/logger"
	"strings"
)

func setupLogrus() error {
	// 配置日志等级
	log.SetLevel(log.InfoLevel)
	logLevel := "debug"
	if l, ok := logger.FlagLToLevel[strings.ToLower(logLevel)]; ok {
		log.SetLevel(l)
	}
	// 日志格式
	log.SetFormatter(&logger.SimpleFormatter{})
	// 日志输出
	log.SetOutput(logger.GetWriter())
	return nil
}

func init() {
	if err := setupLogrus(); err != nil {
		log.Fatal(err)
	}
}

// 初始化数据库
func setupGorm() {

	db := config.DatabaseConnection()
	err := db.Table("manager").AutoMigrate(&model.Manager{})

	myerrors.ErrorPanic(err)
	global.DB = db

}
