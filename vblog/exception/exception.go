package exception

import (
	"fmt"

	"github.com/infraboard/mcube/tools/pretty"
)

func NewAPIException(code int, reason string) *APIExcetion {
	return &APIExcetion{
		Code:    code,
		Reason: reason,
	}
}

// error的自定义实现
// 通过 API 直接序列化
type APIExcetion struct {
	HttpCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func (e *APIExcetion) Error() string {
	return fmt.Sprintf("%s, %s", e.Reason, e.Message)
}

func (e *APIExcetion) String() string {
	return pretty.ToJSON(e)
}

// 设计为链式调用 New().WithMessage()
func (e *APIExcetion) WithMessage(msg string) *APIExcetion {
	e.Message = msg
	return e
}

func (e *APIExcetion) WithMessagef(format string, a ...any) *APIExcetion {
	e.Message = fmt.Sprintf(format, a...)
	return e
}

// 给一个异常判断的方法
func IsException(err error, e *APIExcetion) bool {
	if target, ok := err.(*APIExcetion); ok {
		return target.Code == e.Code
	}
	return false
}