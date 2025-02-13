// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityComment is the golang structure for table activity_comment.
type ActivityComment struct {
	ActivityUuid string      `json:"activityUuid" orm:"activity_uuid" description:"活动ID"`
	UserUuid     string      `json:"userUuid"     orm:"user_uuid"     description:"用户ID"`
	Comment      string      `json:"comment"      orm:"comment"       description:"用户评论"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"发表时间"`
	DeleteAt     *gtime.Time `json:"deleteAt"     orm:"delete_at"     description:"删除时间"`
}
