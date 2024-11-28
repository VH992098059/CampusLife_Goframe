package backapi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationInfo struct {
	PicUrl       string `json:"pic_url"   description:"轮播图路径" v:"required#图片路径不能为空"`
	RotationName string `json:"rotation_name" description:"轮播图图片名"  v:"required#活动名称不能为空"`
}
type RotationCreateReq struct {
	g.Meta `path:"/createRotation" method:"post"`
	RotationInfo
}
type RotationCreateRes struct {
	RotationId int `json:"rotation_id"`
}
type RotationDeleteReq struct {
	g.Meta     `path:"/deleteRotation" method:"delete"`
	RotationId int `json:"rotation_id" description:"轮播图ID" v:"required#ID不能为空"`
}
type RotationDeleteRes struct{}
type RotationListReq struct {
	g.Meta `path:"/listRotation" method:"get"`
	CommonPaginationReq
}
type RotationListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type RotationUpdateReq struct {
	g.Meta     `path:"/updateRotation" method:"put"`
	RotationId int `json:"rotation_id"  description:"轮播图ID" v:"required#ID不能为空"`
	RotationInfo
}
type RotationUpdateRes struct {
}
