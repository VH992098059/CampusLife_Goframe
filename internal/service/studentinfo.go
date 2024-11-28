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
	IStudentInfo interface {
		Create(ctx context.Context, in model.StudentInfoCreateInput) (out model.StudentInfoCreateOutput, err error)
		Update(ctx context.Context, in model.StudentInfoUpdateInput) (out model.StudentInfoUpdateOutput, err error)
		Delete(ctx context.Context, in model.StudentInfoDeleteInput) (out model.StudentInfoDeleteOutput, err error)
		GetList(ctx context.Context, in model.StudentInfoListInput) (out *model.StudentInfoListOutput, err error)
	}
)

var (
	localStudentInfo IStudentInfo
)

func StudentInfo() IStudentInfo {
	if localStudentInfo == nil {
		panic("implement not found for interface IStudentInfo, forgot register?")
	}
	return localStudentInfo
}

func RegisterStudentInfo(i IStudentInfo) {
	localStudentInfo = i
}
