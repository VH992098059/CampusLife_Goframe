package backapi

import "github.com/gogf/gf/v2/frame/g"

type ActivityOrder struct {
	//ID           int `json:"id" v:"required#ID不能为空" description:"ID"`
	Status         int    `json:"status" v:"required#状态不能为空" description:"状态（1：待审核 2：待参加 3：已检票 4：已取消）"`
	ActivityNumber string `json:"activity_number" v:"required#活动ID不能为空" description:"活动ID"`
	CommonUserID   int    `json:"common_user_id" v:"required#普通用户ID不能为空" description:"普通用户ID"`
}
type ActivityOrderAddReq struct {
	g.Meta `path:"/activityOrderAdd" method:"post"`
	ActivityOrder
}
type ActivityOrderAddRes struct{}
type ActivityOrderUpdateReq struct {
	g.Meta `path:"/activityOrderUpdate" method:"put"`
	Status int `json:"status"`
}
type ActivityOrderUpdateRes struct{}
type ActivityOrderDeleteReq struct {
	g.Meta `path:"/activityOrderDelete" method:"delete"`
	Id     int `json:"id"`
}
type ActivityOrderDeleteRes struct{}
type ActivityOrderListReq struct {
	g.Meta `path:"/activityOrderList" method:"get"`
	//CommonPaginationReq
	Page   int `json:"page" description:"分页码" `
	Size   int `json:"size" description:"分页数量"`
	Offset int `json:"offset" dc:"初始容量"`
	Limit  int `json:"limit" dc:"最大容量"`
}
type ActivityOrderListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
