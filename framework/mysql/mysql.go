package mysql

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zg5/z304/framework/config"
)

var DB *gorm.DB

type NacosMysql struct {
	Mysql struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		Database string `json:"Database"`
	} `json:"Mysql"`
}

func InitMysql(serviceName string) error {
	err := config.GetClient()
	if err != nil {
		return err
	}

	configs, err := config.GetConfig(serviceName, "wzy")
	if err != nil {
		return err
	}
	logs.Info(configs)
	var nacos NacosMysql
	err = json.Unmarshal([]byte(configs), &nacos)
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		nacos.Mysql.Username,
		nacos.Mysql.Password,
		nacos.Mysql.Host,
		nacos.Mysql.Port,
		nacos.Mysql.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
