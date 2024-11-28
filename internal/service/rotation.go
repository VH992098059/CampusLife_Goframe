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
	IRotitionInfo interface {
		CreateRotation(ctx context.Context, in model.RotationModelCreateInput) (out model.RotationModelCreateOutput, err error)
		UpdateRotation(ctx context.Context, in model.RotationModelUpdateInput) (out model.RotationModelUpdateOutput, err error)
		GetListRotation(ctx context.Context, in model.RotationModelListInput) (out *model.RotationModelListOutput, err error)
		DeleteRotation(ctx context.Context, in model.RotationModelDeleteInput) (err error)
	}
)

var (
	localRotitionInfo IRotitionInfo
)

func RotitionInfo() IRotitionInfo {
	if localRotitionInfo == nil {
		panic("implement not found for interface IRotitionInfo, forgot register?")
	}
	return localRotitionInfo
}

func RegisterRotitionInfo(i IRotitionInfo) {
	localRotitionInfo = i
}
