package user_test

import (
	"testing"

	"github.com/go_projects/vblog/apps/user"
)

// $2a$10$EoGVEGJL3HUnptnN/Nc0ZOyPKJQ91x3IOlx6d5aeDRw.UHhFfOUlK
// $2a$10$x9nCIHuWadFW2WRsnLu1JO8X2XytaR/FWrYy4q0sSWuk4ps0iEY/y
// https://gitee.com/infraboard/go-course/blob/master/day09/go-hash.md#bcrypt
func TestHashedPassword(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Password = "123456"
	req.HashedPassword()
	t.Log(req.Password)

	t.Log(req.CheckPassword("1234561"))
}
