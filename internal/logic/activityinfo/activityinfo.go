package activityinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"fmt"
)

type sActivityInfo struct{}

func (s sActivityInfo) ActivitySearch(ctx context.Context, in model.ActivitySearchModelInput) (out model.ActivitySearchModelOutput, err error) {
	Page, Previous, Next := 1, 0, in.Size
	if in.Page > Page {
		Previous += in.Size
		Next += in.Size
		Page = in.Page
	} else if in.Page < Page {
		Previous -= in.Size
		Next -= in.Size
		Page = in.Page
	} else if in.Page <= 1 {
		Previous = 0
		Next = in.Size
		Page = 1
	}
	list, count, err := dao.ActivityViewContext.Ctx(ctx).WhereLike("activity_title", "%"+in.ActivityTitle+"%").Limit(Previous, Next).AllAndCount(true)
	if err != nil {
		return out, err
	}
	fmt.Println()
	return model.ActivitySearchModelOutput{List: list, Total: count}, nil
}

func (s sActivityInfo) GetActivity(ctx context.Context, in model.GetActivityModelInput) (out model.GetActivityModelOutput, err error) {

	list, err := dao.ActivityViewContext.Ctx(ctx).Where("uuid", in.Uuid).All()
	if err != nil {
		return out, err
	}
	return model.GetActivityModelOutput{List: list}, nil
}

func (s sActivityInfo) List(ctx context.Context, in model.ActivityListModelInput) (out model.ActivityListModelOutput, err error) {
	//TODO implement me
	panic("implement me")
}

func init() {
	service.RegisterActivityInfo(New())
}
func New() *sActivityInfo {
	return &sActivityInfo{}
}
