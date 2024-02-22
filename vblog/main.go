package main

import (
	"github.com/go_projects/vblog/cmd"
	

	// 通过import方法 完成注册
	// 这样写，就不用再动main这里的代码了。将注册放到apps包下的init.go文件中
	_ "github.com/go_projects/vblog/apps"
)

func main() {
	cmd.Execute()
}
