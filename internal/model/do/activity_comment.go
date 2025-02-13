// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityComment is the golang structure of table activity_comment for DAO operations like Where/Data.
type ActivityComment struct {
	g.Meta       `orm:"table:activity_comment, do:true"`
	ActivityUuid interface{} // 活动ID
	UserUuid     interface{} // 用户ID
	Comment      interface{} // 用户评论
	CreateAt     *gtime.Time // 发表时间
	DeleteAt     *gtime.Time // 删除时间
}
