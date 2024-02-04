package exception

import "net/http"

//
var (
	// 请求不合法
	ErrBadRequest = NewAPIException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
	// 未认证，没有登录，Token没传递
	ErrUnauthorized = NewAPIException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).WithHttpCode(http.StatusUnauthorized)
)