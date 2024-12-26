package activityorder

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sActivityOrder struct{}

func init() {
	service.RegisterActivityOrder(New())
}
func New() *sActivityOrder {
	return &sActivityOrder{}
}
func (s sActivityOrder) Create(ctx context.Context, in model.ActivityOrderAddModelInput) (out model.ActivityOrderAddModelOutput, err error) {
	_, err = dao.ActivityOrder.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ActivityOrderAddModelOutput{}, err
}
func (s sActivityOrder) Delete(ctx context.Context, in model.ActivityOrderDeleteModelInput) (out model.ActivityOrderDeleteModelOutput, err error) {
	_, err = dao.ActivityOrder.Ctx(ctx).Where(dao.ActivityOrder.Columns().Uuid, in.ActivityUUID).Delete()
	if err != nil {
		return out, err
	}
	return model.ActivityOrderDeleteModelOutput{}, err
}

func (s sActivityOrder) GetList(ctx context.Context, in model.ActivityOrderListModelInput) (out model.ActivityOrderListModelOutput, err error) {
	_, err = dao.ActivityOrder.Ctx(ctx).Limit(in.Offset, in.Limit).Unscoped().All()
	if err != nil {
		return out, err
	}
	return model.ActivityOrderListModelOutput{}, err
}

func (s sActivityOrder) Update(ctx context.Context, in model.ActivityOrderUpdateModelInput) (out model.ActivityOrderUpdateModelOutput, err error) {
	_, err = dao.ActivityOrder.Ctx(ctx).Where(dao.ActivityOrder.Columns().Uuid, in.ActivityUUID).Update()
	if err != nil {
		return out, err
	}
	return model.ActivityOrderUpdateModelOutput{}, err
}
