package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/exception"
)

// API成功，返回数据
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}


// API失败，返回错误，API Exception
func Failed(c *gin.Context, err error) {

	// 构造异常数据
	var resp *exception.APIExcetion
	if e, ok := err.(*exception.APIExcetion); ok {
		resp = e
	} else {
		resp = exception.NewAPIException(
			500,
			http.StatusText(http.StatusInternalServerError),
		).WithMessage(err.Error()).WithHttpCode(500)
		// 老师写错过，e.Error()，出现空指针。注意这个问题
	}

	// 返回异常
	c.JSON(resp.HttpCode, resp)

	// 再中断
	c.Abort()
}

