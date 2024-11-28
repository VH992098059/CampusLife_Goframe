package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

var RotationInfo = cRotationInfo{}

type cRotationInfo struct{}

func (c *cRotationInfo) Create(ctx context.Context, req *backapi.RotationCreateReq) (res *backapi.RotationCreateRes, err error) {
	GetId, err := service.RotitionInfo().CreateRotation(ctx, model.RotationModelCreateInput{
		RotationName: req.RotationName,
		PicUrl:       req.PicUrl,
	})
	if err != nil {
		return nil, err
	}

	return &backapi.RotationCreateRes{RotationId: GetId.RotationId}, nil
}
func (c *cRotationInfo) GetList(ctx context.Context, req *backapi.RotationListReq) (res *backapi.RotationListRes, err error) {
	//不过滤软删除查询数据并计算总量（Unscoped）
	List, count, err := dao.RotationInfo.Ctx(ctx).Unscoped().AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.RotationListRes{
		List:  List,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, err
}
func (c *cRotationInfo) Delete(ctx context.Context, req *backapi.RotationDeleteReq) (res *backapi.RotationDeleteRes, err error) {
	err = service.RotitionInfo().DeleteRotation(ctx, model.RotationModelDeleteInput{
		RotationId: req.RotationId,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.RotationDeleteRes{}, nil
}
func (c *cRotationInfo) Update(ctx context.Context, req *backapi.RotationUpdateReq) (res *backapi.RotationUpdateRes, err error) {
	_, err = service.RotitionInfo().UpdateRotation(ctx, model.RotationModelUpdateInput{
		RotationId:   req.RotationId,
		RotationName: req.RotationName,
		PicUrl:       req.PicUrl,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.RotationUpdateRes{}, nil
}
