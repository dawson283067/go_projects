package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/exception"
	"github.com/go_projects/vblog/response"
)

// Required(user.ROLE_ADMIN) ----> gin.HandlerFunc
func Required(roles ...user.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检验用户的角色
		// 直接通过上下文取出Token，通过Token获取用户角色，来定义访问这个接口的角色
		// 后面请求如何获取 中间信息
		if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
			// 遍历判断 用户是否在运行的角色列表里面
			hasPerm := false
			for _, r := range roles {
				if r == v.(*token.Token).Role {
					hasPerm = true
				}
			}
			if !hasPerm {
				response.Failed(c, exception.ErrPermissionDeny.WithMessagef("允许访问的角色：%v", roles))
				return
			}						
		} else {
			response.Failed(c, exception.ErrUnauthorized)
			return
		}
	}
}