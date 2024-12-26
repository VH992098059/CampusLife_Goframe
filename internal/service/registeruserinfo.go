// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo3/internal/model"
)

type (
	IRegisterUserInfo interface {
		RegisterUser(ctx context.Context, in model.RegisterUserModelInput) (out model.RegisterUserModelOutput, err error)
	}
)

var (
	localRegisterUserInfo IRegisterUserInfo
)

func RegisterInfo() IRegisterUserInfo {
	if localRegisterUserInfo == nil {
		panic("implement not found for interface IRegisterUserInfo, forgot register?")
	}
	return localRegisterUserInfo
}

func RegisterUserInfo(i IRegisterUserInfo) {
	localRegisterUserInfo = i
}
