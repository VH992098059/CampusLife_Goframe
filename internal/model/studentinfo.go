package model

type StudentInfo struct {
	StudentId string `json:"student_id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name"`
	IsOn      string `json:"is_on"`
	SchoolId  int    `json:"school_id"`
}
type StudentInfoListInput struct {
	//g.Meta `path:"/studentList" method:"get"`
	Page  int `json:"page" description:"分页码" `
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type StudentInfoListOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type StudentInfoCreateInput struct {
	//g.Meta    `path:"/studentCreate" method:"post"`
	StudentId string `json:"student_id"`
	UserId    string `json:"user_id"`
	Name      string `json:"name" `
	IsOn      string `json:"is_on" `
	SchoolId  int    `json:"school_id" `
}
type StudentInfoCreateOutput struct {
	StudentId int `json:"student_id"`
}
type StudentInfoUpdateInput struct {
	//g.Meta    `path:"/studentUpdate" method:"post"`
	StudentId string `json:"student_id"`
	UserId    string `json:"user_id" `
	Name      string `json:"name"`
	IsOn      string `json:"is_on"`
	SchoolId  int    `json:"school_id"`
}
type StudentInfoUpdateOutput struct {
}
type StudentInfoDeleteInput struct {
	//g.Meta    `path:"/studentDelete" method:"post"`
	StudentId int `json:"student_id"`
}
type StudentInfoDeleteOutput struct{}
