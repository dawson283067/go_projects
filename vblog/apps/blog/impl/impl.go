package impl

import (
	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/conf"
	"github.com/go_projects/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(blog.AppName, &blogServiceImpl{})
}

// blog.Service接口，是直接注册给ioc，不需要对暴露
type blogServiceImpl struct {
	// 依赖了一个数据库操作的链接池对象
	db *gorm.DB
}

func (i *blogServiceImpl) Init() error {
	i.db = conf.C().DB()
	return nil
}

func (i *blogServiceImpl) Destroy() error {
	return nil
}