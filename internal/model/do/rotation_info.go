// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure of table rotation_info for DAO operations like Where/Data.
type RotationInfo struct {
	g.Meta       `orm:"table:rotation_info, do:true"`
	RotationId   interface{} // 轮播图编号
	PicUrl       interface{} // 轮播图路径
	RotationName interface{} // 轮播图图片名
	CreateAt     *gtime.Time // 轮播图创建时间
	UpdateAt     *gtime.Time // 轮播图修改时间
	DeleteAt     *gtime.Time // 轮播图软删除
}
