// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TeacherInfo is the golang structure of table teacher_info for DAO operations like Where/Data.
type TeacherInfo struct {
	g.Meta     `orm:"table:teacher_info, do:true"`
	TeacherId  interface{} // 职工编号
	Name       interface{} // 教师名称
	Age        interface{} // 年龄
	Department interface{} // 所在系
	Phone      interface{} // 教师电话
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 修改时间
	DeleteAt   *gtime.Time // 软删除
}
