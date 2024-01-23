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
	
	// 1.1 确认用户密码是否正确
	req := user.NewQueryUserRequest()
	req.Username = in.Username
	// 面向接口，没有具体对象存在。面向具体的业务逻辑，进行抽象编程
	us, err := i.user.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(us.Items) == 0 {
		return nil, fmt.Errorf("用户名或者密码错误")
	}

	// 1.2 校验密码是否正确
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

// 退出：撤销令牌，把这个令牌删除
// 明确结果返回
func (i *TokenServiceImpl) RevokeToken(
	ctx context.Context,
	in *token.RevokeTokenRequest) (
	*token.Token, error) {
	
	// 查询Token，到数据库里查询，返回一个token对象
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// refresh 确认
	err = tk.CheckRefreshToken(in.RefreshToken)
	if err != nil {
		return nil, err
	}

	// 删除Token
	// DELETE FROM `tokens` WHERE access_token = 'cmlcakka0uti117ngqp0' AND refresh_token = 'cmlcakka0uti117ngqpg'
	err = i.db.WithContext(ctx).
		Where("access_token = ?", in.AccessToken).
		Where("refresh_token = ?", in.RefreshToken).
		Delete(&token.Token{}).
		Error
	if err != nil {
		return nil, err
	}
	return tk, nil
}

// 校验令牌
// 需要确认这个令牌是颁发给谁的
func (i *TokenServiceImpl) ValidateToken(
	ctx context.Context,
	in *token.ValidateTokenRequest) (
	*token.Token, error) {
	
	// 1. 查询Token，判断令牌是否存在
	tk, err := i.getToken(ctx, in.AccessToken)
	if err != nil {
		return nil, err
	}

	// 2. 判断令牌是否过期
	if err := tk.ValidateExpired(); err != nil {
		return nil, err
	}

	// 3. 令牌合法返回令牌
	return tk, nil
}
