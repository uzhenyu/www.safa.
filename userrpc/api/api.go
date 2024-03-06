package api

import (
	"google.golang.org/grpc"
	"zg5/z304/message/user"
)

func Reg(c grpc.ServiceRegistrar) {
	user.RegisterUserServer(c, UserServer{})
}
