package impl

import (
	"context"

	"github.com/go_projects/vblog/apps/user"
)

// 实现 user.Service
// 怎么判断这个服务有没有实现这个接口呢？
// &UserServiceImpl{} 是会分配内存，怎么才能不分配内存?
// var _ user.Service = &UserServiceImpl{}
// nil 如何声明 *UserServiceImpl 的nil
// (*UserServiceImpl)(nil) --> int8 1 int32(1)  (int32)(1)
// nil就是一个*UserServiceIpl的空指针
var _ user.Service = (*UserServiceImpl)(nil)

// 用户创建
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	in *user.CreateUserRequest) (
	*user.User, error) {
	// 1. 校验用户请求
	if err := in.Validate(); err != nil {
		return nil, err
	}

	// 2. 创建用户实例对象
	u := user.NewUser(in)

	// Hash完成后入库

	// 3. 把对象持久化（放到数据库里）
	// orm: orm 需要定义这个对象 存放在哪个表里面，以及struct和数据库表中字段的映射关系
	// object ----> row
	// INSERT INTO `users` (`created_at`,`updated_at`,`username`,`password`,`role`,`label`) VALUES (1705630637,1705630637,'admin','123456','1','{}')
	// 比如 craate user , 4秒的收，请求还没返回，用户就取消了请求，后端会因为请求退出而结束吗？
	// 程序里 并没有中断数据库操作的能力，通过WithContext携带上ctx
	if err := i.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}

	// 4. 返回创建好的对象
	return u, nil
}

// 查询用户列表，对象列表 [{}]
// 这里返回了*UserSet，目的是返回一个对象，里面可以添加更多参数，方便分页等业务操作
// 查询数据库里 多行记录
func (i *UserServiceImpl) QueryUser(
	ctx context.Context,
	in *user.QueryUserRequest) (
	*user.UserSet, error) {
	// 构造一个MySQL 条件查询语句 select * from users where ...
	query := i.db.WithContext(ctx).Model(&user.User{})

	// 构造条件 where username = ""
	if in.Username != "" {
		// query 会生成一个新的语句，不会修改query对象本身
		query = query.Where("username = ?", in.Username)
	}

	set := user.NewUserSet()

	// 统计当前有多个
	// select COUNT(*) from user where ...
	err := query.
		Count(&set.Total).
		Error
	if err != nil {
		return nil, err
	}

	// 做真正的分页查询：sql LIMIT OFFSET
	// LIMIT 20,20 这个是查询的是第2页
	// 使用Find把多行数据查询出来，使用[]User 来接收返回
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

// 查询用户详情，通过Id查询
func (i *UserServiceImpl) DescribeUser(
	ctx context.Context,
	in *user.DescribeUserRequest) (
	*user.User, error) {
	// 构造一个MySQL 条件查询语句 select * from users where ...
	query := i.db.WithContext(ctx).Model(&user.User{}).Where("id = ?", in.UserId)

	// 准备一个对象 接收数据库的返回
	u := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
