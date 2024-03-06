package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Req struct {
	Code int64
	Msg  string
	Data interface{}
}

func Ress(c *gin.Context, code int64, msg string, data interface{}) {
	httpStatus := http.StatusOK
	if code > 10000 {
		httpStatus = http.StatusBadGateway
	}
	c.JSON(httpStatus, &Req{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
