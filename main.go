package main

import (
	log "github.com/sirupsen/logrus"
	"neptune/logic/router"
	myerrors "neptune/utils/errors"
	"neptune/utils/hash"
	"net/http"
)

// 编译时要编译整个包 package，否则init.go无法执行
func main() {
	//routers := gin.Default()
	//routers.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: logger.GinLogFormatter}), gin.Recovery())
	//routers = router.CollectRoute(routers)
	hashPasswd := hash.SHA256DoubleString("123456", false)
	log.Info(hashPasswd)
	routers := router.NewRouter(router.NewConfigRouterGroup())
	server := http.Server{Addr: "127.0.0.1:9001", Handler: routers}
	err := server.ListenAndServe()
	myerrors.ErrorPanic(err)
}
