package impl

import (
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/conf"
	"gorm.io/gorm"
)

var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func NewTokenServiceImpl(userServiceImpl user.Service) *TokenServiceImpl {
	return &TokenServiceImpl{
		// 获取全局的DB对象
		// 前提：配置对象准备完成
		db:   conf.C().DB(),
		user: userServiceImpl,
	}
}

// 怎么实现token.Service接口？
// 定义TokenServiceImpl来实现接口
type TokenServiceImpl struct {
	// 依赖了一个数据库操作的连接池对象
	db *gorm.DB

	// 额外需要依赖user.Service
	// 没有 UserServiceImpl 的具体实现，这样就不跟具体的实现挂钩
	// 依赖接口，不要依赖接口的具体实现
	// 实现的时候是由组装的程序给你实现是什么，比如说main程序
	user user.Service
}
