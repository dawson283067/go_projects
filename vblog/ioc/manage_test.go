package ioc_test

import (
	"testing"

	"github.com/go_projects/vblog/ioc"
)

func TestManageGetAndRegistry(t *testing.T) {
	ioc.Controller().Registry("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))

	// ioc.Controller().Registry(token.AppName, &TestStruct{})
	// t.Log(ioc.Controller().Get(token.AppName))
	// 对象内部 自己去ioc获取依赖
	// i.tokenSvc = ioc.Controller().Get(token.AppName)

	// 断言使用
	ioc.Controller().Get("TestStruct").(*TestStruct).XXX()

	// ioc管理
	// ioc.Init()
	// ioc.Destroy()
}