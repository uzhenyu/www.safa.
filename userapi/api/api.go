package api

import (
	"github.com/gin-gonic/gin"
	"zg5/z304/userapi/api/user"
)

func Reg(r *gin.Engine) {
	user.Reg(r)
}
