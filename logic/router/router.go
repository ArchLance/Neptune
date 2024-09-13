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
	managerRepository := repository.NewManagerRepositoryImpl(global.DB)
	managerService := service.NewManagerServiceImpl(managerRepository, validate)
	managerController := controller.NewManagerController(managerService)

	userRepository := repository.NewUserRepositoryImpl(global.DB)
	userService := service.NewUserServiceImpl(userRepository, validate)
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

		}
	}
	return routers
}

//func CollectRoute(routers *gin.Engine) *gin.Engine {
//
//	baseRouter := routers.Group("/api/v1")
//
//	userGroup := baseRouter.Group("/user")
//	{
//		//用户登陆
//		userGroup.GET("/login", controller.Login)
//		//用户登出
//		//userGroup.POST("/logout", controller.Logout)
//		//用户信息修改
//		userGroup.Use(token.JWTAuth())
//		{
//			////用户修改密码
//			//userGroup.POST("/changePassword", controller.ChangePassword)
//			////获取用户列表
//			//userGroup.POST("/list", controller.GetUserList)
//			////上传文件（给某个用户发送文件）
//			//userGroup.POST("/uploadFile", controller.UploadFile)
//			////查看文件（查看所有用户给自己发送的文件）
//			//userGroup.POST("/fileList", controller.GetFileList)
//			////下载文件（下载别的用户发给自己的文件）
//			//userGroup.POST("/downloadFile", controller.DownloadFile)
//
//		}
//	}
//
//	return routers
//}
