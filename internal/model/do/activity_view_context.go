// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityViewContext is the golang structure of table activity_view_context for DAO operations like Where/Data.
type ActivityViewContext struct {
	g.Meta                `orm:"table:activity_view_context, do:true"`
	Uuid                  interface{} // id
	ActivityTitle         interface{} // 活动标题
	ActivityAbout         interface{} // 活动介绍
	Popular               interface{} // 活动热门度
	RegistrationStartTime *gtime.Time // 报名开始时间
	RegistrationEndTime   *gtime.Time // 报名结束时间
	ActivityStartTime     *gtime.Time // 活动开始时间
	ActivityEndTime       *gtime.Time // 活动结束时间
	PersonLimit           interface{} // 人数限制
	PersonCurrent         interface{} // 当前报名人数
}
