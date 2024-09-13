package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jordan-wright/email"
	"github.com/redis/go-redis/v9"
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
	"net/smtp"
	"os"
	"strings"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", global.ServerConfig.MysqlConfig.User,
		global.ServerConfig.MysqlConfig.Password, global.ServerConfig.MysqlConfig.Host, global.ServerConfig.MysqlConfig.Port, global.ServerConfig.MysqlConfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.NewGormLogger(), NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	myerrors.ErrorPanic(err)
	return db
}

func RedisConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.ServerConfig.RedisConfig.Host + ":" + global.ServerConfig.RedisConfig.Port,
		Password: global.ServerConfig.RedisConfig.Password, // no password set
		DB:       global.ServerConfig.RedisConfig.DbName,   // use default DB
	})
	return rdb
}

func initEmail() {
	address := fmt.Sprintf("%s:%s", global.ServerConfig.MailConfig.Host, global.ServerConfig.MailConfig.Port)
	p, err := email.NewPool(
		address,
		4,
		smtp.PlainAuth("", global.ServerConfig.MailConfig.User, global.ServerConfig.MailConfig.AuthCode, global.ServerConfig.MailConfig.Host),
	)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
	global.EmailPool = p
}

// 初始化数据库
func setupGorm() {

	db := DatabaseConnection()
	err := db.Table("manager").AutoMigrate(&model.Manager{})
	myerrors.ErrorPanic(err)
	err = db.Table("user").AutoMigrate(&model.User{})
	myerrors.ErrorPanic(err)
	global.Redis = RedisConnection()
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
	//initEmail()
}
