# ioc 简单实现

![ioc设计](../docs/ioc.drawio)

功能：
1. 对象托管：
    + 对象的注册
    + 对象获取
2. 对象统一管理
    + 配置（不做实现）
    + 初始化
    + 销毁


## 基于Map来实现一个这样的对象托管容器

### 单容器设计
基于Map来实现一个这样的对象托管容器
```go
container = map[string]any
```

注册的对象必须进行接口约束，要求它必须实现 对象统一管理的方法：
```go



// 对象统一管理
type IocObject interface {
    Init() error
    Destroy() error
}

type TestStruct struct {

}

func (t *TestStruct) Init() error {

}

func (t *TestStruct) Destroy() error {

}

func (t *TestStruct) XXX() error {

}

// container = map[string]IocObject
ioc.Regirstry("service name", &TestStruct)

// 启动的时候，完成对象的统一管理，循环容器里的所有对象，调用他们的Init()方法
ioc.Init()
```

一个map，不允许重名的，
比如：有一个模块叫token：
+ TokenServiceImpl
+ TokenApiHandler

```go
// 注册控制器
ioc.Controller.Registry("module name", &TokenServiceImpl{})
ioc.Api.Registry("module name", &TokenApiHandler{})
// 这样的话，就用一个模块的名字来注册多个模块的不同组成部分
```

根据程序设计，对这些对象的职责约束，将容器进行分区：
+ Api: 负责注册 Api实现类型的对象
+ Controller：负责注册服务区实现类的对象
+ Config: 配置对象，db，Kafka，redis
+ Default: 预留区域

### 多容器

```GO
api_container = map[string]IocObject
controller_container = map[string]IocObject
```

### 实现ioc

封装Container
```go
func TestContainerGetAndRegistry(t *testing.T) {
	c := ioc.NewContainer()
	c.Registry("TestStruct", &TestStruct{})
	t.Log(c.Get("TestStruct"))

	// 通过断言来使用
	c.Get("TestStruct").(*TestStruct).XXX()
}
```


封装Manager
```go
func TestManageGetAndRegistry(t *testing.T) {
	ioc.Controller().Registry("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))

	// 断言使用
	ioc.Controller().Get("TestStruct").(*TestStruct).XXX()
}
```




