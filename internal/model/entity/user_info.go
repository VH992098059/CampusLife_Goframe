// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	UserId   string      `json:"userId"   orm:"user_id"   description:"用户ID"`
	Nickname string      `json:"nickname" orm:"nickname"  description:"用户昵称"`
	Username string      `json:"username" orm:"username"  description:"用户账号"`
	Password string      `json:"password" orm:"password"  description:"密码"`
	Phone    string      `json:"phone"    orm:"phone"     description:"用户电话"`
	Wechat   string      `json:"wechat"   orm:"wechat"    description:"用户微信"`
	UserSalt string      `json:"userSalt" orm:"user_salt" description:"加密盐"`
	CreateAt *gtime.Time `json:"createAt" orm:"create_at" description:"创建时间"`
	UpdateAt *gtime.Time `json:"updateAt" orm:"update_at" description:"修改时间"`
	DeleteAt *gtime.Time `json:"deleteAt" orm:"delete_at" description:"软删除"`
}
