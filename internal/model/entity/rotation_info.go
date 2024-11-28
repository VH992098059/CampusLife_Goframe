// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RotationInfo is the golang structure for table rotation_info.
type RotationInfo struct {
	RotationId   int         `json:"rotationId"   orm:"rotation_id"   description:"轮播图编号"`
	PicUrl       string      `json:"picUrl"       orm:"pic_url"       description:"轮播图路径"`
	RotationName string      `json:"rotationName" orm:"rotation_name" description:"轮播图图片名"`
	CreateAt     *gtime.Time `json:"createAt"     orm:"create_at"     description:"轮播图创建时间"`
	UpdateAt     *gtime.Time `json:"updateAt"     orm:"update_at"     description:"轮播图修改时间"`
	DeleteAt     *gtime.Time `json:"deleteAt"     orm:"delete_at"     description:"轮播图软删除"`
}
