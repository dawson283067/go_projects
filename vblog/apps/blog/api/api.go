package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/ioc"
	"github.com/go_projects/vblog/middleware"
)

func init() {
	ioc.Api().Registry(blog.AppName, &blogApiHandler{})
}

// blog.Service接口，是直接注册给ioc，不需要对暴露
type blogApiHandler struct {
	svc blog.Service
}

func (i *blogApiHandler) Init() error {
	i.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (i *blogApiHandler) Destroy() error {
	return nil
}

// 让ioc帮我们完成接口的路由注册 ioc.GinApi
//
//	type GinApi interface {
//		// 基础约束
//		Object
//		// 额外约束
//		// ioc.Api().Get(token.AppName).(*api.TokenApiHandler).Registry(rr)
//		Registry(rr gin.IRouter)
//	}
//
// ioc调用这个接口来完成注册，ioc会把这个gin.IRouter传给我们，拿到这个router后，就能做事情了
func (i *blogApiHandler) Registry(rr gin.IRouter) {

	r := rr.Group(blog.AppName)

	// 普通接口，允许访客使用，不需要权限
	r.GET("/", i.QueryBlog)
	r.GET("/:id", i.DescribeBlog)

	// 整个模块后面的请求，都需要认证
	r.Use(middleware.Auth)

	// 后台股那里的接口 需要权限
	r.POST("/", i.CreateBlog)
	r.PATCH("/:id", i.PatchBlog)
	r.PUT("/:id", i.UpdateBlog)
	// 只允许管理员才能删除（认证后才能鉴权）
	r.DELETE("/:id", middleware.Required(user.ROLE_ADMIN), i.DeleteBlog)
}
