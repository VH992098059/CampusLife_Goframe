// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminInfo is the golang structure for table admin_info.
type AdminInfo struct {
	AdminId  int         `json:"adminId"  orm:"admin_id"  description:"管理员ID"`
	Admin    string      `json:"admin"    orm:"admin"     description:"管理员账号"`
	Password string      `json:"password" orm:"password"  description:"管理员密码"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"修改时间"`
}
