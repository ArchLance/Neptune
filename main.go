package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"neptune/config"
	"neptune/logic/controller"
	"neptune/logic/model"
	"neptune/logic/repository"
	"neptune/logic/router"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"net/http"
)

func main() {
	log.Info().Msg("Started Server!")
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
