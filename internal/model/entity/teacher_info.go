// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TeacherInfo is the golang structure for table teacher_info.
type TeacherInfo struct {
	TeacherId  int         `json:"teacherId"  orm:"teacher_id" description:"职工编号"`
	Name       string      `json:"name"       orm:"name"       description:"教师名称"`
	Age        int         `json:"age"        orm:"age"        description:"年龄"`
	Department string      `json:"department" orm:"department" description:"所在系"`
	Phone      string      `json:"phone"      orm:"phone"      description:"教师电话"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"  description:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"  description:"修改时间"`
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"  description:"软删除"`
}
