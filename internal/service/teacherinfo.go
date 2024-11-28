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
	ITeacherInfo interface {
		Create(ctx context.Context, in model.TeacherInfoModelAddInput) (out model.TeacherInfoModelAddOutput, err error)
	}
)

var (
	localTeacherInfo ITeacherInfo
)

func TeacherInfo() ITeacherInfo {
	if localTeacherInfo == nil {
		panic("implement not found for interface ITeacherInfo, forgot register?")
	}
	return localTeacherInfo
}

func RegisterTeacherInfo(i ITeacherInfo) {
	localTeacherInfo = i
}
