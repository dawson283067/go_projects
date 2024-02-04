package apps

// 注册业务实现：API + Controller
import (
	// 通过import方法 完成注册
	_ "github.com/go_projects/vblog/apps/token/api"
	_ "github.com/go_projects/vblog/apps/token/impl"
	_ "github.com/go_projects/vblog/apps/user/impl"
	_ "github.com/go_projects/vblog/apps/blog/impl"
	_ "github.com/go_projects/vblog/apps/blog/api"
)