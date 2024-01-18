# 单元测试

## 如何编写单元测试

针对某个功能（函数），来验证功能的正确性

单元测试要符合条件：编写一个单元测试函数来进行目标函数的测试（unittest.Sum）

+ 单元测试的函数需要放到一个包里，独立一个测试包，unittest_test。目标测试函数的名称TestSum。unittest_test.TestSum
+ 默认规则：
  + 单元测试文件的名称：xxx_test.go 
    + xxx.go 是被测试的文件名称
  + 测试包的名称：xxxx_test
    + 同一个目录下面，可以再独立存在关于这个包的测试包，xxxx是被测试的包名（这里是unittest）。这里的测试包名：unittest_test
  + 函数签名：TestXxxxx(t testing.T)
    + Xxxxx是被测试的函数（功能），这里是TestSum(t testing.T)

```go
func TestSum(t *testing.T) {
	
}
```

## 单元测试如何配置

1. vscode 单元测试读取环境变量的配置文件路径配置

File --> Preferences --> Settings --> (User)Extensions --> Go --> Edit in settings.json

Test Env File --> ${workspaceFolder}/etc/unit_test.env

```sh
# 在settings.json中会出现如下内容
{
  "go.testEnvFile": "${workspaceFolder}/etc/unit_test.env",
}
```

2. 在文件里补充需要注入的环境变量
```env
ARG1=1
ARG2=2
```

3. 单元测试里 使用该环境变量
```go
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
```


