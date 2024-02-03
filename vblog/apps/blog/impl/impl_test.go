package impl_test

import (
	"context"

	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/ioc"

	// 1. 加载对象
	_ "github.com/go_projects/vblog/apps"
)

// blog service 的实现的具体对象是在ioc中
// 需要在ioc中获取具体的svc 用来测试

var (
	impl blog.Service
	ctx = context.Background()
)

func init() {
	// 2. ioc获取对象
	impl = ioc.Controller().Get(blog.AppName).(blog.Service)

	// ioc需要初始化才能填充db属性
	if err := ioc.Init(); err != nil {
		panic(err)
	}
}





