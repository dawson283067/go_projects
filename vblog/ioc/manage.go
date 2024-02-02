package ioc

import "github.com/gin-gonic/gin"

// 对象的统一管理方法
// 全局唯一的，全局变量
// ioc.NC.Init() 读起来不优雅，没必要把这个概念暴露出去
var nc = &NamespaceContainer{
	ns: map[string]*Container{
		"api": NewContainer(),
		"controller":NewContainer(),
		"config":NewContainer(),
		"default":NewContainer(),
	},
}

func Init() error {
	return nc.Init()
}

func Destroy() error {
	return nc.Destroy()
}

// ioc.Controller().Registry()
// ioc.Controller().Get()
func Controller() *Container {
	return nc.ns["controller"]
}

// ioc.Api().Registry()
// ioc.Api().Get()
func Api() *Container {
	return nc.ns["api"]
}

func RegistryGinApi(rr gin.IRouter) {
	nc.RegistryGinApi(rr)
}

// 基于这个构建多空间的container
type NamespaceContainer struct {
	ns map[string]*Container
}

func (c *NamespaceContainer) Init() error {

	// 遍历Namespace
	for key := range c.ns {
		c := c.ns[key]
		// 遍历Namespace里面的对象
		for objectName := range c.storage {
			if err := c.storage[objectName].Init(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *NamespaceContainer) Destroy() error {

	// 遍历Namespace
	for key := range c.ns {
		c := c.ns[key]
		// 遍历Namespace里面的对象
		for objectName := range c.storage {
			if err := c.storage[objectName].Destroy(); err != nil {
				return err
			}
		}
	}
	return nil
}

type GinApi interface {
	// 基础约束
	Object
	// 额外约束
	// ioc.Api().Get(token.AppName).(*api.TokenApiHandler).Registry(rr)
	Registry(rr gin.IRouter)
}

// 遍历所有的对象，帮忙完成Api的注册
// 由ioc调用对象提供的Registry方法，来把模块的api 注册给gin root router
func (c *NamespaceContainer) RegistryGinApi(rr gin.IRouter) {
	// 遍历Namespace
	for key := range c.ns {
		c := c.ns[key]
		// 遍历Namespace里面的对象
		for objectName := range c.storage {
			obj := c.storage[objectName]
			// 如果判断一个对象不是GinApi对象（约束）
			// 判断对象有没有Registry(rr gin.IRouter)
			// 断言该对象 满足GinApi接口，实现了Registry函数
			if v, ok := obj.(GinApi); ok {
				v.Registry(rr)
			}
		}
	}
}
