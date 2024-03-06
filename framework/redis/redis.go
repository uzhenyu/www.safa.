package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"zg5/z304/framework/config"
)

var (
	RE *redis.Client
)

type NacosRedis struct {
	Redis struct {
		Host string `json:"Host"`
		Port string `json:"Port"`
	} `json:"Redis"`
}

func InitRedis(serviceName string) error {
	err := config.GetClient()
	if err != nil {
		return err
	}

	configs, err := config.GetConfig(serviceName, "wzy")
	if err != nil {
		return err
	}

	var nacos NacosRedis
	err = json.Unmarshal([]byte(configs), &nacos)
	if err != nil {
		return err
	}

	RE = redis.NewClient(&redis.Options{Addr: nacos.Redis.Host + ":" + nacos.Redis.Port})
	return err
}
