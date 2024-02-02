package ioc

// ioc 容器具体实现

func NewContainer() *Container {
	return &Container{
		storage: map[string]Object{},
	}
}

type Container struct {
	storage map[string]Object
}

func (c *Container) Registry(name string, obj Object){
	c.storage[name] = obj
}

// 获取的值，由使用者进行约束，或者断言
// ioc.Get("module name").(*TestService)
func (c *Container) Get(name string) any {
	return c.storage[name]
}

// 提供一个对象遍历的方法
// 使用者传递一个函数进来，把遍历的对象作为参数交给这个函数（回调函数）
// func (c *Container) ForEach(cb func(ObjectName string, objectValue Object)) {
// 	for key := range c.storage {
// 		cb(key, c.storage[key])
// 	}
// }
