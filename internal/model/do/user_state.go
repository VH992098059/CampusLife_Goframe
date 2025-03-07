// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserState is the golang structure of table user_state for DAO operations like Where/Data.
type UserState struct {
	g.Meta   `orm:"table:user_state, do:true"`
	UserId   interface{} // 用户ID
	Nickname interface{} // 用户昵称
	Username interface{} // 用户账号
	Type     interface{} // 用户类型（活动方，普通用户）
	Avater   interface{} // 用户头像
	Phone    interface{} // 用户电话
	Wechat   interface{} // 用户微信
	Birthday *gtime.Time // 用户生日
	Email    interface{} // 邮箱
	Sex      interface{} // 用户性别
}
