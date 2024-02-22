package cmd

import (
	"github.com/spf13/cobra"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"
	"github.com/gin-gonic/gin"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务器",
	Run: func(cmd *cobra.Command, args []string) {
		// 什么都不做的时候打印帮助信息

		// 1. 初始化程序配置，这里没有配置，使用默认值
		cobra.CheckErr(conf.LoadFromEnv())
		
		// 2. 程序对象管理
		cobra.CheckErr(ioc.Init())
		
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
		cobra.CheckErr(engine.Run(":8080"))		
	},
}

