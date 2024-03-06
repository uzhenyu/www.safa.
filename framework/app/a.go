package app

import (
	"zg5/z304/framework/config"
	"zg5/z304/framework/mysql"
	"zg5/z304/framework/redis"
)

func Inits(group string, apps ...string) error {
	var err error
	err = config.GetClient()
	if err != nil {
		return err
	}
	for _, v := range apps {
		switch v {
		case "mysql":
			err = mysql.InitMysql(group)
			if err != nil {
				return err
			}
		case "redis":
			err = redis.InitRedis(group)
			if err != nil {
				return err
			}
		}
	}
	return err
}
