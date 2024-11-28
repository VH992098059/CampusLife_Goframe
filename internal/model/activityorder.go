package model

type ActivityOrder struct {
	Status       int `json:"status"`         // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityID   int `json:"activity_id"`    // 活动ID
	CommonUserID int `json:"common_user_id"` // 普通用户ID
}
type ActivityOrderAddModelInput struct {
	Status       int `json:"status"`         // 状态（1：待审核 2：待参加 3：已检票 4：已取消）
	ActivityID   int `json:"activity_id"`    // 活动ID
	CommonUserID int `json:"common_user_id"` // 普通用户ID
}
type ActivityOrderAddModelOutput struct{}
