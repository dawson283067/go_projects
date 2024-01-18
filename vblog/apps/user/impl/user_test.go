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
	// 单元测试异常怎么处理
	u, err := i.CreateUser(ctx, nil)
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

}

func TestDescribeUser(t *testing.T) {

}

func init() {
	// 加载被测试对象，i 就是User Service接口的具体实现对象
	i = &impl.UserServiceImpl{}
}
