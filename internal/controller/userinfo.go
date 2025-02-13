package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/consts"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
	fmt.Println("获取redis", get)
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(get.String()), &resultMap)
	return &backapi.UserGetInfoRes{UserId: userId}, nil
}
