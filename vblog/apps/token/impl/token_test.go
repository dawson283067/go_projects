package impl_test

import (
	"context"
	"testing"

	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/token/impl"
	ui "github.com/go_projects/vblog/apps/user/impl"
	"github.com/go_projects/vblog/exception"
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

func TestRevokeToken(t *testing.T) {
	req := token.NewRevokeTokenRequest(
		"cmlcakka0uti117ngqp0",
		"cmlcakka0uti117ngqpg",
	)
	tk, err := i.RevokeToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}


// refresh token expired 5233.530429 minutes
/* 
	{
          "user_id": "11",
          "username": "admin",
          "access_token": "cmlcbuca0uti92286di0",
          "access_token_expired_at": 604800,
          "refresh_token": "cmlcbuca0uti92286dig",
          "refresh_token_expired_at": 2419200,
          "created_at": 1705690617,
          "updated_at": 1705690617,
          "role": 0
    }
*/ 
func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("cmlcbuca0uti92286di0")
	tk, err := i.ValidateToken(ctx, req)
	// 通过断言来获取一个exception
	if e, ok := err.(*exception.APIExcetion); ok {
		t.Log(e.String())
		// 判断该异常是不是 TokenExpired 异常
		if e.Code == token.ErrAccessTokenExpired.Code {
			t.Log(e.String())
		}
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	// 加载被测试对象， i 就是Token Service接口的具体实现对象
	i = impl.NewTokenServiceImpl(ui.NewUserServiceImpl())
}