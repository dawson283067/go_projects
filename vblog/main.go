package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"

	// 通过import方法 完成注册
	_ "github.com/go_projects/vblog/apps/token/api"
	_ "github.com/go_projects/vblog/apps/token/impl"
	_ "github.com/go_projects/vblog/apps/user/impl"
)

func main() {

	// 1. 初始化程序配置，这里没有配置，使用默认值
	if err := conf.LoadFromEnv(); err != nil {
		panic(err)
	}

	// 2.1 先注册对象


	// 2.2 程序对象管理
	if err := ioc.Init(); err != nil {
		panic(err)
	}

	/* v2版，这里都不要了
	// user service impl
	usvc := user_impl.NewUserServiceImpl()

	// token service impl
    tsvc := token_impl.NewTokenServiceImpl(usvc)

	// api
	TokenApiHandler := api.NewTokenApiHandler(tsvc)
	*/

	// Protocol
	engine := gin.Default()

	rr := engine.Group("/vblog/api/v1")
	// 改造之前：TokenApiHandler.Registry(rr)
	// 改造之后第一版：ioc.Api().Get(token.AppName).(*api.TokenApiHandler).Registry(rr)
	ioc.RegistryGinApi(rr)

	// 把Http协议服务器启动起来
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
