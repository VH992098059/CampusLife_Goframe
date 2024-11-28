// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StudentInfo is the golang structure of table student_info for DAO operations like Where/Data.
type StudentInfo struct {
	g.Meta     `orm:"table:student_info, do:true"`
	StudentId  interface{} // 学号
	Name       interface{} // 学生姓名
	Age        interface{} // 年龄
	Class      interface{} // 学生班级
	Department interface{} // 所在系
	Phone      interface{} // 学生电话
	CreateAt   *gtime.Time // 创建时间
	UpdateAt   *gtime.Time // 修改时间
	DeleteAt   *gtime.Time // 软删除
}
