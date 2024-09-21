package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"neptune/global"
	"neptune/logic/controller"
	"neptune/logic/repository"
	"neptune/logic/service"
	"neptune/utils/logger"
	middlewares "neptune/utils/middleware"
	"neptune/utils/token"
	"net/http"
)

type ConfigRouterGroup struct {
	BasePath          string
	ManagerController *controller.ManagerController
	UserController    *controller.UserController
}

func NewConfigRouterGroup() *ConfigRouterGroup {
	validate := validator.New()
	managerRepository := repository.NewManagerRepository(global.DB)
	managerService := service.NewManagerService(managerRepository, validate)
	managerController := controller.NewManagerController(managerService)

	userRepository := repository.NewUserRepository(global.DB)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	return &ConfigRouterGroup{
		BasePath:          "/api",
		ManagerController: managerController,
		UserController:    userController,
	}
}

func NewRouter(config *ConfigRouterGroup) *gin.Engine {
	routers := gin.Default()
	// 解决跨域问题
	routers.Use(middlewares.Cors())
	//
	//routers.StaticFS("/media/upload/avatar", http.Dir("static/upload/avatar"))
	// 自定义log
	routers.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: logger.GinLogFormatter}), gin.Recovery())
	// TODO: 没有验证
	routers.StaticFS("/static/upload/avatar", http.Dir("static/upload/avatar"))
	baseRouter := routers.Group(config.BasePath)

	managerRouter := baseRouter.Group("/manager")
	{
		managerRouter.GET("", config.ManagerController.GetAll)
		managerRouter.GET("/:id", config.ManagerController.GetById)
		managerRouter.POST("/create", config.ManagerController.Create)
		managerRouter.POST("", config.ManagerController.Update)
		managerRouter.DELETE("/:id", config.ManagerController.Delete)
	}
	// token authentation
	userRouter := baseRouter.Group("/user")
	{
		userRouter.POST("/login", config.UserController.Login)
		userRouter.Use(token.JWTAuth())
		{
			userRouter.POST("/avatar", config.UserController.UploadAvatar)
			////用户修改密码
			userRouter.PUT("/changePassword", config.UserController.ChangePassword)
			//// 更新
			userRouter.PUT("/update", config.UserController.Update)
			//// 创建验证码
			userRouter.GET("/sendEmail", config.UserController.GenerateCode)
			//// 校验验证码是否存在或者过期
			userRouter.GET("/verifyCode", config.UserController.CheckCode)
			//// 更换邮箱
			userRouter.PUT("/updateEmail", config.UserController.UpdateEmail)
		}
	}
	return routers
}
