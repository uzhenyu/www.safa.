package main

import (
	"github.com/gin-gonic/gin"
	"zg5/z304/userapi/api"
)

func main() {
	//err := app.Init(conts.ServiceName, "mysql")
	//if err != nil {
	//	return
	//}
	r := gin.Default()
	api.Reg(r)
	r.Run(":8082")
}
