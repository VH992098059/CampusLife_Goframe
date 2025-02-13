package joinactivity

import (
	"context"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
)

type sJoinActivity struct{}

func (s sJoinActivity) InsertJoin(ctx context.Context, in model.JoinActivityModelInput) (out model.JoinActivityModelOutput, err error) {
	_, err = dao.ActivityJoin.Ctx(ctx).Data(in).Insert()

	if err != nil {
		return out, err
	}
	return model.JoinActivityModelOutput{}, nil
}

func init() {
	service.RegisterJoinActivity(New())
}
func New() *sJoinActivity { return &sJoinActivity{} }
