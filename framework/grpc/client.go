package grpc

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"zg5/z304/framework/config"
)

type T struct {
	App struct {
		Ip     string `json:"ip"`
		Port   string `json:"port"`
		Secret string `json:"secret"`
	} `json:"app"`
}

func Client(toService string) (*grpc.ClientConn, error) {
	service, i, err := config.NacosService()
	if err != nil {
		return nil, err
	}
	cnf := new(T)
	err = json.Unmarshal([]byte(service), &cnf)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("%v:%v", service, i), grpc.WithTransportCredentials(insecure.NewCredentials()))
}
