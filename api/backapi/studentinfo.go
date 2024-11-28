package backapi

import "github.com/gogf/gf/v2/frame/g"

type StudentInfo struct {
	StudentId string `json:"student_id" v:"required|length:12,12#学号不能为空|学生ID必须在{max}位"`
	UserId    string `json:"user_id" v:"required|length:14,14#用户编号不能为空|用户ID必须在{max}位"`
	Name      string `json:"name" v:"required|length:0,10#姓名不能为空|姓名ID必须在{max}位"`
	IsOn      string `json:"is_on" v:"required#请填写学生在校状态"`
	SchoolId  int    `json:"school_id" v:"required#请填写学校编号"`
}
type StudentListReq struct {
	g.Meta `path:"/studentList" method:"get"`
	CommonPaginationReq
}
type StudentListRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type StudentCreateReq struct {
	g.Meta `path:"/studentCreate" method:"post"`
	StudentInfo
}
type StudentCreateRes struct {
	StudentId int `json:"student_id"`
}
type StudentUpdateReq struct {
	g.Meta `path:"/studentUpdate" method:"put"`
	StudentInfo
}
type StudentUpdateRes struct {
}
type StudentDeleteReq struct {
	g.Meta    `path:"/studentDelete" method:"delete"`
	StudentId int `json:"student_id" v:"required|length:12,12#学号不能为空|学生ID必须在{max}位"`
}
type StudentDeleteRes struct{}
