// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StudentIsOn is the golang structure of table student_is_on for DAO operations like Where/Data.
type StudentIsOn struct {
	g.Meta    `orm:"table:student_is_on, do:true"`
	StudentId interface{} // 学号
	UserId    interface{} // 用户编号
	IsOn      interface{} // 是否在校
	SchoolId  interface{} // 学校编号
	CreateAt  *gtime.Time // 创建时间
	UpdateAt  *gtime.Time // 修改时间
	DeleteAt  *gtime.Time // 软删除
}
