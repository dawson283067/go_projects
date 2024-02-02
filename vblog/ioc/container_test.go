package ioc_test

import (
	"fmt"
	"testing"

	"github.com/go_projects/vblog/ioc"
)

func TestContainerGetAndRegistry(t *testing.T) {
	c := ioc.NewContainer()
	c.Registry("TestStruct", &TestStruct{})
	t.Log(c.Get("TestStruct"))

	// 通过断言来使用
	c.Get("TestStruct").(*TestStruct).XXX()
}

type TestStruct struct {

}

func (t *TestStruct) Init() error {
	return nil
}

func (t *TestStruct) Destroy() error {
	return nil
}

func (t *TestStruct) XXX() error {
	fmt.Println("xxx log")
	return nil
}