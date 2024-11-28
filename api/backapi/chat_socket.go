package backapi

import "github.com/gogf/gf/v2/frame/g"

type ChatSocket struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}
type ChatSendMessageReq struct {
	g.Meta `path:"/ws?{Id}" method:"post"`
	ChatSocket
}
type ChatSendMessageRes struct {
}
type ChatGetMessageReq struct {
	g.Meta `path:"/ws/get" method:"get"`
}
type ChatGetMessageRes struct {
	List interface{} `json:"list"`
}
