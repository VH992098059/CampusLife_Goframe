package backapi

import "github.com/gogf/gf/v2/frame/g"

/*参加活动*/
type JoinActivityReq struct {
	g.Meta       `path:"/join" method:"post"`
	UserId       string `json:"user_id"`
	ActivityUuid string `json:"activity_uuid"`
}
type JoinActivityRes struct {
	Msg string `json:"msg"`
}
