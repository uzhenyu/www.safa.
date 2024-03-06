package user

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"zg5/z304/framework/http"
	"zg5/z304/userapi/conts"
	"zg5/z304/userapi/service"
)

func Login(c *gin.Context) {
	var loginReq struct {
		ID       int64  `json:"ID"`
		Tel      string `json:"Tel"`
		Password string `json:"Password"`
		Code     int64  `json:"Code"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Res(c, conts.PPM_ERROR, nil, err.Error())
		return
	}
	info, err := service.Login(c, loginReq.Code, loginReq.Tel, loginReq.Password)
	if err != nil {
		http.Res(c, conts.PPM_ERROR, nil, err.Error())
		return
	}
	logs.Info(info)
	http.Res(c, conts.SUCCESS, info, "登陆成功")
	return
}

func Message(c *gin.Context) {
	var MessageReq struct {
		Tel string `json:"Tel"`
	}
	if err := c.ShouldBind(&MessageReq); err != nil {
		http.Res(c, conts.PPM_ERROR, nil, err.Error())
	}
	err := service.Message(c, MessageReq.Tel)
	if err != nil {
		http.Res(c, conts.PPM_ERROR, nil, err.Error())
		return
	}
	http.Res(c, conts.SUCCESS, nil, "发送成功")
	return
}
