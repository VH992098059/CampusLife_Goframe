// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StudentInfo is the golang structure for table student_info.
type StudentInfo struct {
	StudentId  string      `json:"studentId"  orm:"student_id" description:"学号"`
	Name       string      `json:"name"       orm:"name"       description:"学生姓名"`
	Age        int         `json:"age"        orm:"age"        description:"年龄"`
	Class      string      `json:"class"      orm:"class"      description:"学生班级"`
	Department string      `json:"department" orm:"department" description:"所在系"`
	Phone      string      `json:"phone"      orm:"phone"      description:"学生电话"`
	CreateAt   *gtime.Time `json:"createAt"   orm:"create_at"  description:"创建时间"`
	UpdateAt   *gtime.Time `json:"updateAt"   orm:"update_at"  description:"修改时间"`
	DeleteAt   *gtime.Time `json:"deleteAt"   orm:"delete_at"  description:"软删除"`
}
