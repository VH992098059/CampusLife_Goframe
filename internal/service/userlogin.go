package service

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type userLogin struct{}

var user = userLogin{}

func User() *userLogin {
	return &user
}

// GetUserByUserNamePassword JWT验证
func (s *userLogin) GetUserByUserNamePassword(ctx context.Context, in model.LoginInfoModelInput) map[string]interface{} {
	UserInfo := model.UserInfo{}
	_ = dao.UserInfo.Ctx(ctx).Where("username", in.Username).Scan(&UserInfo)
	if in.Username == UserInfo.UserName && utility.EncryptPassword(in.Password, UserInfo.Usersalt) == UserInfo.Password {
		return g.Map{
			"id":       UserInfo.UserId,
			"username": UserInfo.UserName,
		}
	}
	return nil
}
