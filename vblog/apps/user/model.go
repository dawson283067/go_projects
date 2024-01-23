package user

import (
	"time"

	"github.com/infraboard/mcube/tools/pretty"
)

// 存放需要入库的数据结构（PO，Persistent Object）
// 构造User对象的时候，就需要把明文密码转化为hash后的密码
func NewUser(req *CreateUserRequest) *User {
	// hash密码
	req.HashedPassword()

	return &User{
		CreatedAt: time.Now().Unix(),
		// 这里不属于更新，所以就不用给UpdateAt字段
		CreateUserRequest: req,
	}
}

// 用户创建成功后返回一个User对象
// CreatedAt 为啥没用time.Time，int64(TimeStamp)，统一标准化，避免时区对程序产生影响
// 在需要对时间进行展示的时候，由前端根据具体展示那个时区的时间
type User struct {
	// 用户Id
	Id int `json:"id" gorm:"column:id"`
	// 创建时间，时间戳 10位，秒
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新时间，时间戳 10位，秒
	UpdatedAt int64 `json:"updateted_at" gorm:"column:updated_at"`

	// 用户参数
	*CreateUserRequest
}

// TableName() string
// 定义对象存储的表的名称
func (u *User) TableName() string {
	return "users"
}

func (u *User) String() string {
	// 调用了老喻的库
	return pretty.ToJSON(u)
}
