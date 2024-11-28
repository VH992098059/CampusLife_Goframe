package model

type RotationModelInfo struct {
	PicUrl       string `json:"pic_url"   description:"轮播图路径" v:"required#图片路径不能为空"`
	RotationName string `json:"rotation_name" description:"轮播图图片名"  v:"required#活动名称不能为空"`
}
type RotationModelCreateInput struct {
	PicUrl       string `json:"pic_url"   description:"轮播图路径" v:"required#图片路径不能为空"`
	RotationName string `json:"rotation_name" description:"轮播图图片名"  v:"required#活动名称不能为空"`
}
type RotationModelCreateOutput struct {
	RotationId int `json:"rotation_id"`
}
type RotationModelListInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type RotationModelListOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type RotationModelUpdateInput struct {
	RotationId   int    `json:"rotation_id" description:"轮播图ID" v:"required#ID不能为空"`
	PicUrl       string `json:"pic_url"   description:"轮播图路径" v:"required#图片路径不能为空"`
	RotationName string `json:"rotation_name" description:"轮播图图片名"  v:"required#活动名称不能为空"`
}
type RotationModelUpdateOutput struct {
}
type RotationModelDeleteInput struct {
	RotationId int `json:"rotation_id"`
}
