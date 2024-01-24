package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/response"
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
	mr.POST("/", h.Login)
	mr.DELETE("/", h.Logout)
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
		response.Failed(c, err)
		return
	}
	
	// 2. 业务逻辑处理
	tk, err := h.svc.IssueToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 Set Cookie
	c.SetCookie(
		token.TOKEN_COOKIE_KEY,
		tk.AccessToken,
		tk.AccessTokenExpiredAt,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3. 返回处理结果
	response.Success(c, tk)
}

// 退出
func (h *TokenApiHandler) Logout(c *gin.Context) {
	// 1. 解析用户请求
	// token为了安全，存放在Cookie获取自定义Header中
	accessToken := token.GetAccessTokenFromHttp(c.Request)
	req := token.NewRevokeTokenRequest(accessToken, c.Query("refresh_token"))

	// 2. 业务逻辑处理
	_, err := h.svc.RevokeToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2.1 删除前端的cookie
	c.SetCookie(
		token.TOKEN_COOKIE_KEY,
		"",
		-1,
		"/",
		conf.C().Application.Domain,
		false,
		true,
	)

	// 3. 返回处理结果
	response.Success(c, "退出成功")
}