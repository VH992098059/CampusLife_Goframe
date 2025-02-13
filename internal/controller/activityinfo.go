package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

var ActivityInfo = cActivityInfo{}

type cActivityInfo struct{}

func (c *cActivityInfo) GetList(ctx context.Context, req *backapi.ActivityInfoListReq) (res *backapi.ActivityInfoListRes, err error) {
	Page, Previous, Next := 1, 0, req.Size
	if req.Page > Page {
		Previous += req.Size
		Next += req.Size
		Page = req.Page
	} else if req.Page < Page {
		Previous -= req.Size
		Next -= req.Size
		Page = req.Page
	} else if req.Page <= 1 {
		Previous = 0
		Next = req.Size
		Page = 1
	}
	list, count, err := dao.ActivityViewContext.Ctx(ctx).Limit(Previous, Next).AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.ActivityInfoListRes{List: list, Total: count, Size: req.Size, Page: req.Page}, nil
}
func (c *cActivityInfo) GetActivity(ctx context.Context, req *backapi.ActivityInfoGetReq) (res *backapi.ActivityInfoGetRes, err error) {
	list, err := service.ActivityInfo().GetActivity(ctx, model.GetActivityModelInput{
		Uuid: req.Uuid,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.ActivityInfoGetRes{List: list.List}, nil
}
func (c *cActivityInfo) ActivitySearch(ctx context.Context, req *backapi.ActivitySearchReq) (res *backapi.ActivitySearchRes, err error) {
	list, err := service.ActivityInfo().ActivitySearch(ctx, model.ActivitySearchModelInput{
		ActivityTitle: req.ActivityTitle,
		Page:          req.Page,
		Size:          req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.ActivitySearchRes{List: list.List, Total: list.Total, Size: req.Size, Page: req.Page}, nil
}
