package impl_test

import (
	"testing"

	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/common"
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
	req := blog.NewDescribeUserRequest("48")
	set, err := impl.DescribeBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

// 之前的数据
// req.Title = "go语言全栈开发"
// req.Author = "oldyu"
// req.Content = "xxx"
// req.Summary = "xx"
// req.Tags["目录"] = "Go语言"
func TestUpdateBlogPatchMode(t *testing.T) {
	req := blog.NewUpdateBlogRequest("48")
	req.UpdateMode = common.UPDATE_MODE_PATCH
	req.Title = "go语言全栈开发V2"
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlogPutMode(t *testing.T) {
	req := blog.NewUpdateBlogRequest("48")
	// 这里Mode用默认值，是全量替换
	req.Title = "go语言全栈开发V3"
	req.Content = "v3"
	req.Author = "v3"
	ins, err := impl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("47")
	ins, err := impl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

