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
	IDormitoryInfo interface {
		GetList(ctx context.Context, in model.DormitoryInfoModelListInput) (out *model.DormitoryInfoModelListOutput, err error)
		CreateDormitoryInfo(ctx context.Context, in model.DormitoryInfoModelCreateInput) (out model.DormitoryInfoModelCreateOutput, err error)
		UpdateDormitoryInfo(ctx context.Context, in model.DormitoryInfoModelUpdateInput) (out model.DormitoryInfoModelUpdateOutput, err error)
		DeleteDormitoryInfo(ctx context.Context, id int) error
	}
)

var (
	localDormitoryInfo IDormitoryInfo
)

func DormitoryInfo() IDormitoryInfo {
	if localDormitoryInfo == nil {
		panic("implement not found for interface IDormitoryInfo, forgot register?")
	}
	return localDormitoryInfo
}

func RegisterDormitoryInfo(i IDormitoryInfo) {
	localDormitoryInfo = i
}
