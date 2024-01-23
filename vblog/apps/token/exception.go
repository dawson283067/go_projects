package token

import "github.com/go_projects/vblog/exception"

// 这个模块定义的业务异常
// token expired %f minutes
// 约定俗成： ErrXXXXXX 来定义自定义异常当以，方便快速在包内搜索
var (
	ErrAccessTokenExpired = exception.NewAPIException(5000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(5001, "refresh token expired")
)
