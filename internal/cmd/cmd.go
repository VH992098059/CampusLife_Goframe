package cmd

import (
	"context"
	"demo3/api/backapi"
	"demo3/internal/consts"
	"demo3/internal/controller"
	"demo3/internal/dao"
	"demo3/internal/model"
	"demo3/internal/service"
	"demo3/utility"
	"demo3/utility/middleware"
	"demo3/utility/response"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"math/rand"
	"time"

	"demo3/utility/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var UserTimeout = 24 * 60 * 60 * 1000
var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			loginFunc := Login
			loginAfter := LoginAfter
			gfToken := gtoken.GfToken{
				Timeout:         UserTimeout,
				ServerName:      "demo",
				LoginPath:       "/login",
				LoginBeforeFunc: loginFunc,
				LoginAfterFunc:  loginAfter,
				LogoutPath:      "/logout",
				MultiLogin:      true,
				CacheMode:       gtoken.CacheModeRedis,
			}

			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Bind(
					controller.RegisterUserInfo,
				)
				s.BindHandler("/sendCode", func(r *ghttp.Request) {
					phone := r.Get("phone").String()
					code := generateCode()
					// 这里假设有一个发送短信的函数sendSMS
					if phone == "" {
						r.Response.WriteJson(g.Map{
							"status":  "false",
							"message": "验证码发送失败",
						})
					} else {
						if err := sendSMS(phone, code); err != nil {
							r.Response.WriteJson(g.Map{
								"status":  "error",
								"message": "发送验证码失败",
							})
						} else {
							r.Response.WriteJson(g.Map{
								"status":  "success",
								"message": "验证码发送成功",
							})
						}
					}

				})
				group.Middleware(
					//是否允许跨域操作
					service.Middleware().CORS,
					middleware.MiddlewareHandlerResponse,
				)
				//验证token令牌 JWT
				/*group.Bind(
					controller.LoginInfo,
					controller.LogoutInfo,

				)*/
				/*group.Middleware(service.Middleware().Auth)*/

				/*gtoken*/
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					controller.Userinfo,
				)

			})
			s.Run()
			return nil
		},
	}
)

// Login 登录操作
func Login(r *ghttp.Request) (string, interface{}) {
	username := r.Get("username").String()
	password := r.Get("password").String()
	ctx := context.TODO()
	if username == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	UserInfo := model.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where("username", username).Scan(&UserInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	if utility.EncryptPassword(password, UserInfo.Usersalt) != UserInfo.Password {
		r.Response.WriteJson(gtoken.Fail("账号或密码错误"))
		r.ExitAll()
	}
	UserInfo.Password = ""
	UserInfo.Usersalt = ""
	return consts.GTokenUserPrefix + UserInfo.UserId, UserInfo
}

// LoginAfter 登录后操作
func LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 500
		r.Response.WriteJson(respData)
		r.ExitAll()
	} else {
		respData.Msg = "登录成功"
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenUserPrefix)
		userInfo := model.UserRedis{}
		err := dao.UserInfo.Ctx(context.TODO()).Where(userId).Scan(&userInfo)
		if err != nil {
			return
		}
		response.SuccessWithData(r, &backapi.LoginRes{
			Type:     "redis",
			Nickname: userInfo.NickName,
			Token:    respData.GetString("token"),
			ExpireIn: UserTimeout / 1000,
		})

	}
}
func generateCode() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", rand.Intn(1000000)) // 生成6位随机数
}
func sendSMS(phone, code string) error {
	// 实现发送短信的逻辑，这里只做一个简单的模拟
	fmt.Printf("Sending SMS to %s: Your verification code is %s\n", phone, code)
	return nil
}
