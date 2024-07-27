package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student_manage/logic/controller"
)

type ConfigRouterGroup struct {
	BasePath          string
	ManagerController *controller.ManagerController
}

func NewRouter(config ConfigRouterGroup) *gin.Engine {
	routers := gin.Default()
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
