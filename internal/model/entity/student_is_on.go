// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StudentIsOn is the golang structure for table student_is_on.
type StudentIsOn struct {
	StudentId string      `json:"studentId" orm:"student_id" description:"学号"`
	UserId    string      `json:"userId"    orm:"user_id"    description:"用户编号"`
	IsOn      string      `json:"isOn"      orm:"is_on"      description:"是否在校"`
	SchoolId  int         `json:"schoolId"  orm:"school_id"  description:"学校编号"`
	CreateAt  *gtime.Time `json:"createAt"  orm:"create_at"  description:"创建时间"`
	UpdateAt  *gtime.Time `json:"updateAt"  orm:"update_at"  description:"修改时间"`
	DeleteAt  *gtime.Time `json:"deleteAt"  orm:"delete_at"  description:"软删除"`
}
