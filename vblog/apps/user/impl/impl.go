package impl

import (
	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"
	"gorm.io/gorm"
)

// 通过 import 自动完成注册
// 为什么不能直接在这里把db对象初始化了？
// 这个逻辑是在import时执行的，程序在import后才执行的配置模块加载，此时的conf.C()为null
// 这个db属性的初始化一定要在配置加载后执行：conf.Load(), ioc.Init()
func init() {
	ioc.Controller().Registry(user.AppName, &UserServiceImpl{})
}

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

func (i *UserServiceImpl) Init() error {
	// import后，在初始化对象的时候再来获取db对象
	i.db = conf.C().DB()
	return nil
}

func (i *UserServiceImpl) Destroy() error {
	// delete map 自己的维护的map
	return nil
}
