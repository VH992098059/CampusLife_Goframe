// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SchoolInfo is the golang structure of table school_info for DAO operations like Where/Data.
type SchoolInfo struct {
	g.Meta     `orm:"table:school_info, do:true"`
	SchoolId   interface{} // 学校编号
	SchoolName interface{} // 学校名称
	Province   interface{} // 学校省份
	City       interface{} // 学校城市
	County     interface{} // 学号市区
	Address    interface{} // 学校详细地址
}
