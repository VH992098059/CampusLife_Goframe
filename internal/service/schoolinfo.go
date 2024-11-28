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
	ISchoolInfo interface {
		// Create 创建学校逻辑代码
		Create(ctx context.Context, in model.SchoolInfoCreateInput) (out model.SchoolInfoCreateOutput, err error)
		// Update 修改学校逻辑代码
		Update(ctx context.Context, in model.SchoolInfoUpdateInput) (out model.SchoolInfoUpdateOutput, err error)
		// Delete 删除学校逻辑代码
		Delete(ctx context.Context, id string) (out model.SchoolInfoDeleteOutput, err error)
		// GetList 添加学校逻辑代码
		GetList(ctx context.Context, in model.SchoolInfoGetInput) (out *model.SchoolInfoGetOutput, err error)
	}
)

var (
	localSchoolInfo ISchoolInfo
)

func SchoolInfo() ISchoolInfo {
	if localSchoolInfo == nil {
		panic("implement not found for interface ISchoolInfo, forgot register?")
	}
	return localSchoolInfo
}

func RegisterSchoolInfo(i ISchoolInfo) {
	localSchoolInfo = i
}
