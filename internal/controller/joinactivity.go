package controller

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"strings"
)

var JoinActivity = cJoinActivity{}

type cJoinActivity struct {
}

func (c *cJoinActivity) InsertJoin(ctx context.Context, req *backapi.JoinActivityReq) (res *backapi.JoinActivityRes, err error) {
	join := dao.ActivityJoin
	result, err := join.Ctx(ctx).Where("activity_uuid=", req.ActivityUuid).Where("user_id=", req.UserId).One()
	if result.IsEmpty() == true {
		_, err = service.JoinActivity().InsertJoin(ctx, model.JoinActivityModelInput{
			JoinUuid:     strings.ReplaceAll(uuid.New().String()[:20], "-", ""), // 移除UUID横线
			UserId:       req.UserId,
			ActivityUuid: req.ActivityUuid,
		})
		if err != nil {
			return nil, err
		}
		_, err = dao.ActivityInfo.Ctx(ctx).Data(g.Map{
			"person_current": gdb.Raw("person_current+1"),
		}).Where("uuid", req.ActivityUuid).Update()
		if err != nil {
			return nil, err
		}

		return &backapi.JoinActivityRes{}, nil
	} else {
		return &backapi.JoinActivityRes{Msg: "已参加，无需重新参加"}, nil
	}

}
