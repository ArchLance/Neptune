package main

import (
	"neptune/logic/router"
	myerrors "neptune/utils/errors"
	"net/http"
)

// 编译时要编译整个包 package，否则init.go无法执行
func main() {
	routers := router.NewRouter(router.NewConfigRouterGroup())
	server := http.Server{Addr: "127.0.0.1:9001", Handler: routers}
	err := server.ListenAndServe()
	myerrors.ErrorPanic(err)
}
