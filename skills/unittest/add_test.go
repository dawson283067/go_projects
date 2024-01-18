package unittest_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/go_projects/skills/unittest"
)

// 这个单元测试需要读取外部变量?
// 这里的ARG1/ARG2如何传递给单元测试
// run test: go.exe test -timeout 30s -run ^TestSum$ github.com/go_projects/skills/unittest -count=1 -v
// 当我们点击 debug test 时，怎么注入自定义环境变量
func TestSum(t *testing.T) {
	// read file，很难统一路径问题：文件有相对路径和绝对路径
	// 使用环境变量
	a1 := os.Getenv("ARG1")
	a2 := os.Getenv("ARG2")
	a1I, _ := strconv.Atoi(a1)
	a2I, _ := strconv.Atoi(a2)
	t.Log(unittest.Sum(a1I, a2I))
}
