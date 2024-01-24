package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/token"
)

// 来实现对外提供 RESTful 接口
type TokenApiHandler struct {
	svc token.Service
}

// 如何为Handler添加路由，如何把路由注册给Http Server
// 需要一个Root Router: path prefix: /vblog/api/v1
func (h *TokenApiHandler) Registry(rr gin.IRouter) {
	// 每个业务模块 都需要往Gin Engine对象注册路由
	// r := gin.Default()
	// rr := r.Group("vblog/api/v1")

	// 模块路径
	// /vblog/api/v1/tokens
	mr := rr.Group("token.AppName")
	mr.POST("tokens", h.Login)
	mr.DELETE("tokens", h.Logout)
}

// 登录
func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1. 解析用户请求
	// http 的请求可以放到哪里，放body，bytes
	// io.ReadAll(c.Request.Body)
	// defer c.Request.Body.Close()
	// json unmarshal json.Unmarshal(body, o)
	
	// Body 必须是json
	req := token.NewIssueTokenRequest("", "")
	if err := c.BindJSON(req); err != nil {
		return 
	}
	

	// 2. 业务逻辑处理

	// 3. 返回处理结果
}

// 退出
func (h *TokenApiHandler) Logout(ctx *gin.Context) {

}