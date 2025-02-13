// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityViewContext is the golang structure for table activity_view_context.
type ActivityViewContext struct {
	Uuid                  string      `json:"uuid"                  orm:"uuid"                    description:"id"`
	ActivityTitle         string      `json:"activityTitle"         orm:"activity_title"          description:"活动标题"`
	ActivityAbout         string      `json:"activityAbout"         orm:"activity_about"          description:"活动介绍"`
	Popular               int64       `json:"popular"               orm:"popular"                 description:"活动热门度"`
	RegistrationStartTime *gtime.Time `json:"registrationStartTime" orm:"registration_start_time" description:"报名开始时间"`
	RegistrationEndTime   *gtime.Time `json:"registrationEndTime"   orm:"registration_end_time"   description:"报名结束时间"`
	ActivityStartTime     *gtime.Time `json:"activityStartTime"     orm:"activity_start_time"     description:"活动开始时间"`
	ActivityEndTime       *gtime.Time `json:"activityEndTime"       orm:"activity_end_time"       description:"活动结束时间"`
	PersonLimit           int         `json:"personLimit"           orm:"person_limit"            description:"人数限制"`
	PersonCurrent         int         `json:"personCurrent"         orm:"person_current"          description:"当前报名人数"`
}
