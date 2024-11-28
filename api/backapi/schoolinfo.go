package backapi

import "github.com/gogf/gf/v2/frame/g"

type SchoolInfo struct {
	SchoolId   string `json:"school_id" v:"required|length:5,5#学号编号不能为空|学校编号长度必须为{max}位"`
	SchoolName string `json:"school_name"  v:"required#学校名称不能为空"`
	Province   string `json:"school_province"  v:"required#省、自治区、直辖市和特别行政区不能为空"`
	City       string `json:"school_city"  v:"required#城市名称不能为空"`
	County     string `json:"school_county"  v:"required#县（市、区）不能为空"`
	Address    string `json:"school_address"  v:"required#详细地址不能为空"`
}

// SchoolInfoCreateReq 创建学校
type SchoolInfoCreateReq struct {
	g.Meta `path:"/schoolcreate" method:"post"`
	SchoolInfo
}
type SchoolInfoCreateRes struct {
	UserId string `json:"user_id"`
}

// SchoolInfoUpdateReq 更新学校信息
type SchoolInfoUpdateReq struct {
	g.Meta `path:"/schoolupdate" method:"put"`
	SchoolInfo
}
type SchoolInfoUpdateRes struct {
}

// SchoolInfoDeleteReq 删除学校信息
type SchoolInfoDeleteReq struct {
	g.Meta   `path:"/schooldelete" method:"delete"`
	SchoolId string `json:"school_id"`
}
type SchoolInfoDeleteRes struct {
}

// SchoolInfoGetListReq 获取学校信息
type SchoolInfoGetListReq struct {
	g.Meta `path:"/schoolgetlist" method:"get"`
	CommonPaginationReq
}
type SchoolInfoGetListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
