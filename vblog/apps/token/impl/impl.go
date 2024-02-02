package impl

import (
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"
	"gorm.io/gorm"
)

var (
	_ token.Service = (*TokenServiceImpl)(nil)
)

func init() {
	ioc.Controller().Registry(token.AppName, &TokenServiceImpl{})
}

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

// 对象属性初始化
func (i *TokenServiceImpl) Init() error {
	// 凡是结构体属性需要补充的，都需要在这里去做
	// 依赖关系先在import里解决，然后这里初始化对象是没有问题的
	i.db = conf.C().DB()
	// 这句的语法没看懂......
	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}

func (i *TokenServiceImpl) Destroy() error {
	return nil
}
