// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure of table order_info for DAO operations like Where/Data.
type OrderInfo struct {
	g.Meta      `orm:"table:order_info, do:true"`
	OrderId     interface{} //
	DormitoryId interface{} //
	UserId      interface{} //
	Type        interface{} // 100为照明，200为空调
	IsPay       interface{} // 0为未支付，1为已支付
	Price       interface{} //
	CreateAt    *gtime.Time //
	UpdateAt    *gtime.Time //
	DeleteAt    *gtime.Time //
}
