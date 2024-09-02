package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"neptune/config"
	"neptune/global"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
	"neptune/utils/logger"
	"os"
	"strings"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", global.ServerConfig.DBconfig.User,
		global.ServerConfig.DBconfig.Password, global.ServerConfig.DBconfig.Host, global.ServerConfig.DBconfig.Port, global.ServerConfig.DBconfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.NewGormLogger(), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	myerrors.ErrorPanic(err)
	return db
}

// 初始化数据库
func setupGorm() {

	db := DatabaseConnection()
	err := db.Table("manager").AutoMigrate(&model.Manager{})
	myerrors.ErrorPanic(err)
	err = db.Table("user").AutoMigrate(&model.User{})
	myerrors.ErrorPanic(err)
	global.DB = db

}

func initConfig() {
	//初始化viper
	workDir, _ := os.Getwd()
	v := viper.New()
	v.SetConfigName("application")
	v.SetConfigType("yaml")
	v.AddConfigPath(workDir + "/config")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed:", e.Name)
	})
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.ServerConfig = serverConfig
}

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
	initConfig()
	setupGorm()
}
