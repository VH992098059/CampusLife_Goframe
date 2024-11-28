package userinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"demo3/utility"
	"github.com/gogf/gf/v2/util/grand"
)

type sUserInfo struct{}

func init() {
	service.RegisterUserInfo(New())
}
func New() *sUserInfo {
	return &sUserInfo{}
}

func (s *sUserInfo) GetList(ctx context.Context, in model.UserInfoModelInput) (out *model.UserInfoModelOutput, err error) {
	out = &model.UserInfoModelOutput{
		Page: in.Page,
		Size: in.Size,
	}
	//不过滤软删除查询数据（Unscoped）
	list, err := dao.UserInfo.Ctx(ctx).Unscoped().All()
	if err != nil {
		return
	}

	return &model.UserInfoModelOutput{List: list}, nil
}

func (s *sUserInfo) CreateUserInfo(ctx context.Context, in model.UserInfoModelCreateInput) (out model.UserInfoModelCreateOutput, err error) {
	UserSalt := grand.S(10)
	in.Usersalt = UserSalt
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	CreateUserId, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserInfoModelCreateOutput{Userid: uint(CreateUserId)}, nil
}

func (s *sUserInfo) UpdateUserInfo(ctx context.Context, in model.UserInfoModelUpdateInput) (out model.UserInfoModelUpdateOutput, err error) {

	if in.Password != "" {
		UserSalt := grand.S(10)
		in.Password = utility.EncryptPassword(in.Password, UserSalt)
		in.Usersalt = UserSalt
	}
	//_, err = dao.UserInfo.Ctx(ctx).Data(g.Map{"username": in.UserName, "phone": in.Phone, "wechat": in.Wechat}).Where("user_id", in.UserId).Update()
	//过滤某个字段修改数据（FieldEX）
	_, err = dao.UserInfo.Ctx(ctx).Data(in).FieldsEx(dao.UserInfo.Columns().UserId).Where(dao.UserInfo.Columns().UserId, in.UserId).Update()
	if err != nil {
		return out, err
	}
	return model.UserInfoModelUpdateOutput{}, err
}
func (s *sUserInfo) DeleteUserInfo(ctx context.Context, id string) error {
	_, err := dao.UserInfo.Ctx(ctx).Where("user_id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
