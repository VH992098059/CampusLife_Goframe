package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/model"
	"demo3/internal/service"
	"demo3/utility/aliyunCode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"strings"
)

var RegisterUserInfo = cRegisterInfo{}

type cRegisterInfo struct{}

// RegisterUser /*注册用户*/
func (c *cRegisterInfo) RegisterUser(ctx context.Context, req *backapi.RegisterUserReq) (res *backapi.RegisterUserRes, err error) {
	IsCodeTrue := aliyunCode.VerifyPhoneCode(req.Phone, req.Code) /*验证手机号*/
	if IsCodeTrue != true {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code": 401,
			"msg":  "验证码不正确",
		})
		return
	}
	create, err := service.RegisterInfo().RegisterUser(ctx, model.RegisterUserModelInput{
		UserId:   strings.ReplaceAll(uuid.New().String()[:20], "-", ""), // 移除UUID横线
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
