package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/model"
	"demo3/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var TeacherInfo = cTeacherInfo{}

type cTeacherInfo struct{}

func (i *cTeacherInfo) Create(ctx context.Context, req *backapi.TeacherInfoAddReq) (res *backapi.TeacherInfoAddRes, err error) {
	_, err = service.TeacherInfo().Create(ctx, model.TeacherInfoModelAddInput{
		TeacherID:  gconv.Int("13714000001"),
		Name:       req.Name,
		Phone:      req.Phone,
		Age:        req.Age,
		Department: req.Department,
	})
	if err != nil {
		return nil, err
	}
	return &backapi.TeacherInfoAddRes{}, nil
}
