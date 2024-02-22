package user

// 什么是枚举？为什么要使用枚举？
// 枚举 把所有的可选项 一一列举出来
// Role，Admin/Member  Role string ----> "", 看代码的人 不知道 Role到底应该传什么值
// 到底有哪些值 可以传递，或者是需要程序列举出来的
// 枚举核心能力：约束 只能传递 列举出来的值，其他的值都不允许传递

// 通过声明一种自定义类型来声明一种类型
type Role int

// 通过定义满足类型的常量，来定义满足这个类型的列表
// ROLE_MEMBER/ROLE_ADMIN
const (
	// 当 值为0的时候，就是默认值
	// 枚举命名风格：
	ROLE_VISITOR Role = iota
	ROLE_ADMIN
)
