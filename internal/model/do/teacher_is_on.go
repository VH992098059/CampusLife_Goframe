// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeacherIsOn is the golang structure of table teacher_is_on for DAO operations like Where/Data.
type TeacherIsOn struct {
	g.Meta    `orm:"table:teacher_is_on, do:true"`
	TeacherId interface{} // 职工编号
	UserId    interface{} // 用户编号
	IsOn      interface{} // 是否在职
	SchoolId  interface{} // 学校编号
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 软删除
}
