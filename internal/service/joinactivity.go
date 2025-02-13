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
	IJoinActivity interface {
		InsertJoin(ctx context.Context, in model.JoinActivityModelInput) (out model.JoinActivityModelOutput, err error)
	}
)

var (
	localJoinActivity IJoinActivity
)

func JoinActivity() IJoinActivity {
	if localJoinActivity == nil {
		panic("implement not found for interface IJoinActivity, forgot register?")
	}
	return localJoinActivity
}

func RegisterJoinActivity(i IJoinActivity) {
	localJoinActivity = i
}
