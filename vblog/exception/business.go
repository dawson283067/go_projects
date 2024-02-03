package exception

import "net/http"

//
var (
	// 请求不合法
	ErrBadRequest = NewAPIException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
)