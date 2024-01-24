package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/token/api"
	token_impl "github.com/go_projects/vblog/apps/token/impl"
	user_impl "github.com/go_projects/vblog/apps/user/impl"
)

func main() {

	// user service impl
	usvc := user_impl.NewUserServiceImpl()

	// token service impl
    tsvc := token_impl.NewTokenServiceImpl(usvc)

	// api
	TokenApiHandler := api.NewTokenApiHandler(tsvc)

	// Protocol
	engine := gin.Default()

	rr := engine.Group("/vblog/api/v1")
	TokenApiHandler.Registry(rr)

	// 把Http协议服务器启动起来
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}

}
