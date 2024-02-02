package ioc

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