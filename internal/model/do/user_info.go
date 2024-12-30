// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta   `orm:"table:user_info, do:true"`
	UserId   interface{} // 用户ID
	Nickname interface{} // 用户昵称
	Username interface{} // 用户账号
	Password interface{} // 密码
	Type     interface{} // 用户类型（活动方，普通用户）
	Phone    interface{} // 用户电话
	Wechat   interface{} // 用户微信
	UserSalt interface{} // 加密盐
	CreateAt *gtime.Time // 创建时间
	UpdateAt *gtime.Time // 修改时间
	DeleteAt *gtime.Time // 软删除
}
