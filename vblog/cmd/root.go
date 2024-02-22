package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "vblog api server",
	Run: func(cmd *cobra.Command, args []string) {
	  	// 什么都不做的时候打印帮助信息
		cmd.Help()
	},
}
  
func Execute() {
	// 注册Root命令的子命令
	// Flags 返回这个子命令的所有Flag
	// StringVarP 获取命令参数的值，转化为字符串 赋值给指定的变量
	// --username -u
	initCmd.Flags().StringVarP(&username, "username", "u", "admin", "管理员用户的名称")
	rootCmd.AddCommand(initCmd, startCmd)

	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
}
