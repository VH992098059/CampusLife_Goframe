package aliyunCode

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"math/rand"
	"strings"
	"time"
)

type PhoneCode struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

var ctx = gctx.New()
var accessKeyId = "LTAI5tBEBBQTh5PjHHGp76Q4"
var accessKeySecret = "Z6M5EQ52tPwoUKX8CtO9tCeJLyhlsS"

// GenerateCode 生成4位随机验证码
func GenerateCode() string {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%04d", rand.Intn(10000)) // 生成4位随机数
}

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
		AccessKeyId: tea.String(accessKeyId),
		// 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
		AccessKeySecret: tea.String(accessKeySecret),
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// ValidMain /*电话，阿里云短信模板编号*/
func ValidMain(phone string, TemplateCode string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}
	code := GenerateCode()
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("校园生活"),
		TemplateCode:  tea.String(TemplateCode),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		success, _err := client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}
		fmt.Println(success)
		///*获取message信息*/
		//message := *success.Body.Message /*将获取到的地址赋值到变量，进行解引用操作*/
		//if message == "触发天级流控Permits:10" {
		//	g.Log().Printf(ctx, "获取已到达上限，无法获取验证码")
		//	return nil
		//}
		RedisCode(phone, code) /*发送至Redis*/
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}

	}
	return _err
}

// RedisCode 将验证码存储到 Redis
func RedisCode(phone string, code string) {
	phoneCode := PhoneCode{Code: code, Phone: phone}
	phoneCodeJSON, err := json.Marshal(phoneCode)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	err = g.Redis().SetEX(ctx, "verification("+phone+")", phoneCodeJSON, 60*5)
	if err != nil {
		g.Log().Fatal(ctx, err)
		return
	}
	fmt.Printf("手机号 %s 和验证码 %s 已存储到 Redis\n", phone, code)
}

// 验证手机号和验证码
func VerifyPhoneCode(phone string, code string) bool {
	verName := "verification(" + phone + ")"
	phoneCodeJSON, err := g.Redis().Get(ctx, verName)
	if phoneCodeJSON.String() == "" {
		g.Log().Printf(ctx, "验证码不存在或已过期")
		return false
	}
	if err != nil {
		return false
	}
	var phoneCode PhoneCode
	err = json.Unmarshal([]byte(phoneCodeJSON.String()), &phoneCode)
	if err != nil {
		g.Log().Print(ctx, err)
		return false
	}
	return phoneCode.Code == code
}

// DeletePhoneCode 删除手机号和验证码
func DeletePhoneCode(phone string) {
	_, err := g.Redis().Del(ctx, "verification("+phone+")")
	if err != nil {
		panic(err)
	}
	fmt.Printf("验证码已从 Redis 中删除\n")
}
