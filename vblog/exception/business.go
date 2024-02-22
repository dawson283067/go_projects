package exception

import "net/http"

//
var (
	// 请求不合法
	ErrBadRequest = NewAPIException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
	// 未认证，没有登录，Token没传递
	ErrUnauthorized = NewAPIException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).WithHttpCode(http.StatusUnauthorized)
	// 鉴权失败：认证通过，但是没有权限操作 该接口
	ErrPermissionDeny = NewAPIException(http.StatusForbidden, http.StatusText(http.StatusForbidden)).WithHttpCode(http.StatusForbidden)
)
