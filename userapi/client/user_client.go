package client

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"zg5/z304/framework/grpc"
	"zg5/z304/message/user"
	"zg5/z304/userapi/conts"
)

type handler func(ctx context.Context, cli user.UserClient) (interface{}, error)

func withClient(ctx context.Context, handler handler) (interface{}, error) {
	conn, err := grpc.Client(conts.ServiceName)
	if err != nil {
		return nil, err
	}
	userCli := user.NewUserClient(conn)
	res, err := handler(ctx, userCli)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return res, nil
}

func GetUser(ctx context.Context, Code int64, Tel, Password string) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.Login(ctx, &user.LoginRequest{
			Code:     Code,
			Tel:      Tel,
			Password: Password,
		})
		if err != nil {
			return nil, err
		}
		logs.Info(res.Info)
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}
func Message(ctx context.Context, Tel string) error {
	_, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		_, err := cli.Message(ctx, &user.MessageRequest{Tel: Tel})
		if err != nil {
			return nil, err
		}
		return nil, err
	})
	if err != nil {
		return err
	}
	return nil
}
