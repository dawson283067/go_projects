package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/tools/pretty"
)

var (
	v = validator.New()
)

// 面向对象
// user.Service，设计这个模块提供的接口
// 接口定义，一定要考虑兼容性，接口的参数不能变
type Service interface {
	// 用户创建
	// CreateUser(username, password, role string, label map[string]string)
	// 设计CreateUserRequest，可以扩展对象，而不影响接口的定义
	// 返回值尽量用对象来包装
	// 1. 这个接口支持取消吗？要支持取消应该怎么办？
	// 2. 这个接口支持Trace，TraceId怎么传递？
	// 中间件参数，取消/Trace/... 怎么产生怎么传递
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 查询用户列表，对象列表 [{}]
	// 这里返回了*UserSet，目的是返回一个对象，里面可以添加更多参数，方便分页等业务操作
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	// 查询用户详情，通过Id查询
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)

	// 作业：
	// 用户修改
	// 用户删除
}

// 为了避免对象内部出现很多空指针，指针对象未初始化，为该对象提供一个构造函数
// 还能做一些相关兼容，补充默认值的功能，New+对象名称()
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_MEMBER,
		Label: map[string]string{},
	}
}

// 用户创建的参数
type CreateUserRequest struct {
	Username string `json:"username" validate:"required" gorm:"column:username"`
	Password string `json:"password" validate:"required" gorm:"column:password"`
	Role     Role   `json:"role" validate:"required" gorm:"column:role"`
	// https://gorm.io/docs/serializer.html
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
	// 把map序列化，然后放到label的字段里。如果没有serializer，数据库是不知道怎么放这种字段的
}

// 校验：用validator和struct tag来完成校验
func (req *CreateUserRequest) Validate() error {
	// if req.Username == "" {
	// 	return fmt.Errorf("username required")
	// }
	// if req.Password == "" {
	// 	return fmt.Errorf("password required")
	// }

	// validator库，validator.New() 校验器对象，全局单例模式
	// 也可以定义validator，比较麻烦。也可以使用validator，再写自己的校验规则，结合
	return v.Struct(req)
}

func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

// 查询用户列表
type QueryUserRequest struct {
	// 分页大小，一页多少个
	PageSize int
	// 当前页，查询那一页的数据
	PageNumber int
	// 根据用户name查找用户
	Username string
}

func (req *QueryUserRequest) Limit() int {
	return req.PageSize
}

// 1, 0
// 2, 20
// 3, 20 * 2
// 4，20 * 3
func (req *QueryUserRequest) Offset() int {
	return req.PageSize * (req.PageNumber - 1)
}

func NewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

type UserSet struct {
	// 总共有多少个
	Total int64 `json:"total"`
	// 当前查询的数据清单
	Items []*User `json:"items"`
}

func (u *UserSet) String() string {
	return pretty.ToJSON(u)
}

func NewDescribeUserRequest(uid int) *DescribeUserRequest {
	return &DescribeUserRequest{
		UserId: uid,
	}
}

type DescribeUserRequest struct {
	UserId int
}
