// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityJoin is the golang structure of table activity_join for DAO operations like Where/Data.
type ActivityJoin struct {
	g.Meta       `orm:"table:activity_join, do:true"`
	JoinUuid     interface{} // 参加活动id
	ActivityUuid interface{} // 活动id
	UserId       interface{} // 用户id
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 修改时间
	DeleteAt     *gtime.Time // 删除时间
}
