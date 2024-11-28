package model

type ActivityOrder struct {
	Status       int `json:"status"`         // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityID   int `json:"activity_id"`    // 活动ID
	CommonUserID int `json:"common_user_id"` // 普通用户ID
}
type ActivityOrderAddModelInput struct {
	Uuid           string `json:"uuid"`
	Status         int    `json:"status"`          // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityNumber string `json:"activity_number"` // 活动ID
	CommonUserID   int    `json:"common_user_id"`  // 普通用户ID
}
type ActivityOrderAddModelOutput struct{}

type ActivityOrderUpdateModelInput struct {
	Status       int `json:"status"`         // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityID   int `json:"activity_id"`    // 活动ID
	CommonUserID int `json:"common_user_id"` // 普通用户ID
}
type ActivityOrderUpdateModelOutput struct{}
type ActivityOrderDeleteModelInput struct {
	ActivityID int `json:"activity_id"`
}
type ActivityOrderDeleteModelOutput struct{}
type ActivityOrderListModelInput struct {
	Page   int `json:"page" description:"分页码" `
	Size   int `json:"size" description:"分页数量"`
	Offset int `json:"offset" dc:"初始容量"`
	Limit  int `json:"limit" dc:"最大容量"`
}
type ActivityOrderListModelOutput struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码" `
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
