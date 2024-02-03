package blog

import "github.com/infraboard/mcube/tools/pretty"

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreateBlogRequest:        req,
		ChangedBlogStatusRequest: NewChangedBlogStatusRequest(),
		AuditInfo:                NewAuditInfo(),
	}
}

// 系统生成
// `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '文章的Id',
// `created_at` int NOT NULL COMMENT '创建时间',
// `updated_at` int NOT NULL COMMENT '更新时间',
type Blog struct {
	// 博客Id
	Id int `json:"id" gorm:"column:id"`
	// 创建时间，时间戳 10位，秒
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新时间，时间戳 10位，秒
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// 用户参数
	*CreateBlogRequest
	// 发布
	*ChangedBlogStatusRequest
	// 审核
	*AuditInfo
}

func (u *Blog) String() string {
	return pretty.ToJSON(u)
}

func NewChangedBlogStatusRequest() *ChangedBlogStatusRequest {
	return &ChangedBlogStatusRequest{
		Status: STATUS_DRAFT,
	}
}

// 发布才能修改文章状态
// `published_at` int NOT NULL COMMENT '发布时间',
// `status` tinyint NOT NULL COMMENT '文章状态：草稿/已发布',
type ChangedBlogStatusRequest struct {
	// 发布时间
	PublishedAt int64 `json:"published_at" gorm:"column:published_at"`
	// 文章状态：草稿/已发布
	Status Status `json:"status" gorm:"column:status"`
}

func NewAuditInfo() *AuditInfo {
	return &AuditInfo{}
}

// 审核相关字段
// `audit_at` int NOT NULL COMMENT '审核时间',
// `is_audit_pass` tinyint NOT NULL COMMENT '是否审核通过',
type AuditInfo struct {
	// 审核时间
	AuditAt int64 `json:"audit_at" gorm:"audit_at"`
	// 是否审核通过
	IsAuditPass bool `json:"is_audit_pass" gorm:"is_audit_pass"`
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

// `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章标题',
// `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
// `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章内容',
// `summary` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文章概要信息',
// `create_by` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
// `tags` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签',
type CreateBlogRequest struct {
	// 文章标题
	Title string `json:"title" gorm:"title" validate:"required"`
	// 文章作者
	Author string `json:"author" gorm:"author" validate:"required"`
	// 文章内容
	Content string `json:"content" gorm:"content" validate:"required"`
	// 文章概要信息
	Summary string `json:"summary" gorm:"summary"`
	// 创建人
	CreateBy string `json:"create_by" gorm:"create_by"`
	// 标签
	// https://gorm.io/docs/serializer.html
	// 把map序列化，然后放到tags的字段里。如果没有serializer，数据库是不知道怎么放这种字段的
	Tags map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

func (req *CreateBlogRequest) Validate() error {
	return v.Struct(req)
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	// 总共有多少个
	Total int64 `json:"total"`
	// 当前查询的数据清单
	Items []*Blog `json:"items"`
}