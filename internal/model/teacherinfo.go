package model

type TeacherInfo struct {
	//TeacherID  int    `json:"teacher_id"` // 职工编号
	Name       string `json:"name"`       // 教师名称
	Age        int    `json:"age"`        // 年龄
	Department string `json:"department"` // 所在系
	Phone      string `json:"phone"`      // 教师电话
}

type TeacherInfoModelAddInput struct {
	TeacherID  int    `json:"teacher_id"` // 职工编号
	Name       string `json:"name"`       // 教师名称
	Age        int    `json:"age"`        // 年龄
	Department string `json:"department"` // 所在系
	Phone      string `json:"phone"`      // 教师电话
}
type TeacherInfoModelAddOutput struct{}
type TeacherInfoModelUpdateInput struct {
	Name       string `json:"name"`       // 教师名称
	Age        int    `json:"age"`        // 年龄
	Department string `json:"department"` // 所在系
	Phone      string `json:"phone"`      // 教师电话
}
type TeacherInfoModelUpdateOutput struct{}
type TeacherInfoModelDeleteInput struct {
	TeacherID int `json:"teacher_id"` // 职工编号
}
type TeacherInfoModelDeleteOutput struct{}
type TeacherInfoModelListInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type TeacherInfoModelListOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
