package teacherinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sTeacherInfo struct {
}

func New() *sTeacherInfo {
	return &sTeacherInfo{}
}
func init() {
	service.RegisterTeacherInfo(New())
}
func (s sTeacherInfo) Create(ctx context.Context, in model.TeacherInfoModelAddInput) (out model.TeacherInfoModelAddOutput, err error) {
	_, err = dao.TeacherInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.TeacherInfoModelAddOutput{}, nil
}
