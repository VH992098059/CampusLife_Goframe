package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/model"
	"demo3/internal/service"
	"demo3/utility/aliyunCode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

var RegisterUserInfo = cRegisterInfo{}

type cRegisterInfo struct{}

func (c *cRegisterInfo) RegisterUser(ctx context.Context, req *backapi.RegisterUserReq) (res *backapi.RegisterUserRes, err error) {
	IsCodeTrue := aliyunCode.VerifyPhoneCode(req.Phone, req.Code)
	if IsCodeTrue != true {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code": 200,
			"msg":  "验证码不正确",
		})
		return
	}
	create, err := service.RegisterInfo().RegisterUser(ctx, model.RegisterUserModelInput{
		UserId:   uuid.New().String(),
		UserName: req.UserName,
		Password: req.Password,
		NickName: req.NickName,
		Phone:    req.Phone,
		Type:     req.Type,
	})

	if err != nil {
		aliyunCode.DeletePhoneCode(req.Phone)
		return nil, err
	}
	aliyunCode.DeletePhoneCode(req.Phone)
	return &backapi.RegisterUserRes{UserId: create.UserId}, nil
}
