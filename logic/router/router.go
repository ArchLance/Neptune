package router

import (
	"github.com/gin-gonic/gin"
	"neptune/logic/controller"
	"neptune/utils/logger"
	"neptune/utils/token"
	"net/http"
)

type ConfigRouterGroup struct {
	BasePath          string
	ManagerController *controller.ManagerController
}

func NewRouter(config ConfigRouterGroup) *gin.Engine {
	routers := gin.Default()
	routers.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: logger.GinLogFormatter}), gin.Recovery())
	routers.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome go")
	})

	baseRouter := routers.Group(config.BasePath)

	managerRouter := baseRouter.Group("/manager")
	{
		managerRouter.GET("", config.ManagerController.GetAll)
		managerRouter.GET("/:id", config.ManagerController.GetById)
		managerRouter.POST("/create", config.ManagerController.Create)
		managerRouter.PATCH("", config.ManagerController.Update)
		managerRouter.DELETE("/:id", config.ManagerController.Delete)

	}
	return routers
}

func CollectRoute(routers *gin.Engine) *gin.Engine {

	baseRouter := routers.Group("/api/v1")

	userGroup := baseRouter.Group("/user")
	{
		//用户登陆
		userGroup.GET("/login", controller.Login)
		//用户登出
		//userGroup.POST("/logout", controller.Logout)
		//用户信息修改
		userGroup.Use(token.JWTAuth())
		{
			////用户修改密码
			//userGroup.POST("/changePassword", controller.ChangePassword)
			////获取用户列表
			//userGroup.POST("/list", controller.GetUserList)
			////上传文件（给某个用户发送文件）
			//userGroup.POST("/uploadFile", controller.UploadFile)
			////查看文件（查看所有用户给自己发送的文件）
			//userGroup.POST("/fileList", controller.GetFileList)
			////下载文件（下载别的用户发给自己的文件）
			//userGroup.POST("/downloadFile", controller.DownloadFile)

		}
	}

	return routers
}
