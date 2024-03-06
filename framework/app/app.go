package app

import (
	"zg5/z304/framework/config"
	"zg5/z304/framework/mysql"
	"zg5/z304/framework/redis"
)

func Init(serviceName string, apps ...string) error {
	var err error
	err = config.GetClient()
	if err != nil {
		return err
	}
	for _, v := range apps {
		switch v {
		case "mysql":
			err = mysql.InitMysql(serviceName)
			if err != nil {
				panic(err)
			}
		case "redis":
			err = redis.InitRedis(serviceName)
			if err != nil {
				panic(err)
			}
		}
	}
	return err
}
