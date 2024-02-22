package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/apps/token"
	"github.com/go_projects/vblog/common"
	"github.com/go_projects/vblog/response"
)

// + 创建博客：POST /vblogs/api/v1/blogs
// required('username','admin') 伪代码，基于用户判断
// required(visitor/admin) 伪代码，RBAC
func (h *blogApiHandler) CreateBlog(c *gin.Context) {

	req := blog.NewCreateBlogRequest()

	// 后面请求如何获取 中间信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}

	if err := c.BindJSON(req); err != nil {
		response.Failed(c, err)
		return
	}
	ins, err := h.svc.CreateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 修改博客（部分）：PATCH /vblogs/api/v1/blogs/:id
// /vblogs/api/v1/blogs/10 --> id = 10
// /vblogs/api/v1/blogs/10 --> id = 20
// c.Param("id") 获取路径变量的值
func (h *blogApiHandler) PatchBlog(c *gin.Context) {
	// 如何解析路径里面的参数
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PATCH


	// 用户传递的数据
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}

	
	// 后面请求如何获取 中间信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}

	ins, err := h.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 修改博客（全量）：PUT /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) UpdateBlog(c *gin.Context) {
	// 如何解析路径里面的参数
	req := blog.NewUpdateBlogRequest(c.Param("id"))
	req.UpdateMode = common.UPDATE_MODE_PUT
	// 用户传递的数据
	if err := c.BindJSON(req.CreateBlogRequest); err != nil {
		response.Failed(c, err)
		return
	}

	
	// 后面请求如何获取 中间信息
	if v, ok := c.Get(token.TOKEN_MIDDLEWARE_KEY); ok {
		req.CreateBy = v.(*token.Token).UserName
	}
	
	ins, err := h.svc.UpdateBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 删除博客：DELETE /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) DeleteBlog(c *gin.Context) {
	req := blog.NewDeleteBlogRequest(c.Param("id"))
	ins, err := h.svc.DeleteBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// + 查询列表：GET /vblogs/api/v1/blogs?page_size=10&page_number=2
func (h *blogApiHandler) QueryBlog(c *gin.Context) {
	req := blog.NewQueryBlogRequestFromGin(c)
	set, err := h.svc.QueryBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, set)
}

// + 查询详情：GET /vblogs/api/v1/blogs/:id
func (h *blogApiHandler) DescribeBlog(c *gin.Context) {
	req := blog.NewDescribeUserRequest(c.Param("id"))
	ins, err := h.svc.DescribeBlog(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}
