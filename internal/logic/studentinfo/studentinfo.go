package studentinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sStudentInfo struct{}

func init() {
	service.RegisterStudentInfo(New())
}
func New() *sStudentInfo {
	return &sStudentInfo{}
}
func (s sStudentInfo) Create(ctx context.Context, in model.StudentInfoCreateInput) (out model.StudentInfoCreateOutput, err error) {
	GetId, err := dao.StudentIsOn.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.StudentInfoCreateOutput{StudentId: int(GetId)}, err
}

func (s sStudentInfo) Update(ctx context.Context, in model.StudentInfoUpdateInput) (out model.StudentInfoUpdateOutput, err error) {
	_, err = dao.StudentIsOn.Ctx(ctx).Data(in).FieldsEx(dao.StudentIsOn.Columns().StudentId).Where(dao.StudentIsOn.Columns().StudentId, in.StudentId).Update()
	if err != nil {
		return out, err
	}
	return model.StudentInfoUpdateOutput{}, err
}

func (s sStudentInfo) Delete(ctx context.Context, in model.StudentInfoDeleteInput) (out model.StudentInfoDeleteOutput, err error) {
	_, err = dao.StudentIsOn.Ctx(ctx).Where(dao.StudentIsOn.Columns().StudentId, in.StudentId).Delete()
	if err != nil {
		return out, err
	}
	return model.StudentInfoDeleteOutput{}, err
}

func (s sStudentInfo) GetList(ctx context.Context, in model.StudentInfoListInput) (out *model.StudentInfoListOutput, err error) {
	out = &model.StudentInfoListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	GetList, err := dao.StudentIsOn.Ctx(ctx).Unscoped().All()
	if err != nil {
		return out, err
	}
	return &model.StudentInfoListOutput{List: GetList}, err
}
