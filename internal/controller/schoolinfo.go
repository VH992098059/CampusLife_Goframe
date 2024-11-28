package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

var SchoolInfo = cSchoolInfo{}

type cSchoolInfo struct{}

// Create 创建学校控制器
func (c *cSchoolInfo) Create(ctx context.Context, req *backapi.SchoolInfoCreateReq) (out *backapi.SchoolInfoCreateRes, err error) {
	list, err := service.SchoolInfo().Create(ctx, model.SchoolInfoCreateInput{
		SchoolId:   req.SchoolId,
		SchoolName: req.SchoolName,
		Province:   req.Province,
		City:       req.City,
		County:     req.County,
		Address:    req.Address,
	})
	if err != nil {
		return out, err
	}
	return &backapi.SchoolInfoCreateRes{UserId: list.SchoolId}, err
}

// Update 修改学校控制器
func (c *cSchoolInfo) Update(ctx context.Context, req *backapi.SchoolInfoUpdateReq) (out *backapi.SchoolInfoUpdateRes, err error) {
	_, err = service.SchoolInfo().Update(ctx, model.SchoolInfoUpdateInput{
		SchoolId:   req.SchoolId,
		SchoolName: req.SchoolName,
		Province:   req.Province,
		City:       req.City,
		County:     req.County,
		Address:    req.Address,
	})
	if err != nil {
		return out, err
	}
	return &backapi.SchoolInfoUpdateRes{}, err
}

// Delete 删除学校控制器
func (c *cSchoolInfo) Delete(ctx context.Context, req *backapi.SchoolInfoDeleteReq) (out *backapi.SchoolInfoDeleteRes, err error) {
	_, err = service.SchoolInfo().Delete(ctx, req.SchoolId)
	if err != nil {
		return out, err
	}
	return &backapi.SchoolInfoDeleteRes{}, err
}

// GetList 查询学校控制器
func (c *cSchoolInfo) GetList(ctx context.Context, req *backapi.SchoolInfoGetListReq) (out *backapi.SchoolInfoGetListRes, err error) {
	list, count, err := dao.SchoolInfo.Ctx(ctx).Unscoped().AllAndCount(true)
	if err != nil {
		return out, err
	}
	return &backapi.SchoolInfoGetListRes{
		List:  list,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, nil

}
