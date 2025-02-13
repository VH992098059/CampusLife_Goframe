package model

type ActivityListModelInput struct {
	Page int `json:"page"`
}
type ActivityListModelOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
type GetActivityModelInput struct {
	Uuid string `json:"uuid"`
}
type GetActivityModelOutput struct {
	List interface{} `json:"list"`
}
type ActivitySearchModelInput struct {
	ActivityTitle string `json:"activity_title"` // 活动标题
	Page          int    `json:"page" description:"分页码" `
	Size          int    `json:"size" description:"分页数量"`
}
type ActivitySearchModelOutput struct {
	List  interface{} `json:"list"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
