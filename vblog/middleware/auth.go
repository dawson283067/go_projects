package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/ioc"
	"github.com/go_projects/vblog/response"
)

func Auth(c *gin.Context) {
	// 获取tk 模块
	svc := ioc.Controller().Get(token.AppName).(token.Service)

	ak := token.GetAccessTokenFromHttp(c.Request)
	req := token.NewValidateTokenRequest(ak)
	tk, err := svc.ValidateToken(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 请求通过，用户的身份信息 携带在 请求的上下文中，传递给后续请求
	// 对于Gin，如何把请求的中间数据传递下去，使用了一个map，这个map在request对象上
	c.Set(token.TOKEN_MIDDLEWARE_KEY, tk)

	// 后面的请求如何获取 中间信息
	// c.Get(token.TOKEN_MIDDLEWARE_KEY).(*token.Token).UserName	
}