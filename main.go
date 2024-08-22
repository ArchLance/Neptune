package main

import (
	"github.com/go-playground/validator/v10"
	"neptune/global"
	"neptune/logic/controller"
	"neptune/logic/repository"
	"neptune/logic/router"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"net/http"
)

func main() {

	validate := validator.New()
	managerRepository := repository.NewManagerRepositoryImpl(global.DB)
	managerService := service.NewManagerServiceImpl(managerRepository, validate)
	managerController := controller.NewManagerController(managerService)
	routerConfig := router.ConfigRouterGroup{
		BasePath:          "/api",
		ManagerController: managerController,
	}
	routers := router.NewRouter(routerConfig)
	server := http.Server{Addr: ":9001", Handler: routers}
	err := server.ListenAndServe()
	myerrors.ErrorPanic(err)
}
