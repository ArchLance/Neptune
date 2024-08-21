package main

import (
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"neptune/config"
	"neptune/logic/controller"
	"neptune/logic/model"
	"neptune/logic/repository"
	"neptune/logic/router"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/logger"
	"net/http"
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

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()
	err := db.Table("manager").AutoMigrate(&model.Manager{})
	myerrors.ErrorPanic(err)
	managerRepository := repository.NewManagerRepositoryImpl(db)
	managerService := service.NewManagerServiceImpl(managerRepository, validate)
	managerController := controller.NewManagerController(managerService)
	routerConfig := router.ConfigRouterGroup{
		BasePath:          "/api",
		ManagerController: managerController,
	}
	routers := router.NewRouter(routerConfig)
	server := http.Server{Addr: ":9001", Handler: routers}
	err = server.ListenAndServe()
	myerrors.ErrorPanic(err)
}
