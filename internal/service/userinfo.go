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
	IUserInfo interface {
		GetList(ctx context.Context, in model.UserInfoModelInput) (out *model.UserInfoModelOutput, err error)
		CreateUserInfo(ctx context.Context, in model.UserInfoModelCreateInput) (out model.UserInfoModelCreateOutput, err error)
		UpdateUserInfo(ctx context.Context, in model.UserInfoModelUpdateInput) (out model.UserInfoModelUpdateOutput, err error)
		DeleteUserInfo(ctx context.Context, id string) error
	}
)

var (
	localUserInfo IUserInfo
)

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement not found for interface IUserInfo, forgot register?")
	}
	return localUserInfo
}

func RegisterUserInfo(i IUserInfo) {
	localUserInfo = i
}
