package impl

import (
	"context"

	"github.com/go_projects/vblog/apps/blog"
	"github.com/go_projects/vblog/common"
	"github.com/go_projects/vblog/exception"
	"dario.cat/mergo"
)

// 创建一个博客
func (i *blogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	// 1. 校验请求
	if err := req.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("创建博客失败，%s", err)
	}
	req.Validate()

	// 2. 构造对象
	ins := blog.NewBlog(req)

	// 3. 对象入库
	// INSERT INTO `blogs` (`created_at`,`updated_at`,`title`,`author`,`content`,`summary`,`create_by`,`tags`,`published_at`,`status`,`audit_at`,`is_audit_pass`) VALUES (1706933718,1706933718,'go语言全栈开发','oldyu','xxx','xx','','{"目录":"Go语言"}',0,'0',0,false)
	err := i.db.WithContext(ctx).Create(ins).Error
	if err != nil {
		return nil, err
	}

	// 4. 返回对象
	return ins, err
	
	// fmt.Println("CreateBlog")
}
	
// 获取博客列表
func (i *blogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	// 
	set := blog.NewBlogSet()

	// 1. 初始化查询对象
	query := i.db.WithContext(ctx).Model(blog.Blog{})

	// 查询总数
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	// 查询具体的数据
	err = query.
		Limit(in.Limit()).
		Offset(in.Offset()).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}
	
	return set, nil
}
// 获取博客详情
func (i *blogServiceImpl) DescribeBlog(ctx context.Context, in *blog.DescribeBlogReqeust) (*blog.Blog, error) {
		
	// 构造一个MySQL 条件查询语句 select * from users where ...
	query := i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", in.Id)

	// 准备一个对象 接收数据库的返回
	ins := blog.NewBlog(blog.NewCreateBlogRequest())
	if err := query.First(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil	
}
// 更新博客
// 1, 全量更新：对象的替换
// 2, 部分更新：(old obj)Patch机制 --> new obj --> save
func (i *blogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	// 查询老的对象，需要被更新的博客对象
	ins, err :=	i.DescribeBlog(ctx, blog.NewDescribeUserRequest(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case common.UPDATE_MODE_PATCH:
		// if req.Author != "" {
		// 	ins.Author = req.Author
		// }
		// if req.Title != "" {
		// 	ins.Title = req.Title
		// }
		// ... 有没有其他的方法 帮我们完成2个结构体的合并 merge(patch)
		// https://github.com/darcio/mergo
		// // WithOverride will make merge override non-empty dst attribues with non-empty src attributes values.
		if err := mergo.MapWithOverwrite(ins.CreateBlogRequest, req.CreateBlogRequest); err != nil {
			return nil, err
		}
	default:
		// 按道理应该是 *CreateBlogRequest = *CreateBlogRequest
		ins.CreateBlogRequest = req.CreateBlogRequest
	}

	// 再次校验对象。有可能更新过后的数据不合法，所以校验。
	if err := ins.Validate(); err != nil {
		return nil, exception.ErrBadRequest.WithMessagef("校验更新请求失败：%s", err)
	}

	// 更新数据库
	// UPDATE `blogs` SET `id`=48,`created_at`=1707013985,`updated_at`=1707015070,`title`='go语言全栈开发V2',`author`='oldyu',`content`='xxx',`summary`='xx',`tags`='{"目录":"Go语言"}' WHERE id = 48
	err = i.db.WithContext(ctx).Model(&blog.Blog{}).Where("id = ?", ins.Id).Updates(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 删除博客
// 为了与GRPC保持一致，返回一个删除的对象
func (i *blogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogReqeust) (*blog.Blog, error) {
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeUserRequest(req.Id))
	if err != nil {
		return nil, err
	}

	err = i.db.
		WithContext(ctx).
		Model(&blog.Blog{}).
		Where("id = ?", req.Id).
		Delete(ins).
		Error
	if err != nil {
		return nil, err
	}
	
	return ins, nil
}
// 文章状态修改，比如发布
func (i *blogServiceImpl) ChangedBlogStatus(ctx context.Context, req *blog.ChangedBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}
// 文章审核
func (i *blogServiceImpl) AuditBlog(ctx context.Context, req *blog.AuditInfo) (*blog.Blog, error) {
	return nil, nil
}