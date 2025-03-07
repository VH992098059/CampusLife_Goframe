// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserState is the golang structure for table user_state.
type UserState struct {
	UserId   string      `json:"userId"   orm:"user_id"  description:"用户ID"`
	Nickname string      `json:"nickname" orm:"nickname" description:"用户昵称"`
	Username string      `json:"username" orm:"username" description:"用户账号"`
	Type     string      `json:"type"     orm:"type"     description:"用户类型（活动方，普通用户）"`
	Avater   string      `json:"avater"   orm:"avater"   description:"用户头像"`
	Phone    string      `json:"phone"    orm:"phone"    description:"用户电话"`
	Wechat   string      `json:"wechat"   orm:"wechat"   description:"用户微信"`
	Birthday *gtime.Time `json:"birthday" orm:"birthday" description:"用户生日"`
	Email    string      `json:"email"    orm:"email"    description:"邮箱"`
	Sex      string      `json:"sex"      orm:"sex"      description:"用户性别"`
}
