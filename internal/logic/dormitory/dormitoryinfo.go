package dormitory

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sDormitoryInfo struct{}

func init() {
	service.RegisterDormitoryInfo(New())
}
func New() *sDormitoryInfo {
	return &sDormitoryInfo{}
}
func (s sDormitoryInfo) GetList(ctx context.Context, in model.DormitoryInfoModelListInput) (out *model.DormitoryInfoModelListOutput, err error) {
	out = &model.DormitoryInfoModelListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	GetList, err := dao.DormitoryInfo.Ctx(ctx).Unscoped().All()
	if err != nil {
		return nil, err
	}
	return &model.DormitoryInfoModelListOutput{List: GetList}, nil
}

func (s sDormitoryInfo) CreateDormitoryInfo(ctx context.Context, in model.DormitoryInfoModelCreateInput) (out model.DormitoryInfoModelCreateOutput, err error) {
	GetId, err := dao.DormitoryInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.DormitoryInfoModelCreateOutput{DormitoryId: int(GetId)}, nil
}

func (s sDormitoryInfo) UpdateDormitoryInfo(ctx context.Context, in model.DormitoryInfoModelUpdateInput) (out model.DormitoryInfoModelUpdateOutput, err error) {
	_, err = dao.DormitoryInfo.Ctx(ctx).Data(in).FieldsEx(dao.DormitoryInfo.Columns().DormitoryId).Where(dao.DormitoryInfo.Columns().DormitoryId, in.DormitoryId).Update()
	if err != nil {
		return out, err
	}
	return model.DormitoryInfoModelUpdateOutput{}, nil
}

func (s sDormitoryInfo) DeleteDormitoryInfo(ctx context.Context, id int) error {
	_, err := dao.DormitoryInfo.Ctx(ctx).Where(dao.DormitoryInfo.Columns().DormitoryId, id).Delete()
	if err != nil {
		return err
	}
	return nil
}
