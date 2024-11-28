package activityinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sActivityInfo struct{}

func init() {
	service.RegisterActivityInfo(New())
}
func New() *sActivityInfo {
	return &sActivityInfo{}
}
func (s sActivityInfo) Create(ctx context.Context, in model.ActivityInfoModelAddInput) (out model.ActivityInfoModelAddOutput, err error) {
	_, err = dao.ActivityInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.ActivityInfoModelAddOutput{}, nil
}
