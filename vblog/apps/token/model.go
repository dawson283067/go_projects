package token

import (
	"fmt"
	"time"

	"github.com/go_projects/vblog/apps/user"
	"github.com/infraboard/mcube/tools/pretty"
	"github.com/rs/xid"
)

const (
	// 以秒为单位
	DEFAULT_EXPIRED_AT = 2 * 60 * 60
	WEEK_EXPIRED_AT = 7 * 24 * 60 * 60
)

func NewToken(remindMe bool) *Token {
    // access token expired at
	// refresh token expired at
	atet := DEFAULT_EXPIRED_AT
	
	if remindMe {
		// 7天的过期时间
		atet = WEEK_EXPIRED_AT
	}

	return &Token{
		// 直接使用uuid库来生成一个随机字符串
		AccessToken: xid.New().String(),
		AccessTokenExpiredAt: atet,
		RefreshToken: xid.New().String(),
		RefreshTokenExpiredAt: atet * 4,
		CreatedAt: time.Now().Unix(),
	}
}

type Token struct {
	// 该Token是颁发
	UserId string `json:"user_id" gorm:"column:user_id"`
	// 人的名称， user_name
	UserName string `json:"username" gorm:"column:username"`
	// 颁发给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"access_token"`
	// 过期时间(2h), 单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"access_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"refresh_token"`
	// 刷新Token过期时间(7d)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"refresh_token_expired_at"`

	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at" gorm:"updated_at"`

	// 额外补充信息, gorm忽略处理
	Role user.Role `json:"role" gorm:"-"`
}

func (t *Token) CheckRefreshToken(refreshToken string) error {
	if t.RefreshToken != refreshToken {
		return fmt.Errorf("refresh token not correct")
	}
	return nil
}

// 校验Token是否过期
// 1. access_token过期
// 2. refress_token过期
func (t *Token) ValidateExpired() error {
	// 颁发时间 + refresh_token过期
	// type Duration int64
	// time.Second是一个time.Duration对象
	refreshExpiredTime := time.
		Unix(t.CreatedAt, 0).
		Add(time.Duration(t.RefreshTokenExpiredAt) * time.Second)
	
	// 和当前时间比较
	// now - refreshExpiredTime
	// 大于0就是过期了
	rDelta := time.Since(refreshExpiredTime).Minutes()
	if rDelta > 0 {
		return ErrRefreshTokenExpired.WithMessagef("refresh token expired %f minutes", rDelta)
		// return fmt.Errorf("refresh token expired %f minutes", rDelta)
	}

	// 颁发时间 + access_token过期
	accessExpiredTime := time.
		Unix(t.CreatedAt, 0).
		Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
	
	aDelta := time.Since(accessExpiredTime).Minutes()
	if  aDelta > 0 {
		return ErrAccessTokenExpired.WithMessagef("access token expired %f minutes", aDelta)
		// return fmt.Errorf("access token expired %f minutes", aDelta)
	}

	return nil
}

func (t *Token) TableName() string {
	return "tokens"
}

func (t *Token) String() string {
	return pretty.ToJSON(t)
}
