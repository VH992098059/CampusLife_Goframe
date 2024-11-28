package model

type DormitoryModelInfo struct {
	//DormitoryId      int    `json:"dormitoryId"`
	DormitoryType   string `json:"dormitory_type"`
	Floor           int    `json:"floor"`
	DormitoryNumber int    `json:"dormitory_number"`
}
type DormitoryInfoModelListInput struct {
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
type DormitoryInfoModelListOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type DormitoryInfoModelCreateInput struct {
	DormitoryType   string `json:"dormitory_type"`
	Floor           int    `json:"floor"`
	DormitoryNumber int    `json:"dormitory_number"`
}
type DormitoryInfoModelCreateOutput struct {
	DormitoryId int `json:"dormitoryId"`
}
type DormitoryInfoModelUpdateInput struct {
	DormitoryId     int    `json:"dormitoryId"`
	DormitoryType   string `json:"dormitory_type"`
	Floor           int    `json:"floor"`
	DormitoryNumber int    `json:"dormitory_number"`
}
type DormitoryInfoModelUpdateOutput struct{}

/*
	type DormitoryInfoModelDeleteInput struct {
		DormitoryId      int    `json:"dormitoryId"`
	}
*/
type DormitoryInfoModelDeleteOutput struct{}
