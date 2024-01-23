# 对user模块的思考

1. interface.go：user包中定义
   + 这里定义了Service接口，里面定义了三个方法
   ```go
   type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)	
	QueryUser(context.Context, *QueryUserRequest) (*UserSet, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
   }
   ```
   + 这些方法规定了user所能进行的操作

2. impl.go: impl包中定义
   + 这里定义了一个【结构体】UserServiceImpl
   + 该结构体中只有一个字段，db *gorm.DB
   + 目的是从全局配置中拿到 db 数据库连接句柄
   + 在测试文件user_test.go中使用init()方法来初始化这个对象
   ```go
   func init() {
    i = impl.NewUserServiceImpl()
   }
   ```

3. user.go: impl包中定义
   + 这个文件中定义了user全部的业务逻辑。
   + 是【结构体】UserServiceImpl和【接口】Service的桥梁
   + 它实现了interface.go中定义的【接口】Service中的三个方法