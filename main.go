package main

import (
	"github.com/gin-gonic/gin"
	"neptune/logic/router"
	myerrors "neptune/utils/errors"
	"neptune/utils/logger"
	"net/http"
)

// 编译时要编译整个包 package
func main() {
	routers := gin.Default()
	routers.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: logger.GinLogFormatter}), gin.Recovery())
	routers = router.CollectRoute(routers)
	server := http.Server{Addr: ":9001", Handler: routers}
	err := server.ListenAndServe()
	myerrors.ErrorPanic(err)
}
