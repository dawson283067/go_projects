package blog

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go_projects/vblog/common"
)

const (
	// 模块名称
	AppName = "blogs"
)

var (
	v = validator.New()
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
	ChangedBlogStatus(context.Context, *ChangedBlogStatusRequest) (*Blog, error)
	// 文章审核
	AuditBlog(context.Context, *AuditInfo) (*Blog, error)
}

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize: 20,
		PageNumber: 1,
	}
}

func NewQueryBlogRequestFromGin(c *gin.Context) *QueryBlogRequest {
	req := NewQueryBlogRequest()
	ps := c.Query("page_size")
	req.CreateBy = c.Query("create_by")
	if ps != "" {
		req.PageSize, _ = strconv.Atoi(ps)
	}
	pn := c.Query("page_number")
	if pn != "" {
		req.PageNumber, _ = strconv.Atoi(pn)
	}
	return req
}

type QueryBlogRequest struct {
	// 分页大小，一个多少个
	PageSize int
	// 当前页，查询哪一页的数据
	PageNumber int
	// 谁创建的文章
	CreateBy string
}

func (req *QueryBlogRequest) Limit() int {
	return req.PageSize
}

// 1, 0          第一页，偏移是0
// 2, 20         第二页，偏移是20 * 1
// 3, 20 * 2     第三页，偏移是20 * 2
// 4，20 * 3     第四页，偏移是20 * 3
func (req *QueryBlogRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

func NewDescribeUserRequest(id string) *DescribeBlogReqeust {
	return &DescribeBlogReqeust{
		Id: id,
	}
}

type DescribeBlogReqeust struct {
	Id string
}

func NewUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		Id: id,
		UpdateMode: common.UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

type UpdateBlogRequest struct {
	// 被更新的博客Id
	Id string `json:"id"`
	// 更新模式
	UpdateMode common.UpdateMode `json:"update_mode"`
	// 更新时的数据
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogReqeust {
	return &DeleteBlogReqeust{
		Id: id, 
	}
}

type DeleteBlogReqeust struct {
	Id string
}

