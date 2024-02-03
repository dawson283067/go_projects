package impl_test

import (
	"testing"

	"github.com/go_projects/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "go语言全栈开发"
    req.Author = "oldyu"
	req.Content = "xxx"
	req.Summary = "xx"
	req.Tags["目录"] = "Go语言"
	ins, err := impl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	set, err := impl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDescribeUserRequest("47")
	set, err := impl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("47")
	ins, err := impl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}