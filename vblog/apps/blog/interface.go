package blog

import "context"

const (
	// 模块名称
	AppName = "blogs"
)

// Blog Service接口定义，CRUD
type Service interface {
	// 创建一个博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 获取博客列表
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// 获取博客详情
	DescribeBlog(context.Context, *DescribeBlogReqeust) (*Blog, error)
	// 更新博客
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	// 删除博客
	// 为了与GRPC保持一致，返回一个删除的对象
	DeleteBlog(context.Context, *DeleteBlogReqeust) (*Blog, error)
	// 文章状态修改，比如发布
	ChangedBlogStatus(context.Context, *ChangedBlogStatus) (*Blog, error)
	// 文章审核
	AuditBlog(context.Context, *AuditInfo) (*Blog, error)
}

type QueryBlogRequest struct {

}

type DescribeBlogReqeust struct {

}

type UpdateBlogRequest struct {

}

type DeleteBlogReqeust struct {

}

