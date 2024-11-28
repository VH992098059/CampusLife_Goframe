package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

var Dormitoryinfo = cDormitoryinfo{}

type cDormitoryinfo struct{}

func (c *cDormitoryinfo) GetList(ctx context.Context, req *backapi.DormitoryInfoListReq) (res *backapi.DormitoryInfoListRes, err error) {
	GetList, count, err := dao.DormitoryInfo.Ctx(ctx).Unscoped().AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.DormitoryInfoListRes{
		List:  GetList,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, nil

}
func (c *cDormitoryinfo) Update(ctx context.Context, req *backapi.DormitoryInfoUpdateReq) (res *backapi.DormitoryInfoUpdateRes, err error) {
	_, err = service.DormitoryInfo().UpdateDormitoryInfo(ctx, model.DormitoryInfoModelUpdateInput{
		DormitoryType:   req.DormitoryType,
		DormitoryNumber: req.DormitoryNumber,
		Floor:           req.Floor,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.DormitoryInfoUpdateRes{}, nil
}
func (c *cDormitoryinfo) Delete(ctx context.Context, req *backapi.DormitoryInfoDeleteReq) (res *backapi.DormitoryInfoDeleteRes, err error) {
	err = service.DormitoryInfo().DeleteDormitoryInfo(ctx, req.DormitoryId)
	if err != nil {
		return nil, err
	}
	return &backapi.DormitoryInfoDeleteRes{}, nil

}
func (c *cDormitoryinfo) Create(ctx context.Context, req *backapi.DormitoryInfoCreateReq) (res *backapi.DormitoryInfoCreateRes, err error) {
	GetId, err := service.DormitoryInfo().CreateDormitoryInfo(ctx, model.DormitoryInfoModelCreateInput{
		DormitoryType:   req.DormitoryType,
		DormitoryNumber: req.DormitoryNumber,
		Floor:           req.Floor,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.DormitoryInfoCreateRes{DormitoryId: GetId.DormitoryId}, nil
}
