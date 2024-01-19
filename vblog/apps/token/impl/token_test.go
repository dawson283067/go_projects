package impl_test

import (
	"context"
	"testing"

	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/token/impl"
	ui "github.com/go_projects/vblog/apps/user/impl"
)



var (
	i token.Service
	ctx = context.Background()
)

/*
	{
          "user_id": "11",
          "username": "admin",
          "access_token": "cmlcaoca0uths53c6j4g",
          "access_token_expired_at": 7200,
          "refresh_token": "cmlcaoca0uths53c6j50",
          "refresh_token_expired_at": 28800,
          "created_at": 1705690465,
          "updated_at": 1705690465,
          "role": 1
    }
*/
func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")
	req.RemindMe = true
	tk, err := i.IssueToken(ctx, req)
	if err != nil {
		t.Fatal()
	}
	t.Log(tk)
}

func init() {
	// 加载被测试对象， i 就是User Service接口的具体实现对象
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}