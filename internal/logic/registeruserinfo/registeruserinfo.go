package registeruserinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"demo3/utility"
	"github.com/gogf/gf/v2/util/grand"
	"strconv"
)

type sRegisterUserInfo struct{}

func init() {
	service.RegisterUserInfo(New())
}
func New() *sRegisterUserInfo {
	return &sRegisterUserInfo{}
}
func (s sRegisterUserInfo) RegisterUser(ctx context.Context, in model.RegisterUserModelInput) (out model.RegisterUserModelOutput, err error) {
	UserSalt := grand.S(10)
	in.Usersalt = UserSalt
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	CreateUserId, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterUserModelOutput{UserId: strconv.FormatInt(CreateUserId, 10)}, nil
}
