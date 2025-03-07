package userinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sUserInfo struct {
}

func (s sUserInfo) Person(ctx context.Context, in model.UserGetPersonMsgModelInput) (out model.UserGetPersonMsgModelOutput, err error) {
	list, err := dao.UserState.Ctx(ctx).Where("user_id", in.UserId).One()
	if err != nil {
		return out, err
	}
	return model.UserGetPersonMsgModelOutput{List: list}, nil
}

func init() {
	service.RegisterGetUser(New())
}
func New() *sUserInfo { return &sUserInfo{} }
