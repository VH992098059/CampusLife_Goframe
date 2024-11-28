// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OrderInfo is the golang structure for table order_info.
type OrderInfo struct {
	OrderId     string      `json:"orderId"     orm:"order_id"     description:""`
	DormitoryId string      `json:"dormitoryId" orm:"dormitory_id" description:""`
	UserId      string      `json:"userId"      orm:"user_id"      description:""`
	Type        string      `json:"type"        orm:"type"         description:"100为照明，200为空调"`
	IsPay       int         `json:"isPay"       orm:"is_pay"       description:"0为未支付，1为已支付"`
	Price       float64     `json:"price"       orm:"price"        description:""`
	CreateAt    *gtime.Time `json:"createAt"    orm:"create_at"    description:""`
	UpdateAt    *gtime.Time `json:"updateAt"    orm:"update_at"    description:""`
	DeleteAt    *gtime.Time `json:"deleteAt"    orm:"delete_at"    description:""`
}
