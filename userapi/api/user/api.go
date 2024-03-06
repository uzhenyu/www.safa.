package user

import "github.com/gin-gonic/gin"

func Reg(r *gin.Engine) {
	r.POST("/login", Login)
	r.POST("/message", Message)
}
