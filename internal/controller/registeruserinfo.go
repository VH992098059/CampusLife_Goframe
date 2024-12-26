package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/model"
	"demo3/internal/service"
	"github.com/google/uuid"
)

var RegisterUserInfo = cRegisterInfo{}

type cRegisterInfo struct{}

func (c *cRegisterInfo) RegisterUser(ctx context.Context, req *backapi.RegisterUserReq) (res *backapi.RegisterUserRes, err error) {
	create, err := service.RegisterInfo().RegisterUser(ctx, model.RegisterUserModelInput{
		UserId:   uuid.New().String(),
		UserName: req.UserName,
		Password: req.Password,
		NickName: req.NickName,
		Phone:    req.Phone,
		Type:     req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.RegisterUserRes{UserId: create.UserId}, nil
}
