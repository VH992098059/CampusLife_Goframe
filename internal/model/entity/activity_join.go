// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityJoin is the golang structure for table activity_join.
type ActivityJoin struct {
	JoinUuid     string      `json:"joinUuid"     orm:"join_uuid"     description:"参加活动id"`
	ActivityUuid string      `json:"activityUuid" orm:"activity_uuid" description:"活动id"`
	UserId       string      `json:"userId"       orm:"user_id"       description:"用户id"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"修改时间"`
	DeleteAt     *gtime.Time `json:"deleteAt"     orm:"delete_at"     description:"删除时间"`
}
