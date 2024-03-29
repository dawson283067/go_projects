package impl_test

import (
	"context"
	"testing"

	"github.com/go_projects/vblog/apps/user"
	"github.com/go_projects/vblog/apps/user/impl"
)

var (
	i   user.Service
	ctx = context.Background()
)

// 怎么引入被测试的对象？
func TestCreateUser(t *testing.T) {
	// 使用构造函数创建请求对象
	req := user.NewCreateUserRequest()
	// user.CreateUserRequest{} 这样写很容易出现空指针
	req.Username = "Tom"
	req.Password = "123456"
	req.Role = user.ROLE_ADMIN

	// 创建用户，PO到数据库中
	u, err := i.CreateUser(ctx, req)

	// 单元测试异常怎么处理
	// 直接报错中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}

	// 可以自己进行期望对比，进行单元测试报错
	if u == nil {
		t.Fatal("user not created")
	}

	// 正常打印对象
	t.Log(u)
}

func TestCreateVisitorUser(t *testing.T) {
	// 使用构造函数创建请求对象
	req := user.NewCreateUserRequest()
	// user.CreateUserRequest{} 这样写很容易出现空指针
	req.Username = "guest"
	req.Password = "123456"
	req.Role = user.ROLE_VISITOR

	// 创建用户，PO到数据库中
	u, err := i.CreateUser(ctx, req)

	// 单元测试异常怎么处理
	// 直接报错中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}

	// 可以自己进行期望对比，进行单元测试报错
	if u == nil {
		t.Fatal("user not created")
	}

	// 正常打印对象
	t.Log(u)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	us, err := i.QueryUser(ctx, req)
	// 直接报错中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}
	t.Log(us)
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescribeUserRequest("12")
	u, err := i.DescribeUser(ctx, req)
	// 直接报错，中断单元流程并且失败
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)

	t.Log(u.CheckPassword("123456"))
}

func init() {
	// 加载被测试对象，i 就是User Service接口的具体实现对象
	i = impl.NewUserServiceImpl()
}
