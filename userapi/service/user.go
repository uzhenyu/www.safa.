package service

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"time"
	"zg5/z304/userapi/client"
	"zg5/z304/userapi/utils"
)

type LoginRes struct {
	Token  string `json:"Token"`
	UserId int64  `json:"UserId"`
}

func Login(ctx context.Context, code int64, tel, password string) (*LoginRes, error) {
	info, err := client.GetUser(ctx, code, tel, password)
	if err != nil {
		return nil, err
	}
	logs.Info(info.ID)
	token, err := utils.GetJwtToken(time.Now().Unix(), int64(time.Hour*1), info.ID)
	return &LoginRes{
		Token:  token,
		UserId: info.ID,
	}, nil
}

func Message(ctx context.Context, tel string) error {
	err := client.Message(ctx, tel)
	if err != nil {
		return err
	}
	return nil
}
