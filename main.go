package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
	"student_manage/config"
	"student_manage/logic/controller"
	"student_manage/logic/model"
	"student_manage/logic/repository"
	"student_manage/logic/router"
	"student_manage/logic/service"
	myerrors "student_manage/utils/errors"
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
