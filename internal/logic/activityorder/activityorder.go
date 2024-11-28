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
	_, err = dao.ActiveOrder.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ActivityOrderAddModelOutput{}, err
}
