package ioc

// 对象方法约束
type Object interface {
	Init()  error
	Destroy() error
}

