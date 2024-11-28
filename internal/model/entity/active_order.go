// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActiveOrder is the golang structure for table active_order.
type ActiveOrder struct {
	Id           int         `json:"id"           orm:"id"             description:"id"`
	Status       int         `json:"status"       orm:"status"         description:"状态（1：待审核 2：待参加 3：已检票 4：已取消）"`
	ActivityId   int         `json:"activityId"   orm:"activity_id"    description:"活动id"`
	CommonUserId int         `json:"commonUserId" orm:"common_user_id" description:"普通用户id"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"      description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"      description:"修改时间"`
	DeleteAt     *gtime.Time `json:"deleteAt"     orm:"delete_at"      description:"软删除"`
}
