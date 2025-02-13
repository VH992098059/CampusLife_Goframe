// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure of table admin_info for DAO operations like Where/Data.
type AdminInfo struct {
	g.Meta   `orm:"table:admin_info, do:true"`
	AdminId  interface{} // 管理员ID
	Admin    interface{} // 管理员账号
	Password interface{} // 管理员密码
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 修改时间
}
