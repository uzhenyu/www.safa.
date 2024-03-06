package service

import (
	"regexp"
	"strconv"
	"zg5/z304/framework/redis"
	"zg5/z304/message/user"
	"zg5/z304/userrpc/common"
	"zg5/z304/userrpc/model"
)

func GetUser() (*user.UserInfo, error) {
	newUser := model.NewUser()
	res, err := newUser.Get(newUser)
	if err != nil {
		return nil, err
	}
	return mysqlToPb(res)
}

func mysqlToPb(info *model.User) (*user.UserInfo, error) {
	return &user.UserInfo{
		ID:       int64(info.ID),
		Tel:      info.Tel,
		Password: info.Password,
	}, nil
}

func pbToMysql(info *model.User) (*user.UserInfo, error) {
	return &user.UserInfo{
		ID:       int64(info.ID),
		Tel:      info.Tel,
		Password: info.Password,
	}, nil
}

func Tel(tel string) bool {
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	if reg.MatchString(tel) {
		return true
	}
	return false
}

func Message(tel string) error {
	err := common.Code(tel)
	if err != nil {
		return err
	}
	return nil
}

func Code(code int64, tel string) bool {
	codes, _ := redis.RE.Get("tel").Result()
	if strconv.FormatInt(code, 10) != codes {
		return false
	}
	return true
}
