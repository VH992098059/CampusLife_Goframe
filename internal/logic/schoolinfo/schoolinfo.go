package schoolinfo

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"strconv"
)

type sSchoolInfo struct{}

func init() {
	service.RegisterSchoolInfo(New())
}

func New() *sSchoolInfo {
	return &sSchoolInfo{}
}

// Create 创建学校逻辑代码
func (s *sSchoolInfo) Create(ctx context.Context, in model.SchoolInfoCreateInput) (out model.SchoolInfoCreateOutput, err error) {
	SchoolCreateId, err := dao.SchoolInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.SchoolInfoCreateOutput{SchoolId: strconv.FormatInt(SchoolCreateId, 10)}, err
}

// Update 修改学校逻辑代码
func (s *sSchoolInfo) Update(ctx context.Context, in model.SchoolInfoUpdateInput) (out model.SchoolInfoUpdateOutput, err error) {
	_, err = dao.SchoolInfo.Ctx(ctx).Data(in).Update()
	if err != nil {
		return out, err
	}
	return model.SchoolInfoUpdateOutput{}, err
}

// Delete 删除学校逻辑代码
func (s *sSchoolInfo) Delete(ctx context.Context, id string) (out model.SchoolInfoDeleteOutput, err error) {
	_, err = dao.SchoolInfo.Ctx(ctx).Where(dao.SchoolInfo.Columns().SchoolId, id).Delete()
	if err != nil {
		return out, err
	}
	return model.SchoolInfoDeleteOutput{}, err
}

// GetList 添加学校逻辑代码
func (s *sSchoolInfo) GetList(ctx context.Context, in model.SchoolInfoGetInput) (out *model.SchoolInfoGetOutput, err error) {
	out = &model.SchoolInfoGetOutput{
		Page: in.Page,
		Size: in.Size,
	}
	list, err := dao.SchoolInfo.Ctx(ctx).Unscoped().All()
	if err != nil {
		return
	}
	return &model.SchoolInfoGetOutput{List: list}, err
}
