package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/consts"
	"demo3/internal/model"
	"demo3/internal/service"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var Userinfo = cUserinfo{}

type cUserinfo struct{}

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
	var resultMap map[string]interface{}
	if err = json.Unmarshal([]byte(get.String()), &resultMap); err != nil {
		return nil, err
	}
	data := resultMap["data"].(map[string]interface{})
	reNickName, ok := data["nickname"]
	if !ok {
		return nil, errors.New("nickname未找到")
	}
	reAvatar, ok := data["avatar"]
	if !ok {
		return nil, errors.New("avatar未找到")
	}
	return &backapi.UserGetInfoRes{
		UserId:   userId,
		NickName: gconv.String(reNickName),
		Avatar:   gconv.String(reAvatar),
	}, nil
}
func (a *cUserinfo) GetPerson(ctx context.Context, req *backapi.UserGetPersonMsgReq) (res *backapi.UserGetPersonMsgRes, err error) {
	list, err := service.UserInfo().Person(ctx, model.UserGetPersonMsgModelInput{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return &backapi.UserGetPersonMsgRes{List: list.List}, nil
}
