package common

import (
	"fmt"
	"github.com/cloopen/go-sms-sdk/cloopen"
	"log"
	"math/rand"
	"time"
	"zg5/z304/framework/redis"
)

func Code(tel string) error {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("2c94811c8b1e335b018b5fb39e5c0c35").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("7be8b1c3b8824f87baae5557ea4b7513")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	str := ""
	for i := 0; i < 4; i++ {
		str += fmt.Sprintf("%v", rand.Intn(10))
	}
	redis.RE.Set("tel", str, time.Minute*1)
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "2c94811c8b1e335b018b5fb39fcc0c3c",
		// 手机号码
		To: tel,
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{str},
	}
	// 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)
	return nil
}
