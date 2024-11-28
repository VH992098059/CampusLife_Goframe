package backapi

import "github.com/gogf/gf/v2/frame/g"

type DormitoryInfo struct {
	//DormitoryId      int    `json:"dormitoryId" v:"required#宿舍ID不能为空"`
	DormitoryType   string `json:"dormitory_type" v:"required#宿舍类型不能为空"`
	Floor           int    `json:"floor" v:"required#楼层不能为空"`
	DormitoryNumber int    `json:"dormitory_number" v:"required#宿舍号不能为空"`
}
type DormitoryInfoListReq struct {
	g.Meta `path:"/dormitoryInfoList" method:"get"`
	CommonPaginationReq
}
type DormitoryInfoListRes struct {
	List  interface{} `dc:"列表数据"`
	Total int         `dc:"总数"`
	Page  int         `dc:"分页号码"`
	Size  int         `dc:"分页数量"`
}

type DormitoryInfoCreateReq struct {
	g.Meta `path:"/dormitoryCreate" method:"post"`
	DormitoryInfo
}
type DormitoryInfoCreateRes struct {
	DormitoryId int `json:"dormitoryId"`
}
type DormitoryInfoUpdateReq struct {
	g.Meta `path:"/dormitoryUpdate" method:"put"`
	DormitoryInfo
}
type DormitoryInfoUpdateRes struct{}
type DormitoryInfoDeleteReq struct {
	g.Meta      `path:"/dormitoryDelete" method:"delete"`
	DormitoryId int `json:"dormitory_id"`
}
type DormitoryInfoDeleteRes struct{}
