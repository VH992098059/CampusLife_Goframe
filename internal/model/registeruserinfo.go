package model

type RegisterUserModelInput struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Usersalt string `json:"user_salt"`
	Type     string `json:"type"`
}
type RegisterUserModelOutput struct {
	UserId string `json:"user_id"`
}
