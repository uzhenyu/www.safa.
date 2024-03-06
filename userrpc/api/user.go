package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zg5/z304/message/user"
	"zg5/z304/userrpc/service"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (u UserServer) Message(ctx context.Context, request *user.MessageRequest) (*user.MessageResponse, error) {
	err := service.Message(request.Tel)
	if err != nil {
		return nil, err
	}
	return &user.MessageResponse{}, nil
}

func (u UserServer) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	getUser, err := service.GetUser()
	if err != nil {
		return nil, err
	}
	if getUser.Tel != request.Tel {
		return nil, status.Error(codes.InvalidArgument, "账号不正确")
	}
	if getUser.Password != request.Password {
		return nil, status.Error(codes.InvalidArgument, "密码不正确")
	}
	boo1 := service.Tel(request.Tel)
	if boo1 == false {
		return nil, status.Error(codes.InvalidArgument, "手机号格式不正确")
	}
	boo2 := service.Code(request.Code, request.Tel)
	if boo2 == false {
		return nil, status.Error(codes.InvalidArgument, "验证码不正确")
	}
	return &user.LoginResponse{Info: getUser}, nil
}
