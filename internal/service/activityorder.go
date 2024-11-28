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
	IActivityOrder interface {
		Create(ctx context.Context, in model.ActivityOrderAddModelInput) (out model.ActivityOrderAddModelOutput, err error)
	}
)

var (
	localActivityOrder IActivityOrder
)

func ActivityOrder() IActivityOrder {
	if localActivityOrder == nil {
		panic("implement not found for interface IActivityOrder, forgot register?")
	}
	return localActivityOrder
}

func RegisterActivityOrder(i IActivityOrder) {
	localActivityOrder = i
}
