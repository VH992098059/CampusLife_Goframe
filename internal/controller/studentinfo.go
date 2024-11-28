package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

var StudentInfo = cStudentInfo{}

type cStudentInfo struct{}

func (c *cStudentInfo) Create(ctx context.Context, req *backapi.StudentCreateReq) (res *backapi.StudentCreateRes, err error) {
	GetId, err := service.StudentInfo().Create(ctx, model.StudentInfoCreateInput{
		StudentId: req.StudentId,
		UserId:    req.UserId,
		Name:      req.Name,
		IsOn:      req.IsOn,
		SchoolId:  req.SchoolId,
	})
	return &backapi.StudentCreateRes{StudentId: GetId.StudentId}, err
}
func (c *cStudentInfo) Update(ctx context.Context, req *backapi.StudentUpdateReq) (res *backapi.StudentUpdateRes, err error) {
	_, err = service.StudentInfo().Update(ctx, model.StudentInfoUpdateInput{
		StudentId: req.StudentId,
		UserId:    req.UserId,
		Name:      req.Name,
		IsOn:      req.IsOn,
		SchoolId:  req.SchoolId,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.StudentUpdateRes{}, nil
}
func (c *cStudentInfo) Delete(ctx context.Context, req *backapi.StudentDeleteReq) (res *backapi.StudentDeleteRes, err error) {
	_, err = service.StudentInfo().Delete(ctx, model.StudentInfoDeleteInput{StudentId: req.StudentId})
	if err != nil {
		return nil, err
	}
	return &backapi.StudentDeleteRes{}, nil
}
func (c *cStudentInfo) GetList(ctx context.Context, req *backapi.StudentListReq) (res *backapi.StudentListRes, err error) {
	GetList, count, err := dao.StudentIsOn.Ctx(ctx).Unscoped().Limit(0, req.Size).AllAndCount(true)
	if err != nil {
		return nil, err
	}
	return &backapi.StudentListRes{
		List:  GetList,
		Page:  req.Page,
		Size:  req.Size,
		Total: count,
	}, nil
}
