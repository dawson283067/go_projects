package cmd

import (
	"fmt"

	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"
	"github.com/rs/xid"
	"github.com/spf13/cobra"
)

var (
	// 命令参数，需要通过用户传入参数： --username root
	username string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "程序初始化",
	Run: func(cmd *cobra.Command, args []string) {
	  	// 什么都不做的时候打印帮助信息

		// 1. 初始化程序配置，这里没有配置，使用默认值
		cobra.CheckErr(conf.LoadFromEnv())
		

		// 2. 程序对象管理
		cobra.CheckErr(ioc.Init())

		// 3. 需要初始化 管理员用户
		
		// 使用构造函数创建请求对象
		req := user.NewCreateUserRequest()
		// user.CreateUserRequest{} 这样写很容易出现空指针
		req.Username = username
		req.Password = xid.New().String()
		req.Role = user.ROLE_ADMIN

		fmt.Println("用户名：", req.Username)
		fmt.Println("密码：", req.Password)

		// 创建用户，PO到数据库中
		u, err := ioc.Controller().Get(user.AppName).(user.Service).CreateUser(
			cmd.Context(),
			req,
		)

		// 单元测试异常怎么处理
		// 直接报错中断单元流程并且失败
		cobra.CheckErr(err)

		// 正常打印对象
		fmt.Println(u)
	},
}

