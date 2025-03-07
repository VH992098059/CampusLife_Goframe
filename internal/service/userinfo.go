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
		Person(ctx context.Context, in model.UserGetPersonMsgModelInput) (out model.UserGetPersonMsgModelOutput, err error)
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

func RegisterGetUser(i IUserInfo) {
	localUserInfo = i
}
