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

	// 还需要依赖user.Service
	// 依赖接口，不要依赖接口的具体实现
	user user.Service
}
