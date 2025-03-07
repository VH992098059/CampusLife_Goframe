package backapi

type LoginRes struct {
	Uuid     string `json:"uuid"`
	Type     string `json:"type"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
}
