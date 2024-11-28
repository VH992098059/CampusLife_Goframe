package model

type SchoolInfoBase struct {
	SchoolId   string `json:"school_id" v:"required|length:5,5#学号编号不能为空|学校编号长度不能少于{max}位也不能超过{max}位"`
	SchoolName string `json:"school_name"  v:"required#学校名称不能为空"`
	Province   string `json:"school_province"  v:"required#省、自治区、直辖市和特别行政区不能为空"`
	City       string `json:"school_city"  v:"required#城市名称不能为空"`
	County     string `json:"school_county"  v:"required#县（市、区）不能为空"`
	Address    string `json:"school_address"  v:"required#详细地址不能为空"`
}
type SchoolInfoCreateInput struct {
	SchoolId   string `json:"school_id" v:"required|length:5,5#学号编号不能为空|学校编号长度不能少于{max}位也不能超过{max}位"`
	SchoolName string `json:"school_name"  v:"required#学校名称不能为空"`
	Province   string `json:"school_province"  v:"required#省、自治区、直辖市和特别行政区不能为空"`
	City       string `json:"school_city"  v:"required#城市名称不能为空"`
	County     string `json:"school_county"  v:"required#县（市、区）不能为空"`
	Address    string `json:"school_address"  v:"required#详细地址不能为空"`
}
type SchoolInfoCreateOutput struct {
	SchoolId string `json:"school_id"`
}
type SchoolInfoUpdateInput struct {
	SchoolId   string `json:"school_id" v:"required|length:5,5#学号编号不能为空|学校编号长度不能少于{max}位也不能超过{max}位"`
	SchoolName string `json:"school_name"  v:"required#学校名称不能为空"`
	Province   string `json:"school_province"  v:"required#省、自治区、直辖市和特别行政区不能为空"`
	City       string `json:"school_city"  v:"required#城市名称不能为空"`
	County     string `json:"school_county"  v:"required#县（市、区）不能为空"`
	Address    string `json:"school_address"  v:"required#详细地址不能为空"`
}
type SchoolInfoUpdateOutput struct{}

type SchoolInfoDeleteOutput struct{}

type SchoolInfoGetInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type SchoolInfoGetOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
