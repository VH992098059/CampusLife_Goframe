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
	IActivityInfo interface {
		Create(ctx context.Context, in model.ActivityInfoModelAddInput) (out model.ActivityInfoModelAddOutput, err error)
	}
)

var (
	localActivityInfo IActivityInfo
)

func ActivityInfo() IActivityInfo {
	if localActivityInfo == nil {
		panic("implement not found for interface IActivityInfo, forgot register?")
	}
	return localActivityInfo
}

func RegisterActivityInfo(i IActivityInfo) {
	localActivityInfo = i
}
