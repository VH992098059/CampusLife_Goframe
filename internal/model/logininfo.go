package model

type LoginModelInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginInfoModelInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInfoModelOutput struct {
}
