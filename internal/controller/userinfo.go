package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/consts"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var Userinfo = cUserinfo{}

type cUserinfo struct{}

func (a *cUserinfo) GetList(ctx context.Context, req *backapi.UserInfoListReq) (res *backapi.UserInfoListRes, err error) {
	listUser, count, err := dao.UserInfo.Ctx(ctx).Unscoped().AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.UserInfoListRes{
		List:  listUser,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, err
}
func (a *cUserinfo) CreateUser(ctx context.Context, req *backapi.UserInfoCreateReq) (res *backapi.UserInfoCreateRes, err error) {
	create, err := service.UserInfo().CreateUserInfo(ctx, model.UserInfoModelCreateInput{
		UserId:   req.UserId,
		UserName: req.UserName,
		NickName: req.NickName,
		Password: req.Password,
		Phone:    req.Phone,
		Wechat:   req.Wechat,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.UserInfoCreateRes{Userid: create.Userid}, nil
}

func (a *cUserinfo) UpdateUser(ctx context.Context, req *backapi.UserInfoUpdateReq) (res *backapi.UserInfoUpdateRes, err error) {
	_, err = service.UserInfo().UpdateUserInfo(ctx, model.UserInfoModelUpdateInput{
		UserId:   req.UserId,
		UserName: req.UserName,
		NickName: req.NickName,
		Phone:    req.Phone,
		Wechat:   req.Wechat,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.UserInfoUpdateRes{}, nil

}

func (a *cUserinfo) DeleteUser(ctx context.Context, req *backapi.UserInfoDeleteReq) (res *backapi.UserInfoDeleteRes, err error) {
	err = service.UserInfo().DeleteUserInfo(ctx, req.Userid)
	if err != nil {
		return nil, err
	}
	return &backapi.UserInfoDeleteRes{}, nil
}

// Info JWT验证
/*func (a *cUserinfo) Info(ctx context.Context, req *backapi.UserGetInfoReq) (res *backapi.UserGetInfoRes, err error) {
	return &backapi.UserGetInfoRes{
		Id:          gconv.Int(service.Auth().GetIdentity(ctx)),
		IdentityKey: service.Auth().IdentityKey,
		Payload:     service.Auth().GetPayload(ctx),
	}, nil
}
*/

func (a *cUserinfo) Info(ctx context.Context, req *backapi.UserGetInfoReq) (res *backapi.UserGetInfoRes, err error) {
	userId := gstr.StrEx(req.UserKey, consts.GTokenUserPrefix)
	get, err := g.Redis().Get(ctx, "GToken:"+consts.GTokenUserPrefix+userId)
	if err != nil {
		return nil, err
	}
	fmt.Println("获取redis", get)
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(get.String()), &resultMap)
	data := resultMap["data"].(map[string]interface{})
	return &backapi.UserGetInfoRes{UserId: userId, Username: data["username"]}, nil
}
