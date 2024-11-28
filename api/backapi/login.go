package backapi

type LoginRes struct {
	Type     string `json:"type"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
}
