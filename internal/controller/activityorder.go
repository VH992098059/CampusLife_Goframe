package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"github.com/google/uuid"
)

var ActivityOrder = cActivityOrder{}

type cActivityOrder struct{}

func (i *cActivityOrder) Create(ctx context.Context, req *backapi.ActivityOrderAddReq) (res *backapi.ActivityOrderAddRes, err error) {
	_, err = service.ActivityOrder().Create(ctx, model.ActivityOrderAddModelInput{
		Uuid:           uuid.New().String(),
		Status:         req.Status,
		ActivityNumber: req.ActivityNumber,
		CommonUserID:   req.CommonUserID,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.ActivityOrderAddRes{}, nil
}
func (i *cActivityOrder) GetList(ctx context.Context, req *backapi.ActivityOrderListReq) (res *backapi.ActivityOrderListRes, err error) {
	list, count, err := dao.ActivityOrder.Ctx(ctx).Limit(req.Offset, req.Limit).AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.ActivityOrderListRes{
		List:  list,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, err
}
