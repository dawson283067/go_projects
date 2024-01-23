package token

import "context"

// Token Service 接口定义

type Service interface {
	// 登录： 颁发令牌
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)

	// 退出：撤销令牌
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)

	// 校验令牌
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

// 颁发令牌的请求
type IssueTokenRequest struct {
	Username string
	Password string
	// 延长Token的有效期为1周
	RemindMe bool
}

func NewRevokeTokenRequest(accessToken, refreshToken string) *RevokeTokenRequest {
	return &RevokeTokenRequest{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}

// 撤销令牌的请求
type RevokeTokenRequest struct {
	// 颁发给用户的访问令牌（用户需要携带Token来访问接口）
	AccessToken string
	// 他们是配对的，撤销时，需要验证他们是不是一堆
	RefreshToken string
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}

type ValidateTokenRequest struct {
	AccessToken string
}
