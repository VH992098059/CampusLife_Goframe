package backapi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginInfo struct {
	UserName string `json:"username" v:"required#用户名不能为空"`
	Password string `json:"password" v:"required|password#密码不能为空"`
	Phone    string `json:"phone" v:"required-without:UserName"`
}
type LoginReq struct {
	g.Meta `path:"/logininfo" method:"post"`
	LoginInfo
}

//JWT验证
/*type LoginRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}*/

type AuthLogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type AuthLogoutRes struct {
}
