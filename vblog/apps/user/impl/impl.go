package impl

import (
	"github.com/go_projects/vblog/conf"
	"gorm.io/gorm"
)

// 构造函数
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{
		// 获取全局的DB对象
		// 前提：配置对象准备完成
		db: conf.C().DB(),
	}
}

// 怎么实现user.Service接口？
// 定义UserServiceImpl来实现接口
type UserServiceImpl struct {
	// 依赖了一个数据库操作的连接池对象
	db *gorm.DB
}
