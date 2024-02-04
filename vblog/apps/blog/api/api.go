package api

import (
	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/ioc"
)

func init() {
	ioc.Api().Registry(blog.AppName, &blogApiHandler{})
}

// blog.Service接口，是直接注册给ioc，不需要对暴露
type blogApiHandler struct {
	svc blog.Service	
}

func (i *blogApiHandler) Init() error {
	i.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}

func (i *blogApiHandler) Destroy() error {
	return nil
}