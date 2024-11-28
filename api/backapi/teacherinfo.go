package backapi

import "github.com/gogf/gf/v2/frame/g"

type TeacherInfo struct {
	//TeacherID  int    `json:"teacher_id" v:"required#职工编号不能为空" description:"职工编号"` // 职工编号
	Name       string `json:"name" v:"required#教师名称不能为空" description:"教师名称"`     // 教师名称
	Age        int    `json:"age" v:"required#年龄不能为空" description:"年龄"`          // 年龄
	Department string `json:"department" v:"required#所在系不能为空" description:"所在系"` // 所在系
	Phone      string `json:"phone" v:"required#教师电话不能为空" description:"教师电话"`    // 教师电话
}
type TeacherInfoAddReq struct {
	g.Meta `path:"/teacherAdd" method:"post"`
	TeacherInfo
}
type TeacherInfoAddRes struct{}
type TeacherInfoListReq struct {
	g.Meta `path:"/teacherList" method:"GET"`
	CommonPaginationReq
}
type TeacherInfoListRes struct {
	List  interface{} `dc:"列表数据"`
	Total int         `dc:"总数"`
	Page  int         `dc:"分页号码"`
	Size  int         `dc:"分页数量"`
}
type TeacherInfoDeleteReq struct {
	g.Meta    `path:"/teacherDelete" method:"delete"`
	TeacherID int `json:"teacher_id" v:"required#职工编号不能为空" description:"职工编号"` // 职工编号
}
type TeacherInfoDeleteRes struct{}
type TeacherInfoUpdateReq struct {
	g.Meta `path:"/teacherInfo" method:"put"`
	TeacherInfo
}
type TeacherInfoUpdateRes struct{}
