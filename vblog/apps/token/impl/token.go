package impl

import (
	"context"
	"fmt"

	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/user"
)

// 登录：颁发令牌
// 依赖User模块来校验 用户密码是否正确
func (i *TokenServiceImpl) IssueToken(
	ctx context.Context,
	in *token.IssueTokenRequest) (
	*token.Token, error) {
	
	// 1. 确认用户密码是否正确
	req := user.NewQueryUserRequest()
	req.Username = in.Username
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("用户名或者密码错误")
	}

	// 校验密码是否正确
	u := us.Items[0]
	if err := us.Items[0].CheckPassword(in.Password); err != nil {
		return nil, err
	}

	// 2. 正确的请求下 颁发用户令牌
	tk := token.NewToken(in.RemindMe)
	// 关联用户信息
	tk.UserId = fmt.Sprintf("%d", u.Id)
	tk.UserName = u.Username
	tk.Role = u.Role

	// 3. 保存用户Token到数据库中
	err = i.db.WithContext(ctx).
		Model(&token.Token{}).
		Create(tk).
		Error
	if err != nil {
		return nil, err
	}
	
	return tk, nil
}

// 退出：撤销令牌
func (i *TokenServiceImpl) RevokeToken(
	ctx context.Context,
	in *token.RevokeTokenRequest) (
	*token.Token, error) {
	return nil, nil
}

// 校验令牌
// 
func (i *TokenServiceImpl) ValidateToken(
	ctx context.Context,
	in *token.ValidateTokenRequest) (
	*token.Token, error) {
	return nil, nil
}
