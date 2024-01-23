package impl

import (
	"context"
	"github.com/go_projects/vblog/apps/token"
)

// 需要复用的数据库操作
// 这里没有用refresh token，按逻辑写就可以
func (i *TokenServiceImpl) getToken(ctx context.Context, accessToken string) (*token.Token, error) {
	tk := token.NewToken(false)
	err := i.db.
		WithContext(ctx).
		Model(&token.Token{}).
		Where("access_token = ?", accessToken).
		First(tk).
		Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}