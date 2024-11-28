package rotation

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/model/entity"
	"demo3/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRotitionInfo struct {
}

func init() {
	service.RegisterRotitionInfo(New())
}
func New() *sRotitionInfo {
	return &sRotitionInfo{}
}
func (s sRotitionInfo) CreateRotation(ctx context.Context, in model.RotationModelCreateInput) (out model.RotationModelCreateOutput, err error) {
	GetId, err := dao.RotationInfo.Ctx(ctx).Data(entity.RotationInfo{
		RotationId:   gconv.Int(gdb.Raw("rotation_id+1000000")),
		RotationName: in.RotationName,
		PicUrl:       in.PicUrl,
	}).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RotationModelCreateOutput{RotationId: int(GetId)}, nil
}
func (s sRotitionInfo) UpdateRotation(ctx context.Context, in model.RotationModelUpdateInput) (out model.RotationModelUpdateOutput, err error) {
	//过滤某个字段修改数据（FieldEX）
	_, err = dao.RotationInfo.Ctx(ctx).Data(in).FieldsEx(dao.RotationInfo.Columns().RotationId).Where(dao.RotationInfo.Columns().RotationId, in.RotationId).Update()
	if err != nil {
		return out, err
	}
	return model.RotationModelUpdateOutput{}, nil
}

func (s sRotitionInfo) GetListRotation(ctx context.Context, in model.RotationModelListInput) (out *model.RotationModelListOutput, err error) {
	out = &model.RotationModelListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	list, err := dao.RotationInfo.Ctx(ctx).Unscoped().All()
	if err != nil {
		return out, err
	}
	return &model.RotationModelListOutput{List: list}, err
}

func (s sRotitionInfo) DeleteRotation(ctx context.Context, in model.RotationModelDeleteInput) (err error) {
	_, err = dao.RotationInfo.Ctx(ctx).Where("rotation_id", in.RotationId).Delete()
	if err != nil {
		return err
	}
	return nil
}
