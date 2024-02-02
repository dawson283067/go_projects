package ioc_test

import (
	"testing"

	"github.com/go_projects/vblog/ioc"
)

func TestManageGetAndRegistry(t *testing.T) {
	ioc.Controller().Registry("TestStruct", &TestStruct{})
	t.Log(ioc.Controller().Get("TestStruct"))

	// 断言使用
	ioc.Controller().Get("TestStruct").(*TestStruct).XXX()
}