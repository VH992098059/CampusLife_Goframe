package model

type JoinActivityModelInput struct {
	JoinUuid     string `json:"join_uuid"`
	UserId       string `json:"user_id"`
	ActivityUuid string `json:"activity_uuid"`
}
type JoinActivityModelOutput struct {
}
